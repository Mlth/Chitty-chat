// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chat

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	SendMessage(ctx context.Context, in *WrittenMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
	JoinServer(ctx context.Context, in *WrittenMessage, opts ...grpc.CallOption) (Chat_JoinServerClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) SendMessage(ctx context.Context, in *WrittenMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/ChittyChat.chat/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) JoinServer(ctx context.Context, in *WrittenMessage, opts ...grpc.CallOption) (Chat_JoinServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/ChittyChat.chat/JoinServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatJoinServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_JoinServerClient interface {
	Recv() (*WrittenMessage, error)
	grpc.ClientStream
}

type chatJoinServerClient struct {
	grpc.ClientStream
}

func (x *chatJoinServerClient) Recv() (*WrittenMessage, error) {
	m := new(WrittenMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	SendMessage(context.Context, *WrittenMessage) (*EmptyMessage, error)
	JoinServer(*WrittenMessage, Chat_JoinServerServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) SendMessage(context.Context, *WrittenMessage) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) JoinServer(*WrittenMessage, Chat_JoinServerServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinServer not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WrittenMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChat.chat/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).SendMessage(ctx, req.(*WrittenMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_JoinServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WrittenMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).JoinServer(m, &chatJoinServerServer{stream})
}

type Chat_JoinServerServer interface {
	Send(*WrittenMessage) error
	grpc.ServerStream
}

type chatJoinServerServer struct {
	grpc.ServerStream
}

func (x *chatJoinServerServer) Send(m *WrittenMessage) error {
	return x.ServerStream.SendMsg(m)
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChittyChat.chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Chat_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinServer",
			Handler:       _Chat_JoinServer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/proto.proto",
}
