// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/account/service/v1/account.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sex   bool   `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Age   int64  `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAccountRequest) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *CreateAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type CreateAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateAccountReply) Reset() {
	*x = CreateAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountReply) ProtoMessage() {}

func (x *CreateAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountReply.ProtoReflect.Descriptor instead.
func (*CreateAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{2}
}

type UpdateAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAccountReply) Reset() {
	*x = UpdateAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountReply) ProtoMessage() {}

func (x *UpdateAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountReply.ProtoReflect.Descriptor instead.
func (*UpdateAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{3}
}

type DeleteAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccountRequest) Reset() {
	*x = DeleteAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountRequest) ProtoMessage() {}

func (x *DeleteAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{4}
}

type DeleteAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccountReply) Reset() {
	*x = DeleteAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountReply) ProtoMessage() {}

func (x *DeleteAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountReply.ProtoReflect.Descriptor instead.
func (*DeleteAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{5}
}

type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{6}
}

func (x *GetAccountRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAccountReply) Reset() {
	*x = GetAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountReply) ProtoMessage() {}

func (x *GetAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountReply.ProtoReflect.Descriptor instead.
func (*GetAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{7}
}

type ListAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAccountRequest) Reset() {
	*x = ListAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountRequest) ProtoMessage() {}

func (x *ListAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountRequest.ProtoReflect.Descriptor instead.
func (*ListAccountRequest) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{8}
}

type ListAccountReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ListAccountReply_Account `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListAccountReply) Reset() {
	*x = ListAccountReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountReply) ProtoMessage() {}

func (x *ListAccountReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountReply.ProtoReflect.Descriptor instead.
func (*ListAccountReply) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{9}
}

func (x *ListAccountReply) GetResults() []*ListAccountReply_Account {
	if x != nil {
		return x.Results
	}
	return nil
}

type ListAccountReply_Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Sex   bool   `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Age   int64  `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *ListAccountReply_Account) Reset() {
	*x = ListAccountReply_Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_account_service_v1_account_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountReply_Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountReply_Account) ProtoMessage() {}

func (x *ListAccountReply_Account) ProtoReflect() protoreflect.Message {
	mi := &file_api_account_service_v1_account_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountReply_Account.ProtoReflect.Descriptor instead.
func (*ListAccountReply_Account) Descriptor() ([]byte, []int) {
	return file_api_account_service_v1_account_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ListAccountReply_Account) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListAccountReply_Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListAccountReply_Account) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *ListAccountReply_Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ListAccountReply_Account) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_api_account_service_v1_account_proto protoreflect.FileDescriptor

var file_api_account_service_v1_account_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x1d,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22,
	0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xc7, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x1a, 0x67, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x65,
	0x78, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x32, 0xdc, 0x04, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x82, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x69, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x7b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x79, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x36, 0x0a, 0x16, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_account_service_v1_account_proto_rawDescOnce sync.Once
	file_api_account_service_v1_account_proto_rawDescData = file_api_account_service_v1_account_proto_rawDesc
)

func file_api_account_service_v1_account_proto_rawDescGZIP() []byte {
	file_api_account_service_v1_account_proto_rawDescOnce.Do(func() {
		file_api_account_service_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_account_service_v1_account_proto_rawDescData)
	})
	return file_api_account_service_v1_account_proto_rawDescData
}

var file_api_account_service_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_account_service_v1_account_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),     // 0: api.account.service.v1.CreateAccountRequest
	(*CreateAccountReply)(nil),       // 1: api.account.service.v1.CreateAccountReply
	(*UpdateAccountRequest)(nil),     // 2: api.account.service.v1.UpdateAccountRequest
	(*UpdateAccountReply)(nil),       // 3: api.account.service.v1.UpdateAccountReply
	(*DeleteAccountRequest)(nil),     // 4: api.account.service.v1.DeleteAccountRequest
	(*DeleteAccountReply)(nil),       // 5: api.account.service.v1.DeleteAccountReply
	(*GetAccountRequest)(nil),        // 6: api.account.service.v1.GetAccountRequest
	(*GetAccountReply)(nil),          // 7: api.account.service.v1.GetAccountReply
	(*ListAccountRequest)(nil),       // 8: api.account.service.v1.ListAccountRequest
	(*ListAccountReply)(nil),         // 9: api.account.service.v1.ListAccountReply
	(*ListAccountReply_Account)(nil), // 10: api.account.service.v1.ListAccountReply.Account
}
var file_api_account_service_v1_account_proto_depIdxs = []int32{
	10, // 0: api.account.service.v1.ListAccountReply.results:type_name -> api.account.service.v1.ListAccountReply.Account
	0,  // 1: api.account.service.v1.Account.CreateAccount:input_type -> api.account.service.v1.CreateAccountRequest
	2,  // 2: api.account.service.v1.Account.UpdateAccount:input_type -> api.account.service.v1.UpdateAccountRequest
	4,  // 3: api.account.service.v1.Account.DeleteAccount:input_type -> api.account.service.v1.DeleteAccountRequest
	6,  // 4: api.account.service.v1.Account.GetAccount:input_type -> api.account.service.v1.GetAccountRequest
	8,  // 5: api.account.service.v1.Account.ListAccount:input_type -> api.account.service.v1.ListAccountRequest
	1,  // 6: api.account.service.v1.Account.CreateAccount:output_type -> api.account.service.v1.CreateAccountReply
	3,  // 7: api.account.service.v1.Account.UpdateAccount:output_type -> api.account.service.v1.UpdateAccountReply
	5,  // 8: api.account.service.v1.Account.DeleteAccount:output_type -> api.account.service.v1.DeleteAccountReply
	7,  // 9: api.account.service.v1.Account.GetAccount:output_type -> api.account.service.v1.GetAccountReply
	9,  // 10: api.account.service.v1.Account.ListAccount:output_type -> api.account.service.v1.ListAccountReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_account_service_v1_account_proto_init() }
func file_api_account_service_v1_account_proto_init() {
	if File_api_account_service_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_account_service_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_api_account_service_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountReply); i {
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
		file_api_account_service_v1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRequest); i {
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
		file_api_account_service_v1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountReply); i {
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
		file_api_account_service_v1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountRequest); i {
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
		file_api_account_service_v1_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountReply); i {
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
		file_api_account_service_v1_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountRequest); i {
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
		file_api_account_service_v1_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountReply); i {
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
		file_api_account_service_v1_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountRequest); i {
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
		file_api_account_service_v1_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountReply); i {
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
		file_api_account_service_v1_account_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccountReply_Account); i {
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
			RawDescriptor: file_api_account_service_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_account_service_v1_account_proto_goTypes,
		DependencyIndexes: file_api_account_service_v1_account_proto_depIdxs,
		MessageInfos:      file_api_account_service_v1_account_proto_msgTypes,
	}.Build()
	File_api_account_service_v1_account_proto = out.File
	file_api_account_service_v1_account_proto_rawDesc = nil
	file_api_account_service_v1_account_proto_goTypes = nil
	file_api_account_service_v1_account_proto_depIdxs = nil
}
