// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetProfileRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProfileRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AssignRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleName      string                 `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignRoleRequest) Reset() {
	*x = AssignRoleRequest{}
	mi := &file_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRoleRequest) ProtoMessage() {}

func (x *AssignRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRoleRequest.ProtoReflect.Descriptor instead.
func (*AssignRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *AssignRoleRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AssignRoleRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type ListRolesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRolesRequest) Reset() {
	*x = ListRolesRequest{}
	mi := &file_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesRequest) ProtoMessage() {}

func (x *ListRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesRequest.ProtoReflect.Descriptor instead.
func (*ListRolesRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *ListRolesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SearchUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchUserRequest) Reset() {
	*x = SearchUserRequest{}
	mi := &file_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserRequest) ProtoMessage() {}

func (x *SearchUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserRequest.ProtoReflect.Descriptor instead.
func (*SearchUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUserRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type RemoveRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RoleName      string                 `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRoleRequest) Reset() {
	*x = RemoveRoleRequest{}
	mi := &file_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoleRequest) ProtoMessage() {}

func (x *RemoveRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoleRequest.ProtoReflect.Descriptor instead.
func (*RemoveRoleRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveRoleRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveRoleRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Profile       *UserProfile           `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	mi := &file_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *AuthResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthResponse) GetProfile() *UserProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type AssignRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignRoleResponse) Reset() {
	*x = AssignRoleResponse{}
	mi := &file_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRoleResponse) ProtoMessage() {}

func (x *AssignRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRoleResponse.ProtoReflect.Descriptor instead.
func (*AssignRoleResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *AssignRoleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListRolesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Roles         []string               `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRolesResponse) Reset() {
	*x = ListRolesResponse{}
	mi := &file_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesResponse) ProtoMessage() {}

func (x *ListRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRolesResponse.ProtoReflect.Descriptor instead.
func (*ListRolesResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *ListRolesResponse) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	mi := &file_user_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RemoveRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRoleResponse) Reset() {
	*x = RemoveRoleResponse{}
	mi := &file_user_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoleResponse) ProtoMessage() {}

func (x *RemoveRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoleResponse.ProtoReflect.Descriptor instead.
func (*RemoveRoleResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{13}
}

func (x *RemoveRoleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UserList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserProfile         `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserList) Reset() {
	*x = UserList{}
	mi := &file_user_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{14}
}

func (x *UserList) GetUsers() []*UserProfile {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Roles         []string               `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_user_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{15}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserProfile) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type LogoutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_user_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{16}
}

func (x *LogoutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type VerifyEmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailRequest) Reset() {
	*x = VerifyEmailRequest{}
	mi := &file_user_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailRequest) ProtoMessage() {}

func (x *VerifyEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailRequest.ProtoReflect.Descriptor instead.
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{17}
}

func (x *VerifyEmailRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyEmailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyEmailResponse) Reset() {
	*x = VerifyEmailResponse{}
	mi := &file_user_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailResponse) ProtoMessage() {}

func (x *VerifyEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailResponse.ProtoReflect.Descriptor instead.
func (*VerifyEmailResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{18}
}

func (x *VerifyEmailResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x12\x04user\x1a\x1bgoogle/protobuf/empty.proto\"W\n" +
	"\x0fRegisterRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\",\n" +
	"\x11GetProfileRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"F\n" +
	"\x14UpdateProfileRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"I\n" +
	"\x11AssignRoleRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\trole_name\x18\x02 \x01(\tR\broleName\"+\n" +
	"\x10ListRolesRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\")\n" +
	"\x11SearchUserRequest\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\",\n" +
	"\x11DeleteUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"I\n" +
	"\x11RemoveRoleRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\trole_name\x18\x02 \x01(\tR\broleName\"Q\n" +
	"\fAuthResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12+\n" +
	"\aprofile\x18\x02 \x01(\v2\x11.user.UserProfileR\aprofile\".\n" +
	"\x12AssignRoleResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\")\n" +
	"\x11ListRolesResponse\x12\x14\n" +
	"\x05roles\x18\x01 \x03(\tR\x05roles\".\n" +
	"\x12DeleteUserResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\".\n" +
	"\x12RemoveRoleResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"3\n" +
	"\bUserList\x12'\n" +
	"\x05users\x18\x01 \x03(\v2\x11.user.UserProfileR\x05users\"|\n" +
	"\vUserProfile\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\x12\x14\n" +
	"\x05roles\x18\x05 \x03(\tR\x05roles\"*\n" +
	"\x0eLogoutResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"*\n" +
	"\x12VerifyEmailRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"/\n" +
	"\x13VerifyEmailResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\x8f\x06\n" +
	"\vUserService\x125\n" +
	"\bRegister\x12\x15.user.RegisterRequest\x1a\x12.user.AuthResponse\x12/\n" +
	"\x05Login\x12\x12.user.LoginRequest\x1a\x12.user.AuthResponse\x126\n" +
	"\x06Logout\x12\x16.google.protobuf.Empty\x1a\x14.user.LogoutResponse\x128\n" +
	"\n" +
	"GetProfile\x12\x17.user.GetProfileRequest\x1a\x11.user.UserProfile\x122\n" +
	"\x05GetMe\x12\x16.google.protobuf.Empty\x1a\x11.user.UserProfile\x12>\n" +
	"\rUpdateProfile\x12\x1a.user.UpdateProfileRequest\x1a\x11.user.UserProfile\x12?\n" +
	"\n" +
	"AssignRole\x12\x17.user.AssignRoleRequest\x1a\x18.user.AssignRoleResponse\x12<\n" +
	"\tListRoles\x12\x16.user.ListRolesRequest\x1a\x17.user.ListRolesResponse\x125\n" +
	"\vGetAllUsers\x12\x16.google.protobuf.Empty\x1a\x0e.user.UserList\x126\n" +
	"\vSearchUsers\x12\x17.user.SearchUserRequest\x1a\x0e.user.UserList\x12?\n" +
	"\n" +
	"DeleteUser\x12\x17.user.DeleteUserRequest\x1a\x18.user.DeleteUserResponse\x12?\n" +
	"\n" +
	"RemoveRole\x12\x17.user.RemoveRoleRequest\x1a\x18.user.RemoveRoleResponse\x12B\n" +
	"\vVerifyEmail\x12\x18.user.VerifyEmailRequest\x1a\x19.user.VerifyEmailResponseB\x0eZ\f./proto;userb\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_user_proto_goTypes = []any{
	(*RegisterRequest)(nil),      // 0: user.RegisterRequest
	(*LoginRequest)(nil),         // 1: user.LoginRequest
	(*GetProfileRequest)(nil),    // 2: user.GetProfileRequest
	(*UpdateProfileRequest)(nil), // 3: user.UpdateProfileRequest
	(*AssignRoleRequest)(nil),    // 4: user.AssignRoleRequest
	(*ListRolesRequest)(nil),     // 5: user.ListRolesRequest
	(*SearchUserRequest)(nil),    // 6: user.SearchUserRequest
	(*DeleteUserRequest)(nil),    // 7: user.DeleteUserRequest
	(*RemoveRoleRequest)(nil),    // 8: user.RemoveRoleRequest
	(*AuthResponse)(nil),         // 9: user.AuthResponse
	(*AssignRoleResponse)(nil),   // 10: user.AssignRoleResponse
	(*ListRolesResponse)(nil),    // 11: user.ListRolesResponse
	(*DeleteUserResponse)(nil),   // 12: user.DeleteUserResponse
	(*RemoveRoleResponse)(nil),   // 13: user.RemoveRoleResponse
	(*UserList)(nil),             // 14: user.UserList
	(*UserProfile)(nil),          // 15: user.UserProfile
	(*LogoutResponse)(nil),       // 16: user.LogoutResponse
	(*VerifyEmailRequest)(nil),   // 17: user.VerifyEmailRequest
	(*VerifyEmailResponse)(nil),  // 18: user.VerifyEmailResponse
	(*emptypb.Empty)(nil),        // 19: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	15, // 0: user.AuthResponse.profile:type_name -> user.UserProfile
	15, // 1: user.UserList.users:type_name -> user.UserProfile
	0,  // 2: user.UserService.Register:input_type -> user.RegisterRequest
	1,  // 3: user.UserService.Login:input_type -> user.LoginRequest
	19, // 4: user.UserService.Logout:input_type -> google.protobuf.Empty
	2,  // 5: user.UserService.GetProfile:input_type -> user.GetProfileRequest
	19, // 6: user.UserService.GetMe:input_type -> google.protobuf.Empty
	3,  // 7: user.UserService.UpdateProfile:input_type -> user.UpdateProfileRequest
	4,  // 8: user.UserService.AssignRole:input_type -> user.AssignRoleRequest
	5,  // 9: user.UserService.ListRoles:input_type -> user.ListRolesRequest
	19, // 10: user.UserService.GetAllUsers:input_type -> google.protobuf.Empty
	6,  // 11: user.UserService.SearchUsers:input_type -> user.SearchUserRequest
	7,  // 12: user.UserService.DeleteUser:input_type -> user.DeleteUserRequest
	8,  // 13: user.UserService.RemoveRole:input_type -> user.RemoveRoleRequest
	17, // 14: user.UserService.VerifyEmail:input_type -> user.VerifyEmailRequest
	9,  // 15: user.UserService.Register:output_type -> user.AuthResponse
	9,  // 16: user.UserService.Login:output_type -> user.AuthResponse
	16, // 17: user.UserService.Logout:output_type -> user.LogoutResponse
	15, // 18: user.UserService.GetProfile:output_type -> user.UserProfile
	15, // 19: user.UserService.GetMe:output_type -> user.UserProfile
	15, // 20: user.UserService.UpdateProfile:output_type -> user.UserProfile
	10, // 21: user.UserService.AssignRole:output_type -> user.AssignRoleResponse
	11, // 22: user.UserService.ListRoles:output_type -> user.ListRolesResponse
	14, // 23: user.UserService.GetAllUsers:output_type -> user.UserList
	14, // 24: user.UserService.SearchUsers:output_type -> user.UserList
	12, // 25: user.UserService.DeleteUser:output_type -> user.DeleteUserResponse
	13, // 26: user.UserService.RemoveRole:output_type -> user.RemoveRoleResponse
	18, // 27: user.UserService.VerifyEmail:output_type -> user.VerifyEmailResponse
	15, // [15:28] is the sub-list for method output_type
	2,  // [2:15] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   19,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
