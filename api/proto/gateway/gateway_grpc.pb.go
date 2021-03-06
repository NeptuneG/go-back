// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: api/proto/gateway/gateway.proto

package gateway

import (
	context "context"
	auth "github.com/NeptuneG/go-back/api/proto/auth"
	live "github.com/NeptuneG/go-back/api/proto/live"
	payment "github.com/NeptuneG/go-back/api/proto/payment"
	scraper "github.com/NeptuneG/go-back/api/proto/scraper"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayServiceClient interface {
	Register(ctx context.Context, in *auth.RegisterRequest, opts ...grpc.CallOption) (*auth.RegisterResponse, error)
	Login(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error)
	GetUserPoints(ctx context.Context, in *payment.GetUserPointsRequest, opts ...grpc.CallOption) (*payment.GetUserPointsResponse, error)
	CreateUserPoints(ctx context.Context, in *payment.CreateUserPointsRequest, opts ...grpc.CallOption) (*payment.CreateUserPointsResponse, error)
	ListLiveHouses(ctx context.Context, in *live.ListLiveHousesRequest, opts ...grpc.CallOption) (*live.ListLiveHousesResponse, error)
	ListLiveEvents(ctx context.Context, in *live.ListLiveEventsRequest, opts ...grpc.CallOption) (*live.ListLiveEventsResponse, error)
	CreateLiveEventOrder(ctx context.Context, in *payment.CreateLiveEventOrderRequest, opts ...grpc.CallOption) (*payment.CreateLiveEventOrderResponse, error)
	CreateScrapeLiveEventsJob(ctx context.Context, in *scraper.CreateScrapeLiveEventsJobRequest, opts ...grpc.CallOption) (*scraper.CreateScrapeLiveEventsJobResponse, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) Register(ctx context.Context, in *auth.RegisterRequest, opts ...grpc.CallOption) (*auth.RegisterResponse, error) {
	out := new(auth.RegisterResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Login(ctx context.Context, in *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error) {
	out := new(auth.LoginResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetUserPoints(ctx context.Context, in *payment.GetUserPointsRequest, opts ...grpc.CallOption) (*payment.GetUserPointsResponse, error) {
	out := new(payment.GetUserPointsResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/GetUserPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) CreateUserPoints(ctx context.Context, in *payment.CreateUserPointsRequest, opts ...grpc.CallOption) (*payment.CreateUserPointsResponse, error) {
	out := new(payment.CreateUserPointsResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/CreateUserPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) ListLiveHouses(ctx context.Context, in *live.ListLiveHousesRequest, opts ...grpc.CallOption) (*live.ListLiveHousesResponse, error) {
	out := new(live.ListLiveHousesResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/ListLiveHouses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) ListLiveEvents(ctx context.Context, in *live.ListLiveEventsRequest, opts ...grpc.CallOption) (*live.ListLiveEventsResponse, error) {
	out := new(live.ListLiveEventsResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/ListLiveEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) CreateLiveEventOrder(ctx context.Context, in *payment.CreateLiveEventOrderRequest, opts ...grpc.CallOption) (*payment.CreateLiveEventOrderResponse, error) {
	out := new(payment.CreateLiveEventOrderResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/CreateLiveEventOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) CreateScrapeLiveEventsJob(ctx context.Context, in *scraper.CreateScrapeLiveEventsJobRequest, opts ...grpc.CallOption) (*scraper.CreateScrapeLiveEventsJobResponse, error) {
	out := new(scraper.CreateScrapeLiveEventsJobResponse)
	err := c.cc.Invoke(ctx, "/api.proto.gateway.GatewayService/CreateScrapeLiveEventsJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
// All implementations must embed UnimplementedGatewayServiceServer
// for forward compatibility
type GatewayServiceServer interface {
	Register(context.Context, *auth.RegisterRequest) (*auth.RegisterResponse, error)
	Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error)
	GetUserPoints(context.Context, *payment.GetUserPointsRequest) (*payment.GetUserPointsResponse, error)
	CreateUserPoints(context.Context, *payment.CreateUserPointsRequest) (*payment.CreateUserPointsResponse, error)
	ListLiveHouses(context.Context, *live.ListLiveHousesRequest) (*live.ListLiveHousesResponse, error)
	ListLiveEvents(context.Context, *live.ListLiveEventsRequest) (*live.ListLiveEventsResponse, error)
	CreateLiveEventOrder(context.Context, *payment.CreateLiveEventOrderRequest) (*payment.CreateLiveEventOrderResponse, error)
	CreateScrapeLiveEventsJob(context.Context, *scraper.CreateScrapeLiveEventsJobRequest) (*scraper.CreateScrapeLiveEventsJobResponse, error)
	mustEmbedUnimplementedGatewayServiceServer()
}

// UnimplementedGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (UnimplementedGatewayServiceServer) Register(context.Context, *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGatewayServiceServer) Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGatewayServiceServer) GetUserPoints(context.Context, *payment.GetUserPointsRequest) (*payment.GetUserPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPoints not implemented")
}
func (UnimplementedGatewayServiceServer) CreateUserPoints(context.Context, *payment.CreateUserPointsRequest) (*payment.CreateUserPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserPoints not implemented")
}
func (UnimplementedGatewayServiceServer) ListLiveHouses(context.Context, *live.ListLiveHousesRequest) (*live.ListLiveHousesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLiveHouses not implemented")
}
func (UnimplementedGatewayServiceServer) ListLiveEvents(context.Context, *live.ListLiveEventsRequest) (*live.ListLiveEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLiveEvents not implemented")
}
func (UnimplementedGatewayServiceServer) CreateLiveEventOrder(context.Context, *payment.CreateLiveEventOrderRequest) (*payment.CreateLiveEventOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLiveEventOrder not implemented")
}
func (UnimplementedGatewayServiceServer) CreateScrapeLiveEventsJob(context.Context, *scraper.CreateScrapeLiveEventsJobRequest) (*scraper.CreateScrapeLiveEventsJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScrapeLiveEventsJob not implemented")
}
func (UnimplementedGatewayServiceServer) mustEmbedUnimplementedGatewayServiceServer() {}

// UnsafeGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServiceServer will
// result in compilation errors.
type UnsafeGatewayServiceServer interface {
	mustEmbedUnimplementedGatewayServiceServer()
}

func RegisterGatewayServiceServer(s grpc.ServiceRegistrar, srv GatewayServiceServer) {
	s.RegisterService(&GatewayService_ServiceDesc, srv)
}

func _GatewayService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Register(ctx, req.(*auth.RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(auth.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Login(ctx, req.(*auth.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetUserPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payment.GetUserPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetUserPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/GetUserPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetUserPoints(ctx, req.(*payment.GetUserPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_CreateUserPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payment.CreateUserPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).CreateUserPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/CreateUserPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).CreateUserPoints(ctx, req.(*payment.CreateUserPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_ListLiveHouses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(live.ListLiveHousesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).ListLiveHouses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/ListLiveHouses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).ListLiveHouses(ctx, req.(*live.ListLiveHousesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_ListLiveEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(live.ListLiveEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).ListLiveEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/ListLiveEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).ListLiveEvents(ctx, req.(*live.ListLiveEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_CreateLiveEventOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(payment.CreateLiveEventOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).CreateLiveEventOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/CreateLiveEventOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).CreateLiveEventOrder(ctx, req.(*payment.CreateLiveEventOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_CreateScrapeLiveEventsJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(scraper.CreateScrapeLiveEventsJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).CreateScrapeLiveEventsJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.proto.gateway.GatewayService/CreateScrapeLiveEventsJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).CreateScrapeLiveEventsJob(ctx, req.(*scraper.CreateScrapeLiveEventsJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayService_ServiceDesc is the grpc.ServiceDesc for GatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.proto.gateway.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GatewayService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GatewayService_Login_Handler,
		},
		{
			MethodName: "GetUserPoints",
			Handler:    _GatewayService_GetUserPoints_Handler,
		},
		{
			MethodName: "CreateUserPoints",
			Handler:    _GatewayService_CreateUserPoints_Handler,
		},
		{
			MethodName: "ListLiveHouses",
			Handler:    _GatewayService_ListLiveHouses_Handler,
		},
		{
			MethodName: "ListLiveEvents",
			Handler:    _GatewayService_ListLiveEvents_Handler,
		},
		{
			MethodName: "CreateLiveEventOrder",
			Handler:    _GatewayService_CreateLiveEventOrder_Handler,
		},
		{
			MethodName: "CreateScrapeLiveEventsJob",
			Handler:    _GatewayService_CreateScrapeLiveEventsJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/gateway/gateway.proto",
}
