package shares

import (
	"encoding/binary"
	"errors"
	"fmt"

	appns "github.com/riemalabs/nubit-app/da/namespace"
	"github.com/riemalabs/nubit-app/utils/appconsts"
	coretypes "github.com/tendermint/tendermint/types"
	"golang.org/x/exp/slices"
)

// CompactShareSplitter will write raw data compactly across a progressively
// increasing set of shares. It is used to lazily split block data such as
// transactions or intermediate state roots into shares.
type CompactShareSplitter struct {
	shares []Share
	// pendingShare Share
	shareBuilder *Builder
	namespace    appns.Namespace
	done         bool
	shareVersion uint8
	// shareRanges is a map from a transaction key to the range of shares it
	// occupies. The range assumes this compact share splitter is the only
	// thing in the data square (e.g. the range for the first tx starts at index
	// 0).
	shareRanges map[coretypes.TxKey]Range
}

// NewCompactShareSplitter returns a CompactShareSplitter using the provided
// namespace and shareVersion.
func NewCompactShareSplitter(ns appns.Namespace, shareVersion uint8) *CompactShareSplitter {
	sb, err := NewBuilder(ns, shareVersion, true).Init()
	if err != nil {
		panic(err)
	}

	return &CompactShareSplitter{
		shares:       []Share{},
		namespace:    ns,
		shareVersion: shareVersion,
		shareRanges:  map[coretypes.TxKey]Range{},
		shareBuilder: sb,
	}
}

// WriteTx adds the delimited data for the provided tx to the underlying compact
// share splitter.
func (css *CompactShareSplitter) WriteTx(tx coretypes.Tx) error {
	rawData, err := MarshalDelimitedTx(tx)
	if err != nil {
		return fmt.Errorf("included Tx in mem-pool that can not be encoded %v", tx)
	}

	startShare := len(css.shares)

	if err := css.write(rawData); err != nil {
		return err
	}
	endShare := css.Count()
	css.shareRanges[tx.Key()] = NewRange(startShare, endShare)

	return nil
}

// write adds the delimited data to the underlying compact shares.
func (css *CompactShareSplitter) write(rawData []byte) error {
	if css.done {
		// remove the last element
		if !css.shareBuilder.IsEmptyShare() {
			css.shares = css.shares[:len(css.shares)-1]
		}
		css.done = false
	}

	if err := css.shareBuilder.MaybeWriteReservedBytes(); err != nil {
		return err
	}

	for {
		rawDataLeftOver := css.shareBuilder.AddData(rawData)
		if rawDataLeftOver == nil {
			break
		}
		if err := css.stackPending(); err != nil {
			return err
		}

		rawData = rawDataLeftOver
	}

	if css.shareBuilder.AvailableBytes() == 0 {
		if err := css.stackPending(); err != nil {
			return err
		}
	}
	return nil
}

// stackPending will build & add the pending share to accumulated shares
func (css *CompactShareSplitter) stackPending() error {
	pendingShare, err := css.shareBuilder.Build()
	if err != nil {
		return err
	}
	css.shares = append(css.shares, *pendingShare)

	// Now we need to create a new builder
	css.shareBuilder, err = NewBuilder(css.namespace, css.shareVersion, false).Init()
	return err
}

// Export returns the underlying compact shares
func (css *CompactShareSplitter) Export() ([]Share, error) {
	if css.isEmpty() {
		return []Share{}, nil
	}

	// in case Export is called multiple times
	if css.done {
		return css.shares, nil
	}

	var bytesOfPadding int
	// add the pending share to the current shares before returning
	if !css.shareBuilder.IsEmptyShare() {
		bytesOfPadding = css.shareBuilder.ZeroPadIfNecessary()
		if err := css.stackPending(); err != nil {
			return []Share{}, err
		}
	}

	sequenceLen := css.sequenceLen(bytesOfPadding)
	if err := css.writeSequenceLen(sequenceLen); err != nil {
		return []Share{}, err
	}
	css.done = true
	return css.shares, nil
}

// ShareRanges returns a map of share ranges to the corresponding tx keys. All
// share ranges in the map of shareRanges will be offset (i.e. incremented) by
// the shareRangeOffset provided. shareRangeOffset should be 0 for the first
// compact share sequence in the data square (transactions) but should be some
// non-zero number for subsequent compact share sequences (e.g. pfb txs).
func (css *CompactShareSplitter) ShareRanges(shareRangeOffset int) map[coretypes.TxKey]Range {
	// apply the shareRangeOffset to all share ranges
	shareRanges := make(map[coretypes.TxKey]Range, len(css.shareRanges))

	for k, v := range css.shareRanges {
		shareRanges[k] = Range{
			Start: v.Start + shareRangeOffset,
			End:   v.End + shareRangeOffset,
		}
	}

	return shareRanges
}

// writeSequenceLen writes the sequence length to the first share.
func (css *CompactShareSplitter) writeSequenceLen(sequenceLen uint32) error {
	if css.isEmpty() {
		return nil
	}

	// We may find a more efficient way to write seqLen
	b, err := NewBuilder(css.namespace, css.shareVersion, true).Init()
	if err != nil {
		return err
	}
	b.ImportRawShare(css.shares[0].ToBytes())
	if err := b.WriteSequenceLen(sequenceLen); err != nil {
		return err
	}

	firstShare, err := b.Build()
	if err != nil {
		return err
	}

	// replace existing first share with new first share
	css.shares[0] = *firstShare

	return nil
}

// sequenceLen returns the total length in bytes of all units (transactions or
// intermediate state roots) written to this splitter. sequenceLen does not
// include the number of bytes occupied by the namespace ID, the share info
// byte, or the reserved bytes. sequenceLen does include the unit length
// delimiter prefixed to each unit.
func (css *CompactShareSplitter) sequenceLen(bytesOfPadding int) uint32 {
	if len(css.shares) == 0 {
		return 0
	}
	if len(css.shares) == 1 {
		return uint32(appconsts.FirstCompactShareContentSize) - uint32(bytesOfPadding)
	}

	continuationSharesCount := len(css.shares) - 1
	continuationSharesSequenceLen := continuationSharesCount * appconsts.ContinuationCompactShareContentSize
	return uint32(appconsts.FirstCompactShareContentSize + continuationSharesSequenceLen - bytesOfPadding)
}

// isEmpty returns whether this compact share splitter is empty.
func (css *CompactShareSplitter) isEmpty() bool {
	return len(css.shares) == 0 && css.shareBuilder.IsEmptyShare()
}

// Count returns the number of shares that would be made if `Export` was invoked
// on this compact share splitter.
func (css *CompactShareSplitter) Count() int {
	if !css.shareBuilder.IsEmptyShare() && !css.done {
		// pending share is non-empty, so it will be zero padded and added to shares during export
		return len(css.shares) + 1
	}
	return len(css.shares)
}

// MarshalDelimitedTx prefixes a transaction with the length of the transaction
// encoded as a varint.
func MarshalDelimitedTx(tx coretypes.Tx) ([]byte, error) {
	lenBuf := make([]byte, binary.MaxVarintLen64)
	length := uint64(len(tx))
	n := binary.PutUvarint(lenBuf, length)
	return append(lenBuf[:n], tx...), nil
}

// SparseShareSplitter lazily splits blobs into shares that will eventually be
// included in a data square. It also has methods to help progressively count
// how many shares the blobs written take up.
type SparseShareSplitter struct {
	shares []Share
}

func NewSparseShareSplitter() *SparseShareSplitter {
	return &SparseShareSplitter{}
}

// Write writes the provided blob to this sparse share splitter. It returns an
// error or nil if no error is encountered.
func (sss *SparseShareSplitter) Write(blob coretypes.Blob) error {
	if !slices.Contains(appconsts.SupportedShareVersions, blob.ShareVersion) {
		return fmt.Errorf("unsupported share version: %d", blob.ShareVersion)
	}

	rawData := blob.Data
	blobNamespace, err := appns.New(blob.NamespaceVersion, blob.NamespaceID)
	if err != nil {
		return err
	}

	// First share
	b, err := NewBuilder(blobNamespace, blob.ShareVersion, true).Init()
	if err != nil {
		return err
	}
	if err := b.WriteSequenceLen(uint32(len(rawData))); err != nil {
		return err
	}

	for rawData != nil {

		rawDataLeftOver := b.AddData(rawData)
		if rawDataLeftOver == nil {
			// Just call it on the latest share
			b.ZeroPadIfNecessary()
		}

		share, err := b.Build()
		if err != nil {
			return err
		}
		sss.shares = append(sss.shares, *share)

		b, err = NewBuilder(blobNamespace, blob.ShareVersion, false).Init()
		if err != nil {
			return err
		}
		rawData = rawDataLeftOver
	}

	return nil
}

// WriteNamespacePaddingShares adds padding shares with the namespace of the
// last written share. This is useful to follow the non-interactive default
// rules. This function assumes that at least one share has already been
// written.
func (sss *SparseShareSplitter) WriteNamespacePaddingShares(count int) error {
	if count < 0 {
		return errors.New("cannot write negative namespaced shares")
	}
	if count == 0 {
		return nil
	}
	if len(sss.shares) == 0 {
		return errors.New("cannot write namespace padding shares on an empty SparseShareSplitter")
	}
	lastBlob := sss.shares[len(sss.shares)-1]
	lastBlobNs, err := lastBlob.Namespace()
	if err != nil {
		return err
	}
	lastBlobInfo, err := lastBlob.InfoByte()
	if err != nil {
		return err
	}
	nsPaddingShares, err := NamespacePaddingShares(lastBlobNs, lastBlobInfo.Version(), count)
	if err != nil {
		return err
	}
	sss.shares = append(sss.shares, nsPaddingShares...)

	return nil
}

// Export finalizes and returns the underlying shares.
func (sss *SparseShareSplitter) Export() []Share {
	return sss.shares
}

// Count returns the current number of shares that will be made if exporting.
func (sss *SparseShareSplitter) Count() int {
	return len(sss.shares)
}