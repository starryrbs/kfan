// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/house/service/v1/house.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CreateHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Community    string `protobuf:"bytes,2,opt,name=community,proto3" json:"community,omitempty"`
	FloorCount   int32  `protobuf:"varint,3,opt,name=floor_count,json=floorCount,proto3" json:"floor_count,omitempty"`
	KitchenCount int32  `protobuf:"varint,4,opt,name=kitchen_count,json=kitchenCount,proto3" json:"kitchen_count,omitempty"`
	ToiletCount  int32  `protobuf:"varint,5,opt,name=toilet_count,json=toiletCount,proto3" json:"toilet_count,omitempty"`
	HallCount    int32  `protobuf:"varint,6,opt,name=hall_count,json=hallCount,proto3" json:"hall_count,omitempty"`
	RoomCount    int32  `protobuf:"varint,7,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
}

func (x *CreateHouseRequest) Reset() {
	*x = CreateHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_house_service_v1_house_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHouseRequest) ProtoMessage() {}

func (x *CreateHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_house_service_v1_house_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHouseRequest.ProtoReflect.Descriptor instead.
func (*CreateHouseRequest) Descriptor() ([]byte, []int) {
	return file_api_house_service_v1_house_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHouseRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateHouseRequest) GetCommunity() string {
	if x != nil {
		return x.Community
	}
	return ""
}

func (x *CreateHouseRequest) GetFloorCount() int32 {
	if x != nil {
		return x.FloorCount
	}
	return 0
}

func (x *CreateHouseRequest) GetKitchenCount() int32 {
	if x != nil {
		return x.KitchenCount
	}
	return 0
}

func (x *CreateHouseRequest) GetToiletCount() int32 {
	if x != nil {
		return x.ToiletCount
	}
	return 0
}

func (x *CreateHouseRequest) GetHallCount() int32 {
	if x != nil {
		return x.HallCount
	}
	return 0
}

func (x *CreateHouseRequest) GetRoomCount() int32 {
	if x != nil {
		return x.RoomCount
	}
	return 0
}

type CreateHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateHouseReply) Reset() {
	*x = CreateHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_house_service_v1_house_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHouseReply) ProtoMessage() {}

func (x *CreateHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_house_service_v1_house_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHouseReply.ProtoReflect.Descriptor instead.
func (*CreateHouseReply) Descriptor() ([]byte, []int) {
	return file_api_house_service_v1_house_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHouseReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHouseRequest) Reset() {
	*x = GetHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_house_service_v1_house_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHouseRequest) ProtoMessage() {}

func (x *GetHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_house_service_v1_house_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHouseRequest.ProtoReflect.Descriptor instead.
func (*GetHouseRequest) Descriptor() ([]byte, []int) {
	return file_api_house_service_v1_house_proto_rawDescGZIP(), []int{2}
}

func (x *GetHouseRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,11,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Community    string `protobuf:"bytes,2,opt,name=community,proto3" json:"community,omitempty"`
	FloorCount   int32  `protobuf:"varint,3,opt,name=floor_count,json=floorCount,proto3" json:"floor_count,omitempty"`
	KitchenCount int32  `protobuf:"varint,4,opt,name=kitchen_count,json=kitchenCount,proto3" json:"kitchen_count,omitempty"`
	ToiletCount  int32  `protobuf:"varint,5,opt,name=toilet_count,json=toiletCount,proto3" json:"toilet_count,omitempty"`
	HallCount    int32  `protobuf:"varint,6,opt,name=hall_count,json=hallCount,proto3" json:"hall_count,omitempty"`
	RoomCount    int32  `protobuf:"varint,7,opt,name=room_count,json=roomCount,proto3" json:"room_count,omitempty"`
}

func (x *GetHouseReply) Reset() {
	*x = GetHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_house_service_v1_house_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHouseReply) ProtoMessage() {}

func (x *GetHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_house_service_v1_house_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHouseReply.ProtoReflect.Descriptor instead.
func (*GetHouseReply) Descriptor() ([]byte, []int) {
	return file_api_house_service_v1_house_proto_rawDescGZIP(), []int{3}
}

func (x *GetHouseReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetHouseReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetHouseReply) GetCommunity() string {
	if x != nil {
		return x.Community
	}
	return ""
}

func (x *GetHouseReply) GetFloorCount() int32 {
	if x != nil {
		return x.FloorCount
	}
	return 0
}

func (x *GetHouseReply) GetKitchenCount() int32 {
	if x != nil {
		return x.KitchenCount
	}
	return 0
}

func (x *GetHouseReply) GetToiletCount() int32 {
	if x != nil {
		return x.ToiletCount
	}
	return 0
}

func (x *GetHouseReply) GetHallCount() int32 {
	if x != nil {
		return x.HallCount
	}
	return 0
}

func (x *GetHouseReply) GetRoomCount() int32 {
	if x != nil {
		return x.RoomCount
	}
	return 0
}

type ListHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHouseRequest) Reset() {
	*x = ListHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_house_service_v1_house_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHouseRequest) ProtoMessage() {}

func (x *ListHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_house_service_v1_house_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHouseRequest.ProtoReflect.Descriptor instead.
func (*ListHouseRequest) Descriptor() ([]byte, []int) {
	return file_api_house_service_v1_house_proto_rawDescGZIP(), []int{4}
}

type ListHouseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ListHouseReply_House `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListHouseReply) Reset() {
	*x = ListHouseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_house_service_v1_house_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHouseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHouseReply) ProtoMessage() {}

func (x *ListHouseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_house_service_v1_house_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHouseReply.ProtoReflect.Descriptor instead.
func (*ListHouseReply) Descriptor() ([]byte, []int) {
	return file_api_house_service_v1_house_proto_rawDescGZIP(), []int{5}
}

func (x *ListHouseReply) GetResults() []*ListHouseReply_House {
	if x != nil {
		return x.Results
	}
	return nil
}

type ListHouseReply_House struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,11,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Community   string `protobuf:"bytes,2,opt,name=community,proto3" json:"community,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ListHouseReply_House) Reset() {
	*x = ListHouseReply_House{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_house_service_v1_house_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHouseReply_House) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHouseReply_House) ProtoMessage() {}

func (x *ListHouseReply_House) ProtoReflect() protoreflect.Message {
	mi := &file_api_house_service_v1_house_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHouseReply_House.ProtoReflect.Descriptor instead.
func (*ListHouseReply_House) Descriptor() ([]byte, []int) {
	return file_api_house_service_v1_house_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListHouseReply_House) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListHouseReply_House) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListHouseReply_House) GetCommunity() string {
	if x != nil {
		return x.Community
	}
	return ""
}

func (x *ListHouseReply_House) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_api_house_service_v1_house_proto protoreflect.FileDescriptor

var file_api_house_service_v1_house_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x85, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x14,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72,
	0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x69, 0x6c, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f,
	0x69, 0x6c, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x6c,
	0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68,
	0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x6f,
	0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xfa,
	0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x6f, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6b, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f,
	0x69, 0x6c, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x74, 0x6f, 0x69, 0x6c, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x68, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x68, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xc5, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x44, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x6d, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xdf, 0x02, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x12, 0x76, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6d, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x76, 0x31, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x42, 0x4b, 0x0a, 0x14, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x61, 0x72, 0x72, 0x79, 0x72, 0x62, 0x73, 0x2f, 0x6b, 0x66, 0x61, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_house_service_v1_house_proto_rawDescOnce sync.Once
	file_api_house_service_v1_house_proto_rawDescData = file_api_house_service_v1_house_proto_rawDesc
)

func file_api_house_service_v1_house_proto_rawDescGZIP() []byte {
	file_api_house_service_v1_house_proto_rawDescOnce.Do(func() {
		file_api_house_service_v1_house_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_house_service_v1_house_proto_rawDescData)
	})
	return file_api_house_service_v1_house_proto_rawDescData
}

var file_api_house_service_v1_house_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_house_service_v1_house_proto_goTypes = []interface{}{
	(*CreateHouseRequest)(nil),   // 0: api.house.service.v1.CreateHouseRequest
	(*CreateHouseReply)(nil),     // 1: api.house.service.v1.CreateHouseReply
	(*GetHouseRequest)(nil),      // 2: api.house.service.v1.GetHouseRequest
	(*GetHouseReply)(nil),        // 3: api.house.service.v1.GetHouseReply
	(*ListHouseRequest)(nil),     // 4: api.house.service.v1.ListHouseRequest
	(*ListHouseReply)(nil),       // 5: api.house.service.v1.ListHouseReply
	(*ListHouseReply_House)(nil), // 6: api.house.service.v1.ListHouseReply.House
}
var file_api_house_service_v1_house_proto_depIdxs = []int32{
	6, // 0: api.house.service.v1.ListHouseReply.results:type_name -> api.house.service.v1.ListHouseReply.House
	0, // 1: api.house.service.v1.House.CreateHouse:input_type -> api.house.service.v1.CreateHouseRequest
	2, // 2: api.house.service.v1.House.GetHouse:input_type -> api.house.service.v1.GetHouseRequest
	4, // 3: api.house.service.v1.House.ListHouse:input_type -> api.house.service.v1.ListHouseRequest
	1, // 4: api.house.service.v1.House.CreateHouse:output_type -> api.house.service.v1.CreateHouseReply
	3, // 5: api.house.service.v1.House.GetHouse:output_type -> api.house.service.v1.GetHouseReply
	5, // 6: api.house.service.v1.House.ListHouse:output_type -> api.house.service.v1.ListHouseReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_house_service_v1_house_proto_init() }
func file_api_house_service_v1_house_proto_init() {
	if File_api_house_service_v1_house_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_house_service_v1_house_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHouseRequest); i {
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
		file_api_house_service_v1_house_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHouseReply); i {
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
		file_api_house_service_v1_house_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHouseRequest); i {
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
		file_api_house_service_v1_house_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHouseReply); i {
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
		file_api_house_service_v1_house_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHouseRequest); i {
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
		file_api_house_service_v1_house_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHouseReply); i {
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
		file_api_house_service_v1_house_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHouseReply_House); i {
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
			RawDescriptor: file_api_house_service_v1_house_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_house_service_v1_house_proto_goTypes,
		DependencyIndexes: file_api_house_service_v1_house_proto_depIdxs,
		MessageInfos:      file_api_house_service_v1_house_proto_msgTypes,
	}.Build()
	File_api_house_service_v1_house_proto = out.File
	file_api_house_service_v1_house_proto_rawDesc = nil
	file_api_house_service_v1_house_proto_goTypes = nil
	file_api_house_service_v1_house_proto_depIdxs = nil
}
