// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: userMessage.proto

package protoc_gen_go

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserRoomsFilter int32

const (
	UserRoomsFilter_None   UserRoomsFilter = 0
	UserRoomsFilter_Online UserRoomsFilter = 1
	UserRoomsFilter_Unread UserRoomsFilter = 2
)

var UserRoomsFilter_name = map[int32]string{
	0: "None",
	1: "Online",
	2: "Unread",
}
var UserRoomsFilter_value = map[string]int32{
	"None":   0,
	"Online": 1,
	"Unread": 2,
}

func (x UserRoomsFilter) Enum() *UserRoomsFilter {
	p := new(UserRoomsFilter)
	*p = x
	return p
}
func (x UserRoomsFilter) String() string {
	return proto.EnumName(UserRoomsFilter_name, int32(x))
}
func (x *UserRoomsFilter) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserRoomsFilter_value, data, "UserRoomsFilter")
	if err != nil {
		return err
	}
	*x = UserRoomsFilter(value)
	return nil
}
func (UserRoomsFilter) EnumDescriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{0} }

type PublicProfileScope int32

const (
	PublicProfileScope_Self PublicProfileScope = 0
	PublicProfileScope_All  PublicProfileScope = 1
)

var PublicProfileScope_name = map[int32]string{
	0: "Self",
	1: "All",
}
var PublicProfileScope_value = map[string]int32{
	"Self": 0,
	"All":  1,
}

func (x PublicProfileScope) Enum() *PublicProfileScope {
	p := new(PublicProfileScope)
	*p = x
	return p
}
func (x PublicProfileScope) String() string {
	return proto.EnumName(PublicProfileScope_name, int32(x))
}
func (x *PublicProfileScope) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PublicProfileScope_value, data, "PublicProfileScope")
	if err != nil {
		return err
	}
	*x = PublicProfileScope(value)
	return nil
}
func (PublicProfileScope) EnumDescriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{1} }

type User struct {
	ID                    uint64             `protobuf:"varint,1,opt,name=id" json:"id" db:"id,primarykey,autoincrement"`
	UserID                string             `protobuf:"bytes,12,opt,name=user_id,json=userId" json:"user_id" db:"user_id,notnull"`
	Name                  string             `protobuf:"bytes,13,opt,name=name" json:"name" db:"name,notnull"`
	PictureURL            string             `protobuf:"bytes,14,opt,name=picture_url,json=pictureUrl" json:"picture_url" db:"picture_url"`
	InformationURL        string             `protobuf:"bytes,15,opt,name=information_url,json=informationUrl" json:"information_url" db:"information_url"`
	UnreadCount           uint64             `protobuf:"varint,16,opt,name=unread_count,json=unreadCount" json:"unread_count" db:"unread_count,notnull"`
	MetaData              []byte             `protobuf:"bytes,17,opt,name=meta_data,json=metaData" json:"meta_data,omitempty" db:"-"`
	PublicProfileScope    PublicProfileScope `protobuf:"varint,18,opt,name=public_profile_scope,json=publicProfileScope,enum=swagchat.protobuf.PublicProfileScope" json:"public_profile_scope" db:"public_profile_scope,notnull"`
	CanBlock              bool               `protobuf:"varint,19,opt,name=can_block,json=canBlock" json:"can_block" db:"can_block,notnull"`
	Lang                  string             `protobuf:"bytes,20,opt,name=lang" json:"lang" db:"lang"`
	AccessToken           string             `protobuf:"bytes,21,opt,name=access_token,json=accessToken" json:"access_token" db:"-"`
	LastAccessRoomID      string             `protobuf:"bytes,22,opt,name=last_access_room_id,json=lastAccessRoomId" json:"last_access_room_id" db:"last_access_room_id"`
	LastAccessedTimestamp int64              `protobuf:"varint,23,opt,name=last_accessed_timestamp,json=lastAccessedTimestamp" json:"last_accessed_timestamp" db:"last_accessed,notnull"`
	LastAccessed          string             `protobuf:"bytes,24,opt,name=last_accessed,json=lastAccessed" json:"last_accessed" db:"-"`
	CreatedTimestamp      int64              `protobuf:"varint,31,opt,name=created_timestamp,json=createdTimestamp" json:"created_timestamp" db:"created,notnull"`
	Created               string             `protobuf:"bytes,32,opt,name=created" json:"created" db:"-"`
	ModifiedTimestamp     int64              `protobuf:"varint,33,opt,name=modified_timestamp,json=modifiedTimestamp" json:"modified_timestamp" db:"modified,notnull"`
	Modified              string             `protobuf:"bytes,34,opt,name=modified" json:"modified" db:"-"`
	DeletedTimestamp      int64              `protobuf:"varint,35,opt,name=deleted_timestamp,json=deletedTimestamp" json:"deleted_timestamp" db:"deleted,notnull"`
	BlockUsers            []string           `protobuf:"bytes,40,rep,name=block_users,json=blockUsers" json:"block_users,omitempty" db:"-"`
	Devices               []*Device          `protobuf:"bytes,41,rep,name=devices" json:"devices,omitempty" db:"-"`
	Roles                 []int32            `protobuf:"varint,42,rep,name=roles" json:"roles,omitempty" db:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{0} }

func (m *User) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPictureURL() string {
	if m != nil {
		return m.PictureURL
	}
	return ""
}

func (m *User) GetInformationURL() string {
	if m != nil {
		return m.InformationURL
	}
	return ""
}

func (m *User) GetUnreadCount() uint64 {
	if m != nil {
		return m.UnreadCount
	}
	return 0
}

func (m *User) GetMetaData() []byte {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *User) GetPublicProfileScope() PublicProfileScope {
	if m != nil {
		return m.PublicProfileScope
	}
	return PublicProfileScope_Self
}

func (m *User) GetCanBlock() bool {
	if m != nil {
		return m.CanBlock
	}
	return false
}

func (m *User) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *User) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *User) GetLastAccessRoomID() string {
	if m != nil {
		return m.LastAccessRoomID
	}
	return ""
}

func (m *User) GetLastAccessedTimestamp() int64 {
	if m != nil {
		return m.LastAccessedTimestamp
	}
	return 0
}

func (m *User) GetLastAccessed() string {
	if m != nil {
		return m.LastAccessed
	}
	return ""
}

func (m *User) GetCreatedTimestamp() int64 {
	if m != nil {
		return m.CreatedTimestamp
	}
	return 0
}

func (m *User) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *User) GetModifiedTimestamp() int64 {
	if m != nil {
		return m.ModifiedTimestamp
	}
	return 0
}

func (m *User) GetModified() string {
	if m != nil {
		return m.Modified
	}
	return ""
}

func (m *User) GetDeletedTimestamp() int64 {
	if m != nil {
		return m.DeletedTimestamp
	}
	return 0
}

func (m *User) GetBlockUsers() []string {
	if m != nil {
		return m.BlockUsers
	}
	return nil
}

func (m *User) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *User) GetRoles() []int32 {
	if m != nil {
		return m.Roles
	}
	return nil
}

type MiniRoom struct {
	RoomID                      string      `protobuf:"bytes,1,opt,name=room_id,json=roomId" json:"room_id" db:"room_id"`
	UserID                      string      `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id" db:"user_id"`
	Name                        string      `protobuf:"bytes,3,opt,name=name" json:"name" db:"name"`
	PictureURL                  string      `protobuf:"bytes,4,opt,name=picture_url,json=pictureUrl" json:"picture_url" db:"picture_url"`
	InformationURL              string      `protobuf:"bytes,5,opt,name=information_url,json=informationUrl" json:"information_url" db:"information_url"`
	MetaData                    []byte      `protobuf:"bytes,6,opt,name=meta_data,json=metaData" json:"meta_data,omitempty" db:"meta_data"`
	Type                        RoomType    `protobuf:"varint,7,opt,name=type,enum=swagchat.protobuf.RoomType" json:"type" db:"type"`
	LastMessage                 string      `protobuf:"bytes,8,opt,name=last_message,json=lastMessage" json:"last_message" db:"last_message"`
	LastMessageUpdatedTimestamp int64       `protobuf:"varint,9,opt,name=last_message_updated_timestamp,json=lastMessageUpdatedTimestamp" json:"last_message_updated_timestamp" db:"last_message_updated"`
	LastMessageUpdated          string      `protobuf:"bytes,10,opt,name=last_message_updated,json=lastMessageUpdated" json:"last_message_updated" db:"-"`
	CanLeft                     bool        `protobuf:"varint,11,opt,name=can_left,json=canLeft" json:"can_left" db:"can_left"`
	CreatedTimestamp            int64       `protobuf:"varint,21,opt,name=created_timestamp,json=createdTimestamp" json:"created_timestamp" db:"created"`
	Created                     string      `protobuf:"bytes,22,opt,name=created" json:"created" db:"-"`
	ModifiedTimestamp           int64       `protobuf:"varint,23,opt,name=modified_timestamp,json=modifiedTimestamp" json:"modified_timestamp" db:"modified"`
	Modified                    string      `protobuf:"bytes,24,opt,name=modified" json:"modified" db:"-"`
	Users                       []*MiniUser `protobuf:"bytes,30,rep,name=users" json:"users,omitempty" db:"-"`
	RuUnreadCount               int64       `protobuf:"varint,31,opt,name=ru_unread_count,json=ruUnreadCount" json:"ru_unread_count" db:"ru_unread_count"`
}

func (m *MiniRoom) Reset()                    { *m = MiniRoom{} }
func (m *MiniRoom) String() string            { return proto.CompactTextString(m) }
func (*MiniRoom) ProtoMessage()               {}
func (*MiniRoom) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{1} }

func (m *MiniRoom) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *MiniRoom) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *MiniRoom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MiniRoom) GetPictureURL() string {
	if m != nil {
		return m.PictureURL
	}
	return ""
}

func (m *MiniRoom) GetInformationURL() string {
	if m != nil {
		return m.InformationURL
	}
	return ""
}

func (m *MiniRoom) GetMetaData() []byte {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *MiniRoom) GetType() RoomType {
	if m != nil {
		return m.Type
	}
	return RoomType_OneOnOneRoom
}

func (m *MiniRoom) GetLastMessage() string {
	if m != nil {
		return m.LastMessage
	}
	return ""
}

func (m *MiniRoom) GetLastMessageUpdatedTimestamp() int64 {
	if m != nil {
		return m.LastMessageUpdatedTimestamp
	}
	return 0
}

func (m *MiniRoom) GetLastMessageUpdated() string {
	if m != nil {
		return m.LastMessageUpdated
	}
	return ""
}

func (m *MiniRoom) GetCanLeft() bool {
	if m != nil {
		return m.CanLeft
	}
	return false
}

func (m *MiniRoom) GetCreatedTimestamp() int64 {
	if m != nil {
		return m.CreatedTimestamp
	}
	return 0
}

func (m *MiniRoom) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *MiniRoom) GetModifiedTimestamp() int64 {
	if m != nil {
		return m.ModifiedTimestamp
	}
	return 0
}

func (m *MiniRoom) GetModified() string {
	if m != nil {
		return m.Modified
	}
	return ""
}

func (m *MiniRoom) GetUsers() []*MiniUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *MiniRoom) GetRuUnreadCount() int64 {
	if m != nil {
		return m.RuUnreadCount
	}
	return 0
}

type CreateUserRequest struct {
	UserID             *string             `protobuf:"bytes,11,opt,name=user_id,json=userId" json:"userId"`
	Name               *string             `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
	PictureURL         *string             `protobuf:"bytes,13,opt,name=picture_url,json=pictureUrl" json:"pictureUrl"`
	InformationURL     *string             `protobuf:"bytes,14,opt,name=information_url,json=informationUrl" json:"informationUrl"`
	MetaData           []byte              `protobuf:"bytes,15,opt,name=meta_data,json=metaData" json:"metaData"`
	PublicProfileScope *PublicProfileScope `protobuf:"varint,16,opt,name=public_profile_scope,json=publicProfileScope,enum=swagchat.protobuf.PublicProfileScope" json:"publicProfileScope"`
	CanBlock           *bool               `protobuf:"varint,17,opt,name=can_block,json=canBlock" json:"canBlock"`
	Lang               *string             `protobuf:"bytes,18,opt,name=lang" json:"lang,omitempty"`
	BlockUsers         []string            `protobuf:"bytes,19,rep,name=block_users,json=blockUsers" json:"blockUsers"`
	Roles              []int32             `protobuf:"varint,21,rep,name=roles" json:"roles"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{2} }

func (m *CreateUserRequest) GetUserID() string {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return ""
}

func (m *CreateUserRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CreateUserRequest) GetPictureURL() string {
	if m != nil && m.PictureURL != nil {
		return *m.PictureURL
	}
	return ""
}

func (m *CreateUserRequest) GetInformationURL() string {
	if m != nil && m.InformationURL != nil {
		return *m.InformationURL
	}
	return ""
}

func (m *CreateUserRequest) GetMetaData() []byte {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *CreateUserRequest) GetPublicProfileScope() PublicProfileScope {
	if m != nil && m.PublicProfileScope != nil {
		return *m.PublicProfileScope
	}
	return PublicProfileScope_Self
}

func (m *CreateUserRequest) GetCanBlock() bool {
	if m != nil && m.CanBlock != nil {
		return *m.CanBlock
	}
	return false
}

func (m *CreateUserRequest) GetLang() string {
	if m != nil && m.Lang != nil {
		return *m.Lang
	}
	return ""
}

func (m *CreateUserRequest) GetBlockUsers() []string {
	if m != nil {
		return m.BlockUsers
	}
	return nil
}

func (m *CreateUserRequest) GetRoles() []int32 {
	if m != nil {
		return m.Roles
	}
	return nil
}

type RetrieveUsersRequest struct {
	Limit  int32        `protobuf:"varint,11,opt,name=limit" json:"limit"`
	Offset int32        `protobuf:"varint,12,opt,name=offset" json:"offset"`
	Orders []*OrderInfo `protobuf:"bytes,13,rep,name=orders" json:"orders,omitempty"`
}

func (m *RetrieveUsersRequest) Reset()                    { *m = RetrieveUsersRequest{} }
func (m *RetrieveUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveUsersRequest) ProtoMessage()               {}
func (*RetrieveUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{3} }

func (m *RetrieveUsersRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RetrieveUsersRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *RetrieveUsersRequest) GetOrders() []*OrderInfo {
	if m != nil {
		return m.Orders
	}
	return nil
}

type UsersResponse struct {
	Users    []*User      `protobuf:"bytes,1,rep,name=users" json:"users"`
	AllCount int64        `protobuf:"varint,2,opt,name=allCount" json:"allCount"`
	Limit    int32        `protobuf:"varint,3,opt,name=limit" json:"limit"`
	Offset   int32        `protobuf:"varint,4,opt,name=offset" json:"offset"`
	Orders   []*OrderInfo `protobuf:"bytes,5,rep,name=orders" json:"orders,omitempty"`
}

func (m *UsersResponse) Reset()                    { *m = UsersResponse{} }
func (m *UsersResponse) String() string            { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()               {}
func (*UsersResponse) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{4} }

func (m *UsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UsersResponse) GetAllCount() int64 {
	if m != nil {
		return m.AllCount
	}
	return 0
}

func (m *UsersResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *UsersResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *UsersResponse) GetOrders() []*OrderInfo {
	if m != nil {
		return m.Orders
	}
	return nil
}

type RetrieveUserRequest struct {
	UserID string `protobuf:"bytes,11,opt,name=user_id,json=userId" json:"userId"`
}

func (m *RetrieveUserRequest) Reset()                    { *m = RetrieveUserRequest{} }
func (m *RetrieveUserRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveUserRequest) ProtoMessage()               {}
func (*RetrieveUserRequest) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{5} }

func (m *RetrieveUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type UpdateUserRequest struct {
	UserID             string              `protobuf:"bytes,11,opt,name=user_id,json=userId" json:"userId"`
	Name               *string             `protobuf:"bytes,13,opt,name=name" json:"name,omitempty"`
	PictureURL         *string             `protobuf:"bytes,14,opt,name=picture_url,json=pictureUrl" json:"pictureUrl"`
	InformationURL     *string             `protobuf:"bytes,15,opt,name=information_url,json=informationUrl" json:"informationUrl"`
	UnreadCount        *uint64             `protobuf:"varint,16,opt,name=unread_count,json=unreadCount" json:"unreadCount"`
	MetaData           []byte              `protobuf:"bytes,17,opt,name=meta_data,json=metaData" json:"metaData"`
	PublicProfileScope *PublicProfileScope `protobuf:"varint,18,opt,name=public_profile_scope,json=publicProfileScope,enum=swagchat.protobuf.PublicProfileScope" json:"publicProfileScope"`
	CanBlock           *bool               `protobuf:"varint,19,opt,name=can_block,json=canBlock" json:"canBlock"`
	Lang               *string             `protobuf:"bytes,20,opt,name=lang" json:"lang,omitempty"`
	BlockUsers         []string            `protobuf:"bytes,21,rep,name=block_users,json=blockUsers" json:"blockUsers"`
	Roles              []int32             `protobuf:"varint,23,rep,name=roles" json:"roles"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{6} }

func (m *UpdateUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateUserRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UpdateUserRequest) GetPictureURL() string {
	if m != nil && m.PictureURL != nil {
		return *m.PictureURL
	}
	return ""
}

func (m *UpdateUserRequest) GetInformationURL() string {
	if m != nil && m.InformationURL != nil {
		return *m.InformationURL
	}
	return ""
}

func (m *UpdateUserRequest) GetUnreadCount() uint64 {
	if m != nil && m.UnreadCount != nil {
		return *m.UnreadCount
	}
	return 0
}

func (m *UpdateUserRequest) GetMetaData() []byte {
	if m != nil {
		return m.MetaData
	}
	return nil
}

func (m *UpdateUserRequest) GetPublicProfileScope() PublicProfileScope {
	if m != nil && m.PublicProfileScope != nil {
		return *m.PublicProfileScope
	}
	return PublicProfileScope_Self
}

func (m *UpdateUserRequest) GetCanBlock() bool {
	if m != nil && m.CanBlock != nil {
		return *m.CanBlock
	}
	return false
}

func (m *UpdateUserRequest) GetLang() string {
	if m != nil && m.Lang != nil {
		return *m.Lang
	}
	return ""
}

func (m *UpdateUserRequest) GetBlockUsers() []string {
	if m != nil {
		return m.BlockUsers
	}
	return nil
}

func (m *UpdateUserRequest) GetRoles() []int32 {
	if m != nil {
		return m.Roles
	}
	return nil
}

type DeleteUserRequest struct {
	UserID string `protobuf:"bytes,11,opt,name=user_id,json=userId" json:"userId"`
}

func (m *DeleteUserRequest) Reset()                    { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()               {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{7} }

func (m *DeleteUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type RetrieveUserRoomsRequest struct {
	Limit  int32           `protobuf:"varint,11,opt,name=limit" json:"limit"`
	Offset int32           `protobuf:"varint,12,opt,name=offset" json:"offset"`
	Orders []*OrderInfo    `protobuf:"bytes,13,rep,name=orders" json:"orders,omitempty"`
	UserID string          `protobuf:"bytes,14,opt,name=user_id,json=userId" json:"user_id"`
	Filter UserRoomsFilter `protobuf:"varint,15,opt,name=filter,enum=swagchat.protobuf.UserRoomsFilter" json:"filter"`
}

func (m *RetrieveUserRoomsRequest) Reset()         { *m = RetrieveUserRoomsRequest{} }
func (m *RetrieveUserRoomsRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveUserRoomsRequest) ProtoMessage()    {}
func (*RetrieveUserRoomsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorUserMessage, []int{8}
}

func (m *RetrieveUserRoomsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RetrieveUserRoomsRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *RetrieveUserRoomsRequest) GetOrders() []*OrderInfo {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *RetrieveUserRoomsRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *RetrieveUserRoomsRequest) GetFilter() UserRoomsFilter {
	if m != nil {
		return m.Filter
	}
	return UserRoomsFilter_None
}

type UserRoomsResponse struct {
	Rooms    []*MiniRoom  `protobuf:"bytes,1,rep,name=rooms" json:"rooms"`
	AllCount int64        `protobuf:"varint,2,opt,name=allCount" json:"allCount"`
	Limit    int32        `protobuf:"varint,3,opt,name=limit" json:"limit"`
	Offset   int32        `protobuf:"varint,4,opt,name=offset" json:"offset"`
	Orders   []*OrderInfo `protobuf:"bytes,5,rep,name=orders" json:"orders,omitempty"`
}

func (m *UserRoomsResponse) Reset()                    { *m = UserRoomsResponse{} }
func (m *UserRoomsResponse) String() string            { return proto.CompactTextString(m) }
func (*UserRoomsResponse) ProtoMessage()               {}
func (*UserRoomsResponse) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{9} }

func (m *UserRoomsResponse) GetRooms() []*MiniRoom {
	if m != nil {
		return m.Rooms
	}
	return nil
}

func (m *UserRoomsResponse) GetAllCount() int64 {
	if m != nil {
		return m.AllCount
	}
	return 0
}

func (m *UserRoomsResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *UserRoomsResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *UserRoomsResponse) GetOrders() []*OrderInfo {
	if m != nil {
		return m.Orders
	}
	return nil
}

type RetrieveContactsRequest struct {
	UserID string       `protobuf:"bytes,11,opt,name=user_id,json=userId" json:"userId" db:"user_id"`
	Limit  int32        `protobuf:"varint,12,opt,name=limit" json:"limit"`
	Offset int32        `protobuf:"varint,13,opt,name=offset" json:"offset"`
	Orders []*OrderInfo `protobuf:"bytes,14,rep,name=orders" json:"orders,omitempty"`
}

func (m *RetrieveContactsRequest) Reset()         { *m = RetrieveContactsRequest{} }
func (m *RetrieveContactsRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveContactsRequest) ProtoMessage()    {}
func (*RetrieveContactsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorUserMessage, []int{10}
}

func (m *RetrieveContactsRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *RetrieveContactsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RetrieveContactsRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *RetrieveContactsRequest) GetOrders() []*OrderInfo {
	if m != nil {
		return m.Orders
	}
	return nil
}

type RetrieveProfileRequest struct {
	UserID string `protobuf:"bytes,11,opt,name=user_id,json=userId" json:"userId"`
}

func (m *RetrieveProfileRequest) Reset()         { *m = RetrieveProfileRequest{} }
func (m *RetrieveProfileRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveProfileRequest) ProtoMessage()    {}
func (*RetrieveProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorUserMessage, []int{11}
}

func (m *RetrieveProfileRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type RetrieveRoleUsersRequest struct {
	RoleID *int32 `protobuf:"varint,11,opt,name=role_id,json=roleId" json:"roleId"`
}

func (m *RetrieveRoleUsersRequest) Reset()         { *m = RetrieveRoleUsersRequest{} }
func (m *RetrieveRoleUsersRequest) String() string { return proto.CompactTextString(m) }
func (*RetrieveRoleUsersRequest) ProtoMessage()    {}
func (*RetrieveRoleUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorUserMessage, []int{12}
}

func (m *RetrieveRoleUsersRequest) GetRoleID() int32 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

type RoleUsersResponse struct {
	Users   []*User  `protobuf:"bytes,11,rep,name=users" json:"users,omitempty"`
	UserIDs []string `protobuf:"bytes,12,rep,name=user_ids,json=userIds" json:"userIds,omitempty"`
}

func (m *RoleUsersResponse) Reset()                    { *m = RoleUsersResponse{} }
func (m *RoleUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*RoleUsersResponse) ProtoMessage()               {}
func (*RoleUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptorUserMessage, []int{13} }

func (m *RoleUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *RoleUsersResponse) GetUserIDs() []string {
	if m != nil {
		return m.UserIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "swagchat.protobuf.User")
	proto.RegisterType((*MiniRoom)(nil), "swagchat.protobuf.MiniRoom")
	proto.RegisterType((*CreateUserRequest)(nil), "swagchat.protobuf.CreateUserRequest")
	proto.RegisterType((*RetrieveUsersRequest)(nil), "swagchat.protobuf.RetrieveUsersRequest")
	proto.RegisterType((*UsersResponse)(nil), "swagchat.protobuf.UsersResponse")
	proto.RegisterType((*RetrieveUserRequest)(nil), "swagchat.protobuf.RetrieveUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "swagchat.protobuf.UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "swagchat.protobuf.DeleteUserRequest")
	proto.RegisterType((*RetrieveUserRoomsRequest)(nil), "swagchat.protobuf.RetrieveUserRoomsRequest")
	proto.RegisterType((*UserRoomsResponse)(nil), "swagchat.protobuf.UserRoomsResponse")
	proto.RegisterType((*RetrieveContactsRequest)(nil), "swagchat.protobuf.RetrieveContactsRequest")
	proto.RegisterType((*RetrieveProfileRequest)(nil), "swagchat.protobuf.RetrieveProfileRequest")
	proto.RegisterType((*RetrieveRoleUsersRequest)(nil), "swagchat.protobuf.RetrieveRoleUsersRequest")
	proto.RegisterType((*RoleUsersResponse)(nil), "swagchat.protobuf.RoleUsersResponse")
	proto.RegisterEnum("swagchat.protobuf.UserRoomsFilter", UserRoomsFilter_name, UserRoomsFilter_value)
	proto.RegisterEnum("swagchat.protobuf.PublicProfileScope", PublicProfileScope_name, PublicProfileScope_value)
}

func init() { proto.RegisterFile("userMessage.proto", fileDescriptorUserMessage) }

var fileDescriptorUserMessage = []byte{
	// 1722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0x4b, 0x6f, 0xdb, 0x56,
	0x16, 0x36, 0xad, 0x87, 0xe5, 0xab, 0x27, 0xaf, 0x64, 0x9b, 0x71, 0x3c, 0xa6, 0xc2, 0x99, 0x4c,
	0x64, 0x8f, 0x1f, 0x13, 0x27, 0x40, 0x06, 0x79, 0x4c, 0x63, 0xc5, 0x68, 0x61, 0xc0, 0x79, 0x94,
	0x89, 0x51, 0x20, 0x1b, 0x82, 0x26, 0xaf, 0x14, 0x22, 0x24, 0xaf, 0x4a, 0x52, 0x29, 0xbc, 0xed,
	0xba, 0xbf, 0xa0, 0xbb, 0x2e, 0xbb, 0xe9, 0x4f, 0xe8, 0xba, 0xcb, 0xa2, 0x3f, 0x80, 0x8b, 0x2c,
	0xba, 0x10, 0xd0, 0x8d, 0x81, 0xae, 0xba, 0x29, 0xee, 0x83, 0x22, 0x29, 0x51, 0x4e, 0x5c, 0x07,
	0x2d, 0xba, 0x92, 0x78, 0x5e, 0xf7, 0x9e, 0xc3, 0xf3, 0x7d, 0xe7, 0x10, 0x88, 0x43, 0x1f, 0x79,
	0x8f, 0x91, 0xef, 0xeb, 0x7d, 0xb4, 0x33, 0xf0, 0x70, 0x80, 0xa1, 0xe8, 0x7f, 0xa1, 0xf7, 0x8d,
	0x57, 0x7a, 0xc0, 0x9e, 0x4f, 0x86, 0xbd, 0xd5, 0x56, 0x1f, 0xf7, 0x31, 0x7d, 0xda, 0x25, 0xff,
	0x98, 0x62, 0xb5, 0x69, 0xa2, 0x37, 0x96, 0x81, 0x52, 0xde, 0xab, 0xa2, 0x87, 0xb1, 0x93, 0x16,
	0x35, 0x0d, 0xec, 0x38, 0xd8, 0x4d, 0x09, 0x95, 0x6f, 0xcb, 0x20, 0x7f, 0xec, 0x23, 0x0f, 0xfe,
	0x1f, 0xcc, 0x5b, 0xa6, 0x24, 0xb4, 0x85, 0x4e, 0xbe, 0xbb, 0xf3, 0x43, 0x28, 0xcf, 0xbd, 0x0d,
	0xe5, 0xf9, 0xc3, 0x83, 0xb3, 0x50, 0x6e, 0x9b, 0x27, 0x77, 0x15, 0xcb, 0xdc, 0x1a, 0x78, 0x96,
	0xa3, 0x7b, 0xa7, 0xaf, 0xd1, 0xe9, 0x96, 0x3e, 0x0c, 0xb0, 0xe5, 0x1a, 0x1e, 0x72, 0x90, 0x1b,
	0x28, 0xea, 0xbc, 0x65, 0xc2, 0x8f, 0xc0, 0x02, 0xc9, 0x41, 0xb3, 0x4c, 0xa9, 0xd2, 0x16, 0x3a,
	0x8b, 0xdd, 0x7f, 0xf3, 0x20, 0x45, 0x12, 0x9e, 0x06, 0x6a, 0x91, 0x40, 0xdc, 0x68, 0xcb, 0xc5,
	0x81, 0x3b, 0xb4, 0x6d, 0x45, 0x2d, 0x12, 0xc9, 0xa1, 0x09, 0xb7, 0x41, 0xde, 0xd5, 0x1d, 0x24,
	0x55, 0xa9, 0xf7, 0x15, 0xe2, 0x7d, 0x16, 0xca, 0x22, 0xf1, 0x21, 0xf2, 0xd8, 0x81, 0x9a, 0xc1,
	0x4f, 0x40, 0x79, 0x60, 0x19, 0xc1, 0xd0, 0x43, 0xda, 0xd0, 0xb3, 0xa5, 0x5a, 0xea, 0x4c, 0xf0,
	0x8c, 0xa9, 0x8e, 0xd5, 0xa3, 0xb3, 0x50, 0x6e, 0x90, 0x18, 0x09, 0x63, 0x45, 0x05, 0xfc, 0xe9,
	0xd8, 0xb3, 0xe1, 0x67, 0xa0, 0x6e, 0xb9, 0x3d, 0xec, 0x39, 0x7a, 0x60, 0x61, 0x97, 0x06, 0xab,
	0xd3, 0x60, 0x51, 0x15, 0x6a, 0x87, 0xb1, 0x9a, 0x05, 0xa4, 0x89, 0x4c, 0x38, 0x29, 0x6a, 0x2d,
	0x21, 0x21, 0x81, 0x0f, 0x40, 0x65, 0xe8, 0x7a, 0x48, 0x37, 0x35, 0x03, 0x0f, 0xdd, 0x40, 0x6a,
	0xd0, 0xda, 0x5e, 0xe3, 0x89, 0x5d, 0xa1, 0xc5, 0x48, 0xe8, 0xe3, 0x04, 0xcb, 0x4c, 0xfc, 0x88,
	0x48, 0xe1, 0x0d, 0xb0, 0xe8, 0xa0, 0x40, 0xd7, 0x4c, 0x3d, 0xd0, 0x25, 0xb1, 0x2d, 0x74, 0x2a,
	0x5d, 0x70, 0x16, 0xca, 0x45, 0xe2, 0xbe, 0xad, 0xa8, 0x25, 0xa2, 0x3c, 0xd0, 0x03, 0x1d, 0x7e,
	0x29, 0x80, 0xd6, 0x60, 0x78, 0x62, 0x5b, 0x86, 0x36, 0xf0, 0x70, 0xcf, 0xb2, 0x91, 0xe6, 0x1b,
	0x78, 0x80, 0x24, 0xd8, 0x16, 0x3a, 0xb5, 0xbd, 0xeb, 0x3b, 0x53, 0xfd, 0xb4, 0xf3, 0x8c, 0x9a,
	0x3f, 0x63, 0xd6, 0xcf, 0x89, 0x71, 0x77, 0x83, 0x5f, 0xef, 0x1a, 0xad, 0x59, 0x46, 0xb8, 0xf8,
	0x9a, 0x70, 0x30, 0xe5, 0x0e, 0xef, 0x81, 0x45, 0x43, 0x77, 0xb5, 0x13, 0x1b, 0x1b, 0xaf, 0xa5,
	0x66, 0x5b, 0xe8, 0x94, 0xba, 0xeb, 0x3c, 0xe2, 0x32, 0x89, 0x38, 0x56, 0xc6, 0x61, 0x4a, 0x86,
	0xee, 0x76, 0x89, 0x08, 0x5e, 0x07, 0x79, 0x5b, 0x77, 0xfb, 0x52, 0x8b, 0x96, 0x5f, 0xe4, 0x7e,
	0x8b, 0xc4, 0x8f, 0xc8, 0x15, 0x95, 0xaa, 0xe1, 0x4d, 0x50, 0xd1, 0x0d, 0x03, 0xf9, 0xbe, 0x16,
	0xe0, 0xd7, 0xc8, 0x95, 0x96, 0xa8, 0x79, 0x8d, 0x9b, 0x47, 0x85, 0x29, 0x33, 0x9b, 0x17, 0xc4,
	0x04, 0x1a, 0xa0, 0x69, 0xeb, 0x7e, 0xa0, 0x71, 0x3f, 0x82, 0x0d, 0xd2, 0xa8, 0xcb, 0xd4, 0xf3,
	0x36, 0x7f, 0xcf, 0x8d, 0x23, 0xdd, 0x0f, 0xf6, 0xa9, 0x85, 0x8a, 0xb1, 0x43, 0x5b, 0x56, 0x62,
	0x87, 0x4f, 0xb9, 0x2a, 0x6a, 0xc3, 0x4e, 0x5b, 0x9b, 0xf0, 0x25, 0x58, 0x49, 0x58, 0x22, 0x53,
	0x0b, 0x2c, 0x07, 0xf9, 0x81, 0xee, 0x0c, 0xa4, 0x95, 0xb6, 0xd0, 0xc9, 0x75, 0x15, 0x7e, 0xc5,
	0xd5, 0x89, 0xa0, 0x28, 0x81, 0x86, 0xa5, 0x38, 0x2c, 0x32, 0x5f, 0x44, 0x01, 0xe0, 0x2d, 0x50,
	0x4d, 0x39, 0x48, 0x52, 0x66, 0xd2, 0x95, 0xa4, 0x37, 0x3c, 0x04, 0xa2, 0xe1, 0x21, 0x3d, 0x48,
	0x5d, 0x45, 0xa6, 0x57, 0x59, 0xe3, 0x8e, 0xb4, 0x93, 0xb9, 0x51, 0x7c, 0x89, 0x06, 0x97, 0xc4,
	0xe7, 0x77, 0xc0, 0x02, 0x97, 0x49, 0xed, 0xcc, 0x93, 0x23, 0x35, 0x3c, 0x02, 0xd0, 0xc1, 0xa6,
	0xd5, 0xb3, 0x52, 0xa7, 0x5e, 0xa3, 0xa7, 0xfe, 0x83, 0x3b, 0x2d, 0x11, 0xa7, 0xc8, 0x2a, 0x3e,
	0x56, 0x8c, 0x44, 0xf1, 0xb9, 0x9b, 0xa0, 0x14, 0x09, 0x25, 0x25, 0xf3, 0xe0, 0xb1, 0x9e, 0xa4,
	0x6b, 0x22, 0x1b, 0xa5, 0xd3, 0xfd, 0xe7, 0x74, 0xba, 0xdc, 0x28, 0x91, 0x2e, 0x97, 0xc4, 0xc7,
	0xfe, 0x07, 0x94, 0x69, 0x97, 0x6a, 0x84, 0x9b, 0x7c, 0xa9, 0xd3, 0xce, 0x75, 0x16, 0x53, 0xb0,
	0x03, 0x54, 0x4d, 0x98, 0xcd, 0x27, 0xcc, 0xc7, 0x18, 0xd8, 0x97, 0x36, 0xda, 0xb9, 0x4e, 0x79,
	0xef, 0x4a, 0x06, 0xd4, 0x0e, 0xa8, 0x45, 0x2a, 0x46, 0xe4, 0x05, 0xdb, 0xa0, 0xe0, 0x61, 0x1b,
	0xf9, 0xd2, 0x66, 0x3b, 0xd7, 0x29, 0xa4, 0x6c, 0x98, 0xe2, 0x6e, 0xfe, 0xc7, 0x6f, 0xe4, 0x39,
	0xe5, 0xfb, 0x12, 0x28, 0x3d, 0xb6, 0x5c, 0x8b, 0xf4, 0x1b, 0xbc, 0x03, 0x16, 0xa2, 0x36, 0x16,
	0x68, 0x61, 0xd6, 0x23, 0xbe, 0x1d, 0x37, 0x6f, 0x85, 0x04, 0x19, 0x37, 0x6c, 0xd1, 0x63, 0x6d,
	0x7a, 0x27, 0x26, 0xea, 0xf9, 0xb4, 0xe3, 0x98, 0xa8, 0x2b, 0x09, 0xa2, 0x8e, 0x09, 0xfa, 0x3a,
	0x27, 0xe8, 0xdc, 0x34, 0x3c, 0x89, 0x7c, 0x06, 0x31, 0xe7, 0x3f, 0x24, 0x31, 0x17, 0x3e, 0x08,
	0x31, 0xef, 0x26, 0x29, 0xb5, 0x48, 0x29, 0x15, 0x9e, 0x85, 0x72, 0x8d, 0x76, 0x65, 0xa4, 0x48,
	0x52, 0x6b, 0x17, 0xe4, 0x83, 0xd3, 0x01, 0x92, 0x16, 0x28, 0x93, 0x5e, 0xcd, 0x78, 0xbd, 0xa4,
	0xe6, 0x2f, 0x4e, 0x07, 0x28, 0x5d, 0x16, 0xe2, 0xa4, 0xa8, 0xd4, 0x17, 0xde, 0x07, 0x14, 0x9c,
	0x9a, 0xc3, 0xc6, 0xaf, 0x54, 0x9a, 0x1e, 0x73, 0x49, 0xbd, 0xa2, 0x96, 0xc9, 0x23, 0x1f, 0xd6,
	0xb0, 0x07, 0xd6, 0x93, 0x5a, 0x6d, 0x38, 0x30, 0x27, 0x70, 0xbd, 0x48, 0x1b, 0x3d, 0x35, 0x5d,
	0xb2, 0x3c, 0x14, 0xf5, 0x6a, 0x22, 0xee, 0x31, 0x13, 0xc6, 0x8d, 0xff, 0x10, 0xb4, 0xb2, 0xbc,
	0x24, 0x90, 0x89, 0x3d, 0x38, 0x1d, 0x0a, 0xfe, 0x17, 0x10, 0x42, 0xd7, 0x6c, 0xd4, 0x0b, 0xa4,
	0x32, 0x1d, 0x00, 0x4b, 0xdc, 0xab, 0x1a, 0x0d, 0x00, 0xa2, 0x23, 0x8c, 0xa1, 0xbb, 0x47, 0xa8,
	0x17, 0xc0, 0xfd, 0x2c, 0x9a, 0x5a, 0xa2, 0xe9, 0xb4, 0xb8, 0x6b, 0x25, 0x41, 0x53, 0xef, 0xa0,
	0xa7, 0xe5, 0xf3, 0xe9, 0xe9, 0x20, 0x93, 0x9e, 0x18, 0x3f, 0xa7, 0x2e, 0x1a, 0x59, 0xbd, 0x93,
	0x96, 0xa4, 0x77, 0xd0, 0xd2, 0x03, 0x50, 0x60, 0x2c, 0xb2, 0x4e, 0xc9, 0x21, 0xab, 0x7b, 0x08,
	0xa8, 0x09, 0xf8, 0xd2, 0xd0, 0xa7, 0x5e, 0xf0, 0x00, 0xd4, 0xbd, 0xa1, 0x96, 0x5a, 0x24, 0x32,
	0x28, 0x7c, 0xc2, 0x44, 0x51, 0xab, 0xde, 0xf0, 0x38, 0xde, 0x22, 0x38, 0x81, 0x7c, 0x97, 0x07,
	0xe2, 0x23, 0x5a, 0x08, 0x72, 0x9a, 0x8a, 0x3e, 0x1f, 0x22, 0x3f, 0x80, 0xdb, 0x31, 0x21, 0x94,
	0x69, 0x2e, 0xad, 0x98, 0x0c, 0x46, 0xa1, 0xcc, 0xe1, 0x3f, 0xa6, 0x01, 0xc8, 0x69, 0x80, 0x6e,
	0x79, 0x1c, 0xf3, 0x0f, 0xd2, 0x98, 0x67, 0x2b, 0xdc, 0x5a, 0x1a, 0xef, 0xa3, 0x50, 0x4e, 0xa0,
	0x3b, 0x85, 0xf4, 0xc7, 0xd3, 0x48, 0x67, 0xfb, 0xdc, 0xbf, 0xa6, 0x51, 0x3e, 0x0a, 0xe5, 0x09,
	0x3c, 0x4f, 0xe1, 0x7b, 0x23, 0x89, 0xef, 0x3a, 0xc5, 0x77, 0x65, 0x14, 0xca, 0x63, 0x3c, 0x27,
	0x90, 0x8d, 0x67, 0xec, 0x4c, 0x8d, 0x8b, 0xec, 0x4c, 0xcb, 0xa3, 0x50, 0xce, 0x58, 0x86, 0x32,
	0x17, 0xa4, 0x8d, 0xe4, 0x82, 0x24, 0x52, 0x7c, 0xd0, 0xbb, 0x45, 0x4b, 0x50, 0x62, 0x1d, 0x82,
	0x7c, 0x1d, 0x82, 0xac, 0xd0, 0x74, 0xf7, 0x79, 0x90, 0x1e, 0x4c, 0x4d, 0x3a, 0x98, 0x68, 0xa1,
	0xbb, 0xe3, 0x81, 0x44, 0x0a, 0x1d, 0x8f, 0xa7, 0xd4, 0xa8, 0xea, 0x44, 0x93, 0x66, 0x89, 0x4e,
	0x1a, 0xf8, 0x36, 0x94, 0x0b, 0x2a, 0x11, 0x8c, 0x42, 0x99, 0x69, 0xd2, 0x13, 0xe7, 0x2b, 0x01,
	0xb4, 0x54, 0x14, 0x78, 0x16, 0x7a, 0x83, 0x58, 0x34, 0xde, 0x33, 0xab, 0xa0, 0x60, 0x5b, 0x8e,
	0xc5, 0x20, 0x5e, 0xe8, 0xe6, 0x49, 0x2f, 0xaa, 0x4c, 0x04, 0xd7, 0x40, 0x11, 0xf7, 0x7a, 0x3e,
	0x0a, 0x68, 0x8b, 0x44, 0x4a, 0x2e, 0x83, 0xb7, 0x41, 0x11, 0x7b, 0x26, 0xb9, 0x7c, 0x95, 0xe2,
	0x61, 0x2d, 0xa3, 0xc6, 0x4f, 0x89, 0x01, 0x79, 0xed, 0x2a, 0xb7, 0xe5, 0xd7, 0xf9, 0x45, 0x00,
	0x55, 0x7e, 0x0d, 0x7f, 0x80, 0x5d, 0x1f, 0xc1, 0xff, 0x45, 0xe0, 0x12, 0x68, 0xb0, 0x95, 0x8c,
	0x60, 0x14, 0x58, 0x8b, 0x24, 0x41, 0x6a, 0x19, 0xe1, 0x6a, 0x0b, 0x94, 0x74, 0xdb, 0xa6, 0xe8,
	0xa0, 0x73, 0x30, 0xd7, 0x6d, 0x90, 0x7b, 0x92, 0x77, 0x11, 0xc9, 0xd5, 0xf1, 0xbf, 0x38, 0xdf,
	0xdc, 0x79, 0xf9, 0xe6, 0xcf, 0xcd, 0xb7, 0x70, 0xe1, 0x7c, 0x9f, 0x80, 0x66, 0xb2, 0xfa, 0x51,
	0xf1, 0x6f, 0x4e, 0x02, 0x56, 0x4a, 0x4f, 0xf0, 0x69, 0xd0, 0xf2, 0x78, 0x3f, 0xe7, 0x81, 0xc8,
	0x78, 0xfa, 0x72, 0xe1, 0xc6, 0x1c, 0x50, 0x9d, 0xcd, 0x01, 0xb5, 0xcb, 0x73, 0x40, 0xfd, 0x12,
	0x1c, 0xb0, 0x97, 0xf9, 0xf1, 0x55, 0x1f, 0x85, 0x72, 0xf2, 0xeb, 0x2a, 0xfd, 0xa9, 0xb5, 0x31,
	0xfd, 0xa9, 0x75, 0x51, 0xde, 0x80, 0x7f, 0x0a, 0x6f, 0x34, 0xdf, 0x8b, 0x37, 0x5a, 0xb3, 0x79,
	0x63, 0xe9, 0x8f, 0xf2, 0xc6, 0xca, 0xfb, 0xf1, 0xc6, 0x11, 0x10, 0x0f, 0xe8, 0x4e, 0xfd, 0x41,
	0xda, 0xf6, 0x37, 0x01, 0x48, 0x29, 0x1c, 0x60, 0xec, 0xfc, 0x55, 0x4c, 0x04, 0x6f, 0xc4, 0x59,
	0xd4, 0xe2, 0xc9, 0x1f, 0x67, 0x31, 0xc6, 0xc8, 0x43, 0x50, 0xec, 0x59, 0x76, 0x80, 0x3c, 0xda,
	0xc7, 0xb5, 0x3d, 0x65, 0x06, 0x37, 0xd1, 0x6c, 0x3e, 0xa6, 0x96, 0xd1, 0x05, 0x99, 0x1f, 0xcf,
	0xfe, 0x57, 0x01, 0x88, 0x89, 0xac, 0x39, 0xf1, 0xdd, 0x27, 0x6f, 0x04, 0x3b, 0x11, 0xf1, 0xcd,
	0xda, 0x2a, 0x88, 0x13, 0x23, 0x3f, 0x6a, 0xad, 0xb2, 0x9f, 0xbf, 0x05, 0xf9, 0xfd, 0x24, 0x80,
	0x95, 0xe8, 0xad, 0x3f, 0xc2, 0x6e, 0xa0, 0x1b, 0xc1, 0xf8, 0xa5, 0xef, 0x4f, 0xb6, 0x52, 0x67,
	0x56, 0x2b, 0xcd, 0xfc, 0x9a, 0x19, 0x27, 0x55, 0x39, 0x2f, 0xa9, 0xea, 0xb9, 0x49, 0xd5, 0x2e,
	0x9c, 0xd4, 0xa7, 0x60, 0x39, 0xca, 0x89, 0xc3, 0xfb, 0xd2, 0xe8, 0x78, 0x1a, 0x83, 0x83, 0x40,
	0x32, 0x35, 0xa6, 0xb7, 0xc9, 0x47, 0xa2, 0x8d, 0xa2, 0xa0, 0x05, 0xb6, 0xda, 0x11, 0x33, 0x16,
	0x90, 0x28, 0x49, 0x40, 0xf6, 0xcb, 0x03, 0x7e, 0x2d, 0x00, 0x31, 0x11, 0x89, 0x37, 0x5c, 0x37,
	0x9a, 0xb4, 0xe5, 0xf3, 0x27, 0x6d, 0x73, 0x14, 0xca, 0x75, 0x6a, 0xb9, 0x85, 0x1d, 0x2b, 0x40,
	0xce, 0x20, 0x38, 0x8d, 0x66, 0xee, 0x3d, 0x50, 0xe2, 0x39, 0xfa, 0x52, 0x85, 0x52, 0x50, 0xfb,
	0x6d, 0x28, 0x2f, 0xb0, 0x04, 0x09, 0x97, 0x88, 0x2c, 0xb3, 0xa4, 0xeb, 0x02, 0x17, 0xb1, 0xcb,
	0x6d, 0xde, 0x02, 0xf5, 0x09, 0xd0, 0xc0, 0x12, 0xc8, 0x3f, 0xc1, 0x2e, 0x6a, 0xcc, 0x41, 0x00,
	0x8a, 0x4f, 0x5d, 0xdb, 0x72, 0x51, 0x43, 0x20, 0xff, 0xd9, 0x02, 0xdc, 0x98, 0xdf, 0xbc, 0x01,
	0xe0, 0x34, 0xfd, 0x12, 0xbf, 0xe7, 0xc8, 0xee, 0x35, 0xe6, 0xe0, 0x02, 0xc8, 0xed, 0xdb, 0x76,
	0x43, 0xe8, 0x6e, 0xbd, 0xdc, 0xec, 0x5b, 0xc1, 0xab, 0xe1, 0xc9, 0x8e, 0x81, 0x9d, 0xdd, 0x28,
	0xc3, 0xdd, 0x28, 0x43, 0xf6, 0xc7, 0xd8, 0xee, 0x23, 0x77, 0xbb, 0x8f, 0x7f, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0xd9, 0x16, 0xe7, 0xb5, 0x15, 0x00, 0x00,
}
