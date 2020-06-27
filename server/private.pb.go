// Code generated by protoc-gen-go. DO NOT EDIT.
// source: private.proto

package server

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Nodes struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2a91b51c7bdc125, []int{0}
}

func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (m *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(m, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type FindRequest struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2a91b51c7bdc125, []int{1}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

func (m *FindRequest) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*Nodes)(nil), "server.Nodes")
	proto.RegisterType((*FindRequest)(nil), "server.FindRequest")
}

func init() {
	proto.RegisterFile("private.proto", fileDescriptor_d2a91b51c7bdc125)
}

var fileDescriptor_d2a91b51c7bdc125 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x41, 0x4b, 0x02, 0x51,
	0x10, 0xc7, 0xd5, 0xd0, 0xc3, 0xa8, 0x05, 0xcf, 0x08, 0xd9, 0x08, 0x6c, 0xbb, 0x08, 0xc1, 0xdb,
	0x50, 0x0c, 0xaa, 0x43, 0x60, 0x14, 0x04, 0x21, 0xa2, 0x9d, 0xba, 0xed, 0xee, 0x1b, 0xd7, 0x07,
	0xeb, 0x8e, 0xbd, 0x79, 0x2b, 0xec, 0x67, 0xee, 0x4b, 0xc4, 0xba, 0x19, 0x5a, 0x08, 0x5e, 0xe7,
	0x3f, 0xbf, 0xff, 0xfc, 0x06, 0x9a, 0x4b, 0xa3, 0x57, 0xbe, 0x45, 0xb9, 0x34, 0x64, 0x49, 0xd4,
	0x18, 0xcd, 0x0a, 0x8d, 0x73, 0x1e, 0x11, 0x45, 0x31, 0x7a, 0xeb, 0x69, 0x90, 0xce, 0x3c, 0x5c,
	0x2c, 0x6d, 0x56, 0x2c, 0x39, 0x90, 0x90, 0xfa, 0x01, 0xdc, 0x6b, 0xa8, 0x8e, 0x48, 0x21, 0x0b,
	0x17, 0xaa, 0xf9, 0x98, 0xdb, 0xe5, 0xce, 0x51, 0xb7, 0xde, 0x6b, 0xc8, 0xa2, 0x49, 0xe6, 0xe9,
	0xa4, 0x88, 0xdc, 0x0b, 0xa8, 0xbf, 0xe8, 0x44, 0x4d, 0xf0, 0x33, 0x45, 0xb6, 0xe2, 0x18, 0x2a,
	0x5a, 0xb5, 0xcb, 0x9d, 0x72, 0xb7, 0x31, 0xa9, 0x68, 0xd5, 0xfb, 0xaa, 0xc0, 0xc9, 0x6b, 0x62,
	0xd1, 0x24, 0x7e, 0x3c, 0x45, 0xb3, 0xd2, 0x21, 0x8a, 0x01, 0xc0, 0x34, 0x0d, 0x43, 0x64, 0x26,
	0xc3, 0xe2, 0x4c, 0x16, 0x5e, 0x72, 0xe3, 0x25, 0x9f, 0x73, 0x2f, 0xa7, 0xb9, 0x7d, 0x8d, 0xdd,
	0x92, 0x18, 0x40, 0x7d, 0x6c, 0x50, 0x61, 0x01, 0xee, 0xe5, 0x76, 0x2c, 0xdd, 0x92, 0x78, 0x80,
	0xd3, 0x5c, 0xf0, 0xf7, 0xe2, 0x30, 0x7b, 0xf7, 0x83, 0x18, 0x45, 0x6b, 0xb3, 0xb7, 0xa5, 0xff,
	0x0f, 0xbe, 0x87, 0xd6, 0x1f, 0xf8, 0x4d, 0xb3, 0x3d, 0x8c, 0x7d, 0x84, 0x76, 0x1e, 0x3f, 0xc5,
	0xc4, 0xc8, 0x76, 0x6c, 0x30, 0x44, 0xa5, 0x93, 0x28, 0x4f, 0x0f, 0x2b, 0xb8, 0x81, 0xda, 0x88,
	0xac, 0x9e, 0x65, 0x62, 0x27, 0x71, 0xf6, 0x7c, 0xee, 0x96, 0x86, 0x57, 0x1f, 0x97, 0x91, 0xb6,
	0xf3, 0x34, 0x90, 0x21, 0x2d, 0x3c, 0xeb, 0x6b, 0x9e, 0xd3, 0x6d, 0xbf, 0x7f, 0xe7, 0x45, 0x64,
	0x94, 0x57, 0x74, 0x04, 0xb5, 0x35, 0xd6, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x73, 0x94, 0x3c,
	0x54, 0x27, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InternalServiceClient is the client API for InternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalServiceClient interface {
	Successors(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Nodes, error)
	Predecessor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error)
	FindSuccessorByTable(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	FindSuccessorByList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	FindClosestPrecedingNode(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error)
	Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error)
}

type internalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalServiceClient(cc grpc.ClientConnInterface) InternalServiceClient {
	return &internalServiceClient{cc}
}

func (c *internalServiceClient) Successors(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Nodes, error) {
	out := new(Nodes)
	err := c.cc.Invoke(ctx, "/server.InternalService/Successors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Predecessor(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/Predecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindSuccessorByTable(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindSuccessorByTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindSuccessorByList(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindSuccessorByList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) FindClosestPrecedingNode(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/server.InternalService/FindClosestPrecedingNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalServiceClient) Notify(ctx context.Context, in *Node, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/server.InternalService/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServiceServer is the server API for InternalService service.
type InternalServiceServer interface {
	Successors(context.Context, *empty.Empty) (*Nodes, error)
	Predecessor(context.Context, *empty.Empty) (*Node, error)
	FindSuccessorByTable(context.Context, *FindRequest) (*Node, error)
	FindSuccessorByList(context.Context, *FindRequest) (*Node, error)
	FindClosestPrecedingNode(context.Context, *FindRequest) (*Node, error)
	Notify(context.Context, *Node) (*empty.Empty, error)
}

// UnimplementedInternalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInternalServiceServer struct {
}

func (*UnimplementedInternalServiceServer) Successors(ctx context.Context, req *empty.Empty) (*Nodes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Successors not implemented")
}
func (*UnimplementedInternalServiceServer) Predecessor(ctx context.Context, req *empty.Empty) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predecessor not implemented")
}
func (*UnimplementedInternalServiceServer) FindSuccessorByTable(ctx context.Context, req *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorByTable not implemented")
}
func (*UnimplementedInternalServiceServer) FindSuccessorByList(ctx context.Context, req *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessorByList not implemented")
}
func (*UnimplementedInternalServiceServer) FindClosestPrecedingNode(ctx context.Context, req *FindRequest) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindClosestPrecedingNode not implemented")
}
func (*UnimplementedInternalServiceServer) Notify(ctx context.Context, req *Node) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}

func RegisterInternalServiceServer(s *grpc.Server, srv InternalServiceServer) {
	s.RegisterService(&_InternalService_serviceDesc, srv)
}

func _InternalService_Successors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Successors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Successors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Successors(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Predecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Predecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Predecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Predecessor(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindSuccessorByTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindSuccessorByTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindSuccessorByTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindSuccessorByTable(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindSuccessorByList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindSuccessorByList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindSuccessorByList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindSuccessorByList(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_FindClosestPrecedingNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).FindClosestPrecedingNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/FindClosestPrecedingNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).FindClosestPrecedingNode(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalService_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServiceServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.InternalService/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServiceServer).Notify(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.InternalService",
	HandlerType: (*InternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Successors",
			Handler:    _InternalService_Successors_Handler,
		},
		{
			MethodName: "Predecessor",
			Handler:    _InternalService_Predecessor_Handler,
		},
		{
			MethodName: "FindSuccessorByTable",
			Handler:    _InternalService_FindSuccessorByTable_Handler,
		},
		{
			MethodName: "FindSuccessorByList",
			Handler:    _InternalService_FindSuccessorByList_Handler,
		},
		{
			MethodName: "FindClosestPrecedingNode",
			Handler:    _InternalService_FindClosestPrecedingNode_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _InternalService_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "private.proto",
}
