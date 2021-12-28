// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// LiveServiceClient is the client API for LiveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiveServiceClient interface {
	CreateLiveHouse(ctx context.Context, in *CreateLiveHouseRequest, opts ...grpc.CallOption) (*CreateLiveHouseResponse, error)
	ListLiveHouses(ctx context.Context, in *ListLiveHousesRequest, opts ...grpc.CallOption) (*ListLiveHousesResponse, error)
	CreateLiveEvent(ctx context.Context, in *CreateLiveEventRequest, opts ...grpc.CallOption) (*CreateLiveEventResponse, error)
	GetLiveEvent(ctx context.Context, in *GetLiveEventRequest, opts ...grpc.CallOption) (*GetLiveEventResponse, error)
	ListLiveEvents(ctx context.Context, in *ListLiveEventsRequest, opts ...grpc.CallOption) (*ListLiveEventsResponse, error)
	ReserveSeat(ctx context.Context, in *ReserveSeatRequest, opts ...grpc.CallOption) (*ReserveSeatResponse, error)
	RollbackSeatReservation(ctx context.Context, in *RollbackSeatReservationRequest, opts ...grpc.CallOption) (*RollbackSeatReservationResponse, error)
}

type liveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiveServiceClient(cc grpc.ClientConnInterface) LiveServiceClient {
	return &liveServiceClient{cc}
}

func (c *liveServiceClient) CreateLiveHouse(ctx context.Context, in *CreateLiveHouseRequest, opts ...grpc.CallOption) (*CreateLiveHouseResponse, error) {
	out := new(CreateLiveHouseResponse)
	err := c.cc.Invoke(ctx, "/neptuneg.go_back.serivces.live.LiveService/CreateLiveHouse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) ListLiveHouses(ctx context.Context, in *ListLiveHousesRequest, opts ...grpc.CallOption) (*ListLiveHousesResponse, error) {
	out := new(ListLiveHousesResponse)
	err := c.cc.Invoke(ctx, "/neptuneg.go_back.serivces.live.LiveService/ListLiveHouses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) CreateLiveEvent(ctx context.Context, in *CreateLiveEventRequest, opts ...grpc.CallOption) (*CreateLiveEventResponse, error) {
	out := new(CreateLiveEventResponse)
	err := c.cc.Invoke(ctx, "/neptuneg.go_back.serivces.live.LiveService/CreateLiveEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) GetLiveEvent(ctx context.Context, in *GetLiveEventRequest, opts ...grpc.CallOption) (*GetLiveEventResponse, error) {
	out := new(GetLiveEventResponse)
	err := c.cc.Invoke(ctx, "/neptuneg.go_back.serivces.live.LiveService/GetLiveEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) ListLiveEvents(ctx context.Context, in *ListLiveEventsRequest, opts ...grpc.CallOption) (*ListLiveEventsResponse, error) {
	out := new(ListLiveEventsResponse)
	err := c.cc.Invoke(ctx, "/neptuneg.go_back.serivces.live.LiveService/ListLiveEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) ReserveSeat(ctx context.Context, in *ReserveSeatRequest, opts ...grpc.CallOption) (*ReserveSeatResponse, error) {
	out := new(ReserveSeatResponse)
	err := c.cc.Invoke(ctx, "/neptuneg.go_back.serivces.live.LiveService/ReserveSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liveServiceClient) RollbackSeatReservation(ctx context.Context, in *RollbackSeatReservationRequest, opts ...grpc.CallOption) (*RollbackSeatReservationResponse, error) {
	out := new(RollbackSeatReservationResponse)
	err := c.cc.Invoke(ctx, "/neptuneg.go_back.serivces.live.LiveService/RollbackSeatReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiveServiceServer is the server API for LiveService service.
// All implementations must embed UnimplementedLiveServiceServer
// for forward compatibility
type LiveServiceServer interface {
	CreateLiveHouse(context.Context, *CreateLiveHouseRequest) (*CreateLiveHouseResponse, error)
	ListLiveHouses(context.Context, *ListLiveHousesRequest) (*ListLiveHousesResponse, error)
	CreateLiveEvent(context.Context, *CreateLiveEventRequest) (*CreateLiveEventResponse, error)
	GetLiveEvent(context.Context, *GetLiveEventRequest) (*GetLiveEventResponse, error)
	ListLiveEvents(context.Context, *ListLiveEventsRequest) (*ListLiveEventsResponse, error)
	ReserveSeat(context.Context, *ReserveSeatRequest) (*ReserveSeatResponse, error)
	RollbackSeatReservation(context.Context, *RollbackSeatReservationRequest) (*RollbackSeatReservationResponse, error)
	mustEmbedUnimplementedLiveServiceServer()
}

// UnimplementedLiveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiveServiceServer struct {
}

func (UnimplementedLiveServiceServer) CreateLiveHouse(context.Context, *CreateLiveHouseRequest) (*CreateLiveHouseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLiveHouse not implemented")
}
func (UnimplementedLiveServiceServer) ListLiveHouses(context.Context, *ListLiveHousesRequest) (*ListLiveHousesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLiveHouses not implemented")
}
func (UnimplementedLiveServiceServer) CreateLiveEvent(context.Context, *CreateLiveEventRequest) (*CreateLiveEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLiveEvent not implemented")
}
func (UnimplementedLiveServiceServer) GetLiveEvent(context.Context, *GetLiveEventRequest) (*GetLiveEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveEvent not implemented")
}
func (UnimplementedLiveServiceServer) ListLiveEvents(context.Context, *ListLiveEventsRequest) (*ListLiveEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLiveEvents not implemented")
}
func (UnimplementedLiveServiceServer) ReserveSeat(context.Context, *ReserveSeatRequest) (*ReserveSeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveSeat not implemented")
}
func (UnimplementedLiveServiceServer) RollbackSeatReservation(context.Context, *RollbackSeatReservationRequest) (*RollbackSeatReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackSeatReservation not implemented")
}
func (UnimplementedLiveServiceServer) mustEmbedUnimplementedLiveServiceServer() {}

// UnsafeLiveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiveServiceServer will
// result in compilation errors.
type UnsafeLiveServiceServer interface {
	mustEmbedUnimplementedLiveServiceServer()
}

func RegisterLiveServiceServer(s grpc.ServiceRegistrar, srv LiveServiceServer) {
	s.RegisterService(&LiveService_ServiceDesc, srv)
}

func _LiveService_CreateLiveHouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLiveHouseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).CreateLiveHouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptuneg.go_back.serivces.live.LiveService/CreateLiveHouse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).CreateLiveHouse(ctx, req.(*CreateLiveHouseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_ListLiveHouses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLiveHousesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).ListLiveHouses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptuneg.go_back.serivces.live.LiveService/ListLiveHouses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).ListLiveHouses(ctx, req.(*ListLiveHousesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_CreateLiveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLiveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).CreateLiveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptuneg.go_back.serivces.live.LiveService/CreateLiveEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).CreateLiveEvent(ctx, req.(*CreateLiveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_GetLiveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLiveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).GetLiveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptuneg.go_back.serivces.live.LiveService/GetLiveEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).GetLiveEvent(ctx, req.(*GetLiveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_ListLiveEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLiveEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).ListLiveEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptuneg.go_back.serivces.live.LiveService/ListLiveEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).ListLiveEvents(ctx, req.(*ListLiveEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_ReserveSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveSeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).ReserveSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptuneg.go_back.serivces.live.LiveService/ReserveSeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).ReserveSeat(ctx, req.(*ReserveSeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiveService_RollbackSeatReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackSeatReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiveServiceServer).RollbackSeatReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neptuneg.go_back.serivces.live.LiveService/RollbackSeatReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiveServiceServer).RollbackSeatReservation(ctx, req.(*RollbackSeatReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LiveService_ServiceDesc is the grpc.ServiceDesc for LiveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "neptuneg.go_back.serivces.live.LiveService",
	HandlerType: (*LiveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLiveHouse",
			Handler:    _LiveService_CreateLiveHouse_Handler,
		},
		{
			MethodName: "ListLiveHouses",
			Handler:    _LiveService_ListLiveHouses_Handler,
		},
		{
			MethodName: "CreateLiveEvent",
			Handler:    _LiveService_CreateLiveEvent_Handler,
		},
		{
			MethodName: "GetLiveEvent",
			Handler:    _LiveService_GetLiveEvent_Handler,
		},
		{
			MethodName: "ListLiveEvents",
			Handler:    _LiveService_ListLiveEvents_Handler,
		},
		{
			MethodName: "ReserveSeat",
			Handler:    _LiveService_ReserveSeat_Handler,
		},
		{
			MethodName: "RollbackSeatReservation",
			Handler:    _LiveService_RollbackSeatReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/live/proto/live.proto",
}
