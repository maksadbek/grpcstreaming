// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: hellogrpc.proto

package __

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

const (
	UploaderService_UploadFile_FullMethodName = "/proto.UploaderService/UploadFile"
)

// UploaderServiceClient is the client API for UploaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploaderServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (UploaderService_UploadFileClient, error)
}

type uploaderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploaderServiceClient(cc grpc.ClientConnInterface) UploaderServiceClient {
	return &uploaderServiceClient{cc}
}

func (c *uploaderServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (UploaderService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploaderService_ServiceDesc.Streams[0], UploaderService_UploadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &uploaderServiceUploadFileClient{stream}
	return x, nil
}

type UploaderService_UploadFileClient interface {
	Send(*UploadRequestType) error
	CloseAndRecv() (*UploadResponseType, error)
	grpc.ClientStream
}

type uploaderServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *uploaderServiceUploadFileClient) Send(m *UploadRequestType) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploaderServiceUploadFileClient) CloseAndRecv() (*UploadResponseType, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponseType)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploaderServiceServer is the server API for UploaderService service.
// All implementations must embed UnimplementedUploaderServiceServer
// for forward compatibility
type UploaderServiceServer interface {
	UploadFile(UploaderService_UploadFileServer) error
	mustEmbedUnimplementedUploaderServiceServer()
}

// UnimplementedUploaderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUploaderServiceServer struct {
}

func (UnimplementedUploaderServiceServer) UploadFile(UploaderService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedUploaderServiceServer) mustEmbedUnimplementedUploaderServiceServer() {}

// UnsafeUploaderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploaderServiceServer will
// result in compilation errors.
type UnsafeUploaderServiceServer interface {
	mustEmbedUnimplementedUploaderServiceServer()
}

func RegisterUploaderServiceServer(s grpc.ServiceRegistrar, srv UploaderServiceServer) {
	s.RegisterService(&UploaderService_ServiceDesc, srv)
}

func _UploaderService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploaderServiceServer).UploadFile(&uploaderServiceUploadFileServer{stream})
}

type UploaderService_UploadFileServer interface {
	SendAndClose(*UploadResponseType) error
	Recv() (*UploadRequestType, error)
	grpc.ServerStream
}

type uploaderServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *uploaderServiceUploadFileServer) SendAndClose(m *UploadResponseType) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploaderServiceUploadFileServer) Recv() (*UploadRequestType, error) {
	m := new(UploadRequestType)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploaderService_ServiceDesc is the grpc.ServiceDesc for UploaderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploaderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UploaderService",
	HandlerType: (*UploaderServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _UploaderService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "hellogrpc.proto",
}