// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: authsvc.proto

package authsvc

import (
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

type AuthCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthCredentials) Reset() {
	*x = AuthCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCredentials) ProtoMessage() {}

func (x *AuthCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCredentials.ProtoReflect.Descriptor instead.
func (*AuthCredentials) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{0}
}

func (x *AuthCredentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AuthCredentials) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credentials *AuthCredentials `protobuf:"bytes,1,opt,name=credentials,proto3" json:"credentials,omitempty"`
}

func (x *AuthLoginRequest) Reset() {
	*x = AuthLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginRequest) ProtoMessage() {}

func (x *AuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{1}
}

func (x *AuthLoginRequest) GetCredentials() *AuthCredentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type AuthLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Err   string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AuthLoginReply) Reset() {
	*x = AuthLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginReply) ProtoMessage() {}

func (x *AuthLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginReply.ProtoReflect.Descriptor instead.
func (*AuthLoginReply) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{2}
}

func (x *AuthLoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthLoginReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type AuthLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthLogoutRequest) Reset() {
	*x = AuthLogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLogoutRequest) ProtoMessage() {}

func (x *AuthLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLogoutRequest.ProtoReflect.Descriptor instead.
func (*AuthLogoutRequest) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{3}
}

func (x *AuthLogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthLogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AuthLogoutReply) Reset() {
	*x = AuthLogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLogoutReply) ProtoMessage() {}

func (x *AuthLogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLogoutReply.ProtoReflect.Descriptor instead.
func (*AuthLogoutReply) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{4}
}

func (x *AuthLogoutReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuthLogoutReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type AuthVerifyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthVerifyTokenRequest) Reset() {
	*x = AuthVerifyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthVerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthVerifyTokenRequest) ProtoMessage() {}

func (x *AuthVerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthVerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*AuthVerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{5}
}

func (x *AuthVerifyTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthVerifyTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AuthVerifyTokenReply) Reset() {
	*x = AuthVerifyTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthVerifyTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthVerifyTokenReply) ProtoMessage() {}

func (x *AuthVerifyTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthVerifyTokenReply.ProtoReflect.Descriptor instead.
func (*AuthVerifyTokenReply) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{6}
}

func (x *AuthVerifyTokenReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type AuthChangeTimezoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Timezone string `protobuf:"bytes,2,opt,name=timezone,proto3" json:"timezone,omitempty"`
}

func (x *AuthChangeTimezoneRequest) Reset() {
	*x = AuthChangeTimezoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthChangeTimezoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthChangeTimezoneRequest) ProtoMessage() {}

func (x *AuthChangeTimezoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthChangeTimezoneRequest.ProtoReflect.Descriptor instead.
func (*AuthChangeTimezoneRequest) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{7}
}

func (x *AuthChangeTimezoneRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthChangeTimezoneRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

type AuthChangeTimezoneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AuthChangeTimezoneReply) Reset() {
	*x = AuthChangeTimezoneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthChangeTimezoneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthChangeTimezoneReply) ProtoMessage() {}

func (x *AuthChangeTimezoneReply) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthChangeTimezoneReply.ProtoReflect.Descriptor instead.
func (*AuthChangeTimezoneReply) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{8}
}

func (x *AuthChangeTimezoneReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuthChangeTimezoneReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type AuthGetTimezoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthGetTimezoneRequest) Reset() {
	*x = AuthGetTimezoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthGetTimezoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthGetTimezoneRequest) ProtoMessage() {}

func (x *AuthGetTimezoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthGetTimezoneRequest.ProtoReflect.Descriptor instead.
func (*AuthGetTimezoneRequest) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{9}
}

func (x *AuthGetTimezoneRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthGetTimezoneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timezone string `protobuf:"bytes,1,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Err      string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AuthGetTimezoneReply) Reset() {
	*x = AuthGetTimezoneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthGetTimezoneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthGetTimezoneReply) ProtoMessage() {}

func (x *AuthGetTimezoneReply) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthGetTimezoneReply.ProtoReflect.Descriptor instead.
func (*AuthGetTimezoneReply) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{10}
}

func (x *AuthGetTimezoneReply) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *AuthGetTimezoneReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type AuthServiceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthServiceStatusRequest) Reset() {
	*x = AuthServiceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthServiceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthServiceStatusRequest) ProtoMessage() {}

func (x *AuthServiceStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthServiceStatusRequest.ProtoReflect.Descriptor instead.
func (*AuthServiceStatusRequest) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{11}
}

type AuthServiceStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AuthServiceStatusReply) Reset() {
	*x = AuthServiceStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authsvc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthServiceStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthServiceStatusReply) ProtoMessage() {}

func (x *AuthServiceStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_authsvc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthServiceStatusReply.ProtoReflect.Descriptor instead.
func (*AuthServiceStatusReply) Descriptor() ([]byte, []int) {
	return file_authsvc_proto_rawDescGZIP(), []int{12}
}

func (x *AuthServiceStatusReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AuthServiceStatusReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_authsvc_proto protoreflect.FileDescriptor

var file_authsvc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x49, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x49,
	0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x38, 0x0a, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x22, 0x29, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3b,
	0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x2e, 0x0a, 0x16, 0x41,
	0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x28, 0x0a, 0x14, 0x41,
	0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x4d, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x43, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x2e, 0x0a, 0x16, 0x41, 0x75, 0x74,
	0x68, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x44, 0x0a, 0x14, 0x41, 0x75, 0x74,
	0x68, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22,
	0x1a, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x16, 0x41,
	0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32,
	0xb6, 0x03, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12,
	0x15, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1d,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0f,
	0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x61,
	0x75, 0x74, 0x68, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authsvc_proto_rawDescOnce sync.Once
	file_authsvc_proto_rawDescData = file_authsvc_proto_rawDesc
)

func file_authsvc_proto_rawDescGZIP() []byte {
	file_authsvc_proto_rawDescOnce.Do(func() {
		file_authsvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_authsvc_proto_rawDescData)
	})
	return file_authsvc_proto_rawDescData
}

var file_authsvc_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_authsvc_proto_goTypes = []interface{}{
	(*AuthCredentials)(nil),           // 0: pb.AuthCredentials
	(*AuthLoginRequest)(nil),          // 1: pb.AuthLoginRequest
	(*AuthLoginReply)(nil),            // 2: pb.AuthLoginReply
	(*AuthLogoutRequest)(nil),         // 3: pb.AuthLogoutRequest
	(*AuthLogoutReply)(nil),           // 4: pb.AuthLogoutReply
	(*AuthVerifyTokenRequest)(nil),    // 5: pb.AuthVerifyTokenRequest
	(*AuthVerifyTokenReply)(nil),      // 6: pb.AuthVerifyTokenReply
	(*AuthChangeTimezoneRequest)(nil), // 7: pb.AuthChangeTimezoneRequest
	(*AuthChangeTimezoneReply)(nil),   // 8: pb.AuthChangeTimezoneReply
	(*AuthGetTimezoneRequest)(nil),    // 9: pb.AuthGetTimezoneRequest
	(*AuthGetTimezoneReply)(nil),      // 10: pb.AuthGetTimezoneReply
	(*AuthServiceStatusRequest)(nil),  // 11: pb.AuthServiceStatusRequest
	(*AuthServiceStatusReply)(nil),    // 12: pb.AuthServiceStatusReply
}
var file_authsvc_proto_depIdxs = []int32{
	0,  // 0: pb.AuthLoginRequest.credentials:type_name -> pb.AuthCredentials
	1,  // 1: pb.Auth.AuthLogin:input_type -> pb.AuthLoginRequest
	3,  // 2: pb.Auth.AuthLogout:input_type -> pb.AuthLogoutRequest
	5,  // 3: pb.Auth.AuthVerifyToken:input_type -> pb.AuthVerifyTokenRequest
	7,  // 4: pb.Auth.AuthChangeTimezone:input_type -> pb.AuthChangeTimezoneRequest
	9,  // 5: pb.Auth.AuthGetTimezone:input_type -> pb.AuthGetTimezoneRequest
	11, // 6: pb.Auth.AuthServiceStatus:input_type -> pb.AuthServiceStatusRequest
	2,  // 7: pb.Auth.AuthLogin:output_type -> pb.AuthLoginReply
	4,  // 8: pb.Auth.AuthLogout:output_type -> pb.AuthLogoutReply
	6,  // 9: pb.Auth.AuthVerifyToken:output_type -> pb.AuthVerifyTokenReply
	8,  // 10: pb.Auth.AuthChangeTimezone:output_type -> pb.AuthChangeTimezoneReply
	10, // 11: pb.Auth.AuthGetTimezone:output_type -> pb.AuthGetTimezoneReply
	12, // 12: pb.Auth.AuthServiceStatus:output_type -> pb.AuthServiceStatusReply
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_authsvc_proto_init() }
func file_authsvc_proto_init() {
	if File_authsvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authsvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCredentials); i {
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
		file_authsvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLoginRequest); i {
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
		file_authsvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLoginReply); i {
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
		file_authsvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLogoutRequest); i {
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
		file_authsvc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLogoutReply); i {
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
		file_authsvc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthVerifyTokenRequest); i {
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
		file_authsvc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthVerifyTokenReply); i {
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
		file_authsvc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthChangeTimezoneRequest); i {
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
		file_authsvc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthChangeTimezoneReply); i {
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
		file_authsvc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthGetTimezoneRequest); i {
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
		file_authsvc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthGetTimezoneReply); i {
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
		file_authsvc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthServiceStatusRequest); i {
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
		file_authsvc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthServiceStatusReply); i {
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
			RawDescriptor: file_authsvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authsvc_proto_goTypes,
		DependencyIndexes: file_authsvc_proto_depIdxs,
		MessageInfos:      file_authsvc_proto_msgTypes,
	}.Build()
	File_authsvc_proto = out.File
	file_authsvc_proto_rawDesc = nil
	file_authsvc_proto_goTypes = nil
	file_authsvc_proto_depIdxs = nil
}
