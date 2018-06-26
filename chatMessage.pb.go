// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatMessage.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	chatMessage.proto
	chatService.proto

It has these top-level messages:
	Room
	Messages
	Message
	MessagePayload
	RoomIDs
	UserIDs
	RoleIDs
	UserRole
	PostUserRoleReq
	GetUserIDsOfUserRoleReq
	GetRoleIDsOfUserRoleReq
	RoomUser
	PostRoomUserReq
	DeleteRoomUserReq
	GetRoomIDsOfRoomUserReq
	GetUserIDsOfRoomUserReq
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Room struct {
	Workspace    string `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	RoomId       string `protobuf:"bytes,11,opt,name=roomId" json:"roomId,omitempty"`
	UserId       string `protobuf:"bytes,12,opt,name=userId" json:"userId,omitempty"`
	Name         string `protobuf:"bytes,13,opt,name=name" json:"name,omitempty"`
	PictureURL   string `protobuf:"bytes,14,opt,name=pictureURL" json:"pictureURL,omitempty"`
	Type         string `protobuf:"bytes,15,opt,name=type" json:"type,omitempty"`
	WebhookToken string `protobuf:"bytes,16,opt,name=WebhookToken" json:"WebhookToken,omitempty"`
	Created      int64  `protobuf:"varint,30,opt,name=created" json:"created,omitempty"`
	Modified     int64  `protobuf:"varint,31,opt,name=modified" json:"modified,omitempty"`
}

func (m *Room) Reset()                    { *m = Room{} }
func (m *Room) String() string            { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()               {}
func (*Room) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Room) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *Room) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *Room) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Room) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Room) GetPictureURL() string {
	if m != nil {
		return m.PictureURL
	}
	return ""
}

func (m *Room) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Room) GetWebhookToken() string {
	if m != nil {
		return m.WebhookToken
	}
	return ""
}

func (m *Room) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Room) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

type Messages struct {
	AllCount int64      `protobuf:"varint,1,opt,name=allCount" json:"allCount,omitempty"`
	Limit    int32      `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Offset   int32      `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Order    string     `protobuf:"bytes,4,opt,name=order" json:"order,omitempty"`
	Messages []*Message `protobuf:"bytes,5,rep,name=messages" json:"messages,omitempty"`
}

func (m *Messages) Reset()                    { *m = Messages{} }
func (m *Messages) String() string            { return proto.CompactTextString(m) }
func (*Messages) ProtoMessage()               {}
func (*Messages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Messages) GetAllCount() int64 {
	if m != nil {
		return m.AllCount
	}
	return 0
}

func (m *Messages) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Messages) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Messages) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *Messages) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Message struct {
	Workspace     string          `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	Endpoint      string          `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
	Authorization string          `protobuf:"bytes,3,opt,name=authorization" json:"authorization,omitempty"`
	UserIds       []string        `protobuf:"bytes,4,rep,name=userIds" json:"userIds,omitempty"`
	RoomId        string          `protobuf:"bytes,11,opt,name=roomId" json:"roomId,omitempty"`
	UserId        string          `protobuf:"bytes,12,opt,name=userId" json:"userId,omitempty"`
	Type          string          `protobuf:"bytes,13,opt,name=type" json:"type,omitempty"`
	EventName     string          `protobuf:"bytes,14,opt,name=eventName" json:"eventName,omitempty"`
	Payload       *MessagePayload `protobuf:"bytes,15,opt,name=payload" json:"payload,omitempty"`
	Role          int32           `protobuf:"varint,16,opt,name=role" json:"role,omitempty"`
	WebhookToken  string          `protobuf:"bytes,17,opt,name=WebhookToken" json:"WebhookToken,omitempty"`
	Created       int64           `protobuf:"varint,30,opt,name=created" json:"created,omitempty"`
	Modified      int64           `protobuf:"varint,31,opt,name=modified" json:"modified,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Message) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *Message) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Message) GetAuthorization() string {
	if m != nil {
		return m.Authorization
	}
	return ""
}

func (m *Message) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *Message) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *Message) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Message) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Message) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *Message) GetPayload() *MessagePayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *Message) GetWebhookToken() string {
	if m != nil {
		return m.WebhookToken
	}
	return ""
}

func (m *Message) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Message) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

type MessagePayload struct {
	// text
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// image
	Mime         string `protobuf:"bytes,11,opt,name=mime" json:"mime,omitempty"`
	Filename     string `protobuf:"bytes,12,opt,name=filename" json:"filename,omitempty"`
	SourceUrl    string `protobuf:"bytes,13,opt,name=sourceUrl" json:"sourceUrl,omitempty"`
	ThumbnailUrl string `protobuf:"bytes,14,opt,name=thumbnailUrl" json:"thumbnailUrl,omitempty"`
	Width        int32  `protobuf:"varint,15,opt,name=width" json:"width,omitempty"`
	Height       int32  `protobuf:"varint,16,opt,name=height" json:"height,omitempty"`
}

func (m *MessagePayload) Reset()                    { *m = MessagePayload{} }
func (m *MessagePayload) String() string            { return proto.CompactTextString(m) }
func (*MessagePayload) ProtoMessage()               {}
func (*MessagePayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MessagePayload) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *MessagePayload) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

func (m *MessagePayload) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *MessagePayload) GetSourceUrl() string {
	if m != nil {
		return m.SourceUrl
	}
	return ""
}

func (m *MessagePayload) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

func (m *MessagePayload) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *MessagePayload) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type RoomIDs struct {
	RoomIDs []string `protobuf:"bytes,1,rep,name=roomIDs" json:"roomIDs,omitempty"`
}

func (m *RoomIDs) Reset()                    { *m = RoomIDs{} }
func (m *RoomIDs) String() string            { return proto.CompactTextString(m) }
func (*RoomIDs) ProtoMessage()               {}
func (*RoomIDs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RoomIDs) GetRoomIDs() []string {
	if m != nil {
		return m.RoomIDs
	}
	return nil
}

type UserIDs struct {
	UserIDs []string `protobuf:"bytes,1,rep,name=userIDs" json:"userIDs,omitempty"`
}

func (m *UserIDs) Reset()                    { *m = UserIDs{} }
func (m *UserIDs) String() string            { return proto.CompactTextString(m) }
func (*UserIDs) ProtoMessage()               {}
func (*UserIDs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserIDs) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type RoleIDs struct {
	RoleIDs []int32 `protobuf:"varint,1,rep,packed,name=roleIDs" json:"roleIDs,omitempty"`
}

func (m *RoleIDs) Reset()                    { *m = RoleIDs{} }
func (m *RoleIDs) String() string            { return proto.CompactTextString(m) }
func (*RoleIDs) ProtoMessage()               {}
func (*RoleIDs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RoleIDs) GetRoleIDs() []int32 {
	if m != nil {
		return m.RoleIDs
	}
	return nil
}

type UserRole struct {
	// @inject_tag: db:"user_id"
	UserID string `protobuf:"bytes,11,opt,name=userID" json:"userID,omitempty"`
	// @inject_tag: db:"role_id"
	RoleID int32 `protobuf:"varint,12,opt,name=roleID" json:"roleID,omitempty"`
}

func (m *UserRole) Reset()                    { *m = UserRole{} }
func (m *UserRole) String() string            { return proto.CompactTextString(m) }
func (*UserRole) ProtoMessage()               {}
func (*UserRole) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserRole) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserRole) GetRoleID() int32 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

type PostUserRoleReq struct {
	Workspace string    `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	UserRole  *UserRole `protobuf:"bytes,11,opt,name=userRole" json:"userRole,omitempty"`
}

func (m *PostUserRoleReq) Reset()                    { *m = PostUserRoleReq{} }
func (m *PostUserRoleReq) String() string            { return proto.CompactTextString(m) }
func (*PostUserRoleReq) ProtoMessage()               {}
func (*PostUserRoleReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *PostUserRoleReq) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *PostUserRoleReq) GetUserRole() *UserRole {
	if m != nil {
		return m.UserRole
	}
	return nil
}

type GetUserIDsOfUserRoleReq struct {
	Workspace string `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	RoleID    int32  `protobuf:"varint,12,opt,name=roleID" json:"roleID,omitempty"`
}

func (m *GetUserIDsOfUserRoleReq) Reset()                    { *m = GetUserIDsOfUserRoleReq{} }
func (m *GetUserIDsOfUserRoleReq) String() string            { return proto.CompactTextString(m) }
func (*GetUserIDsOfUserRoleReq) ProtoMessage()               {}
func (*GetUserIDsOfUserRoleReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetUserIDsOfUserRoleReq) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *GetUserIDsOfUserRoleReq) GetRoleID() int32 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

type GetRoleIDsOfUserRoleReq struct {
	Workspace string `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	UserID    string `protobuf:"bytes,11,opt,name=userID" json:"userID,omitempty"`
}

func (m *GetRoleIDsOfUserRoleReq) Reset()                    { *m = GetRoleIDsOfUserRoleReq{} }
func (m *GetRoleIDsOfUserRoleReq) String() string            { return proto.CompactTextString(m) }
func (*GetRoleIDsOfUserRoleReq) ProtoMessage()               {}
func (*GetRoleIDsOfUserRoleReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetRoleIDsOfUserRoleReq) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *GetRoleIDsOfUserRoleReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type RoomUser struct {
	// @inject_tag: db:"room_id"
	RoomID string `protobuf:"bytes,11,opt,name=roomID" json:"roomID,omitempty"`
	// @inject_tag: db:"user_id"
	UserID string `protobuf:"bytes,12,opt,name=userID" json:"userID,omitempty"`
	// @inject_tag: db:"unread_count"
	UnreadCount int32 `protobuf:"varint,13,opt,name=unreadCount" json:"unreadCount,omitempty"`
	// @inject_tag: db:"display"
	Display bool `protobuf:"varint,14,opt,name=display" json:"display,omitempty"`
}

func (m *RoomUser) Reset()                    { *m = RoomUser{} }
func (m *RoomUser) String() string            { return proto.CompactTextString(m) }
func (*RoomUser) ProtoMessage()               {}
func (*RoomUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RoomUser) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *RoomUser) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *RoomUser) GetUnreadCount() int32 {
	if m != nil {
		return m.UnreadCount
	}
	return 0
}

func (m *RoomUser) GetDisplay() bool {
	if m != nil {
		return m.Display
	}
	return false
}

type PostRoomUserReq struct {
	Workspace string   `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	RoomID    string   `protobuf:"bytes,11,opt,name=roomID" json:"roomID,omitempty"`
	UserIDs   []string `protobuf:"bytes,12,rep,name=userIDs" json:"userIDs,omitempty"`
	Display   bool     `protobuf:"varint,13,opt,name=display" json:"display,omitempty"`
}

func (m *PostRoomUserReq) Reset()                    { *m = PostRoomUserReq{} }
func (m *PostRoomUserReq) String() string            { return proto.CompactTextString(m) }
func (*PostRoomUserReq) ProtoMessage()               {}
func (*PostRoomUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PostRoomUserReq) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *PostRoomUserReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *PostRoomUserReq) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

func (m *PostRoomUserReq) GetDisplay() bool {
	if m != nil {
		return m.Display
	}
	return false
}

type DeleteRoomUserReq struct {
	Workspace string   `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	RoomID    string   `protobuf:"bytes,11,opt,name=roomID" json:"roomID,omitempty"`
	UserIDs   []string `protobuf:"bytes,12,rep,name=userIDs" json:"userIDs,omitempty"`
}

func (m *DeleteRoomUserReq) Reset()                    { *m = DeleteRoomUserReq{} }
func (m *DeleteRoomUserReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteRoomUserReq) ProtoMessage()               {}
func (*DeleteRoomUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *DeleteRoomUserReq) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *DeleteRoomUserReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *DeleteRoomUserReq) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

type GetRoomIDsOfRoomUserReq struct {
	Workspace string `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	UserID    string `protobuf:"bytes,11,opt,name=userID" json:"userID,omitempty"`
}

func (m *GetRoomIDsOfRoomUserReq) Reset()                    { *m = GetRoomIDsOfRoomUserReq{} }
func (m *GetRoomIDsOfRoomUserReq) String() string            { return proto.CompactTextString(m) }
func (*GetRoomIDsOfRoomUserReq) ProtoMessage()               {}
func (*GetRoomIDsOfRoomUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetRoomIDsOfRoomUserReq) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *GetRoomIDsOfRoomUserReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type GetUserIDsOfRoomUserReq struct {
	Workspace string `protobuf:"bytes,1,opt,name=workspace" json:"workspace,omitempty"`
	RoomID    string `protobuf:"bytes,11,opt,name=roomID" json:"roomID,omitempty"`
}

func (m *GetUserIDsOfRoomUserReq) Reset()                    { *m = GetUserIDsOfRoomUserReq{} }
func (m *GetUserIDsOfRoomUserReq) String() string            { return proto.CompactTextString(m) }
func (*GetUserIDsOfRoomUserReq) ProtoMessage()               {}
func (*GetUserIDsOfRoomUserReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetUserIDsOfRoomUserReq) GetWorkspace() string {
	if m != nil {
		return m.Workspace
	}
	return ""
}

func (m *GetUserIDsOfRoomUserReq) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func init() {
	proto.RegisterType((*Room)(nil), "swagchat.protobuf.Room")
	proto.RegisterType((*Messages)(nil), "swagchat.protobuf.Messages")
	proto.RegisterType((*Message)(nil), "swagchat.protobuf.Message")
	proto.RegisterType((*MessagePayload)(nil), "swagchat.protobuf.MessagePayload")
	proto.RegisterType((*RoomIDs)(nil), "swagchat.protobuf.RoomIDs")
	proto.RegisterType((*UserIDs)(nil), "swagchat.protobuf.UserIDs")
	proto.RegisterType((*RoleIDs)(nil), "swagchat.protobuf.RoleIDs")
	proto.RegisterType((*UserRole)(nil), "swagchat.protobuf.UserRole")
	proto.RegisterType((*PostUserRoleReq)(nil), "swagchat.protobuf.PostUserRoleReq")
	proto.RegisterType((*GetUserIDsOfUserRoleReq)(nil), "swagchat.protobuf.GetUserIDsOfUserRoleReq")
	proto.RegisterType((*GetRoleIDsOfUserRoleReq)(nil), "swagchat.protobuf.GetRoleIDsOfUserRoleReq")
	proto.RegisterType((*RoomUser)(nil), "swagchat.protobuf.RoomUser")
	proto.RegisterType((*PostRoomUserReq)(nil), "swagchat.protobuf.PostRoomUserReq")
	proto.RegisterType((*DeleteRoomUserReq)(nil), "swagchat.protobuf.DeleteRoomUserReq")
	proto.RegisterType((*GetRoomIDsOfRoomUserReq)(nil), "swagchat.protobuf.GetRoomIDsOfRoomUserReq")
	proto.RegisterType((*GetUserIDsOfRoomUserReq)(nil), "swagchat.protobuf.GetUserIDsOfRoomUserReq")
}

func init() { proto.RegisterFile("chatMessage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x96, 0x4f, 0xe2, 0xc6, 0xd9, 0x24, 0xed, 0x89, 0x75, 0x74, 0x58, 0x95, 0xaa, 0x04, 0xc3,
	0x45, 0xae, 0x52, 0xa9, 0x48, 0x20, 0xc1, 0x1d, 0x44, 0x42, 0x95, 0x80, 0x56, 0x2b, 0x2a, 0x24,
	0xee, 0x36, 0xf6, 0x38, 0x5e, 0xd5, 0xf6, 0x9a, 0xf5, 0xba, 0xa5, 0xf0, 0x12, 0x3c, 0x03, 0xcf,
	0xc2, 0x33, 0x71, 0x8d, 0xf6, 0xc7, 0x6e, 0x4c, 0x53, 0xb5, 0xe5, 0xe7, 0x6e, 0xbe, 0xf1, 0xec,
	0xcc, 0xec, 0xf7, 0x7d, 0x6b, 0x34, 0x0e, 0x13, 0x2a, 0x5f, 0x43, 0x59, 0xd2, 0x25, 0xcc, 0x0a,
	0xc1, 0x25, 0xf7, 0xc7, 0xe5, 0x19, 0x5d, 0xaa, 0xb4, 0xc1, 0x8b, 0x2a, 0x0e, 0xbe, 0x3b, 0xa8,
	0x4b, 0x38, 0xcf, 0xfc, 0x1d, 0xd4, 0x3f, 0xe3, 0xe2, 0xa4, 0x2c, 0x68, 0x08, 0xd8, 0x99, 0x38,
	0xd3, 0x3e, 0xb9, 0x48, 0xf8, 0xff, 0xa3, 0x0d, 0xc1, 0x79, 0x76, 0x10, 0xe1, 0x81, 0xfe, 0x64,
	0x91, 0xca, 0x57, 0x25, 0x88, 0x83, 0x08, 0x0f, 0x4d, 0xde, 0x20, 0xdf, 0x47, 0xdd, 0x9c, 0x66,
	0x80, 0x47, 0x3a, 0xab, 0x63, 0x7f, 0x17, 0xa1, 0x82, 0x85, 0xb2, 0x12, 0x70, 0x4c, 0x5e, 0xe1,
	0x4d, 0xfd, 0x65, 0x25, 0xa3, 0xce, 0xc8, 0xf3, 0x02, 0xf0, 0x96, 0x39, 0xa3, 0x62, 0x3f, 0x40,
	0xc3, 0x77, 0xb0, 0x48, 0x38, 0x3f, 0x79, 0xcb, 0x4f, 0x20, 0xc7, 0xff, 0xea, 0x6f, 0xad, 0x9c,
	0x8f, 0x51, 0x2f, 0x14, 0x40, 0x25, 0x44, 0x78, 0x77, 0xe2, 0x4c, 0x3b, 0xa4, 0x86, 0xfe, 0x36,
	0xf2, 0x32, 0x1e, 0xb1, 0x98, 0x41, 0x84, 0xef, 0xe9, 0x4f, 0x0d, 0x0e, 0xbe, 0x3a, 0xc8, 0xb3,
	0xec, 0x94, 0xaa, 0x90, 0xa6, 0xe9, 0x0b, 0x5e, 0xe5, 0x52, 0xdf, 0xbd, 0x43, 0x1a, 0xec, 0xff,
	0x87, 0xdc, 0x94, 0x65, 0x4c, 0xe2, 0x7f, 0x26, 0xce, 0xd4, 0x25, 0x06, 0xa8, 0x8b, 0xf3, 0x38,
	0x2e, 0x41, 0xe2, 0x8e, 0x4e, 0x5b, 0xa4, 0xaa, 0xb9, 0x88, 0x40, 0xe0, 0xae, 0xde, 0xd4, 0x00,
	0xff, 0x31, 0xf2, 0x32, 0x3b, 0x0b, 0xbb, 0x93, 0xce, 0x74, 0xb0, 0xbf, 0x3d, 0xbb, 0xa4, 0xc5,
	0xcc, 0xae, 0x43, 0x9a, 0xda, 0xe0, 0x4b, 0x07, 0xf5, 0x6c, 0xf6, 0x1a, 0x81, 0xb6, 0x91, 0x07,
	0x79, 0x54, 0x70, 0x96, 0x9b, 0x45, 0xfb, 0xa4, 0xc1, 0xfe, 0x43, 0x34, 0xa2, 0x95, 0x4c, 0xb8,
	0x60, 0x9f, 0xa8, 0x64, 0x3c, 0xd7, 0x2b, 0xf7, 0x49, 0x3b, 0xa9, 0x68, 0x34, 0xe2, 0x95, 0xb8,
	0x3b, 0xe9, 0x4c, 0xfb, 0xa4, 0x86, 0xbf, 0x22, 0xbe, 0x16, 0x72, 0xb4, 0x22, 0xe4, 0x0e, 0xea,
	0xc3, 0x29, 0xe4, 0xf2, 0x8d, 0x72, 0x85, 0xd1, 0xfe, 0x22, 0xe1, 0x3f, 0x43, 0xbd, 0x82, 0x9e,
	0xa7, 0x9c, 0x46, 0x5a, 0xfd, 0xc1, 0xfe, 0xfd, 0xab, 0xe9, 0x39, 0x32, 0x85, 0xa4, 0x3e, 0xa1,
	0xc6, 0x09, 0x9e, 0x82, 0xf6, 0x86, 0x4b, 0x74, 0x7c, 0xc9, 0x37, 0xe3, 0x3f, 0xe6, 0x9b, 0x6f,
	0x0e, 0xda, 0x6c, 0x6f, 0xa2, 0xef, 0x0b, 0x1f, 0xa5, 0x15, 0x45, 0xc7, 0x2a, 0x97, 0xb1, 0x0c,
	0x2c, 0x63, 0x3a, 0x56, 0x6d, 0x63, 0x96, 0x82, 0x7e, 0x18, 0x86, 0xb1, 0x06, 0x2b, 0x7e, 0x4a,
	0x5e, 0x89, 0x10, 0x8e, 0x45, 0x6a, 0x89, 0xbb, 0x48, 0xa8, 0xeb, 0xc8, 0xa4, 0xca, 0x16, 0x39,
	0x65, 0xa9, 0x2a, 0x30, 0x04, 0xb6, 0x72, 0xca, 0x79, 0x67, 0x2c, 0x92, 0x89, 0x66, 0xd0, 0x25,
	0x06, 0x28, 0x8d, 0x12, 0x60, 0xcb, 0x44, 0x5a, 0x7a, 0x2c, 0x0a, 0x1e, 0xa0, 0x9e, 0x7a, 0xf6,
	0x07, 0xf3, 0x52, 0xf1, 0x20, 0x4c, 0x88, 0x1d, 0x23, 0xbc, 0x85, 0xaa, 0xe8, 0x58, 0x49, 0x6a,
	0x8a, 0x2a, 0x13, 0xd6, 0x45, 0x16, 0x9a, 0x4e, 0x29, 0x34, 0x9d, 0x74, 0xa8, 0x8b, 0x5c, 0x52,
	0xc3, 0xe0, 0x29, 0xf2, 0x54, 0x27, 0x55, 0xd8, 0xd8, 0x66, 0x5e, 0xdb, 0xc9, 0x20, 0x63, 0x33,
	0x55, 0xae, 0xc9, 0x71, 0x89, 0x45, 0x41, 0x82, 0xb6, 0x8e, 0x78, 0x29, 0xeb, 0xf3, 0x04, 0x3e,
	0x5c, 0xf3, 0x16, 0x9e, 0x20, 0xaf, 0xb2, 0xc5, 0x7a, 0xc4, 0x60, 0xff, 0xee, 0x1a, 0x3b, 0x35,
	0xfd, 0x9a, 0xe2, 0xe0, 0x10, 0xdd, 0x79, 0x09, 0xd2, 0x5e, 0xf9, 0x30, 0xbe, 0xf9, 0xc4, 0xab,
	0x56, 0x37, 0x0d, 0x2d, 0x3d, 0xb7, 0x6c, 0xb8, 0x8e, 0xa3, 0xe0, 0x14, 0x79, 0x4a, 0x36, 0xd5,
	0xa8, 0x79, 0x96, 0xf3, 0xd6, 0xb3, 0x9c, 0xaf, 0x9c, 0x1d, 0xb6, 0xf8, 0x9d, 0xa0, 0x41, 0x95,
	0x0b, 0xa0, 0x91, 0xf9, 0xcf, 0x8d, 0xf4, 0xa6, 0xab, 0x29, 0xa5, 0x5f, 0xc4, 0xca, 0x22, 0xa5,
	0xe7, 0xda, 0x61, 0x1e, 0xa9, 0x61, 0xf0, 0xd9, 0x68, 0x50, 0xcf, 0xbe, 0x21, 0x23, 0x6b, 0x96,
	0x5b, 0xf1, 0xd1, 0xb0, 0xe5, 0xa3, 0xd5, 0xe1, 0xa3, 0xf6, 0xf0, 0x10, 0x8d, 0xe7, 0x90, 0x82,
	0x84, 0xbf, 0x38, 0xbe, 0x91, 0x4a, 0x3b, 0xff, 0x30, 0xbe, 0xd5, 0xa8, 0xb5, 0x52, 0xfd, 0x64,
	0xa6, 0xdf, 0xde, 0xfd, 0xf9, 0xee, 0xfb, 0x9d, 0x25, 0x93, 0x49, 0xb5, 0x98, 0x85, 0x3c, 0xdb,
	0xab, 0x0d, 0xbd, 0x57, 0x1b, 0x7a, 0xb1, 0xa1, 0xa3, 0x47, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x1f, 0x79, 0x92, 0x54, 0xf9, 0x07, 0x00, 0x00,
}
