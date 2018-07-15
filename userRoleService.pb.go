// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: userRoleService.proto

package protobuf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserRoleService service

type UserRoleServiceClient interface {
	CreateUserRoles(ctx context.Context, in *CreateUserRolesRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetRoleIdsOfUserRole(ctx context.Context, in *GetRoleIdsOfUserRoleRequest, opts ...grpc.CallOption) (*RoleIds, error)
	GetUserIdsOfUserRole(ctx context.Context, in *GetUserIdsOfUserRoleRequest, opts ...grpc.CallOption) (*UserIds, error)
	DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type userRoleServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserRoleServiceClient(cc *grpc.ClientConn) UserRoleServiceClient {
	return &userRoleServiceClient{cc}
}

func (c *userRoleServiceClient) CreateUserRoles(ctx context.Context, in *CreateUserRolesRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.UserRoleService/CreateUserRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) GetRoleIdsOfUserRole(ctx context.Context, in *GetRoleIdsOfUserRoleRequest, opts ...grpc.CallOption) (*RoleIds, error) {
	out := new(RoleIds)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.UserRoleService/GetRoleIdsOfUserRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) GetUserIdsOfUserRole(ctx context.Context, in *GetUserIdsOfUserRoleRequest, opts ...grpc.CallOption) (*UserIds, error) {
	out := new(UserIds)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.UserRoleService/GetUserIdsOfUserRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) DeleteUserRole(ctx context.Context, in *DeleteUserRoleRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.UserRoleService/DeleteUserRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserRoleService service

type UserRoleServiceServer interface {
	CreateUserRoles(context.Context, *CreateUserRolesRequest) (*google_protobuf1.Empty, error)
	GetRoleIdsOfUserRole(context.Context, *GetRoleIdsOfUserRoleRequest) (*RoleIds, error)
	GetUserIdsOfUserRole(context.Context, *GetUserIdsOfUserRoleRequest) (*UserIds, error)
	DeleteUserRole(context.Context, *DeleteUserRoleRequest) (*google_protobuf1.Empty, error)
}

func RegisterUserRoleServiceServer(s *grpc.Server, srv UserRoleServiceServer) {
	s.RegisterService(&_UserRoleService_serviceDesc, srv)
}

func _UserRoleService_CreateUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).CreateUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.UserRoleService/CreateUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).CreateUserRoles(ctx, req.(*CreateUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_GetRoleIdsOfUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleIdsOfUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).GetRoleIdsOfUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.UserRoleService/GetRoleIdsOfUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).GetRoleIdsOfUserRole(ctx, req.(*GetRoleIdsOfUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_GetUserIdsOfUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdsOfUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).GetUserIdsOfUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.UserRoleService/GetUserIdsOfUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).GetUserIdsOfUserRole(ctx, req.(*GetUserIdsOfUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_DeleteUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).DeleteUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.UserRoleService/DeleteUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).DeleteUserRole(ctx, req.(*DeleteUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "swagchat.protobuf.UserRoleService",
	HandlerType: (*UserRoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserRoles",
			Handler:    _UserRoleService_CreateUserRoles_Handler,
		},
		{
			MethodName: "GetRoleIdsOfUserRole",
			Handler:    _UserRoleService_GetRoleIdsOfUserRole_Handler,
		},
		{
			MethodName: "GetUserIdsOfUserRole",
			Handler:    _UserRoleService_GetUserIdsOfUserRole_Handler,
		},
		{
			MethodName: "DeleteUserRole",
			Handler:    _UserRoleService_DeleteUserRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userRoleService.proto",
}

func init() { proto.RegisterFile("userRoleService.proto", fileDescriptorUserRoleService) }

var fileDescriptorUserRoleService = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x17, 0x04, 0x0f, 0x39, 0xb8, 0x18, 0x3f, 0x0e, 0x71, 0xf1, 0xb0, 0x27, 0xbd, 0xa4,
	0xa0, 0x6f, 0xe0, 0x07, 0xe2, 0x41, 0x84, 0xf5, 0xe3, 0xe0, 0x2d, 0xed, 0xce, 0x66, 0x0b, 0x6d,
	0xa7, 0x76, 0x26, 0x8a, 0xcf, 0xe0, 0x4b, 0x4b, 0x6d, 0xd2, 0xa5, 0x36, 0xa8, 0xb7, 0x24, 0xf3,
	0x9b, 0xf9, 0xfd, 0x93, 0x88, 0x03, 0x47, 0xd0, 0x2c, 0xb0, 0x80, 0x07, 0x68, 0xde, 0xf2, 0x0c,
	0x74, 0xdd, 0x20, 0xa3, 0xdc, 0xa5, 0x77, 0x63, 0xb3, 0xb5, 0xe1, 0x6e, 0x9f, 0xba, 0x95, 0x9a,
	0x59, 0x44, 0x5b, 0x40, 0x62, 0xea, 0x3c, 0x31, 0x55, 0x85, 0x6c, 0x38, 0xc7, 0x8a, 0x3a, 0x40,
	0x1d, 0xf9, 0x6a, 0xc0, 0x13, 0x28, 0x6b, 0xfe, 0xf0, 0xc5, 0x5e, 0x72, 0x07, 0x44, 0xc6, 0x7a,
	0x89, 0xda, 0xcb, 0xb0, 0x2c, 0xb1, 0x1a, 0x1c, 0x9e, 0x7d, 0x6e, 0x89, 0xe9, 0xd3, 0x30, 0x93,
	0x7c, 0x16, 0xd3, 0xcb, 0x06, 0x0c, 0x43, 0x28, 0x90, 0x3c, 0xd5, 0xa3, 0x84, 0xfa, 0x07, 0xb3,
	0x80, 0x57, 0x07, 0xc4, 0xea, 0x50, 0x77, 0xd9, 0x36, 0xe0, 0x75, 0x9b, 0x6d, 0x3e, 0x91, 0x4b,
	0xb1, 0x7f, 0x03, 0xdc, 0xc2, 0xb7, 0x4b, 0xba, 0x5f, 0x85, 0x4e, 0xa9, 0x23, 0xc3, 0x63, 0x60,
	0x30, 0xa8, 0x08, 0xef, 0xe1, 0xde, 0xd2, 0xf6, 0xfc, 0xcb, 0x32, 0x02, 0x7f, 0xb3, 0x78, 0x78,
	0x3e, 0x91, 0x8f, 0x62, 0xe7, 0x0a, 0x0a, 0xd8, 0xdc, 0x5f, 0x9e, 0x44, 0xf8, 0x21, 0xf2, 0xe7,
	0x0b, 0x5d, 0x1c, 0xbf, 0xcc, 0x6c, 0xce, 0x6b, 0x97, 0xea, 0x0c, 0xcb, 0x24, 0xcc, 0xeb, 0x7f,
	0x39, 0xdd, 0xfe, 0x5e, 0x9d, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x18, 0x2c, 0x62, 0x61, 0x47,
	0x02, 0x00, 0x00,
}