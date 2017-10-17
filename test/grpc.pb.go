// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc.proto

/*
Package com_grpc is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	MyRequest
	MyResponse
	MyMsg
	MyMsg2
*/
package com_grpc

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyRequest struct {
	Value  int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	Value2 int32 `protobuf:"varint,2,opt,name=Value2,proto3" json:"Value2,omitempty"`
}

func (m *MyRequest) Reset()                    { *m = MyRequest{} }
func (m *MyRequest) String() string            { return proto.CompactTextString(m) }
func (*MyRequest) ProtoMessage()               {}
func (*MyRequest) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{0} }

func (m *MyRequest) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MyRequest) GetValue2() int32 {
	if m != nil {
		return m.Value2
	}
	return 0
}

type MyResponse struct {
	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *MyResponse) Reset()                    { *m = MyResponse{} }
func (m *MyResponse) String() string            { return proto.CompactTextString(m) }
func (*MyResponse) ProtoMessage()               {}
func (*MyResponse) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{1} }

func (m *MyResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MyMsg struct {
	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *MyMsg) Reset()                    { *m = MyMsg{} }
func (m *MyMsg) String() string            { return proto.CompactTextString(m) }
func (*MyMsg) ProtoMessage()               {}
func (*MyMsg) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{2} }

func (m *MyMsg) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MyMsg2 struct {
	Value int64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *MyMsg2) Reset()                    { *m = MyMsg2{} }
func (m *MyMsg2) String() string            { return proto.CompactTextString(m) }
func (*MyMsg2) ProtoMessage()               {}
func (*MyMsg2) Descriptor() ([]byte, []int) { return fileDescriptorGrpc, []int{3} }

func (m *MyMsg2) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*MyRequest)(nil), "com.grpc.MyRequest")
	proto.RegisterType((*MyResponse)(nil), "com.grpc.MyResponse")
	proto.RegisterType((*MyMsg)(nil), "com.grpc.MyMsg")
	proto.RegisterType((*MyMsg2)(nil), "com.grpc.MyMsg2")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MyTest service

type MyTestClient interface {
	UnaryCall(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
	// This RPC streams from the server only.
	Downstream(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (MyTest_DownstreamClient, error)
	// This RPC streams from the client.
	Upstreamy(ctx context.Context, opts ...grpc.CallOption) (MyTest_UpstreamyClient, error)
	// This one streams in both directions.
	Bidi(ctx context.Context, opts ...grpc.CallOption) (MyTest_BidiClient, error)
}

type myTestClient struct {
	cc *grpc.ClientConn
}

func NewMyTestClient(cc *grpc.ClientConn) MyTestClient {
	return &myTestClient{cc}
}

func (c *myTestClient) UnaryCall(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := grpc.Invoke(ctx, "/com.grpc.MyTest/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myTestClient) Downstream(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (MyTest_DownstreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MyTest_serviceDesc.Streams[0], c.cc, "/com.grpc.MyTest/Downstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &myTestDownstreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MyTest_DownstreamClient interface {
	Recv() (*MyMsg, error)
	grpc.ClientStream
}

type myTestDownstreamClient struct {
	grpc.ClientStream
}

func (x *myTestDownstreamClient) Recv() (*MyMsg, error) {
	m := new(MyMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *myTestClient) Upstreamy(ctx context.Context, opts ...grpc.CallOption) (MyTest_UpstreamyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MyTest_serviceDesc.Streams[1], c.cc, "/com.grpc.MyTest/Upstreamy", opts...)
	if err != nil {
		return nil, err
	}
	x := &myTestUpstreamyClient{stream}
	return x, nil
}

type MyTest_UpstreamyClient interface {
	Send(*MyMsg) error
	CloseAndRecv() (*MyResponse, error)
	grpc.ClientStream
}

type myTestUpstreamyClient struct {
	grpc.ClientStream
}

func (x *myTestUpstreamyClient) Send(m *MyMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myTestUpstreamyClient) CloseAndRecv() (*MyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *myTestClient) Bidi(ctx context.Context, opts ...grpc.CallOption) (MyTest_BidiClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MyTest_serviceDesc.Streams[2], c.cc, "/com.grpc.MyTest/Bidi", opts...)
	if err != nil {
		return nil, err
	}
	x := &myTestBidiClient{stream}
	return x, nil
}

type MyTest_BidiClient interface {
	Send(*MyMsg) error
	Recv() (*MyMsg2, error)
	grpc.ClientStream
}

type myTestBidiClient struct {
	grpc.ClientStream
}

func (x *myTestBidiClient) Send(m *MyMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myTestBidiClient) Recv() (*MyMsg2, error) {
	m := new(MyMsg2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MyTest service

type MyTestServer interface {
	UnaryCall(context.Context, *MyRequest) (*MyResponse, error)
	// This RPC streams from the server only.
	Downstream(*MyRequest, MyTest_DownstreamServer) error
	// This RPC streams from the client.
	Upstreamy(MyTest_UpstreamyServer) error
	// This one streams in both directions.
	Bidi(MyTest_BidiServer) error
}

func RegisterMyTestServer(s *grpc.Server, srv MyTestServer) {
	s.RegisterService(&_MyTest_serviceDesc, srv)
}

func _MyTest_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyTestServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.grpc.MyTest/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyTestServer).UnaryCall(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyTest_Downstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyTestServer).Downstream(m, &myTestDownstreamServer{stream})
}

type MyTest_DownstreamServer interface {
	Send(*MyMsg) error
	grpc.ServerStream
}

type myTestDownstreamServer struct {
	grpc.ServerStream
}

func (x *myTestDownstreamServer) Send(m *MyMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _MyTest_Upstreamy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyTestServer).Upstreamy(&myTestUpstreamyServer{stream})
}

type MyTest_UpstreamyServer interface {
	SendAndClose(*MyResponse) error
	Recv() (*MyMsg, error)
	grpc.ServerStream
}

type myTestUpstreamyServer struct {
	grpc.ServerStream
}

func (x *myTestUpstreamyServer) SendAndClose(m *MyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myTestUpstreamyServer) Recv() (*MyMsg, error) {
	m := new(MyMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MyTest_Bidi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyTestServer).Bidi(&myTestBidiServer{stream})
}

type MyTest_BidiServer interface {
	Send(*MyMsg2) error
	Recv() (*MyMsg, error)
	grpc.ServerStream
}

type myTestBidiServer struct {
	grpc.ServerStream
}

func (x *myTestBidiServer) Send(m *MyMsg2) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myTestBidiServer) Recv() (*MyMsg, error) {
	m := new(MyMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MyTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.grpc.MyTest",
	HandlerType: (*MyTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _MyTest_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Downstream",
			Handler:       _MyTest_Downstream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upstreamy",
			Handler:       _MyTest_Upstreamy_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Bidi",
			Handler:       _MyTest_Bidi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptorGrpc) }

var fileDescriptorGrpc = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xce, 0xcf, 0xd5, 0x03, 0xf1, 0x95, 0x2c,
	0xb9, 0x38, 0x7d, 0x2b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0xc3,
	0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x1c, 0x21, 0x31, 0x2e,
	0x36, 0x30, 0xc3, 0x48, 0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x53, 0x52, 0xe2, 0xe2,
	0x02, 0x69, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xc5, 0xae, 0x57, 0x49, 0x96, 0x8b, 0xd5, 0xb7,
	0xd2, 0xb7, 0x38, 0x1d, 0x87, 0xb4, 0x1c, 0x17, 0x1b, 0x58, 0xda, 0x08, 0xbb, 0xbc, 0xd1, 0x6d,
	0x46, 0x90, 0x82, 0x10, 0x90, 0xdb, 0xcc, 0xb8, 0x38, 0x43, 0xf3, 0x12, 0x8b, 0x2a, 0x9d, 0x13,
	0x73, 0x72, 0x84, 0x84, 0xf5, 0x60, 0x1e, 0xd0, 0x83, 0xbb, 0x5e, 0x4a, 0x04, 0x55, 0x10, 0xea,
	0x2e, 0x13, 0x2e, 0x2e, 0x97, 0xfc, 0xf2, 0xbc, 0xe2, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0xec, 0x1a,
	0xf9, 0x91, 0x05, 0x7d, 0x8b, 0xd3, 0x0d, 0x18, 0x85, 0x4c, 0xb8, 0x38, 0x43, 0x0b, 0x20, 0x7a,
	0x2a, 0x85, 0xd0, 0xe5, 0xb1, 0xdb, 0xa4, 0xc1, 0x28, 0xa4, 0xcb, 0xc5, 0xe2, 0x94, 0x99, 0x92,
	0x89, 0xa9, 0x41, 0x00, 0x4d, 0xc0, 0x48, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x19, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x57, 0xb7, 0xf4, 0x9a, 0x01, 0x00, 0x00,
}
