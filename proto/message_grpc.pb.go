// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/message.proto

package billing

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FlowsMesssageClient is the client API for FlowsMesssage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlowsMesssageClient interface {
	CreateMessage(ctx context.Context, in *CreateFlowsMessageRequest, opts ...grpc.CallOption) (*CreateFlowsMessageResponse, error)
}

type flowsMesssageClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowsMesssageClient(cc grpc.ClientConnInterface) FlowsMesssageClient {
	return &flowsMesssageClient{cc}
}

func (c *flowsMesssageClient) CreateMessage(ctx context.Context, in *CreateFlowsMessageRequest, opts ...grpc.CallOption) (*CreateFlowsMessageResponse, error) {
	out := new(CreateFlowsMessageResponse)
	err := c.cc.Invoke(ctx, "/message.FlowsMesssage/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowsMesssageServer is the server API for FlowsMesssage service.
// All implementations must embed UnimplementedFlowsMesssageServer
// for forward compatibility
type FlowsMesssageServer interface {
	CreateMessage(context.Context, *CreateFlowsMessageRequest) (*CreateFlowsMessageResponse, error)
	mustEmbedUnimplementedFlowsMesssageServer()
}

// UnimplementedFlowsMesssageServer must be embedded to have forward compatible implementations.
type UnimplementedFlowsMesssageServer struct {
}

func (UnimplementedFlowsMesssageServer) CreateMessage(context.Context, *CreateFlowsMessageRequest) (*CreateFlowsMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedFlowsMesssageServer) mustEmbedUnimplementedFlowsMesssageServer() {}

// UnsafeFlowsMesssageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlowsMesssageServer will
// result in compilation errors.
type UnsafeFlowsMesssageServer interface {
	mustEmbedUnimplementedFlowsMesssageServer()
}

func RegisterFlowsMesssageServer(s grpc.ServiceRegistrar, srv FlowsMesssageServer) {
	s.RegisterService(&FlowsMesssage_ServiceDesc, srv)
}

func _FlowsMesssage_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlowsMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowsMesssageServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.FlowsMesssage/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowsMesssageServer).CreateMessage(ctx, req.(*CreateFlowsMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlowsMesssage_ServiceDesc is the grpc.ServiceDesc for FlowsMesssage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlowsMesssage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.FlowsMesssage",
	HandlerType: (*FlowsMesssageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _FlowsMesssage_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message.proto",
}
