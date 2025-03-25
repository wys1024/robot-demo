// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: aichat/v1/aichat.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Aichat_SendMessage_FullMethodName = "/aichat.v1.Aichat/SendMessage"
)

// AichatClient is the client API for Aichat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Chat服务定义
type AichatClient interface {
	// SendMessage 发送聊天消息
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type aichatClient struct {
	cc grpc.ClientConnInterface
}

func NewAichatClient(cc grpc.ClientConnInterface) AichatClient {
	return &aichatClient{cc}
}

func (c *aichatClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, Aichat_SendMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AichatServer is the server API for Aichat service.
// All implementations must embed UnimplementedAichatServer
// for forward compatibility.
//
// Chat服务定义
type AichatServer interface {
	// SendMessage 发送聊天消息
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	mustEmbedUnimplementedAichatServer()
}

// UnimplementedAichatServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAichatServer struct{}

func (UnimplementedAichatServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedAichatServer) mustEmbedUnimplementedAichatServer() {}
func (UnimplementedAichatServer) testEmbeddedByValue()                {}

// UnsafeAichatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AichatServer will
// result in compilation errors.
type UnsafeAichatServer interface {
	mustEmbedUnimplementedAichatServer()
}

func RegisterAichatServer(s grpc.ServiceRegistrar, srv AichatServer) {
	// If the following call pancis, it indicates UnimplementedAichatServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Aichat_ServiceDesc, srv)
}

func _Aichat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AichatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aichat_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AichatServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Aichat_ServiceDesc is the grpc.ServiceDesc for Aichat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aichat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aichat.v1.Aichat",
	HandlerType: (*AichatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Aichat_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "aichat/v1/aichat.proto",
}
