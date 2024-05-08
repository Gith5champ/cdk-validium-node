// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nubit/mint/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryInflationRateRequest is the request type for the Query/InflationRate RPC
// method.
type QueryInflationRateRequest struct {
}

func (m *QueryInflationRateRequest) Reset()         { *m = QueryInflationRateRequest{} }
func (m *QueryInflationRateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryInflationRateRequest) ProtoMessage()    {}
func (*QueryInflationRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_665023cffa247a9d, []int{0}
}
func (m *QueryInflationRateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInflationRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInflationRateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInflationRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInflationRateRequest.Merge(m, src)
}
func (m *QueryInflationRateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryInflationRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInflationRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInflationRateRequest proto.InternalMessageInfo

// QueryInflationRateResponse is the response type for the Query/InflationRate
// RPC method.
type QueryInflationRateResponse struct {
	// InflationRate is the current inflation rate.
	InflationRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation_rate,json=inflationRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rate"`
}

func (m *QueryInflationRateResponse) Reset()         { *m = QueryInflationRateResponse{} }
func (m *QueryInflationRateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryInflationRateResponse) ProtoMessage()    {}
func (*QueryInflationRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_665023cffa247a9d, []int{1}
}
func (m *QueryInflationRateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryInflationRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryInflationRateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryInflationRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInflationRateResponse.Merge(m, src)
}
func (m *QueryInflationRateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryInflationRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInflationRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInflationRateResponse proto.InternalMessageInfo

// QueryAnnualProvisionsRequest is the request type for the
// Query/AnnualProvisions RPC method.
type QueryAnnualProvisionsRequest struct {
}

func (m *QueryAnnualProvisionsRequest) Reset()         { *m = QueryAnnualProvisionsRequest{} }
func (m *QueryAnnualProvisionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAnnualProvisionsRequest) ProtoMessage()    {}
func (*QueryAnnualProvisionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_665023cffa247a9d, []int{2}
}
func (m *QueryAnnualProvisionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAnnualProvisionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAnnualProvisionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAnnualProvisionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAnnualProvisionsRequest.Merge(m, src)
}
func (m *QueryAnnualProvisionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAnnualProvisionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAnnualProvisionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAnnualProvisionsRequest proto.InternalMessageInfo

// QueryAnnualProvisionsResponse is the response type for the
// Query/AnnualProvisions RPC method.
type QueryAnnualProvisionsResponse struct {
	// AnnualProvisions is the current annual provisions.
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annual_provisions"`
}

func (m *QueryAnnualProvisionsResponse) Reset()         { *m = QueryAnnualProvisionsResponse{} }
func (m *QueryAnnualProvisionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAnnualProvisionsResponse) ProtoMessage()    {}
func (*QueryAnnualProvisionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_665023cffa247a9d, []int{3}
}
func (m *QueryAnnualProvisionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAnnualProvisionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAnnualProvisionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAnnualProvisionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAnnualProvisionsResponse.Merge(m, src)
}
func (m *QueryAnnualProvisionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAnnualProvisionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAnnualProvisionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAnnualProvisionsResponse proto.InternalMessageInfo

// QueryGenesisTimeRequest is the request type for the Query/GenesisTime RPC
// method.
type QueryGenesisTimeRequest struct {
}

func (m *QueryGenesisTimeRequest) Reset()         { *m = QueryGenesisTimeRequest{} }
func (m *QueryGenesisTimeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGenesisTimeRequest) ProtoMessage()    {}
func (*QueryGenesisTimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_665023cffa247a9d, []int{4}
}
func (m *QueryGenesisTimeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGenesisTimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGenesisTimeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGenesisTimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGenesisTimeRequest.Merge(m, src)
}
func (m *QueryGenesisTimeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGenesisTimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGenesisTimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGenesisTimeRequest proto.InternalMessageInfo

// QueryGenesisTimeResponse is the response type for the Query/GenesisTime RPC
// method.
type QueryGenesisTimeResponse struct {
	// GenesisTime is the timestamp associated with the first block.
	GenesisTime *time.Time `protobuf:"bytes,1,opt,name=genesis_time,json=genesisTime,proto3,stdtime" json:"genesis_time,omitempty"`
}

func (m *QueryGenesisTimeResponse) Reset()         { *m = QueryGenesisTimeResponse{} }
func (m *QueryGenesisTimeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGenesisTimeResponse) ProtoMessage()    {}
func (*QueryGenesisTimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_665023cffa247a9d, []int{5}
}
func (m *QueryGenesisTimeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGenesisTimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGenesisTimeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGenesisTimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGenesisTimeResponse.Merge(m, src)
}
func (m *QueryGenesisTimeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGenesisTimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGenesisTimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGenesisTimeResponse proto.InternalMessageInfo

func (m *QueryGenesisTimeResponse) GetGenesisTime() *time.Time {
	if m != nil {
		return m.GenesisTime
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryInflationRateRequest)(nil), "nubit.mint.v1.QueryInflationRateRequest")
	proto.RegisterType((*QueryInflationRateResponse)(nil), "nubit.mint.v1.QueryInflationRateResponse")
	proto.RegisterType((*QueryAnnualProvisionsRequest)(nil), "nubit.mint.v1.QueryAnnualProvisionsRequest")
	proto.RegisterType((*QueryAnnualProvisionsResponse)(nil), "nubit.mint.v1.QueryAnnualProvisionsResponse")
	proto.RegisterType((*QueryGenesisTimeRequest)(nil), "nubit.mint.v1.QueryGenesisTimeRequest")
	proto.RegisterType((*QueryGenesisTimeResponse)(nil), "nubit.mint.v1.QueryGenesisTimeResponse")
}

func init() { proto.RegisterFile("nubit/mint/v1/query.proto", fileDescriptor_665023cffa247a9d) }

var fileDescriptor_665023cffa247a9d = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x6b, 0x7e, 0x1d, 0xdc, 0x15, 0x0d, 0x0b, 0x89, 0x36, 0x8c, 0x14, 0x32, 0x51, 0x3a,
	0x8d, 0xd9, 0xea, 0x38, 0x72, 0xa2, 0x20, 0x21, 0x24, 0x0e, 0x10, 0x6d, 0x17, 0x38, 0x54, 0x4e,
	0xe7, 0x65, 0xd6, 0x12, 0x3b, 0x8b, 0x9d, 0x4a, 0x93, 0x38, 0x71, 0xe3, 0x36, 0xc4, 0x99, 0xbf,
	0x87, 0x1d, 0x27, 0x71, 0x41, 0x1c, 0x06, 0x6a, 0xf9, 0x43, 0x50, 0x1c, 0x67, 0xb4, 0x6b, 0x2a,
	0xa6, 0x9d, 0xf2, 0xe3, 0x7d, 0xf3, 0x7d, 0x9f, 0xf8, 0x3d, 0xd8, 0x12, 0x59, 0xc0, 0x35, 0x89,
	0xb9, 0xd0, 0x64, 0xd4, 0x23, 0x07, 0x19, 0x4b, 0x0f, 0x71, 0x92, 0x4a, 0x2d, 0x51, 0xc3, 0x48,
	0x38, 0x97, 0xf0, 0xa8, 0xe7, 0xdc, 0x0e, 0x65, 0x28, 0x8d, 0x42, 0xf2, 0xbb, 0x62, 0xc8, 0x59,
	0x09, 0xa5, 0x0c, 0x23, 0x46, 0x68, 0xc2, 0x09, 0x15, 0x42, 0x6a, 0xaa, 0xb9, 0x14, 0xca, 0xaa,
	0xcd, 0xd9, 0xed, 0x66, 0x55, 0xa1, 0xb4, 0xed, 0x77, 0xe6, 0x29, 0xc8, 0x76, 0x89, 0xe6, 0x31,
	0x53, 0x9a, 0xc6, 0x49, 0x31, 0xe0, 0xdd, 0x85, 0xad, 0xb7, 0x39, 0xcc, 0x2b, 0xb1, 0x1b, 0x99,
	0x9d, 0x3e, 0xd5, 0xcc, 0x67, 0x07, 0x19, 0x53, 0xda, 0x53, 0xd0, 0xa9, 0x12, 0x55, 0x22, 0x85,
	0x62, 0x68, 0x1b, 0xde, 0xe4, 0xa5, 0x30, 0x48, 0xa9, 0x66, 0x4d, 0x70, 0x1f, 0x74, 0x97, 0xfa,
	0xf8, 0xf8, 0xb4, 0x5d, 0xfb, 0x79, 0xda, 0xee, 0x84, 0x5c, 0xef, 0x65, 0x01, 0x1e, 0xca, 0x98,
	0x0c, 0xa5, 0x8a, 0xa5, 0xb2, 0x97, 0x0d, 0xb5, 0xb3, 0x4f, 0xf4, 0x61, 0xc2, 0x14, 0x7e, 0xc1,
	0x86, 0x7e, 0x83, 0x4f, 0xaf, 0xf7, 0x5c, 0xb8, 0x62, 0x4c, 0x9f, 0x09, 0x91, 0xd1, 0xe8, 0x4d,
	0x2a, 0x47, 0x5c, 0xe5, 0xff, 0x5a, 0x42, 0x7d, 0x80, 0xf7, 0x16, 0xe8, 0x96, 0xeb, 0x3d, 0xbc,
	0x45, 0x8d, 0x36, 0x48, 0xce, 0xc4, 0x4b, 0xa2, 0x2d, 0xd3, 0x73, 0x26, 0x5e, 0x0b, 0xde, 0x31,
	0xee, 0x2f, 0x99, 0x60, 0x8a, 0xab, 0x2d, 0x1e, 0x9f, 0x9d, 0xd6, 0x00, 0x36, 0xe7, 0x25, 0xcb,
	0xf4, 0x1c, 0x2e, 0x85, 0xc5, 0xeb, 0x41, 0x9e, 0x80, 0xc1, 0xa9, 0x6f, 0x3a, 0xb8, 0x88, 0x07,
	0x97, 0xf1, 0xe0, 0xad, 0x32, 0x9e, 0xfe, 0xb5, 0xa3, 0x5f, 0x6d, 0xe0, 0xd7, 0xc3, 0x7f, 0xcb,
	0x36, 0xbf, 0x5d, 0x85, 0xd7, 0x8d, 0x03, 0xfa, 0x0c, 0x60, 0x63, 0x26, 0x14, 0xd4, 0xc5, 0x33,
	0x35, 0xc2, 0x0b, 0x43, 0x75, 0xd6, 0x2e, 0x30, 0x59, 0x50, 0x7b, 0xeb, 0x1f, 0xbf, 0xff, 0xf9,
	0x72, 0xe5, 0x21, 0x5a, 0x2d, 0x0f, 0xc9, 0x36, 0x2c, 0x60, 0x9a, 0xf6, 0xc8, 0x6c, 0xf8, 0xe8,
	0x2b, 0x80, 0xcb, 0xe7, 0x33, 0x41, 0xeb, 0x55, 0x66, 0x0b, 0x92, 0x75, 0x1e, 0x5f, 0x6c, 0xd8,
	0xc2, 0x61, 0x03, 0xd7, 0x45, 0x9d, 0x4a, 0xb8, 0xb9, 0x06, 0xa0, 0x4f, 0x00, 0xd6, 0xa7, 0xa2,
	0x41, 0x9d, 0x2a, 0xb7, 0xf9, 0x58, 0x9d, 0x47, 0xff, 0x9d, 0xb3, 0x40, 0x6b, 0x06, 0x68, 0x15,
	0x3d, 0xa8, 0x04, 0x9a, 0x8e, 0xbf, 0xbf, 0x7d, 0x3c, 0x76, 0xc1, 0xc9, 0xd8, 0x05, 0xbf, 0xc7,
	0x2e, 0x38, 0x9a, 0xb8, 0xb5, 0x93, 0x89, 0x5b, 0xfb, 0x31, 0x71, 0x6b, 0xef, 0x9e, 0x4e, 0x35,
	0xd3, 0xe7, 0x2c, 0xa6, 0xaf, 0x69, 0xa0, 0x88, 0x21, 0xd8, 0x18, 0xd1, 0x88, 0xef, 0x50, 0x2d,
	0x53, 0x92, 0x69, 0x1e, 0x29, 0xa2, 0xf7, 0x89, 0xda, 0x2b, 0x5c, 0x4c, 0x65, 0x83, 0x1b, 0xa6,
	0x47, 0x4f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x60, 0x35, 0xb9, 0x6e, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// InflationRate returns the current inflation rate.
	InflationRate(ctx context.Context, in *QueryInflationRateRequest, opts ...grpc.CallOption) (*QueryInflationRateResponse, error)
	// AnnualProvisions returns the current annual provisions.
	AnnualProvisions(ctx context.Context, in *QueryAnnualProvisionsRequest, opts ...grpc.CallOption) (*QueryAnnualProvisionsResponse, error)
	// GenesisTime returns the genesis time.
	GenesisTime(ctx context.Context, in *QueryGenesisTimeRequest, opts ...grpc.CallOption) (*QueryGenesisTimeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) InflationRate(ctx context.Context, in *QueryInflationRateRequest, opts ...grpc.CallOption) (*QueryInflationRateResponse, error) {
	out := new(QueryInflationRateResponse)
	err := c.cc.Invoke(ctx, "/nubit.mint.v1.Query/InflationRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AnnualProvisions(ctx context.Context, in *QueryAnnualProvisionsRequest, opts ...grpc.CallOption) (*QueryAnnualProvisionsResponse, error) {
	out := new(QueryAnnualProvisionsResponse)
	err := c.cc.Invoke(ctx, "/nubit.mint.v1.Query/AnnualProvisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GenesisTime(ctx context.Context, in *QueryGenesisTimeRequest, opts ...grpc.CallOption) (*QueryGenesisTimeResponse, error) {
	out := new(QueryGenesisTimeResponse)
	err := c.cc.Invoke(ctx, "/nubit.mint.v1.Query/GenesisTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// InflationRate returns the current inflation rate.
	InflationRate(context.Context, *QueryInflationRateRequest) (*QueryInflationRateResponse, error)
	// AnnualProvisions returns the current annual provisions.
	AnnualProvisions(context.Context, *QueryAnnualProvisionsRequest) (*QueryAnnualProvisionsResponse, error)
	// GenesisTime returns the genesis time.
	GenesisTime(context.Context, *QueryGenesisTimeRequest) (*QueryGenesisTimeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) InflationRate(ctx context.Context, req *QueryInflationRateRequest) (*QueryInflationRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InflationRate not implemented")
}
func (*UnimplementedQueryServer) AnnualProvisions(ctx context.Context, req *QueryAnnualProvisionsRequest) (*QueryAnnualProvisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnualProvisions not implemented")
}
func (*UnimplementedQueryServer) GenesisTime(ctx context.Context, req *QueryGenesisTimeRequest) (*QueryGenesisTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenesisTime not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_InflationRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryInflationRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).InflationRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nubit.mint.v1.Query/InflationRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).InflationRate(ctx, req.(*QueryInflationRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AnnualProvisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAnnualProvisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AnnualProvisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nubit.mint.v1.Query/AnnualProvisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AnnualProvisions(ctx, req.(*QueryAnnualProvisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GenesisTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGenesisTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GenesisTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nubit.mint.v1.Query/GenesisTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GenesisTime(ctx, req.(*QueryGenesisTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nubit.mint.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InflationRate",
			Handler:    _Query_InflationRate_Handler,
		},
		{
			MethodName: "AnnualProvisions",
			Handler:    _Query_AnnualProvisions_Handler,
		},
		{
			MethodName: "GenesisTime",
			Handler:    _Query_GenesisTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nubit/mint/v1/query.proto",
}

func (m *QueryInflationRateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInflationRateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInflationRateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryInflationRateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryInflationRateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryInflationRateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InflationRate.Size()
		i -= size
		if _, err := m.InflationRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAnnualProvisionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAnnualProvisionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAnnualProvisionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAnnualProvisionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAnnualProvisionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAnnualProvisionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGenesisTimeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGenesisTimeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGenesisTimeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGenesisTimeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGenesisTimeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGenesisTimeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GenesisTime != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.GenesisTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.GenesisTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintQuery(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryInflationRateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryInflationRateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InflationRate.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAnnualProvisionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAnnualProvisionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGenesisTimeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGenesisTimeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GenesisTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.GenesisTime)
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryInflationRateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInflationRateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInflationRateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryInflationRateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryInflationRateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryInflationRateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAnnualProvisionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAnnualProvisionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAnnualProvisionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAnnualProvisionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAnnualProvisionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAnnualProvisionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGenesisTimeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGenesisTimeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGenesisTimeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGenesisTimeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGenesisTimeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGenesisTimeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GenesisTime == nil {
				m.GenesisTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.GenesisTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)