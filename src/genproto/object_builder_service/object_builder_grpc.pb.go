// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: object_builder.proto

package object_builder_service

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

// ObjectBuilderServiceClient is the client API for ObjectBuilderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectBuilderServiceClient interface {
	Create(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	GetSingle(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	GetList(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	GetRecursiveList(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	Update(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	Delete(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	ManyToManyAppend(ctx context.Context, in *ManyToManyMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ManyToManyDelete(ctx context.Context, in *ManyToManyMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetObjectDetails(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	GetListInExcel(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error)
	Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*CommonMessage, error)
}

type objectBuilderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectBuilderServiceClient(cc grpc.ClientConnInterface) ObjectBuilderServiceClient {
	return &objectBuilderServiceClient{cc}
}

func (c *objectBuilderServiceClient) Create(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) GetSingle(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/GetSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) GetList(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) GetRecursiveList(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/GetRecursiveList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) Update(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) Delete(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) ManyToManyAppend(ctx context.Context, in *ManyToManyMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/ManyToManyAppend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) ManyToManyDelete(ctx context.Context, in *ManyToManyMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/ManyToManyDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) GetObjectDetails(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/GetObjectDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) GetListInExcel(ctx context.Context, in *CommonMessage, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/GetListInExcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectBuilderServiceClient) Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*CommonMessage, error) {
	out := new(CommonMessage)
	err := c.cc.Invoke(ctx, "/object_builder_service.ObjectBuilderService/Batch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectBuilderServiceServer is the server API for ObjectBuilderService service.
// All implementations must embed UnimplementedObjectBuilderServiceServer
// for forward compatibility
type ObjectBuilderServiceServer interface {
	Create(context.Context, *CommonMessage) (*CommonMessage, error)
	GetSingle(context.Context, *CommonMessage) (*CommonMessage, error)
	GetList(context.Context, *CommonMessage) (*CommonMessage, error)
	GetRecursiveList(context.Context, *CommonMessage) (*CommonMessage, error)
	Update(context.Context, *CommonMessage) (*CommonMessage, error)
	Delete(context.Context, *CommonMessage) (*CommonMessage, error)
	ManyToManyAppend(context.Context, *ManyToManyMessage) (*emptypb.Empty, error)
	ManyToManyDelete(context.Context, *ManyToManyMessage) (*emptypb.Empty, error)
	GetObjectDetails(context.Context, *CommonMessage) (*CommonMessage, error)
	GetListInExcel(context.Context, *CommonMessage) (*CommonMessage, error)
	Batch(context.Context, *BatchRequest) (*CommonMessage, error)
	mustEmbedUnimplementedObjectBuilderServiceServer()
}

// UnimplementedObjectBuilderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedObjectBuilderServiceServer struct {
}

func (UnimplementedObjectBuilderServiceServer) Create(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedObjectBuilderServiceServer) GetSingle(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingle not implemented")
}
func (UnimplementedObjectBuilderServiceServer) GetList(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedObjectBuilderServiceServer) GetRecursiveList(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecursiveList not implemented")
}
func (UnimplementedObjectBuilderServiceServer) Update(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedObjectBuilderServiceServer) Delete(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedObjectBuilderServiceServer) ManyToManyAppend(context.Context, *ManyToManyMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManyToManyAppend not implemented")
}
func (UnimplementedObjectBuilderServiceServer) ManyToManyDelete(context.Context, *ManyToManyMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManyToManyDelete not implemented")
}
func (UnimplementedObjectBuilderServiceServer) GetObjectDetails(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectDetails not implemented")
}
func (UnimplementedObjectBuilderServiceServer) GetListInExcel(context.Context, *CommonMessage) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListInExcel not implemented")
}
func (UnimplementedObjectBuilderServiceServer) Batch(context.Context, *BatchRequest) (*CommonMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Batch not implemented")
}
func (UnimplementedObjectBuilderServiceServer) mustEmbedUnimplementedObjectBuilderServiceServer() {}

// UnsafeObjectBuilderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectBuilderServiceServer will
// result in compilation errors.
type UnsafeObjectBuilderServiceServer interface {
	mustEmbedUnimplementedObjectBuilderServiceServer()
}

func RegisterObjectBuilderServiceServer(s grpc.ServiceRegistrar, srv ObjectBuilderServiceServer) {
	s.RegisterService(&ObjectBuilderService_ServiceDesc, srv)
}

func _ObjectBuilderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).Create(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_GetSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).GetSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/GetSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).GetSingle(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).GetList(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_GetRecursiveList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).GetRecursiveList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/GetRecursiveList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).GetRecursiveList(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).Update(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).Delete(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_ManyToManyAppend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManyToManyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).ManyToManyAppend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/ManyToManyAppend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).ManyToManyAppend(ctx, req.(*ManyToManyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_ManyToManyDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManyToManyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).ManyToManyDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/ManyToManyDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).ManyToManyDelete(ctx, req.(*ManyToManyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_GetObjectDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).GetObjectDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/GetObjectDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).GetObjectDetails(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_GetListInExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).GetListInExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/GetListInExcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).GetListInExcel(ctx, req.(*CommonMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectBuilderService_Batch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectBuilderServiceServer).Batch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.ObjectBuilderService/Batch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectBuilderServiceServer).Batch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectBuilderService_ServiceDesc is the grpc.ServiceDesc for ObjectBuilderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectBuilderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.ObjectBuilderService",
	HandlerType: (*ObjectBuilderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ObjectBuilderService_Create_Handler,
		},
		{
			MethodName: "GetSingle",
			Handler:    _ObjectBuilderService_GetSingle_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ObjectBuilderService_GetList_Handler,
		},
		{
			MethodName: "GetRecursiveList",
			Handler:    _ObjectBuilderService_GetRecursiveList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ObjectBuilderService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ObjectBuilderService_Delete_Handler,
		},
		{
			MethodName: "ManyToManyAppend",
			Handler:    _ObjectBuilderService_ManyToManyAppend_Handler,
		},
		{
			MethodName: "ManyToManyDelete",
			Handler:    _ObjectBuilderService_ManyToManyDelete_Handler,
		},
		{
			MethodName: "GetObjectDetails",
			Handler:    _ObjectBuilderService_GetObjectDetails_Handler,
		},
		{
			MethodName: "GetListInExcel",
			Handler:    _ObjectBuilderService_GetListInExcel_Handler,
		},
		{
			MethodName: "Batch",
			Handler:    _ObjectBuilderService_Batch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "object_builder.proto",
}
