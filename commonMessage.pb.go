// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commonMessage.proto

package protobuf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RoomIds struct {
	RoomIDs []string `protobuf:"bytes,1,rep,name=room_ids,json=roomIds" json:"roomIds,omitempty"`
}

func (m *RoomIds) Reset()                    { *m = RoomIds{} }
func (m *RoomIds) String() string            { return proto.CompactTextString(m) }
func (*RoomIds) ProtoMessage()               {}
func (*RoomIds) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{0} }

func (m *RoomIds) GetRoomIDs() []string {
	if m != nil {
		return m.RoomIDs
	}
	return nil
}

type UserIds struct {
	UserIDs []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds" json:"userIds,omitempty"`
}

func (m *UserIds) Reset()                    { *m = UserIds{} }
func (m *UserIds) String() string            { return proto.CompactTextString(m) }
func (*UserIds) ProtoMessage()               {}
func (*UserIds) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{1} }

func (m *UserIds) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type RoleIds struct {
	RoleIDs []int32 `protobuf:"varint,1,rep,packed,name=role_ids,json=roleIds" json:"roleIds"`
}

func (m *RoleIds) Reset()                    { *m = RoleIds{} }
func (m *RoleIds) String() string            { return proto.CompactTextString(m) }
func (*RoleIds) ProtoMessage()               {}
func (*RoleIds) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{2} }

func (m *RoleIds) GetRoleIDs() []int32 {
	if m != nil {
		return m.RoleIDs
	}
	return nil
}

type ErrorResponse struct {
	Message          string          `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	DeveloperMessage string          `protobuf:"bytes,2,opt,name=developer_message,json=developerMessage,proto3" json:"developerMessage"`
	Info             string          `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	InvalidParams    []*InvalidParam `protobuf:"bytes,4,rep,name=invalid_params,json=invalidParams" json:"invalidParams"`
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{3} }

func (m *ErrorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ErrorResponse) GetDeveloperMessage() string {
	if m != nil {
		return m.DeveloperMessage
	}
	return ""
}

func (m *ErrorResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *ErrorResponse) GetInvalidParams() []*InvalidParam {
	if m != nil {
		return m.InvalidParams
	}
	return nil
}

type InvalidParam struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (m *InvalidParam) Reset()                    { *m = InvalidParam{} }
func (m *InvalidParam) String() string            { return proto.CompactTextString(m) }
func (*InvalidParam) ProtoMessage()               {}
func (*InvalidParam) Descriptor() ([]byte, []int) { return fileDescriptorCommonMessage, []int{4} }

func (m *InvalidParam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InvalidParam) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*RoomIds)(nil), "swagchat.protobuf.RoomIds")
	proto.RegisterType((*UserIds)(nil), "swagchat.protobuf.UserIds")
	proto.RegisterType((*RoleIds)(nil), "swagchat.protobuf.RoleIds")
	proto.RegisterType((*ErrorResponse)(nil), "swagchat.protobuf.ErrorResponse")
	proto.RegisterType((*InvalidParam)(nil), "swagchat.protobuf.InvalidParam")
}

func init() { proto.RegisterFile("commonMessage.proto", fileDescriptorCommonMessage) }

var fileDescriptorCommonMessage = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4b, 0xf3, 0x30,
	0x18, 0xc6, 0xe9, 0xb7, 0x7d, 0xeb, 0xb7, 0x7c, 0x4e, 0x6c, 0x1c, 0xa3, 0x88, 0xd8, 0xd2, 0xd3,
	0x0e, 0xd2, 0xa1, 0xde, 0xd4, 0x8b, 0x45, 0x85, 0x1d, 0x04, 0x09, 0x78, 0xd0, 0xcb, 0xc8, 0xd6,
	0xac, 0x0b, 0x34, 0x7d, 0x4b, 0xd2, 0x4e, 0xfc, 0x5b, 0x85, 0x1e, 0x3c, 0xee, 0xaf, 0x90, 0x34,
	0xed, 0x98, 0xee, 0xf6, 0x3c, 0xef, 0xfb, 0xf4, 0xd7, 0x27, 0x09, 0x3a, 0x5e, 0x80, 0x10, 0x90,
	0x3d, 0x31, 0xa5, 0x68, 0xc2, 0xc2, 0x5c, 0x42, 0x01, 0xd8, 0x51, 0xef, 0x34, 0x59, 0xac, 0x68,
	0x61, 0xfc, 0xbc, 0x5c, 0x9e, 0x0c, 0x13, 0x48, 0xa0, 0x76, 0x13, 0xad, 0xcc, 0x22, 0x78, 0x44,
	0x36, 0x01, 0x10, 0xd3, 0x58, 0xe1, 0x1b, 0xf4, 0x4f, 0x02, 0x88, 0x19, 0x8f, 0x95, 0x6b, 0xf9,
	0x9d, 0x71, 0x3f, 0xf2, 0xbf, 0x2a, 0xcf, 0xac, 0xef, 0xd5, 0xa6, 0xf2, 0x1c, 0x69, 0x92, 0xe7,
	0x20, 0x78, 0xc1, 0x44, 0x5e, 0x7c, 0x10, 0xbb, 0x19, 0x69, 0xce, 0x8b, 0x62, 0xb2, 0xe1, 0x94,
	0x8a, 0xc9, 0xdf, 0x9c, 0x7a, 0x6d, 0x38, 0xa5, 0x49, 0xee, 0x72, 0x9a, 0x51, 0x70, 0xab, 0xfb,
	0xa4, 0x4c, 0x73, 0x2e, 0x74, 0x9f, 0x94, 0x6d, 0x39, 0x7f, 0xa3, 0x91, 0xe9, 0x93, 0x32, 0xc3,
	0xb1, 0xa5, 0x49, 0x92, 0x56, 0x04, 0x9f, 0x16, 0x1a, 0x3c, 0x48, 0x09, 0x92, 0x30, 0x95, 0x43,
	0xa6, 0x18, 0x76, 0x91, 0x2d, 0xcc, 0xcd, 0xb8, 0x96, 0x6f, 0x8d, 0xfb, 0xa4, 0xb5, 0xf8, 0x0e,
	0x39, 0x31, 0x5b, 0xb3, 0x14, 0x72, 0x26, 0x67, 0x6d, 0xe6, 0x8f, 0xce, 0x44, 0xc3, 0x4d, 0xe5,
	0x1d, 0x6d, 0x97, 0xcd, 0xcd, 0x92, 0xbd, 0x09, 0xc6, 0xa8, 0xcb, 0xb3, 0x25, 0xb8, 0x9d, 0x9a,
	0x5c, 0x6b, 0xfc, 0x8a, 0x0e, 0x79, 0xb6, 0xa6, 0x29, 0x8f, 0x67, 0x39, 0x95, 0x54, 0x28, 0xb7,
	0xeb, 0x77, 0xc6, 0xff, 0x2f, 0xbd, 0x70, 0xef, 0x49, 0xc2, 0xa9, 0x09, 0x3e, 0xeb, 0x5c, 0xe4,
	0x6c, 0x2a, 0x6f, 0xc0, 0x77, 0x26, 0x8a, 0xfc, 0xb4, 0xc1, 0x35, 0x3a, 0xd8, 0xfd, 0x42, 0xff,
	0x3e, 0xa3, 0xa2, 0x3d, 0x58, 0xad, 0xf1, 0x08, 0xf5, 0x24, 0xa3, 0x0a, 0x32, 0x73, 0x14, 0xd2,
	0xb8, 0xe8, 0xec, 0xed, 0x34, 0xe1, 0xc5, 0xaa, 0x9c, 0x87, 0x0b, 0x10, 0x93, 0xb6, 0xca, 0xa4,
	0xad, 0x32, 0xef, 0xd5, 0xea, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xdd, 0x60, 0xe1, 0x4e,
	0x02, 0x00, 0x00,
}
