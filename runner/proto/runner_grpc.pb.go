// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: runner/proto/runner.proto

package runner

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

// ServerProtoClient is the client API for ServerProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerProtoClient interface {
	SendServer(ctx context.Context, opts ...grpc.CallOption) (ServerProto_SendServerClient, error)
}

type serverProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewServerProtoClient(cc grpc.ClientConnInterface) ServerProtoClient {
	return &serverProtoClient{cc}
}

func (c *serverProtoClient) SendServer(ctx context.Context, opts ...grpc.CallOption) (ServerProto_SendServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerProto_ServiceDesc.Streams[0], "/runner.ServerProto/SendServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverProtoSendServerClient{stream}
	return x, nil
}

type ServerProto_SendServerClient interface {
	Send(*ServerRequest) error
	Recv() (*ServerReply, error)
	grpc.ClientStream
}

type serverProtoSendServerClient struct {
	grpc.ClientStream
}

func (x *serverProtoSendServerClient) Send(m *ServerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *serverProtoSendServerClient) Recv() (*ServerReply, error) {
	m := new(ServerReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerProtoServer is the server API for ServerProto service.
// All implementations must embed UnimplementedServerProtoServer
// for forward compatibility
type ServerProtoServer interface {
	SendServer(ServerProto_SendServerServer) error
	mustEmbedUnimplementedServerProtoServer()
}

// UnimplementedServerProtoServer must be embedded to have forward compatible implementations.
type UnimplementedServerProtoServer struct {
}

func (UnimplementedServerProtoServer) SendServer(ServerProto_SendServerServer) error {
	return status.Errorf(codes.Unimplemented, "method SendServer not implemented")
}
func (UnimplementedServerProtoServer) mustEmbedUnimplementedServerProtoServer() {}

// UnsafeServerProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerProtoServer will
// result in compilation errors.
type UnsafeServerProtoServer interface {
	mustEmbedUnimplementedServerProtoServer()
}

func RegisterServerProtoServer(s grpc.ServiceRegistrar, srv ServerProtoServer) {
	s.RegisterService(&ServerProto_ServiceDesc, srv)
}

func _ServerProto_SendServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ServerProtoServer).SendServer(&serverProtoSendServerServer{stream})
}

type ServerProto_SendServerServer interface {
	Send(*ServerReply) error
	Recv() (*ServerRequest, error)
	grpc.ServerStream
}

type serverProtoSendServerServer struct {
	grpc.ServerStream
}

func (x *serverProtoSendServerServer) Send(m *ServerReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *serverProtoSendServerServer) Recv() (*ServerRequest, error) {
	m := new(ServerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerProto_ServiceDesc is the grpc.ServiceDesc for ServerProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "runner.ServerProto",
	HandlerType: (*ServerProtoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendServer",
			Handler:       _ServerProto_SendServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "runner/proto/runner.proto",
}
