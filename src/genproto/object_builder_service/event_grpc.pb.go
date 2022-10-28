// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: event.proto

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

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	Create(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error)
	GetList(ctx context.Context, in *GetEventsListRequest, opts ...grpc.CallOption) (*GetEventsListResponse, error)
	GetSingle(ctx context.Context, in *EventPrimaryKey, opts ...grpc.CallOption) (*Event, error)
	Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *EventPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) Create(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/object_builder_service.EventService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetList(ctx context.Context, in *GetEventsListRequest, opts ...grpc.CallOption) (*GetEventsListResponse, error) {
	out := new(GetEventsListResponse)
	err := c.cc.Invoke(ctx, "/object_builder_service.EventService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetSingle(ctx context.Context, in *EventPrimaryKey, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/object_builder_service.EventService/GetSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.EventService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) Delete(ctx context.Context, in *EventPrimaryKey, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/object_builder_service.EventService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	Create(context.Context, *CreateEventRequest) (*Event, error)
	GetList(context.Context, *GetEventsListRequest) (*GetEventsListResponse, error)
	GetSingle(context.Context, *EventPrimaryKey) (*Event, error)
	Update(context.Context, *Event) (*emptypb.Empty, error)
	Delete(context.Context, *EventPrimaryKey) (*emptypb.Empty, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) Create(context.Context, *CreateEventRequest) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEventServiceServer) GetList(context.Context, *GetEventsListRequest) (*GetEventsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedEventServiceServer) GetSingle(context.Context, *EventPrimaryKey) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingle not implemented")
}
func (UnimplementedEventServiceServer) Update(context.Context, *Event) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEventServiceServer) Delete(context.Context, *EventPrimaryKey) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.EventService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Create(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.EventService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetList(ctx, req.(*GetEventsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.EventService/GetSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetSingle(ctx, req.(*EventPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.EventService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Update(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_builder_service.EventService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).Delete(ctx, req.(*EventPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_builder_service.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EventService_Create_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _EventService_GetList_Handler,
		},
		{
			MethodName: "GetSingle",
			Handler:    _EventService_GetSingle_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EventService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EventService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "event.proto",
}