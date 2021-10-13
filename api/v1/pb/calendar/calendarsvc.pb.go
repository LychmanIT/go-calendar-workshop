// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: calendarsvc.proto

package calendarsvc

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Time        string   `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Timezone    string   `protobuf:"bytes,5,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Duration    int32    `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Notes       []string `protobuf:"bytes,7,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Event) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Event) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Event) GetNotes() []string {
	if x != nil {
		return x.Notes
	}
	return nil
}

type IndexEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []*IndexEventRequest_Filters `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *IndexEventRequest) Reset() {
	*x = IndexEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEventRequest) ProtoMessage() {}

func (x *IndexEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEventRequest.ProtoReflect.Descriptor instead.
func (*IndexEventRequest) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{1}
}

func (x *IndexEventRequest) GetFilters() []*IndexEventRequest_Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

type IndexEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	Err    string   `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *IndexEventReply) Reset() {
	*x = IndexEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEventReply) ProtoMessage() {}

func (x *IndexEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEventReply.ProtoReflect.Descriptor instead.
func (*IndexEventReply) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{2}
}

func (x *IndexEventReply) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *IndexEventReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type StoreEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *StoreEventRequest) Reset() {
	*x = StoreEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreEventRequest) ProtoMessage() {}

func (x *StoreEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreEventRequest.ProtoReflect.Descriptor instead.
func (*StoreEventRequest) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{3}
}

func (x *StoreEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type StoreEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *StoreEventReply) Reset() {
	*x = StoreEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreEventReply) ProtoMessage() {}

func (x *StoreEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreEventReply.ProtoReflect.Descriptor instead.
func (*StoreEventReply) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{4}
}

func (x *StoreEventReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StoreEventReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type ShowEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ShowEventRequest) Reset() {
	*x = ShowEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowEventRequest) ProtoMessage() {}

func (x *ShowEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowEventRequest.ProtoReflect.Descriptor instead.
func (*ShowEventRequest) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{5}
}

func (x *ShowEventRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ShowEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Err   string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ShowEventReply) Reset() {
	*x = ShowEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowEventReply) ProtoMessage() {}

func (x *ShowEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowEventReply.ProtoReflect.Descriptor instead.
func (*ShowEventReply) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{6}
}

func (x *ShowEventReply) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ShowEventReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event *Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type UpdateEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *UpdateEventReply) Reset() {
	*x = UpdateEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventReply) ProtoMessage() {}

func (x *UpdateEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventReply.ProtoReflect.Descriptor instead.
func (*UpdateEventReply) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateEventReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateEventReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventRequest.ProtoReflect.Descriptor instead.
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DeleteEventReply) Reset() {
	*x = DeleteEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventReply) ProtoMessage() {}

func (x *DeleteEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEventReply.ProtoReflect.Descriptor instead.
func (*DeleteEventReply) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteEventReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteEventReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type ServiceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceStatusRequest) Reset() {
	*x = ServiceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatusRequest) ProtoMessage() {}

func (x *ServiceStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatusRequest.ProtoReflect.Descriptor instead.
func (*ServiceStatusRequest) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{11}
}

type ServiceStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ServiceStatusReply) Reset() {
	*x = ServiceStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatusReply) ProtoMessage() {}

func (x *ServiceStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatusReply.ProtoReflect.Descriptor instead.
func (*ServiceStatusReply) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{12}
}

func (x *ServiceStatusReply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ServiceStatusReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type IndexEventRequest_Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IndexEventRequest_Filters) Reset() {
	*x = IndexEventRequest_Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendarsvc_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexEventRequest_Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexEventRequest_Filters) ProtoMessage() {}

func (x *IndexEventRequest_Filters) ProtoReflect() protoreflect.Message {
	mi := &file_calendarsvc_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexEventRequest_Filters.ProtoReflect.Descriptor instead.
func (*IndexEventRequest_Filters) Descriptor() ([]byte, []int) {
	return file_calendarsvc_proto_rawDescGZIP(), []int{1, 0}
}

func (x *IndexEventRequest_Filters) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IndexEventRequest_Filters) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_calendarsvc_proto protoreflect.FileDescriptor

var file_calendarsvc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xb1, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x7f, 0x0a, 0x11, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x37, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x31, 0x0a, 0x07, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x0f,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x21, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x45, 0x72, 0x72, 0x22, 0x34, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x77, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x0e, 0x53,
	0x68, 0x6f, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72,
	0x22, 0x45, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x3c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3e, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72,
	0x72, 0x32, 0xfe, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x3a,
	0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x68, 0x6f, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calendarsvc_proto_rawDescOnce sync.Once
	file_calendarsvc_proto_rawDescData = file_calendarsvc_proto_rawDesc
)

func file_calendarsvc_proto_rawDescGZIP() []byte {
	file_calendarsvc_proto_rawDescOnce.Do(func() {
		file_calendarsvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calendarsvc_proto_rawDescData)
	})
	return file_calendarsvc_proto_rawDescData
}

var file_calendarsvc_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_calendarsvc_proto_goTypes = []interface{}{
	(*Event)(nil),                     // 0: pb.Event
	(*IndexEventRequest)(nil),         // 1: pb.IndexEventRequest
	(*IndexEventReply)(nil),           // 2: pb.IndexEventReply
	(*StoreEventRequest)(nil),         // 3: pb.StoreEventRequest
	(*StoreEventReply)(nil),           // 4: pb.StoreEventReply
	(*ShowEventRequest)(nil),          // 5: pb.ShowEventRequest
	(*ShowEventReply)(nil),            // 6: pb.ShowEventReply
	(*UpdateEventRequest)(nil),        // 7: pb.UpdateEventRequest
	(*UpdateEventReply)(nil),          // 8: pb.UpdateEventReply
	(*DeleteEventRequest)(nil),        // 9: pb.DeleteEventRequest
	(*DeleteEventReply)(nil),          // 10: pb.DeleteEventReply
	(*ServiceStatusRequest)(nil),      // 11: pb.ServiceStatusRequest
	(*ServiceStatusReply)(nil),        // 12: pb.ServiceStatusReply
	(*IndexEventRequest_Filters)(nil), // 13: pb.IndexEventRequest.Filters
}
var file_calendarsvc_proto_depIdxs = []int32{
	13, // 0: pb.IndexEventRequest.filters:type_name -> pb.IndexEventRequest.Filters
	0,  // 1: pb.IndexEventReply.events:type_name -> pb.Event
	0,  // 2: pb.StoreEventRequest.event:type_name -> pb.Event
	0,  // 3: pb.ShowEventReply.event:type_name -> pb.Event
	0,  // 4: pb.UpdateEventRequest.event:type_name -> pb.Event
	1,  // 5: pb.Calendar.IndexEvent:input_type -> pb.IndexEventRequest
	3,  // 6: pb.Calendar.StoreEvent:input_type -> pb.StoreEventRequest
	5,  // 7: pb.Calendar.ShowEvent:input_type -> pb.ShowEventRequest
	7,  // 8: pb.Calendar.UpdateEvent:input_type -> pb.UpdateEventRequest
	9,  // 9: pb.Calendar.DeleteEvent:input_type -> pb.DeleteEventRequest
	11, // 10: pb.Calendar.ServiceStatus:input_type -> pb.ServiceStatusRequest
	2,  // 11: pb.Calendar.IndexEvent:output_type -> pb.IndexEventReply
	4,  // 12: pb.Calendar.StoreEvent:output_type -> pb.StoreEventReply
	6,  // 13: pb.Calendar.ShowEvent:output_type -> pb.ShowEventReply
	8,  // 14: pb.Calendar.UpdateEvent:output_type -> pb.UpdateEventReply
	10, // 15: pb.Calendar.DeleteEvent:output_type -> pb.DeleteEventReply
	12, // 16: pb.Calendar.ServiceStatus:output_type -> pb.ServiceStatusReply
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_calendarsvc_proto_init() }
func file_calendarsvc_proto_init() {
	if File_calendarsvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calendarsvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_calendarsvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexEventRequest); i {
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
		file_calendarsvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexEventReply); i {
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
		file_calendarsvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreEventRequest); i {
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
		file_calendarsvc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreEventReply); i {
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
		file_calendarsvc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowEventRequest); i {
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
		file_calendarsvc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowEventReply); i {
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
		file_calendarsvc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventRequest); i {
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
		file_calendarsvc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventReply); i {
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
		file_calendarsvc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventRequest); i {
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
		file_calendarsvc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEventReply); i {
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
		file_calendarsvc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatusRequest); i {
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
		file_calendarsvc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatusReply); i {
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
		file_calendarsvc_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexEventRequest_Filters); i {
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
			RawDescriptor: file_calendarsvc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calendarsvc_proto_goTypes,
		DependencyIndexes: file_calendarsvc_proto_depIdxs,
		MessageInfos:      file_calendarsvc_proto_msgTypes,
	}.Build()
	File_calendarsvc_proto = out.File
	file_calendarsvc_proto_rawDesc = nil
	file_calendarsvc_proto_goTypes = nil
	file_calendarsvc_proto_depIdxs = nil
}
