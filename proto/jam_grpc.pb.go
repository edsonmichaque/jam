// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: proto/jam.proto

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

const (
	Jam_CreateJail_FullMethodName = "/Jam/CreateJail"
	Jam_ListJails_FullMethodName  = "/Jam/ListJails"
)

// JamClient is the client API for Jam service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JamClient interface {
	CreateJail(ctx context.Context, in *CreateJailRequest, opts ...grpc.CallOption) (*CreateJailResponse, error)
	ListJails(ctx context.Context, in *ListJailsRequest, opts ...grpc.CallOption) (*ListJailsResponse, error)
}

type jamClient struct {
	cc grpc.ClientConnInterface
}

func NewJamClient(cc grpc.ClientConnInterface) JamClient {
	return &jamClient{cc}
}

func (c *jamClient) CreateJail(ctx context.Context, in *CreateJailRequest, opts ...grpc.CallOption) (*CreateJailResponse, error) {
	out := new(CreateJailResponse)
	err := c.cc.Invoke(ctx, Jam_CreateJail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jamClient) ListJails(ctx context.Context, in *ListJailsRequest, opts ...grpc.CallOption) (*ListJailsResponse, error) {
	out := new(ListJailsResponse)
	err := c.cc.Invoke(ctx, Jam_ListJails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JamServer is the server API for Jam service.
// All implementations must embed UnimplementedJamServer
// for forward compatibility
type JamServer interface {
	CreateJail(context.Context, *CreateJailRequest) (*CreateJailResponse, error)
	ListJails(context.Context, *ListJailsRequest) (*ListJailsResponse, error)
	mustEmbedUnimplementedJamServer()
}

// UnimplementedJamServer must be embedded to have forward compatible implementations.
type UnimplementedJamServer struct {
}

func (UnimplementedJamServer) CreateJail(context.Context, *CreateJailRequest) (*CreateJailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJail not implemented")
}
func (UnimplementedJamServer) ListJails(context.Context, *ListJailsRequest) (*ListJailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJails not implemented")
}
func (UnimplementedJamServer) mustEmbedUnimplementedJamServer() {}

// UnsafeJamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JamServer will
// result in compilation errors.
type UnsafeJamServer interface {
	mustEmbedUnimplementedJamServer()
}

func RegisterJamServer(s grpc.ServiceRegistrar, srv JamServer) {
	s.RegisterService(&Jam_ServiceDesc, srv)
}

func _Jam_CreateJail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JamServer).CreateJail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jam_CreateJail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JamServer).CreateJail(ctx, req.(*CreateJailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jam_ListJails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JamServer).ListJails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jam_ListJails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JamServer).ListJails(ctx, req.(*ListJailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jam_ServiceDesc is the grpc.ServiceDesc for Jam service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jam_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Jam",
	HandlerType: (*JamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJail",
			Handler:    _Jam_CreateJail_Handler,
		},
		{
			MethodName: "ListJails",
			Handler:    _Jam_ListJails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/jam.proto",
}
