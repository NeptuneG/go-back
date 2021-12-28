// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: services/user/proto/user.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Points int64  `protobuf:"varint,3,opt,name=points,proto3" json:"points,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_services_user_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPoints() int64 {
	if x != nil {
		return x.Points
	}
	return 0
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_services_user_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_services_user_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_proto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_proto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_services_user_proto_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_proto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_proto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_services_user_proto_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ConsumeUserPointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Points      int32  `protobuf:"varint,2,opt,name=points,proto3" json:"points,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ConsumeUserPointsRequest) Reset() {
	*x = ConsumeUserPointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_proto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeUserPointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeUserPointsRequest) ProtoMessage() {}

func (x *ConsumeUserPointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_proto_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeUserPointsRequest.ProtoReflect.Descriptor instead.
func (*ConsumeUserPointsRequest) Descriptor() ([]byte, []int) {
	return file_services_user_proto_user_proto_rawDescGZIP(), []int{5}
}

func (x *ConsumeUserPointsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsumeUserPointsRequest) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *ConsumeUserPointsRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ConsumeUserPointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ConsumeUserPointsResponse) Reset() {
	*x = ConsumeUserPointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_user_proto_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeUserPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeUserPointsResponse) ProtoMessage() {}

func (x *ConsumeUserPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_user_proto_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeUserPointsResponse.ProtoReflect.Descriptor instead.
func (*ConsumeUserPointsResponse) Descriptor() ([]byte, []int) {
	return file_services_user_proto_user_proto_rawDescGZIP(), []int{6}
}

func (x *ConsumeUserPointsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_services_user_proto_user_proto protoreflect.FileDescriptor

var file_services_user_proto_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4e, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x18, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x55, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x65,
	0x70, 0x74, 0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xa9, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65,
	0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63,
	0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6e, 0x65, 0x70, 0x74,
	0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x82, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x2e, 0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x88, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x38,
	0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6e, 0x65, 0x70, 0x74, 0x75,
	0x6e, 0x65, 0x67, 0x2e, 0x67, 0x6f, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x69,
	0x76, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x65, 0x70, 0x74, 0x75, 0x6e, 0x65, 0x47, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61,
	0x63, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_user_proto_user_proto_rawDescOnce sync.Once
	file_services_user_proto_user_proto_rawDescData = file_services_user_proto_user_proto_rawDesc
)

func file_services_user_proto_user_proto_rawDescGZIP() []byte {
	file_services_user_proto_user_proto_rawDescOnce.Do(func() {
		file_services_user_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_user_proto_user_proto_rawDescData)
	})
	return file_services_user_proto_user_proto_rawDescData
}

var file_services_user_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_user_proto_user_proto_goTypes = []interface{}{
	(*User)(nil),                      // 0: neptuneg.go_back.serivces.user.User
	(*CreateUserRequest)(nil),         // 1: neptuneg.go_back.serivces.user.CreateUserRequest
	(*CreateUserResponse)(nil),        // 2: neptuneg.go_back.serivces.user.CreateUserResponse
	(*GetUserRequest)(nil),            // 3: neptuneg.go_back.serivces.user.GetUserRequest
	(*GetUserResponse)(nil),           // 4: neptuneg.go_back.serivces.user.GetUserResponse
	(*ConsumeUserPointsRequest)(nil),  // 5: neptuneg.go_back.serivces.user.ConsumeUserPointsRequest
	(*ConsumeUserPointsResponse)(nil), // 6: neptuneg.go_back.serivces.user.ConsumeUserPointsResponse
}
var file_services_user_proto_user_proto_depIdxs = []int32{
	0, // 0: neptuneg.go_back.serivces.user.CreateUserResponse.user:type_name -> neptuneg.go_back.serivces.user.User
	0, // 1: neptuneg.go_back.serivces.user.GetUserResponse.user:type_name -> neptuneg.go_back.serivces.user.User
	0, // 2: neptuneg.go_back.serivces.user.ConsumeUserPointsResponse.user:type_name -> neptuneg.go_back.serivces.user.User
	1, // 3: neptuneg.go_back.serivces.user.UserService.CreateUser:input_type -> neptuneg.go_back.serivces.user.CreateUserRequest
	3, // 4: neptuneg.go_back.serivces.user.UserService.GetUser:input_type -> neptuneg.go_back.serivces.user.GetUserRequest
	5, // 5: neptuneg.go_back.serivces.user.UserService.ConsumeUserPoints:input_type -> neptuneg.go_back.serivces.user.ConsumeUserPointsRequest
	2, // 6: neptuneg.go_back.serivces.user.UserService.CreateUser:output_type -> neptuneg.go_back.serivces.user.CreateUserResponse
	4, // 7: neptuneg.go_back.serivces.user.UserService.GetUser:output_type -> neptuneg.go_back.serivces.user.GetUserResponse
	6, // 8: neptuneg.go_back.serivces.user.UserService.ConsumeUserPoints:output_type -> neptuneg.go_back.serivces.user.ConsumeUserPointsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_user_proto_user_proto_init() }
func file_services_user_proto_user_proto_init() {
	if File_services_user_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_user_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_user_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_user_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_user_proto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_user_proto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_user_proto_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeUserPointsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_services_user_proto_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeUserPointsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_user_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_user_proto_user_proto_goTypes,
		DependencyIndexes: file_services_user_proto_user_proto_depIdxs,
		MessageInfos:      file_services_user_proto_user_proto_msgTypes,
	}.Build()
	File_services_user_proto_user_proto = out.File
	file_services_user_proto_user_proto_rawDesc = nil
	file_services_user_proto_user_proto_goTypes = nil
	file_services_user_proto_user_proto_depIdxs = nil
}