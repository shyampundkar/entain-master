// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.19.4
// source: sports/sports.proto

package sports

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible values for filtering events are, visible, hidden or show all (default)
// enum approach is taken to detect value presence, which is not possible with boolean
// as it has false as default value. The service does not know if client wants hidden
// events or its just default false value. With show_all as default value, any other value
// than show_all indicates that it has been set by client.
type ListEventsRequestFilter_Visibility int32

const (
	ListEventsRequestFilter_SHOW_ALL ListEventsRequestFilter_Visibility = 0 // default
	ListEventsRequestFilter_VISIBLE  ListEventsRequestFilter_Visibility = 1
	ListEventsRequestFilter_HIDDEN   ListEventsRequestFilter_Visibility = 2
)

// Enum value maps for ListEventsRequestFilter_Visibility.
var (
	ListEventsRequestFilter_Visibility_name = map[int32]string{
		0: "SHOW_ALL",
		1: "VISIBLE",
		2: "HIDDEN",
	}
	ListEventsRequestFilter_Visibility_value = map[string]int32{
		"SHOW_ALL": 0,
		"VISIBLE":  1,
		"HIDDEN":   2,
	}
)

func (x ListEventsRequestFilter_Visibility) Enum() *ListEventsRequestFilter_Visibility {
	p := new(ListEventsRequestFilter_Visibility)
	*p = x
	return p
}

func (x ListEventsRequestFilter_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListEventsRequestFilter_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_sports_sports_proto_enumTypes[0].Descriptor()
}

func (ListEventsRequestFilter_Visibility) Type() protoreflect.EnumType {
	return &file_sports_sports_proto_enumTypes[0]
}

func (x ListEventsRequestFilter_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListEventsRequestFilter_Visibility.Descriptor instead.
func (ListEventsRequestFilter_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{2, 0}
}

type ListEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter  *ListEventsRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Orderby string                   `protobuf:"bytes,2,opt,name=orderby,proto3" json:"orderby,omitempty"`
}

func (x *ListEventsRequest) Reset() {
	*x = ListEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequest) ProtoMessage() {}

func (x *ListEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[0]
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
	return file_sports_sports_proto_rawDescGZIP(), []int{0}
}

func (x *ListEventsRequest) GetFilter() *ListEventsRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListEventsRequest) GetOrderby() string {
	if x != nil {
		return x.Orderby
	}
	return ""
}

// Response to ListEvents call.
type ListEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ListEventsResponse) Reset() {
	*x = ListEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsResponse) ProtoMessage() {}

func (x *ListEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsResponse.ProtoReflect.Descriptor instead.
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{1}
}

func (x *ListEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// Filter for listing events.
type ListEventsRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeetingIds []int64 `protobuf:"varint,1,rep,packed,name=meeting_ids,json=meetingIds,proto3" json:"meeting_ids,omitempty"`
	//optional event visibilty field
	OptionalEventVisibility ListEventsRequestFilter_Visibility `protobuf:"varint,2,opt,name=optional_event_visibility,json=optionalEventVisibility,proto3,enum=sports.ListEventsRequestFilter_Visibility" json:"optional_event_visibility,omitempty"`
}

func (x *ListEventsRequestFilter) Reset() {
	*x = ListEventsRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequestFilter) ProtoMessage() {}

func (x *ListEventsRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequestFilter.ProtoReflect.Descriptor instead.
func (*ListEventsRequestFilter) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventsRequestFilter) GetMeetingIds() []int64 {
	if x != nil {
		return x.MeetingIds
	}
	return nil
}

func (x *ListEventsRequestFilter) GetOptionalEventVisibility() ListEventsRequestFilter_Visibility {
	if x != nil {
		return x.OptionalEventVisibility
	}
	return ListEventsRequestFilter_SHOW_ALL
}

// A event resource.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the event.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// MeetingID represents a unique identifier for the events meeting.
	MeetingId int64 `protobuf:"varint,2,opt,name=meeting_id,json=meetingId,proto3" json:"meeting_id,omitempty"`
	// Name is the official name given to the event.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Number represents the number of the event.
	Number int64 `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	// Visible represents whether or not the event is visible.
	Visible bool `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
	// AdvertisedStartTime is the time the event is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
	// The status is simply, `OPEN` if advertised_start_time > current date else `CLOSED`
	Status string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[3]
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
	return file_sports_sports_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetMeetingId() int64 {
	if x != nil {
		return x.MeetingId
	}
	return 0
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Event) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Event) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

func (x *Event) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_sports_sports_proto protoreflect.FileDescriptor

var file_sports_sports_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x62, 0x79, 0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xd7, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x12, 0x66,
	0x0a, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x17, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x41, 0x4c, 0x4c,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x53, 0x49, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x02, 0x22, 0xe4, 0x01, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x32, 0x67, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x5e, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2f,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sports_sports_proto_rawDescOnce sync.Once
	file_sports_sports_proto_rawDescData = file_sports_sports_proto_rawDesc
)

func file_sports_sports_proto_rawDescGZIP() []byte {
	file_sports_sports_proto_rawDescOnce.Do(func() {
		file_sports_sports_proto_rawDescData = protoimpl.X.CompressGZIP(file_sports_sports_proto_rawDescData)
	})
	return file_sports_sports_proto_rawDescData
}

var file_sports_sports_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sports_sports_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sports_sports_proto_goTypes = []interface{}{
	(ListEventsRequestFilter_Visibility)(0), // 0: sports.ListEventsRequestFilter.Visibility
	(*ListEventsRequest)(nil),               // 1: sports.ListEventsRequest
	(*ListEventsResponse)(nil),              // 2: sports.ListEventsResponse
	(*ListEventsRequestFilter)(nil),         // 3: sports.ListEventsRequestFilter
	(*Event)(nil),                           // 4: sports.Event
	(*timestamppb.Timestamp)(nil),           // 5: google.protobuf.Timestamp
}
var file_sports_sports_proto_depIdxs = []int32{
	3, // 0: sports.ListEventsRequest.filter:type_name -> sports.ListEventsRequestFilter
	4, // 1: sports.ListEventsResponse.events:type_name -> sports.Event
	0, // 2: sports.ListEventsRequestFilter.optional_event_visibility:type_name -> sports.ListEventsRequestFilter.Visibility
	5, // 3: sports.Event.advertised_start_time:type_name -> google.protobuf.Timestamp
	1, // 4: sports.Sport.ListEvents:input_type -> sports.ListEventsRequest
	2, // 5: sports.Sport.ListEvents:output_type -> sports.ListEventsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sports_sports_proto_init() }
func file_sports_sports_proto_init() {
	if File_sports_sports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sports_sports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_sports_sports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsResponse); i {
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
		file_sports_sports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequestFilter); i {
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
		file_sports_sports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sports_sports_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sports_sports_proto_goTypes,
		DependencyIndexes: file_sports_sports_proto_depIdxs,
		EnumInfos:         file_sports_sports_proto_enumTypes,
		MessageInfos:      file_sports_sports_proto_msgTypes,
	}.Build()
	File_sports_sports_proto = out.File
	file_sports_sports_proto_rawDesc = nil
	file_sports_sports_proto_goTypes = nil
	file_sports_sports_proto_depIdxs = nil
}