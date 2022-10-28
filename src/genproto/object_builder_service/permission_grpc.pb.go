// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: permission.proto

package object_builder_service

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

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	UpsertPermissionsByAppId(ctx context.Context, in *UpsertPermissionsByAppIdRequest, opts ...grpc.CallOption) (*UpsertPermissionsByAppIdResponse, error)
	GetAllPermissionsByRoleId(ctx context.Context, in *GetAllPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) UpsertPermissionsByAppId(ctx context.Context, in *UpsertPermissionsByAppIdRequest, opts ...grpc.CallOption) (*UpsertPermissionsByAppIdResponse, error) {
	out := new(UpsertPermissionsByAppIdResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.PermissionService/UpsertPermissionsByAppId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) GetAllPermissionsByRoleId(ctx context.Context, in *GetAllPermissionRequest, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.PermissionService/GetAllPermissionsByRoleId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations must embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	UpsertPermissionsByAppId(context.Context, *UpsertPermissionsByAppIdRequest) (*UpsertPermissionsByAppIdResponse, error)
	GetAllPermissionsByRoleId(context.Context, *GetAllPermissionRequest) (*CommonMessage, error)
	mustEmbedUnimplementedPermissionServiceServer()
}

// UnimplementedPermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) UpsertPermissionsByAppId(context.Context, *UpsertPermissionsByAppIdRequest) (*UpsertPermissionsByAppIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertPermissionsByAppId not implemented")
}
func (UnimplementedPermissionServiceServer) GetAllPermissionsByRoleId(context.Context, *GetAllPermissionRequest) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPermissionsByRoleId not implemented")
}
func (UnimplementedPermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_UpsertPermissionsByAppId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPermissionsByAppIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpsertPermissionsByAppId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PermissionService/UpsertPermissionsByAppId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpsertPermissionsByAppId(ctx, req.(*UpsertPermissionsByAppIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_GetAllPermissionsByRoleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetAllPermissionsByRoleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.PermissionService/GetAllPermissionsByRoleId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetAllPermissionsByRoleId(ctx, req.(*GetAllPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertPermissionsByAppId",
			Handler:    _PermissionService_UpsertPermissionsByAppId_Handler,
		},
		{
			MethodName: "GetAllPermissionsByRoleId",
			Handler:    _PermissionService_GetAllPermissionsByRoleId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission.proto",
}
