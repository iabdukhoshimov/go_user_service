// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: payme.proto

package payment_service

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

// PaymeServiceClient is the client API for PaymeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymeServiceClient interface {
	GeneratePaymeLink(ctx context.Context, in *PaymeLinkRequest, opts ...grpc.CallOption) (*PaymeLinkResponse, error)
	GeneratePaymeLinkSvod(ctx context.Context, in *PaymeLinkRequestSvod, opts ...grpc.CallOption) (*PaymeLinkResponse, error)
}

type paymeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymeServiceClient(cc grpc.ClientConnInterface) PaymeServiceClient {
	return &paymeServiceClient{cc}
}

func (c *paymeServiceClient) GeneratePaymeLink(ctx context.Context, in *PaymeLinkRequest, opts ...grpc.CallOption) (*PaymeLinkResponse, error) {
	out := new(PaymeLinkResponse)
	err := c.cc.Invoke(ctx, "/payments.PaymeService/GeneratePaymeLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymeServiceClient) GeneratePaymeLinkSvod(ctx context.Context, in *PaymeLinkRequestSvod, opts ...grpc.CallOption) (*PaymeLinkResponse, error) {
	out := new(PaymeLinkResponse)
	err := c.cc.Invoke(ctx, "/payments.PaymeService/GeneratePaymeLinkSvod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymeServiceServer is the server API for PaymeService service.
// All implementations must embed UnimplementedPaymeServiceServer
// for forward compatibility
type PaymeServiceServer interface {
	GeneratePaymeLink(context.Context, *PaymeLinkRequest) (*PaymeLinkResponse, error)
	GeneratePaymeLinkSvod(context.Context, *PaymeLinkRequestSvod) (*PaymeLinkResponse, error)
	mustEmbedUnimplementedPaymeServiceServer()
}

// UnimplementedPaymeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymeServiceServer struct {
}

func (UnimplementedPaymeServiceServer) GeneratePaymeLink(context.Context, *PaymeLinkRequest) (*PaymeLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePaymeLink not implemented")
}
func (UnimplementedPaymeServiceServer) GeneratePaymeLinkSvod(context.Context, *PaymeLinkRequestSvod) (*PaymeLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePaymeLinkSvod not implemented")
}
func (UnimplementedPaymeServiceServer) mustEmbedUnimplementedPaymeServiceServer() {}

// UnsafePaymeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymeServiceServer will
// result in compilation errors.
type UnsafePaymeServiceServer interface {
	mustEmbedUnimplementedPaymeServiceServer()
}

func RegisterPaymeServiceServer(s grpc.ServiceRegistrar, srv PaymeServiceServer) {
	s.RegisterService(&PaymeService_ServiceDesc, srv)
}

func _PaymeService_GeneratePaymeLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymeLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymeServiceServer).GeneratePaymeLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payments.PaymeService/GeneratePaymeLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymeServiceServer).GeneratePaymeLink(ctx, req.(*PaymeLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymeService_GeneratePaymeLinkSvod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymeLinkRequestSvod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymeServiceServer).GeneratePaymeLinkSvod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payments.PaymeService/GeneratePaymeLinkSvod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymeServiceServer).GeneratePaymeLinkSvod(ctx, req.(*PaymeLinkRequestSvod))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymeService_ServiceDesc is the grpc.ServiceDesc for PaymeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payments.PaymeService",
	HandlerType: (*PaymeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GeneratePaymeLink",
			Handler:    _PaymeService_GeneratePaymeLink_Handler,
		},
		{
			MethodName: "GeneratePaymeLinkSvod",
			Handler:    _PaymeService_GeneratePaymeLinkSvod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payme.proto",
}
