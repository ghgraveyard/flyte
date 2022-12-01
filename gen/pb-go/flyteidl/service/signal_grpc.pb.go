// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	admin "github.com/flyteorg/flyte/gen/pb-go/flyteidl/admin"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SignalServiceClient is the client API for SignalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignalServiceClient interface {
	// Fetches or creates a :ref:`ref_flyteidl.admin.Signal`.
	GetOrCreateSignal(ctx context.Context, in *admin.SignalGetOrCreateRequest, opts ...grpc.CallOption) (*admin.Signal, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Signal` definitions.
	ListSignals(ctx context.Context, in *admin.SignalListRequest, opts ...grpc.CallOption) (*admin.SignalList, error)
	// Sets the value on a :ref:`ref_flyteidl.admin.Signal` definition
	SetSignal(ctx context.Context, in *admin.SignalSetRequest, opts ...grpc.CallOption) (*admin.SignalSetResponse, error)
}

type signalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignalServiceClient(cc grpc.ClientConnInterface) SignalServiceClient {
	return &signalServiceClient{cc}
}

func (c *signalServiceClient) GetOrCreateSignal(ctx context.Context, in *admin.SignalGetOrCreateRequest, opts ...grpc.CallOption) (*admin.Signal, error) {
	out := new(admin.Signal)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/GetOrCreateSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) ListSignals(ctx context.Context, in *admin.SignalListRequest, opts ...grpc.CallOption) (*admin.SignalList, error) {
	out := new(admin.SignalList)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/ListSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) SetSignal(ctx context.Context, in *admin.SignalSetRequest, opts ...grpc.CallOption) (*admin.SignalSetResponse, error) {
	out := new(admin.SignalSetResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/SetSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalServiceServer is the server API for SignalService service.
// All implementations should embed UnimplementedSignalServiceServer
// for forward compatibility
type SignalServiceServer interface {
	// Fetches or creates a :ref:`ref_flyteidl.admin.Signal`.
	GetOrCreateSignal(context.Context, *admin.SignalGetOrCreateRequest) (*admin.Signal, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Signal` definitions.
	ListSignals(context.Context, *admin.SignalListRequest) (*admin.SignalList, error)
	// Sets the value on a :ref:`ref_flyteidl.admin.Signal` definition
	SetSignal(context.Context, *admin.SignalSetRequest) (*admin.SignalSetResponse, error)
}

// UnimplementedSignalServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSignalServiceServer struct {
}

func (UnimplementedSignalServiceServer) GetOrCreateSignal(context.Context, *admin.SignalGetOrCreateRequest) (*admin.Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateSignal not implemented")
}
func (UnimplementedSignalServiceServer) ListSignals(context.Context, *admin.SignalListRequest) (*admin.SignalList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSignals not implemented")
}
func (UnimplementedSignalServiceServer) SetSignal(context.Context, *admin.SignalSetRequest) (*admin.SignalSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSignal not implemented")
}

// UnsafeSignalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignalServiceServer will
// result in compilation errors.
type UnsafeSignalServiceServer interface {
	mustEmbedUnimplementedSignalServiceServer()
}

func RegisterSignalServiceServer(s grpc.ServiceRegistrar, srv SignalServiceServer) {
	s.RegisterService(&SignalService_ServiceDesc, srv)
}

func _SignalService_GetOrCreateSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalGetOrCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetOrCreateSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/GetOrCreateSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetOrCreateSignal(ctx, req.(*admin.SignalGetOrCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_ListSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).ListSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/ListSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).ListSignals(ctx, req.(*admin.SignalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_SetSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).SetSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/SetSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).SetSignal(ctx, req.(*admin.SignalSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SignalService_ServiceDesc is the grpc.ServiceDesc for SignalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.SignalService",
	HandlerType: (*SignalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrCreateSignal",
			Handler:    _SignalService_GetOrCreateSignal_Handler,
		},
		{
			MethodName: "ListSignals",
			Handler:    _SignalService_ListSignals_Handler,
		},
		{
			MethodName: "SetSignal",
			Handler:    _SignalService_SetSignal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/signal.proto",
}
