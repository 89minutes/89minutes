// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: story_service.proto

package pb

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

// StoryServiceClient is the client API for StoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoryServiceClient interface {
	PingPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
	UploadStoryAndFiles(ctx context.Context, opts ...grpc.CallOption) (StoryService_UploadStoryAndFilesClient, error)
}

type storyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoryServiceClient(cc grpc.ClientConnInterface) StoryServiceClient {
	return &storyServiceClient{cc}
}

func (c *storyServiceClient) PingPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/StoryService/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storyServiceClient) UploadStoryAndFiles(ctx context.Context, opts ...grpc.CallOption) (StoryService_UploadStoryAndFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &StoryService_ServiceDesc.Streams[0], "/StoryService/UploadStoryAndFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &storyServiceUploadStoryAndFilesClient{stream}
	return x, nil
}

type StoryService_UploadStoryAndFilesClient interface {
	Send(*UploadStoryAndFilesReq) error
	CloseAndRecv() (*UploadStoryAndFilesRes, error)
	grpc.ClientStream
}

type storyServiceUploadStoryAndFilesClient struct {
	grpc.ClientStream
}

func (x *storyServiceUploadStoryAndFilesClient) Send(m *UploadStoryAndFilesReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storyServiceUploadStoryAndFilesClient) CloseAndRecv() (*UploadStoryAndFilesRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStoryAndFilesRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StoryServiceServer is the server API for StoryService service.
// All implementations must embed UnimplementedStoryServiceServer
// for forward compatibility
type StoryServiceServer interface {
	PingPong(context.Context, *Ping) (*Pong, error)
	UploadStoryAndFiles(StoryService_UploadStoryAndFilesServer) error
	mustEmbedUnimplementedStoryServiceServer()
}

// UnimplementedStoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoryServiceServer struct {
}

func (UnimplementedStoryServiceServer) PingPong(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (UnimplementedStoryServiceServer) UploadStoryAndFiles(StoryService_UploadStoryAndFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadStoryAndFiles not implemented")
}
func (UnimplementedStoryServiceServer) mustEmbedUnimplementedStoryServiceServer() {}

// UnsafeStoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoryServiceServer will
// result in compilation errors.
type UnsafeStoryServiceServer interface {
	mustEmbedUnimplementedStoryServiceServer()
}

func RegisterStoryServiceServer(s grpc.ServiceRegistrar, srv StoryServiceServer) {
	s.RegisterService(&StoryService_ServiceDesc, srv)
}

func _StoryService_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoryServiceServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StoryService/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoryServiceServer).PingPong(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoryService_UploadStoryAndFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StoryServiceServer).UploadStoryAndFiles(&storyServiceUploadStoryAndFilesServer{stream})
}

type StoryService_UploadStoryAndFilesServer interface {
	SendAndClose(*UploadStoryAndFilesRes) error
	Recv() (*UploadStoryAndFilesReq, error)
	grpc.ServerStream
}

type storyServiceUploadStoryAndFilesServer struct {
	grpc.ServerStream
}

func (x *storyServiceUploadStoryAndFilesServer) SendAndClose(m *UploadStoryAndFilesRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storyServiceUploadStoryAndFilesServer) Recv() (*UploadStoryAndFilesReq, error) {
	m := new(UploadStoryAndFilesReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StoryService_ServiceDesc is the grpc.ServiceDesc for StoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StoryService",
	HandlerType: (*StoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingPong",
			Handler:    _StoryService_PingPong_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadStoryAndFiles",
			Handler:       _StoryService_UploadStoryAndFiles_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "story_service.proto",
}
