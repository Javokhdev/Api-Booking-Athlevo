// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: protos/booking.proto

package booking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BookingPersonalService_CreateBookingPersonal_FullMethodName = "/gym.BookingPersonalService/CreateBookingPersonal"
	BookingPersonalService_GetBookingPersonal_FullMethodName    = "/gym.BookingPersonalService/GetBookingPersonal"
	BookingPersonalService_UpdateBookingPersonal_FullMethodName = "/gym.BookingPersonalService/UpdateBookingPersonal"
	BookingPersonalService_DeleteBookingPersonal_FullMethodName = "/gym.BookingPersonalService/DeleteBookingPersonal"
	BookingPersonalService_ListBookingPersonal_FullMethodName   = "/gym.BookingPersonalService/ListBookingPersonal"
)

// BookingPersonalServiceClient is the client API for BookingPersonalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingPersonalServiceClient interface {
	CreateBookingPersonal(ctx context.Context, in *CreateBookingPersonalRequest, opts ...grpc.CallOption) (*BookingPersonal, error)
	GetBookingPersonal(ctx context.Context, in *GetBookingPersonalRequest, opts ...grpc.CallOption) (*BookingPersonal, error)
	UpdateBookingPersonal(ctx context.Context, in *UpdateBookingPersonalRequest, opts ...grpc.CallOption) (*BookingPersonal, error)
	DeleteBookingPersonal(ctx context.Context, in *DeleteBookingPersonalRequest, opts ...grpc.CallOption) (*Empty, error)
	ListBookingPersonal(ctx context.Context, in *ListBookingPersonalRequest, opts ...grpc.CallOption) (*ListBookingPersonalResponse, error)
}

type bookingPersonalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingPersonalServiceClient(cc grpc.ClientConnInterface) BookingPersonalServiceClient {
	return &bookingPersonalServiceClient{cc}
}

func (c *bookingPersonalServiceClient) CreateBookingPersonal(ctx context.Context, in *CreateBookingPersonalRequest, opts ...grpc.CallOption) (*BookingPersonal, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingPersonal)
	err := c.cc.Invoke(ctx, BookingPersonalService_CreateBookingPersonal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingPersonalServiceClient) GetBookingPersonal(ctx context.Context, in *GetBookingPersonalRequest, opts ...grpc.CallOption) (*BookingPersonal, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingPersonal)
	err := c.cc.Invoke(ctx, BookingPersonalService_GetBookingPersonal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingPersonalServiceClient) UpdateBookingPersonal(ctx context.Context, in *UpdateBookingPersonalRequest, opts ...grpc.CallOption) (*BookingPersonal, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingPersonal)
	err := c.cc.Invoke(ctx, BookingPersonalService_UpdateBookingPersonal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingPersonalServiceClient) DeleteBookingPersonal(ctx context.Context, in *DeleteBookingPersonalRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, BookingPersonalService_DeleteBookingPersonal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingPersonalServiceClient) ListBookingPersonal(ctx context.Context, in *ListBookingPersonalRequest, opts ...grpc.CallOption) (*ListBookingPersonalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBookingPersonalResponse)
	err := c.cc.Invoke(ctx, BookingPersonalService_ListBookingPersonal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingPersonalServiceServer is the server API for BookingPersonalService service.
// All implementations must embed UnimplementedBookingPersonalServiceServer
// for forward compatibility.
type BookingPersonalServiceServer interface {
	CreateBookingPersonal(context.Context, *CreateBookingPersonalRequest) (*BookingPersonal, error)
	GetBookingPersonal(context.Context, *GetBookingPersonalRequest) (*BookingPersonal, error)
	UpdateBookingPersonal(context.Context, *UpdateBookingPersonalRequest) (*BookingPersonal, error)
	DeleteBookingPersonal(context.Context, *DeleteBookingPersonalRequest) (*Empty, error)
	ListBookingPersonal(context.Context, *ListBookingPersonalRequest) (*ListBookingPersonalResponse, error)
	mustEmbedUnimplementedBookingPersonalServiceServer()
}

// UnimplementedBookingPersonalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookingPersonalServiceServer struct{}

func (UnimplementedBookingPersonalServiceServer) CreateBookingPersonal(context.Context, *CreateBookingPersonalRequest) (*BookingPersonal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookingPersonal not implemented")
}
func (UnimplementedBookingPersonalServiceServer) GetBookingPersonal(context.Context, *GetBookingPersonalRequest) (*BookingPersonal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingPersonal not implemented")
}
func (UnimplementedBookingPersonalServiceServer) UpdateBookingPersonal(context.Context, *UpdateBookingPersonalRequest) (*BookingPersonal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookingPersonal not implemented")
}
func (UnimplementedBookingPersonalServiceServer) DeleteBookingPersonal(context.Context, *DeleteBookingPersonalRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookingPersonal not implemented")
}
func (UnimplementedBookingPersonalServiceServer) ListBookingPersonal(context.Context, *ListBookingPersonalRequest) (*ListBookingPersonalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookingPersonal not implemented")
}
func (UnimplementedBookingPersonalServiceServer) mustEmbedUnimplementedBookingPersonalServiceServer() {
}
func (UnimplementedBookingPersonalServiceServer) testEmbeddedByValue() {}

// UnsafeBookingPersonalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingPersonalServiceServer will
// result in compilation errors.
type UnsafeBookingPersonalServiceServer interface {
	mustEmbedUnimplementedBookingPersonalServiceServer()
}

func RegisterBookingPersonalServiceServer(s grpc.ServiceRegistrar, srv BookingPersonalServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookingPersonalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookingPersonalService_ServiceDesc, srv)
}

func _BookingPersonalService_CreateBookingPersonal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingPersonalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingPersonalServiceServer).CreateBookingPersonal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingPersonalService_CreateBookingPersonal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingPersonalServiceServer).CreateBookingPersonal(ctx, req.(*CreateBookingPersonalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingPersonalService_GetBookingPersonal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingPersonalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingPersonalServiceServer).GetBookingPersonal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingPersonalService_GetBookingPersonal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingPersonalServiceServer).GetBookingPersonal(ctx, req.(*GetBookingPersonalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingPersonalService_UpdateBookingPersonal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookingPersonalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingPersonalServiceServer).UpdateBookingPersonal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingPersonalService_UpdateBookingPersonal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingPersonalServiceServer).UpdateBookingPersonal(ctx, req.(*UpdateBookingPersonalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingPersonalService_DeleteBookingPersonal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingPersonalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingPersonalServiceServer).DeleteBookingPersonal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingPersonalService_DeleteBookingPersonal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingPersonalServiceServer).DeleteBookingPersonal(ctx, req.(*DeleteBookingPersonalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingPersonalService_ListBookingPersonal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookingPersonalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingPersonalServiceServer).ListBookingPersonal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingPersonalService_ListBookingPersonal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingPersonalServiceServer).ListBookingPersonal(ctx, req.(*ListBookingPersonalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingPersonalService_ServiceDesc is the grpc.ServiceDesc for BookingPersonalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingPersonalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gym.BookingPersonalService",
	HandlerType: (*BookingPersonalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBookingPersonal",
			Handler:    _BookingPersonalService_CreateBookingPersonal_Handler,
		},
		{
			MethodName: "GetBookingPersonal",
			Handler:    _BookingPersonalService_GetBookingPersonal_Handler,
		},
		{
			MethodName: "UpdateBookingPersonal",
			Handler:    _BookingPersonalService_UpdateBookingPersonal_Handler,
		},
		{
			MethodName: "DeleteBookingPersonal",
			Handler:    _BookingPersonalService_DeleteBookingPersonal_Handler,
		},
		{
			MethodName: "ListBookingPersonal",
			Handler:    _BookingPersonalService_ListBookingPersonal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/booking.proto",
}

const (
	BookingGroupService_CreateBookingGroup_FullMethodName = "/gym.BookingGroupService/CreateBookingGroup"
	BookingGroupService_GetBookingGroup_FullMethodName    = "/gym.BookingGroupService/GetBookingGroup"
	BookingGroupService_UpdateBookingGroup_FullMethodName = "/gym.BookingGroupService/UpdateBookingGroup"
	BookingGroupService_DeleteBookingGroup_FullMethodName = "/gym.BookingGroupService/DeleteBookingGroup"
	BookingGroupService_ListBookingGroup_FullMethodName   = "/gym.BookingGroupService/ListBookingGroup"
)

// BookingGroupServiceClient is the client API for BookingGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingGroupServiceClient interface {
	CreateBookingGroup(ctx context.Context, in *CreateBookingGroupRequest, opts ...grpc.CallOption) (*BookingGroup, error)
	GetBookingGroup(ctx context.Context, in *GetBookingGroupRequest, opts ...grpc.CallOption) (*BookingGroup, error)
	UpdateBookingGroup(ctx context.Context, in *UpdateBookingGroupRequest, opts ...grpc.CallOption) (*BookingGroup, error)
	DeleteBookingGroup(ctx context.Context, in *DeleteBookingGroupRequest, opts ...grpc.CallOption) (*Empty, error)
	ListBookingGroup(ctx context.Context, in *ListBookingGroupRequest, opts ...grpc.CallOption) (*ListBookingGroupResponse, error)
}

type bookingGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingGroupServiceClient(cc grpc.ClientConnInterface) BookingGroupServiceClient {
	return &bookingGroupServiceClient{cc}
}

func (c *bookingGroupServiceClient) CreateBookingGroup(ctx context.Context, in *CreateBookingGroupRequest, opts ...grpc.CallOption) (*BookingGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingGroup)
	err := c.cc.Invoke(ctx, BookingGroupService_CreateBookingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingGroupServiceClient) GetBookingGroup(ctx context.Context, in *GetBookingGroupRequest, opts ...grpc.CallOption) (*BookingGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingGroup)
	err := c.cc.Invoke(ctx, BookingGroupService_GetBookingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingGroupServiceClient) UpdateBookingGroup(ctx context.Context, in *UpdateBookingGroupRequest, opts ...grpc.CallOption) (*BookingGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingGroup)
	err := c.cc.Invoke(ctx, BookingGroupService_UpdateBookingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingGroupServiceClient) DeleteBookingGroup(ctx context.Context, in *DeleteBookingGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, BookingGroupService_DeleteBookingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingGroupServiceClient) ListBookingGroup(ctx context.Context, in *ListBookingGroupRequest, opts ...grpc.CallOption) (*ListBookingGroupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBookingGroupResponse)
	err := c.cc.Invoke(ctx, BookingGroupService_ListBookingGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingGroupServiceServer is the server API for BookingGroupService service.
// All implementations must embed UnimplementedBookingGroupServiceServer
// for forward compatibility.
type BookingGroupServiceServer interface {
	CreateBookingGroup(context.Context, *CreateBookingGroupRequest) (*BookingGroup, error)
	GetBookingGroup(context.Context, *GetBookingGroupRequest) (*BookingGroup, error)
	UpdateBookingGroup(context.Context, *UpdateBookingGroupRequest) (*BookingGroup, error)
	DeleteBookingGroup(context.Context, *DeleteBookingGroupRequest) (*Empty, error)
	ListBookingGroup(context.Context, *ListBookingGroupRequest) (*ListBookingGroupResponse, error)
	mustEmbedUnimplementedBookingGroupServiceServer()
}

// UnimplementedBookingGroupServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookingGroupServiceServer struct{}

func (UnimplementedBookingGroupServiceServer) CreateBookingGroup(context.Context, *CreateBookingGroupRequest) (*BookingGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookingGroup not implemented")
}
func (UnimplementedBookingGroupServiceServer) GetBookingGroup(context.Context, *GetBookingGroupRequest) (*BookingGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingGroup not implemented")
}
func (UnimplementedBookingGroupServiceServer) UpdateBookingGroup(context.Context, *UpdateBookingGroupRequest) (*BookingGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookingGroup not implemented")
}
func (UnimplementedBookingGroupServiceServer) DeleteBookingGroup(context.Context, *DeleteBookingGroupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookingGroup not implemented")
}
func (UnimplementedBookingGroupServiceServer) ListBookingGroup(context.Context, *ListBookingGroupRequest) (*ListBookingGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookingGroup not implemented")
}
func (UnimplementedBookingGroupServiceServer) mustEmbedUnimplementedBookingGroupServiceServer() {}
func (UnimplementedBookingGroupServiceServer) testEmbeddedByValue()                             {}

// UnsafeBookingGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingGroupServiceServer will
// result in compilation errors.
type UnsafeBookingGroupServiceServer interface {
	mustEmbedUnimplementedBookingGroupServiceServer()
}

func RegisterBookingGroupServiceServer(s grpc.ServiceRegistrar, srv BookingGroupServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookingGroupServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookingGroupService_ServiceDesc, srv)
}

func _BookingGroupService_CreateBookingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGroupServiceServer).CreateBookingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingGroupService_CreateBookingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGroupServiceServer).CreateBookingGroup(ctx, req.(*CreateBookingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingGroupService_GetBookingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGroupServiceServer).GetBookingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingGroupService_GetBookingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGroupServiceServer).GetBookingGroup(ctx, req.(*GetBookingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingGroupService_UpdateBookingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGroupServiceServer).UpdateBookingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingGroupService_UpdateBookingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGroupServiceServer).UpdateBookingGroup(ctx, req.(*UpdateBookingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingGroupService_DeleteBookingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGroupServiceServer).DeleteBookingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingGroupService_DeleteBookingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGroupServiceServer).DeleteBookingGroup(ctx, req.(*DeleteBookingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingGroupService_ListBookingGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookingGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingGroupServiceServer).ListBookingGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingGroupService_ListBookingGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingGroupServiceServer).ListBookingGroup(ctx, req.(*ListBookingGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingGroupService_ServiceDesc is the grpc.ServiceDesc for BookingGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gym.BookingGroupService",
	HandlerType: (*BookingGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBookingGroup",
			Handler:    _BookingGroupService_CreateBookingGroup_Handler,
		},
		{
			MethodName: "GetBookingGroup",
			Handler:    _BookingGroupService_GetBookingGroup_Handler,
		},
		{
			MethodName: "UpdateBookingGroup",
			Handler:    _BookingGroupService_UpdateBookingGroup_Handler,
		},
		{
			MethodName: "DeleteBookingGroup",
			Handler:    _BookingGroupService_DeleteBookingGroup_Handler,
		},
		{
			MethodName: "ListBookingGroup",
			Handler:    _BookingGroupService_ListBookingGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/booking.proto",
}

const (
	BookingCoachService_CreateBookingCoach_FullMethodName = "/gym.BookingCoachService/CreateBookingCoach"
	BookingCoachService_GetBookingCoach_FullMethodName    = "/gym.BookingCoachService/GetBookingCoach"
	BookingCoachService_UpdateBookingCoach_FullMethodName = "/gym.BookingCoachService/UpdateBookingCoach"
	BookingCoachService_DeleteBookingCoach_FullMethodName = "/gym.BookingCoachService/DeleteBookingCoach"
	BookingCoachService_ListBookingCoach_FullMethodName   = "/gym.BookingCoachService/ListBookingCoach"
)

// BookingCoachServiceClient is the client API for BookingCoachService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingCoachServiceClient interface {
	CreateBookingCoach(ctx context.Context, in *CreateBookingCoachRequest, opts ...grpc.CallOption) (*BookingCoach, error)
	GetBookingCoach(ctx context.Context, in *GetBookingCoachRequest, opts ...grpc.CallOption) (*BookingCoach, error)
	UpdateBookingCoach(ctx context.Context, in *UpdateBookingCoachRequest, opts ...grpc.CallOption) (*BookingCoach, error)
	DeleteBookingCoach(ctx context.Context, in *DeleteBookingCoachRequest, opts ...grpc.CallOption) (*Empty, error)
	ListBookingCoach(ctx context.Context, in *ListBookingCoachRequest, opts ...grpc.CallOption) (*ListBookingCoachResponse, error)
}

type bookingCoachServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingCoachServiceClient(cc grpc.ClientConnInterface) BookingCoachServiceClient {
	return &bookingCoachServiceClient{cc}
}

func (c *bookingCoachServiceClient) CreateBookingCoach(ctx context.Context, in *CreateBookingCoachRequest, opts ...grpc.CallOption) (*BookingCoach, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingCoach)
	err := c.cc.Invoke(ctx, BookingCoachService_CreateBookingCoach_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingCoachServiceClient) GetBookingCoach(ctx context.Context, in *GetBookingCoachRequest, opts ...grpc.CallOption) (*BookingCoach, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingCoach)
	err := c.cc.Invoke(ctx, BookingCoachService_GetBookingCoach_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingCoachServiceClient) UpdateBookingCoach(ctx context.Context, in *UpdateBookingCoachRequest, opts ...grpc.CallOption) (*BookingCoach, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookingCoach)
	err := c.cc.Invoke(ctx, BookingCoachService_UpdateBookingCoach_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingCoachServiceClient) DeleteBookingCoach(ctx context.Context, in *DeleteBookingCoachRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, BookingCoachService_DeleteBookingCoach_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingCoachServiceClient) ListBookingCoach(ctx context.Context, in *ListBookingCoachRequest, opts ...grpc.CallOption) (*ListBookingCoachResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBookingCoachResponse)
	err := c.cc.Invoke(ctx, BookingCoachService_ListBookingCoach_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingCoachServiceServer is the server API for BookingCoachService service.
// All implementations must embed UnimplementedBookingCoachServiceServer
// for forward compatibility.
type BookingCoachServiceServer interface {
	CreateBookingCoach(context.Context, *CreateBookingCoachRequest) (*BookingCoach, error)
	GetBookingCoach(context.Context, *GetBookingCoachRequest) (*BookingCoach, error)
	UpdateBookingCoach(context.Context, *UpdateBookingCoachRequest) (*BookingCoach, error)
	DeleteBookingCoach(context.Context, *DeleteBookingCoachRequest) (*Empty, error)
	ListBookingCoach(context.Context, *ListBookingCoachRequest) (*ListBookingCoachResponse, error)
	mustEmbedUnimplementedBookingCoachServiceServer()
}

// UnimplementedBookingCoachServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookingCoachServiceServer struct{}

func (UnimplementedBookingCoachServiceServer) CreateBookingCoach(context.Context, *CreateBookingCoachRequest) (*BookingCoach, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookingCoach not implemented")
}
func (UnimplementedBookingCoachServiceServer) GetBookingCoach(context.Context, *GetBookingCoachRequest) (*BookingCoach, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingCoach not implemented")
}
func (UnimplementedBookingCoachServiceServer) UpdateBookingCoach(context.Context, *UpdateBookingCoachRequest) (*BookingCoach, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookingCoach not implemented")
}
func (UnimplementedBookingCoachServiceServer) DeleteBookingCoach(context.Context, *DeleteBookingCoachRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookingCoach not implemented")
}
func (UnimplementedBookingCoachServiceServer) ListBookingCoach(context.Context, *ListBookingCoachRequest) (*ListBookingCoachResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookingCoach not implemented")
}
func (UnimplementedBookingCoachServiceServer) mustEmbedUnimplementedBookingCoachServiceServer() {}
func (UnimplementedBookingCoachServiceServer) testEmbeddedByValue()                             {}

// UnsafeBookingCoachServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingCoachServiceServer will
// result in compilation errors.
type UnsafeBookingCoachServiceServer interface {
	mustEmbedUnimplementedBookingCoachServiceServer()
}

func RegisterBookingCoachServiceServer(s grpc.ServiceRegistrar, srv BookingCoachServiceServer) {
	// If the following call pancis, it indicates UnimplementedBookingCoachServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BookingCoachService_ServiceDesc, srv)
}

func _BookingCoachService_CreateBookingCoach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingCoachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingCoachServiceServer).CreateBookingCoach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingCoachService_CreateBookingCoach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingCoachServiceServer).CreateBookingCoach(ctx, req.(*CreateBookingCoachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingCoachService_GetBookingCoach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingCoachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingCoachServiceServer).GetBookingCoach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingCoachService_GetBookingCoach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingCoachServiceServer).GetBookingCoach(ctx, req.(*GetBookingCoachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingCoachService_UpdateBookingCoach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookingCoachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingCoachServiceServer).UpdateBookingCoach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingCoachService_UpdateBookingCoach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingCoachServiceServer).UpdateBookingCoach(ctx, req.(*UpdateBookingCoachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingCoachService_DeleteBookingCoach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingCoachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingCoachServiceServer).DeleteBookingCoach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingCoachService_DeleteBookingCoach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingCoachServiceServer).DeleteBookingCoach(ctx, req.(*DeleteBookingCoachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingCoachService_ListBookingCoach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookingCoachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingCoachServiceServer).ListBookingCoach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingCoachService_ListBookingCoach_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingCoachServiceServer).ListBookingCoach(ctx, req.(*ListBookingCoachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingCoachService_ServiceDesc is the grpc.ServiceDesc for BookingCoachService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingCoachService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gym.BookingCoachService",
	HandlerType: (*BookingCoachServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBookingCoach",
			Handler:    _BookingCoachService_CreateBookingCoach_Handler,
		},
		{
			MethodName: "GetBookingCoach",
			Handler:    _BookingCoachService_GetBookingCoach_Handler,
		},
		{
			MethodName: "UpdateBookingCoach",
			Handler:    _BookingCoachService_UpdateBookingCoach_Handler,
		},
		{
			MethodName: "DeleteBookingCoach",
			Handler:    _BookingCoachService_DeleteBookingCoach_Handler,
		},
		{
			MethodName: "ListBookingCoach",
			Handler:    _BookingCoachService_ListBookingCoach_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/booking.proto",
}
