// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: processor/api/processor_v1/service.proto

package processor_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProcessorServiceClient is the client API for ProcessorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessorServiceClient interface {
	GetMockTemplate(ctx context.Context, in *GetMockTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type processorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessorServiceClient(cc grpc.ClientConnInterface) ProcessorServiceClient {
	return &processorServiceClient{cc}
}

func (c *processorServiceClient) GetMockTemplate(ctx context.Context, in *GetMockTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mock.v1.ProcessorService/GetMockTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessorServiceServer is the server API for ProcessorService service.
// All implementations must embed UnimplementedProcessorServiceServer
// for forward compatibility
type ProcessorServiceServer interface {
	GetMockTemplate(context.Context, *GetMockTemplateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedProcessorServiceServer()
}

// UnimplementedProcessorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessorServiceServer struct {
}

func (UnimplementedProcessorServiceServer) GetMockTemplate(context.Context, *GetMockTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMockTemplate not implemented")
}
func (UnimplementedProcessorServiceServer) mustEmbedUnimplementedProcessorServiceServer() {}

// UnsafeProcessorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessorServiceServer will
// result in compilation errors.
type UnsafeProcessorServiceServer interface {
	mustEmbedUnimplementedProcessorServiceServer()
}

func RegisterProcessorServiceServer(s grpc.ServiceRegistrar, srv ProcessorServiceServer) {
	s.RegisterService(&ProcessorService_ServiceDesc, srv)
}

func _ProcessorService_GetMockTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMockTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessorServiceServer).GetMockTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mock.v1.ProcessorService/GetMockTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessorServiceServer).GetMockTemplate(ctx, req.(*GetMockTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessorService_ServiceDesc is the grpc.ServiceDesc for ProcessorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mock.v1.ProcessorService",
	HandlerType: (*ProcessorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMockTemplate",
			Handler:    _ProcessorService_GetMockTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "processor/api/processor_v1/service.proto",
}
