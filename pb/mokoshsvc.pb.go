// Code generated by protoc-gen-go.
// source: mokoshsvc.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	mokoshsvc.proto

It has these top-level messages:
	CapabilitiesRequest
	CapabilitiesReply
	MergePart
	MergeResult
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Capability int32

const (
	Capability_MERGE Capability = 0
)

var Capability_name = map[int32]string{
	0: "MERGE",
}
var Capability_value = map[string]int32{
	"MERGE": 0,
}

func (x Capability) String() string {
	return proto.EnumName(Capability_name, int32(x))
}
func (Capability) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The request message
type CapabilitiesRequest struct {
}

func (m *CapabilitiesRequest) Reset()                    { *m = CapabilitiesRequest{} }
func (m *CapabilitiesRequest) String() string            { return proto.CompactTextString(m) }
func (*CapabilitiesRequest) ProtoMessage()               {}
func (*CapabilitiesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing capabilites of this server
type CapabilitiesReply struct {
	Capabilities []Capability `protobuf:"varint,1,rep,packed,name=capabilities,enum=pb.Capability" json:"capabilities,omitempty"`
}

func (m *CapabilitiesReply) Reset()                    { *m = CapabilitiesReply{} }
func (m *CapabilitiesReply) String() string            { return proto.CompactTextString(m) }
func (*CapabilitiesReply) ProtoMessage()               {}
func (*CapabilitiesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CapabilitiesReply) GetCapabilities() []Capability {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

type MergePart struct {
	Data          []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	AddEmptyIfOdd bool   `protobuf:"varint,2,opt,name=addEmptyIfOdd" json:"addEmptyIfOdd,omitempty"`
}

func (m *MergePart) Reset()                    { *m = MergePart{} }
func (m *MergePart) String() string            { return proto.CompactTextString(m) }
func (*MergePart) ProtoMessage()               {}
func (*MergePart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MergePart) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MergePart) GetAddEmptyIfOdd() bool {
	if m != nil {
		return m.AddEmptyIfOdd
	}
	return false
}

type MergeResult struct {
	Data  []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Pages int32  `protobuf:"varint,2,opt,name=pages" json:"pages,omitempty"`
}

func (m *MergeResult) Reset()                    { *m = MergeResult{} }
func (m *MergeResult) String() string            { return proto.CompactTextString(m) }
func (*MergeResult) ProtoMessage()               {}
func (*MergeResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MergeResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MergeResult) GetPages() int32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func init() {
	proto.RegisterType((*CapabilitiesRequest)(nil), "pb.CapabilitiesRequest")
	proto.RegisterType((*CapabilitiesReply)(nil), "pb.CapabilitiesReply")
	proto.RegisterType((*MergePart)(nil), "pb.MergePart")
	proto.RegisterType((*MergeResult)(nil), "pb.MergeResult")
	proto.RegisterEnum("pb.Capability", Capability_name, Capability_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Merger service

type MergerClient interface {
	Merge(ctx context.Context, opts ...grpc.CallOption) (Merger_MergeClient, error)
}

type mergerClient struct {
	cc *grpc.ClientConn
}

func NewMergerClient(cc *grpc.ClientConn) MergerClient {
	return &mergerClient{cc}
}

func (c *mergerClient) Merge(ctx context.Context, opts ...grpc.CallOption) (Merger_MergeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Merger_serviceDesc.Streams[0], c.cc, "/pb.Merger/Merge", opts...)
	if err != nil {
		return nil, err
	}
	x := &mergerMergeClient{stream}
	return x, nil
}

type Merger_MergeClient interface {
	Send(*MergePart) error
	CloseAndRecv() (*MergeResult, error)
	grpc.ClientStream
}

type mergerMergeClient struct {
	grpc.ClientStream
}

func (x *mergerMergeClient) Send(m *MergePart) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mergerMergeClient) CloseAndRecv() (*MergeResult, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MergeResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Merger service

type MergerServer interface {
	Merge(Merger_MergeServer) error
}

func RegisterMergerServer(s *grpc.Server, srv MergerServer) {
	s.RegisterService(&_Merger_serviceDesc, srv)
}

func _Merger_Merge_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MergerServer).Merge(&mergerMergeServer{stream})
}

type Merger_MergeServer interface {
	SendAndClose(*MergeResult) error
	Recv() (*MergePart, error)
	grpc.ServerStream
}

type mergerMergeServer struct {
	grpc.ServerStream
}

func (x *mergerMergeServer) SendAndClose(m *MergeResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mergerMergeServer) Recv() (*MergePart, error) {
	m := new(MergePart)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Merger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Merger",
	HandlerType: (*MergerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Merge",
			Handler:       _Merger_Merge_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "mokoshsvc.proto",
}

// Client API for Mokosh service

type MokoshClient interface {
	// Request all capabilities supported by this instance
	Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesReply, error)
}

type mokoshClient struct {
	cc *grpc.ClientConn
}

func NewMokoshClient(cc *grpc.ClientConn) MokoshClient {
	return &mokoshClient{cc}
}

func (c *mokoshClient) Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesReply, error) {
	out := new(CapabilitiesReply)
	err := grpc.Invoke(ctx, "/pb.Mokosh/Capabilities", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mokosh service

type MokoshServer interface {
	// Request all capabilities supported by this instance
	Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesReply, error)
}

func RegisterMokoshServer(s *grpc.Server, srv MokoshServer) {
	s.RegisterService(&_Mokosh_serviceDesc, srv)
}

func _Mokosh_Capabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MokoshServer).Capabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mokosh/Capabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MokoshServer).Capabilities(ctx, req.(*CapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mokosh_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Mokosh",
	HandlerType: (*MokoshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Capabilities",
			Handler:    _Mokosh_Capabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mokoshsvc.proto",
}

func init() { proto.RegisterFile("mokoshsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0x9b, 0x69, 0x8b, 0x3b, 0xbb, 0x4d, 0x4f, 0xc7, 0xca, 0x9e, 0x4a, 0xf0, 0xa1, 0x28,
	0xf4, 0xa1, 0x22, 0xbe, 0x0a, 0x52, 0x86, 0x42, 0x51, 0xf2, 0x0d, 0xd2, 0x25, 0xce, 0x62, 0x47,
	0x63, 0x93, 0x09, 0xfd, 0xf6, 0x62, 0x06, 0x5d, 0x83, 0x7b, 0xbb, 0xfb, 0xe7, 0xf8, 0xe5, 0xf2,
	0x0b, 0xcc, 0xb6, 0xcd, 0x57, 0xa3, 0x3f, 0xf5, 0xcf, 0x3a, 0x55, 0x6d, 0x63, 0x1a, 0x1c, 0xa9,
	0x92, 0xce, 0xe1, 0xea, 0x99, 0x2b, 0x5e, 0x56, 0x75, 0x65, 0x2a, 0xa9, 0x99, 0xfc, 0xde, 0x49,
	0x6d, 0xe8, 0x0a, 0x2e, 0xdd, 0x58, 0xd5, 0x1d, 0x66, 0x10, 0xae, 0x07, 0x61, 0x44, 0xe2, 0x93,
	0x64, 0x9a, 0x4d, 0x53, 0x55, 0xa6, 0xfd, 0x70, 0xc7, 0x9c, 0x19, 0x9a, 0xc3, 0xb8, 0x90, 0xed,
	0x46, 0xbe, 0xf3, 0xd6, 0x20, 0xc2, 0xa9, 0xe0, 0x86, 0x47, 0x24, 0x26, 0x49, 0xc8, 0x6c, 0x8d,
	0x37, 0x30, 0xe1, 0x42, 0xe4, 0x5b, 0x65, 0xba, 0x97, 0x8f, 0x37, 0x21, 0xa2, 0x51, 0x4c, 0x92,
	0x33, 0xe6, 0x86, 0xf4, 0x11, 0xce, 0x2d, 0x86, 0x49, 0xbd, 0xab, 0x8f, 0x83, 0xae, 0xc1, 0x57,
	0x7c, 0x23, 0xb5, 0x05, 0xf8, 0x6c, 0xdf, 0xdc, 0x2e, 0x00, 0x0e, 0xbb, 0xe1, 0x18, 0xfc, 0x22,
	0x67, 0xab, 0xfc, 0xc2, 0xcb, 0x1e, 0x20, 0xb0, 0xc4, 0x16, 0xef, 0xc0, 0xb7, 0x15, 0x4e, 0xfe,
	0x5e, 0xd2, 0x6f, 0xbb, 0x9c, 0xf5, 0xed, 0xfe, 0x56, 0xea, 0x25, 0x24, 0x7b, 0x85, 0xa0, 0xb0,
	0x1a, 0xf1, 0x09, 0xc2, 0xa1, 0x22, 0x5c, 0x38, 0x1e, 0x0e, 0x2e, 0x97, 0xf3, 0xff, 0x07, 0xaa,
	0xee, 0xa8, 0x57, 0x06, 0xf6, 0x1b, 0xee, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xfa, 0xc5,
	0x07, 0x99, 0x01, 0x00, 0x00,
}
