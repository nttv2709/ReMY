// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: calendar.proto

package remy

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListEventsRequest_DATE_TYPE int32

const (
	ListEventsRequest_DAY   ListEventsRequest_DATE_TYPE = 0
	ListEventsRequest_MONTH ListEventsRequest_DATE_TYPE = 1
	ListEventsRequest_YEAR  ListEventsRequest_DATE_TYPE = 2
)

// Enum value maps for ListEventsRequest_DATE_TYPE.
var (
	ListEventsRequest_DATE_TYPE_name = map[int32]string{
		0: "DAY",
		1: "MONTH",
		2: "YEAR",
	}
	ListEventsRequest_DATE_TYPE_value = map[string]int32{
		"DAY":   0,
		"MONTH": 1,
		"YEAR":  2,
	}
)

func (x ListEventsRequest_DATE_TYPE) Enum() *ListEventsRequest_DATE_TYPE {
	p := new(ListEventsRequest_DATE_TYPE)
	*p = x
	return p
}

func (x ListEventsRequest_DATE_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListEventsRequest_DATE_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_calendar_proto_enumTypes[0].Descriptor()
}

func (ListEventsRequest_DATE_TYPE) Type() protoreflect.EnumType {
	return &file_calendar_proto_enumTypes[0]
}

func (x ListEventsRequest_DATE_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListEventsRequest_DATE_TYPE.Descriptor instead.
func (ListEventsRequest_DATE_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{6, 0}
}

type RangeTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *RangeTime) Reset() {
	*x = RangeTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeTime) ProtoMessage() {}

func (x *RangeTime) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeTime.ProtoReflect.Descriptor instead.
func (*RangeTime) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{0}
}

func (x *RangeTime) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *RangeTime) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	RangeTime *RangeTime `protobuf:"bytes,3,opt,name=range_time,json=rangeTime,proto3" json:"range_time,omitempty"`
	Location  *Location  `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[1]
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
	return file_calendar_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetRangeTime() *RangeTime {
	if x != nil {
		return x.RangeTime
	}
	return nil
}

func (x *Event) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type ListEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event []*Event               `protobuf:"bytes,1,rep,name=event,proto3" json:"event,omitempty"`
	Time  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ListEvents) Reset() {
	*x = ListEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEvents) ProtoMessage() {}

func (x *ListEvents) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEvents.ProtoReflect.Descriptor instead.
func (*ListEvents) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{2}
}

func (x *ListEvents) GetEvent() []*Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *ListEvents) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Location) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title     string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	RangeTime *RangeTime `protobuf:"bytes,2,opt,name=range_time,json=rangeTime,proto3" json:"range_time,omitempty"`
	Location  *Location  `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{4}
}

func (x *CreateEventRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEventRequest) GetRangeTime() *RangeTime {
	if x != nil {
		return x.RangeTime
	}
	return nil
}

func (x *CreateEventRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type CreateEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateEventReply) Reset() {
	*x = CreateEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventReply) ProtoMessage() {}

func (x *CreateEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventReply.ProtoReflect.Descriptor instead.
func (*CreateEventReply) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{5}
}

func (x *CreateEventReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RangeTime *RangeTime                  `protobuf:"bytes,1,opt,name=range_time,json=rangeTime,proto3" json:"range_time,omitempty"`
	DateType  ListEventsRequest_DATE_TYPE `protobuf:"varint,2,opt,name=date_type,json=dateType,proto3,enum=api.ListEventsRequest_DATE_TYPE" json:"date_type,omitempty"`
}

func (x *ListEventsRequest) Reset() {
	*x = ListEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequest) ProtoMessage() {}

func (x *ListEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequest.ProtoReflect.Descriptor instead.
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{6}
}

func (x *ListEventsRequest) GetRangeTime() *RangeTime {
	if x != nil {
		return x.RangeTime
	}
	return nil
}

func (x *ListEventsRequest) GetDateType() ListEventsRequest_DATE_TYPE {
	if x != nil {
		return x.DateType
	}
	return ListEventsRequest_DAY
}

type ListEventsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListEvents []*ListEvents `protobuf:"bytes,1,rep,name=list_events,json=listEvents,proto3" json:"list_events,omitempty"`
}

func (x *ListEventsReply) Reset() {
	*x = ListEventsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsReply) ProtoMessage() {}

func (x *ListEventsReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsReply.ProtoReflect.Descriptor instead.
func (*ListEventsReply) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{7}
}

func (x *ListEventsReply) GetListEvents() []*ListEvents {
	if x != nil {
		return x.ListEvents
	}
	return nil
}

type GetRemindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GetRemindRequest) Reset() {
	*x = GetRemindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemindRequest) ProtoMessage() {}

func (x *GetRemindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemindRequest.ProtoReflect.Descriptor instead.
func (*GetRemindRequest) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{8}
}

func (x *GetRemindRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type GetRemindReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Okay bool                   `protobuf:"varint,2,opt,name=okay,proto3" json:"okay,omitempty"`
}

func (x *GetRemindReply) Reset() {
	*x = GetRemindReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRemindReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemindReply) ProtoMessage() {}

func (x *GetRemindReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemindReply.ProtoReflect.Descriptor instead.
func (*GetRemindReply) Descriptor() ([]byte, []int) {
	return file_calendar_proto_rawDescGZIP(), []int{9}
}

func (x *GetRemindReply) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *GetRemindReply) GetOkay() bool {
	if x != nil {
		return x.Okay
	}
	return false
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[10]
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
	return file_calendar_proto_rawDescGZIP(), []int{10}
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
}

func (x *UpdateEventReply) Reset() {
	*x = UpdateEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventReply) ProtoMessage() {}

func (x *UpdateEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[11]
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
	return file_calendar_proto_rawDescGZIP(), []int{11}
}

type DeleteEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEventRequest) Reset() {
	*x = DeleteEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventRequest) ProtoMessage() {}

func (x *DeleteEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[12]
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
	return file_calendar_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteEventRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteEventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEventReply) Reset() {
	*x = DeleteEventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calendar_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEventReply) ProtoMessage() {}

func (x *DeleteEventReply) ProtoReflect() protoreflect.Message {
	mi := &file_calendar_proto_msgTypes[13]
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
	return file_calendar_proto_rawDescGZIP(), []int{13}
}

var File_calendar_proto protoreflect.FileDescriptor

var file_calendar_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x09, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x84, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xac, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x29, 0x0a, 0x09, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f,
	0x4e, 0x54, 0x48, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x45, 0x41, 0x52, 0x10, 0x02, 0x22,
	0x43, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x30, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x6b, 0x61, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x6b, 0x61, 0x79, 0x22, 0x36, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32,
	0xcd, 0x02, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x74,
	0x74, 0x76, 0x32, 0x37, 0x30, 0x39, 0x2f, 0x72, 0x65, 0x6d, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_calendar_proto_rawDescOnce sync.Once
	file_calendar_proto_rawDescData = file_calendar_proto_rawDesc
)

func file_calendar_proto_rawDescGZIP() []byte {
	file_calendar_proto_rawDescOnce.Do(func() {
		file_calendar_proto_rawDescData = protoimpl.X.CompressGZIP(file_calendar_proto_rawDescData)
	})
	return file_calendar_proto_rawDescData
}

var file_calendar_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_calendar_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_calendar_proto_goTypes = []interface{}{
	(ListEventsRequest_DATE_TYPE)(0), // 0: api.ListEventsRequest.DATE_TYPE
	(*RangeTime)(nil),                // 1: api.RangeTime
	(*Event)(nil),                    // 2: api.Event
	(*ListEvents)(nil),               // 3: api.ListEvents
	(*Location)(nil),                 // 4: api.Location
	(*CreateEventRequest)(nil),       // 5: api.CreateEventRequest
	(*CreateEventReply)(nil),         // 6: api.CreateEventReply
	(*ListEventsRequest)(nil),        // 7: api.ListEventsRequest
	(*ListEventsReply)(nil),          // 8: api.ListEventsReply
	(*GetRemindRequest)(nil),         // 9: api.GetRemindRequest
	(*GetRemindReply)(nil),           // 10: api.GetRemindReply
	(*UpdateEventRequest)(nil),       // 11: api.UpdateEventRequest
	(*UpdateEventReply)(nil),         // 12: api.UpdateEventReply
	(*DeleteEventRequest)(nil),       // 13: api.DeleteEventRequest
	(*DeleteEventReply)(nil),         // 14: api.DeleteEventReply
	(*timestamppb.Timestamp)(nil),    // 15: google.protobuf.Timestamp
}
var file_calendar_proto_depIdxs = []int32{
	15, // 0: api.RangeTime.start:type_name -> google.protobuf.Timestamp
	15, // 1: api.RangeTime.end:type_name -> google.protobuf.Timestamp
	1,  // 2: api.Event.range_time:type_name -> api.RangeTime
	4,  // 3: api.Event.location:type_name -> api.Location
	2,  // 4: api.ListEvents.event:type_name -> api.Event
	15, // 5: api.ListEvents.time:type_name -> google.protobuf.Timestamp
	1,  // 6: api.CreateEventRequest.range_time:type_name -> api.RangeTime
	4,  // 7: api.CreateEventRequest.location:type_name -> api.Location
	1,  // 8: api.ListEventsRequest.range_time:type_name -> api.RangeTime
	0,  // 9: api.ListEventsRequest.date_type:type_name -> api.ListEventsRequest.DATE_TYPE
	3,  // 10: api.ListEventsReply.list_events:type_name -> api.ListEvents
	4,  // 11: api.GetRemindRequest.location:type_name -> api.Location
	15, // 12: api.GetRemindReply.time:type_name -> google.protobuf.Timestamp
	2,  // 13: api.UpdateEventRequest.event:type_name -> api.Event
	5,  // 14: api.CalendarService.CreateEvent:input_type -> api.CreateEventRequest
	11, // 15: api.CalendarService.UpdateEvent:input_type -> api.UpdateEventRequest
	13, // 16: api.CalendarService.DeleteEvent:input_type -> api.DeleteEventRequest
	7,  // 17: api.CalendarService.ListEvents:input_type -> api.ListEventsRequest
	9,  // 18: api.CalendarService.GetRemind:input_type -> api.GetRemindRequest
	6,  // 19: api.CalendarService.CreateEvent:output_type -> api.CreateEventReply
	12, // 20: api.CalendarService.UpdateEvent:output_type -> api.UpdateEventReply
	14, // 21: api.CalendarService.DeleteEvent:output_type -> api.DeleteEventReply
	8,  // 22: api.CalendarService.ListEvents:output_type -> api.ListEventsReply
	10, // 23: api.CalendarService.GetRemind:output_type -> api.GetRemindReply
	19, // [19:24] is the sub-list for method output_type
	14, // [14:19] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_calendar_proto_init() }
func file_calendar_proto_init() {
	if File_calendar_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calendar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeTime); i {
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
		file_calendar_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_calendar_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEvents); i {
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
		file_calendar_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_calendar_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_calendar_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventReply); i {
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
		file_calendar_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequest); i {
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
		file_calendar_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsReply); i {
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
		file_calendar_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemindRequest); i {
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
		file_calendar_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRemindReply); i {
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
		file_calendar_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
		file_calendar_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
		file_calendar_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
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
		file_calendar_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calendar_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calendar_proto_goTypes,
		DependencyIndexes: file_calendar_proto_depIdxs,
		EnumInfos:         file_calendar_proto_enumTypes,
		MessageInfos:      file_calendar_proto_msgTypes,
	}.Build()
	File_calendar_proto = out.File
	file_calendar_proto_rawDesc = nil
	file_calendar_proto_goTypes = nil
	file_calendar_proto_depIdxs = nil
}
