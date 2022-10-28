// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: empowermentMongo.proto

package mongo_document_service

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

// EmpowermentServiceClient is the client API for EmpowermentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmpowermentServiceClient interface {
	Upsert(ctx context.Context, in *Empowerment, opts ...grpc.CallOption) (*TimeResp, error)
	Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Empowerment, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*TimeResp, error)
}

type empowermentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmpowermentServiceClient(cc grpc.ClientConnInterface) EmpowermentServiceClient {
	return &empowermentServiceClient{cc}
}

func (c *empowermentServiceClient) Upsert(ctx context.Context, in *Empowerment, opts ...grpc.CallOption) (*TimeResp, error) {
	out := new(TimeResp)
	err := c.cc.Invoke(ctx, "/mongo_document_service.EmpowermentService/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empowermentServiceClient) Get(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Empowerment, error) {
	out := new(Empowerment)
	err := c.cc.Invoke(ctx, "/mongo_document_service.EmpowermentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empowermentServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mongo_document_service.EmpowermentService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empowermentServiceClient) UpdateStatus(ctx context.Context, in *UpdateStatusReq, opts ...grpc.CallOption) (*TimeResp, error) {
	out := new(TimeResp)
	err := c.cc.Invoke(ctx, "/mongo_document_service.EmpowermentService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmpowermentServiceServer is the server API for EmpowermentService service.
// All implementations must embed UnimplementedEmpowermentServiceServer
// for forward compatibility
type EmpowermentServiceServer interface {
	Upsert(context.Context, *Empowerment) (*TimeResp, error)
	Get(context.Context, *ById) (*Empowerment, error)
	Delete(context.Context, *ById) (*emptypb.Empty, error)
	UpdateStatus(context.Context, *UpdateStatusReq) (*TimeResp, error)
	mustEmbedUnimplementedEmpowermentServiceServer()
}

// UnimplementedEmpowermentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmpowermentServiceServer struct {
}

func (UnimplementedEmpowermentServiceServer) Upsert(context.Context, *Empowerment) (*TimeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedEmpowermentServiceServer) Get(context.Context, *ById) (*Empowerment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEmpowermentServiceServer) Delete(context.Context, *ById) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEmpowermentServiceServer) UpdateStatus(context.Context, *UpdateStatusReq) (*TimeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedEmpowermentServiceServer) mustEmbedUnimplementedEmpowermentServiceServer() {}

// UnsafeEmpowermentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmpowermentServiceServer will
// result in compilation errors.
type UnsafeEmpowermentServiceServer interface {
	mustEmbedUnimplementedEmpowermentServiceServer()
}

func RegisterEmpowermentServiceServer(s grpc.ServiceRegistrar, srv EmpowermentServiceServer) {
	s.RegisterService(&EmpowermentService_ServiceDesc, srv)
}

func _EmpowermentService_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empowerment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpowermentServiceServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.EmpowermentService/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpowermentServiceServer).Upsert(ctx, req.(*Empowerment))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpowermentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpowermentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.EmpowermentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpowermentServiceServer).Get(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpowermentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpowermentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.EmpowermentService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpowermentServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpowermentService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpowermentServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongo_document_service.EmpowermentService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpowermentServiceServer).UpdateStatus(ctx, req.(*UpdateStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EmpowermentService_ServiceDesc is the grpc.ServiceDesc for EmpowermentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmpowermentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongo_document_service.EmpowermentService",
	HandlerType: (*EmpowermentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upsert",
			Handler:    _EmpowermentService_Upsert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EmpowermentService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EmpowermentService_Delete_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _EmpowermentService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "empowermentMongo.proto",
}