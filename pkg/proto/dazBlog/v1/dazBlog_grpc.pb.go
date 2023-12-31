// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: dazBlog/v1/dazBlog.proto

package v1

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
	DazBlog_ListUsers_FullMethodName = "/v1.DazBlog/ListUsers"
)

// DazBlogClient is the client API for DazBlog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DazBlogClient interface {
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type dazBlogClient struct {
	cc grpc.ClientConnInterface
}

func NewDazBlogClient(cc grpc.ClientConnInterface) DazBlogClient {
	return &dazBlogClient{cc}
}

func (c *dazBlogClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, DazBlog_ListUsers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DazBlogServer is the server API for DazBlog service.
// All implementations must embed UnimplementedDazBlogServer
// for forward compatibility
type DazBlogServer interface {
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	mustEmbedUnimplementedDazBlogServer()
}

// UnimplementedDazBlogServer must be embedded to have forward compatible implementations.
type UnimplementedDazBlogServer struct {
}

func (UnimplementedDazBlogServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedDazBlogServer) mustEmbedUnimplementedDazBlogServer() {}

// UnsafeDazBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DazBlogServer will
// result in compilation errors.
type UnsafeDazBlogServer interface {
	mustEmbedUnimplementedDazBlogServer()
}

func RegisterDazBlogServer(s grpc.ServiceRegistrar, srv DazBlogServer) {
	s.RegisterService(&DazBlog_ServiceDesc, srv)
}

func _DazBlog_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DazBlogServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DazBlog_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DazBlogServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DazBlog_ServiceDesc is the grpc.ServiceDesc for DazBlog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DazBlog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DazBlog",
	HandlerType: (*DazBlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsers",
			Handler:    _DazBlog_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dazBlog/v1/dazBlog.proto",
}
