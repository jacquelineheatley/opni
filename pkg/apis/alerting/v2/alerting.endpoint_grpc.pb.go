// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/apis/alerting/v2/alerting.endpoint.proto

package v2

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AlertEndpointsV2Client is the client API for AlertEndpointsV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertEndpointsV2Client interface {
	CreateAlertEndpoint(ctx context.Context, in *AlertEndpoint, opts ...grpc.CallOption) (*v1.Reference, error)
	GetAlertEndpoint(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*AlertEndpoint, error)
	ListAlertEndpoints(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AlertEndpointList, error)
	UpdateAlertEndpoint(ctx context.Context, in *UpdateAlertEndpointRequest, opts ...grpc.CallOption) (*UpdateAlertEndpointResponse, error)
	DeleteAlertEndpoint(ctx context.Context, in *DeleteAlertEndpointRequest, opts ...grpc.CallOption) (*DeleteAlertEndpointResponse, error)
	TestAlertEndpoint(ctx context.Context, in *AlertEndpoint, opts ...grpc.CallOption) (*TestAlertEndpointResponse, error)
	EphemeralDispatcher(ctx context.Context, in *EphemeralDispatcherRequest, opts ...grpc.CallOption) (*EphemeralDispatcherResponse, error)
}

type alertEndpointsV2Client struct {
	cc grpc.ClientConnInterface
}

func NewAlertEndpointsV2Client(cc grpc.ClientConnInterface) AlertEndpointsV2Client {
	return &alertEndpointsV2Client{cc}
}

func (c *alertEndpointsV2Client) CreateAlertEndpoint(ctx context.Context, in *AlertEndpoint, opts ...grpc.CallOption) (*v1.Reference, error) {
	out := new(v1.Reference)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsV2/CreateAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsV2Client) GetAlertEndpoint(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*AlertEndpoint, error) {
	out := new(AlertEndpoint)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsV2/GetAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsV2Client) ListAlertEndpoints(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AlertEndpointList, error) {
	out := new(AlertEndpointList)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsV2/ListAlertEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsV2Client) UpdateAlertEndpoint(ctx context.Context, in *UpdateAlertEndpointRequest, opts ...grpc.CallOption) (*UpdateAlertEndpointResponse, error) {
	out := new(UpdateAlertEndpointResponse)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsV2/UpdateAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsV2Client) DeleteAlertEndpoint(ctx context.Context, in *DeleteAlertEndpointRequest, opts ...grpc.CallOption) (*DeleteAlertEndpointResponse, error) {
	out := new(DeleteAlertEndpointResponse)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsV2/DeleteAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsV2Client) TestAlertEndpoint(ctx context.Context, in *AlertEndpoint, opts ...grpc.CallOption) (*TestAlertEndpointResponse, error) {
	out := new(TestAlertEndpointResponse)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsV2/TestAlertEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsV2Client) EphemeralDispatcher(ctx context.Context, in *EphemeralDispatcherRequest, opts ...grpc.CallOption) (*EphemeralDispatcherResponse, error) {
	out := new(EphemeralDispatcherResponse)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsV2/EphemeralDispatcher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertEndpointsV2Server is the server API for AlertEndpointsV2 service.
// All implementations must embed UnimplementedAlertEndpointsV2Server
// for forward compatibility
type AlertEndpointsV2Server interface {
	CreateAlertEndpoint(context.Context, *AlertEndpoint) (*v1.Reference, error)
	GetAlertEndpoint(context.Context, *v1.Reference) (*AlertEndpoint, error)
	ListAlertEndpoints(context.Context, *emptypb.Empty) (*AlertEndpointList, error)
	UpdateAlertEndpoint(context.Context, *UpdateAlertEndpointRequest) (*UpdateAlertEndpointResponse, error)
	DeleteAlertEndpoint(context.Context, *DeleteAlertEndpointRequest) (*DeleteAlertEndpointResponse, error)
	TestAlertEndpoint(context.Context, *AlertEndpoint) (*TestAlertEndpointResponse, error)
	EphemeralDispatcher(context.Context, *EphemeralDispatcherRequest) (*EphemeralDispatcherResponse, error)
	mustEmbedUnimplementedAlertEndpointsV2Server()
}

// UnimplementedAlertEndpointsV2Server must be embedded to have forward compatible implementations.
type UnimplementedAlertEndpointsV2Server struct {
}

func (UnimplementedAlertEndpointsV2Server) CreateAlertEndpoint(context.Context, *AlertEndpoint) (*v1.Reference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlertEndpoint not implemented")
}
func (UnimplementedAlertEndpointsV2Server) GetAlertEndpoint(context.Context, *v1.Reference) (*AlertEndpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlertEndpoint not implemented")
}
func (UnimplementedAlertEndpointsV2Server) ListAlertEndpoints(context.Context, *emptypb.Empty) (*AlertEndpointList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlertEndpoints not implemented")
}
func (UnimplementedAlertEndpointsV2Server) UpdateAlertEndpoint(context.Context, *UpdateAlertEndpointRequest) (*UpdateAlertEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlertEndpoint not implemented")
}
func (UnimplementedAlertEndpointsV2Server) DeleteAlertEndpoint(context.Context, *DeleteAlertEndpointRequest) (*DeleteAlertEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlertEndpoint not implemented")
}
func (UnimplementedAlertEndpointsV2Server) TestAlertEndpoint(context.Context, *AlertEndpoint) (*TestAlertEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAlertEndpoint not implemented")
}
func (UnimplementedAlertEndpointsV2Server) EphemeralDispatcher(context.Context, *EphemeralDispatcherRequest) (*EphemeralDispatcherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EphemeralDispatcher not implemented")
}
func (UnimplementedAlertEndpointsV2Server) mustEmbedUnimplementedAlertEndpointsV2Server() {}

// UnsafeAlertEndpointsV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertEndpointsV2Server will
// result in compilation errors.
type UnsafeAlertEndpointsV2Server interface {
	mustEmbedUnimplementedAlertEndpointsV2Server()
}

func RegisterAlertEndpointsV2Server(s grpc.ServiceRegistrar, srv AlertEndpointsV2Server) {
	s.RegisterService(&AlertEndpointsV2_ServiceDesc, srv)
}

func _AlertEndpointsV2_CreateAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsV2Server).CreateAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsV2/CreateAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsV2Server).CreateAlertEndpoint(ctx, req.(*AlertEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsV2_GetAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsV2Server).GetAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsV2/GetAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsV2Server).GetAlertEndpoint(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsV2_ListAlertEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsV2Server).ListAlertEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsV2/ListAlertEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsV2Server).ListAlertEndpoints(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsV2_UpdateAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsV2Server).UpdateAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsV2/UpdateAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsV2Server).UpdateAlertEndpoint(ctx, req.(*UpdateAlertEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsV2_DeleteAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsV2Server).DeleteAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsV2/DeleteAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsV2Server).DeleteAlertEndpoint(ctx, req.(*DeleteAlertEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsV2_TestAlertEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertEndpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsV2Server).TestAlertEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsV2/TestAlertEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsV2Server).TestAlertEndpoint(ctx, req.(*AlertEndpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsV2_EphemeralDispatcher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EphemeralDispatcherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsV2Server).EphemeralDispatcher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsV2/EphemeralDispatcher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsV2Server).EphemeralDispatcher(ctx, req.(*EphemeralDispatcherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertEndpointsV2_ServiceDesc is the grpc.ServiceDesc for AlertEndpointsV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertEndpointsV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alertingv2.AlertEndpointsV2",
	HandlerType: (*AlertEndpointsV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlertEndpoint",
			Handler:    _AlertEndpointsV2_CreateAlertEndpoint_Handler,
		},
		{
			MethodName: "GetAlertEndpoint",
			Handler:    _AlertEndpointsV2_GetAlertEndpoint_Handler,
		},
		{
			MethodName: "ListAlertEndpoints",
			Handler:    _AlertEndpointsV2_ListAlertEndpoints_Handler,
		},
		{
			MethodName: "UpdateAlertEndpoint",
			Handler:    _AlertEndpointsV2_UpdateAlertEndpoint_Handler,
		},
		{
			MethodName: "DeleteAlertEndpoint",
			Handler:    _AlertEndpointsV2_DeleteAlertEndpoint_Handler,
		},
		{
			MethodName: "TestAlertEndpoint",
			Handler:    _AlertEndpointsV2_TestAlertEndpoint_Handler,
		},
		{
			MethodName: "EphemeralDispatcher",
			Handler:    _AlertEndpointsV2_EphemeralDispatcher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/alerting/v2/alerting.endpoint.proto",
}

// AlertEndpointsSyncV2Client is the client API for AlertEndpointsSyncV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertEndpointsSyncV2Client interface {
	// ImportConfig
	ImportConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Downstream Syncs config
	SyncConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFullConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type alertEndpointsSyncV2Client struct {
	cc grpc.ClientConnInterface
}

func NewAlertEndpointsSyncV2Client(cc grpc.ClientConnInterface) AlertEndpointsSyncV2Client {
	return &alertEndpointsSyncV2Client{cc}
}

func (c *alertEndpointsSyncV2Client) ImportConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsSyncV2/ImportConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsSyncV2Client) SyncConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsSyncV2/SyncConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertEndpointsSyncV2Client) GetFullConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/alertingv2.AlertEndpointsSyncV2/GetFullConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertEndpointsSyncV2Server is the server API for AlertEndpointsSyncV2 service.
// All implementations must embed UnimplementedAlertEndpointsSyncV2Server
// for forward compatibility
type AlertEndpointsSyncV2Server interface {
	// ImportConfig
	ImportConfig(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// Downstream Syncs config
	SyncConfig(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetFullConfig(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedAlertEndpointsSyncV2Server()
}

// UnimplementedAlertEndpointsSyncV2Server must be embedded to have forward compatible implementations.
type UnimplementedAlertEndpointsSyncV2Server struct {
}

func (UnimplementedAlertEndpointsSyncV2Server) ImportConfig(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportConfig not implemented")
}
func (UnimplementedAlertEndpointsSyncV2Server) SyncConfig(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncConfig not implemented")
}
func (UnimplementedAlertEndpointsSyncV2Server) GetFullConfig(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullConfig not implemented")
}
func (UnimplementedAlertEndpointsSyncV2Server) mustEmbedUnimplementedAlertEndpointsSyncV2Server() {}

// UnsafeAlertEndpointsSyncV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertEndpointsSyncV2Server will
// result in compilation errors.
type UnsafeAlertEndpointsSyncV2Server interface {
	mustEmbedUnimplementedAlertEndpointsSyncV2Server()
}

func RegisterAlertEndpointsSyncV2Server(s grpc.ServiceRegistrar, srv AlertEndpointsSyncV2Server) {
	s.RegisterService(&AlertEndpointsSyncV2_ServiceDesc, srv)
}

func _AlertEndpointsSyncV2_ImportConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsSyncV2Server).ImportConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsSyncV2/ImportConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsSyncV2Server).ImportConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsSyncV2_SyncConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsSyncV2Server).SyncConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsSyncV2/SyncConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsSyncV2Server).SyncConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertEndpointsSyncV2_GetFullConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertEndpointsSyncV2Server).GetFullConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alertingv2.AlertEndpointsSyncV2/GetFullConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertEndpointsSyncV2Server).GetFullConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertEndpointsSyncV2_ServiceDesc is the grpc.ServiceDesc for AlertEndpointsSyncV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertEndpointsSyncV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alertingv2.AlertEndpointsSyncV2",
	HandlerType: (*AlertEndpointsSyncV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportConfig",
			Handler:    _AlertEndpointsSyncV2_ImportConfig_Handler,
		},
		{
			MethodName: "SyncConfig",
			Handler:    _AlertEndpointsSyncV2_SyncConfig_Handler,
		},
		{
			MethodName: "GetFullConfig",
			Handler:    _AlertEndpointsSyncV2_GetFullConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/apis/alerting/v2/alerting.endpoint.proto",
}
