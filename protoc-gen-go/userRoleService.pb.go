// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: userRoleService.proto

package protoc_gen_go

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
	AddUserRoles(ctx context.Context, in *AddUserRolesRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	DeleteUserRoles(ctx context.Context, in *DeleteUserRolesRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type userRoleServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserRoleServiceClient(cc *grpc.ClientConn) UserRoleServiceClient {
	return &userRoleServiceClient{cc}
}

func (c *userRoleServiceClient) AddUserRoles(ctx context.Context, in *AddUserRolesRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.UserRoleService/AddUserRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRoleServiceClient) DeleteUserRoles(ctx context.Context, in *DeleteUserRolesRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/swagchat.protobuf.UserRoleService/DeleteUserRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserRoleService service

type UserRoleServiceServer interface {
	AddUserRoles(context.Context, *AddUserRolesRequest) (*google_protobuf1.Empty, error)
	DeleteUserRoles(context.Context, *DeleteUserRolesRequest) (*google_protobuf1.Empty, error)
}

func RegisterUserRoleServiceServer(s *grpc.Server, srv UserRoleServiceServer) {
	s.RegisterService(&_UserRoleService_serviceDesc, srv)
}

func _UserRoleService_AddUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).AddUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.UserRoleService/AddUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).AddUserRoles(ctx, req.(*AddUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRoleService_DeleteUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRoleServiceServer).DeleteUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/swagchat.protobuf.UserRoleService/DeleteUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRoleServiceServer).DeleteUserRoles(ctx, req.(*DeleteUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "swagchat.protobuf.UserRoleService",
	HandlerType: (*UserRoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUserRoles",
			Handler:    _UserRoleService_AddUserRoles_Handler,
		},
		{
			MethodName: "DeleteUserRoles",
			Handler:    _UserRoleService_DeleteUserRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userRoleService.proto",
}

func init() { proto.RegisterFile("userRoleService.proto", fileDescriptorUserRoleService) }

var fileDescriptorUserRoleService = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2d, 0x4e, 0x2d,
	0x0a, 0xca, 0xcf, 0x49, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x2c, 0x2e, 0x4f, 0x4c, 0x4f, 0xce, 0x48, 0x2c, 0x81, 0xf0, 0x93, 0x4a, 0xd3,
	0xa4, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2,
	0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x0a, 0xa4, 0xa4, 0xa1, 0xb2, 0x30, 0xe5,
	0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x50, 0x49, 0xb8, 0x25, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9,
	0x50, 0x4b, 0x8c, 0x76, 0x33, 0x72, 0xf1, 0x87, 0xa2, 0x5a, 0x2f, 0x14, 0xc0, 0xc5, 0xe3, 0x98,
	0x92, 0x02, 0x13, 0x2d, 0x16, 0x52, 0xd3, 0xc3, 0x70, 0x89, 0x1e, 0xb2, 0x82, 0xa0, 0xd4, 0xc2,
	0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x31, 0x3d, 0x88, 0x03, 0x10, 0xaa, 0x5c, 0x41, 0x0e, 0x50, 0x62,
	0x10, 0x0a, 0xe3, 0xe2, 0x77, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0x45, 0x18, 0xaa, 0x89, 0xc5, 0x50,
	0x34, 0x35, 0x04, 0xcd, 0x75, 0xd2, 0x89, 0xd2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x87, 0x19, 0x88, 0x08, 0x00, 0x30, 0x23, 0x59, 0x37, 0x3d, 0x35, 0x4f, 0x37,
	0x3d, 0x3f, 0x89, 0x0d, 0xcc, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x77, 0xd9, 0x4b, 0xa3,
	0x70, 0x01, 0x00, 0x00,
}
