// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: internal/api/common.proto

package grpc_test

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
	Common_Health_FullMethodName  = "/Common/Health"
	Common_Version_FullMethodName = "/Common/Version"
	Common_Testing_FullMethodName = "/Common/Testing"
)

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommonClient interface {
	// Health check
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	// Fetch version
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Simple testing
	Testing(ctx context.Context, in *TestingRequest, opts ...grpc.CallOption) (*TestingResponse, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, Common_Health_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, Common_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) Testing(ctx context.Context, in *TestingRequest, opts ...grpc.CallOption) (*TestingResponse, error) {
	out := new(TestingResponse)
	err := c.cc.Invoke(ctx, Common_Testing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
// All implementations must embed UnimplementedCommonServer
// for forward compatibility
type CommonServer interface {
	// Health check
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	// Fetch version
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Simple testing
	Testing(context.Context, *TestingRequest) (*TestingResponse, error)
	mustEmbedUnimplementedCommonServer()
}

// UnimplementedCommonServer must be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (UnimplementedCommonServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedCommonServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedCommonServer) Testing(context.Context, *TestingRequest) (*TestingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Testing not implemented")
}
func (UnimplementedCommonServer) mustEmbedUnimplementedCommonServer() {}

// UnsafeCommonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommonServer will
// result in compilation errors.
type UnsafeCommonServer interface {
	mustEmbedUnimplementedCommonServer()
}

func RegisterCommonServer(s grpc.ServiceRegistrar, srv CommonServer) {
	s.RegisterService(&Common_ServiceDesc, srv)
}

func _Common_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_Testing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).Testing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Common_Testing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).Testing(ctx, req.(*TestingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Common_ServiceDesc is the grpc.ServiceDesc for Common service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Common_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _Common_Health_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Common_Version_Handler,
		},
		{
			MethodName: "Testing",
			Handler:    _Common_Testing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/common.proto",
}
