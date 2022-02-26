// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.2
// source: google/cloud/saasaccelerator/management/logs/v1/notification_service_payload.proto

package logs

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Types of Notification Status.
type NotificationStage_Stage int32

const (
	// Default.
	NotificationStage_STAGE_UNSPECIFIED NotificationStage_Stage = 0
	// Notification was sent.
	NotificationStage_SENT NotificationStage_Stage = 1
	// Notification failed to send.
	NotificationStage_SEND_FAILURE NotificationStage_Stage = 2
	// Notification was dropped.
	NotificationStage_DROPPED NotificationStage_Stage = 3
)

// Enum value maps for NotificationStage_Stage.
var (
	NotificationStage_Stage_name = map[int32]string{
		0: "STAGE_UNSPECIFIED",
		1: "SENT",
		2: "SEND_FAILURE",
		3: "DROPPED",
	}
	NotificationStage_Stage_value = map[string]int32{
		"STAGE_UNSPECIFIED": 0,
		"SENT":              1,
		"SEND_FAILURE":      2,
		"DROPPED":           3,
	}
)

func (x NotificationStage_Stage) Enum() *NotificationStage_Stage {
	p := new(NotificationStage_Stage)
	*p = x
	return p
}

func (x NotificationStage_Stage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationStage_Stage) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_enumTypes[0].Descriptor()
}

func (NotificationStage_Stage) Type() protoreflect.EnumType {
	return &file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_enumTypes[0]
}

func (x NotificationStage_Stage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationStage_Stage.Descriptor instead.
func (NotificationStage_Stage) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescGZIP(), []int{0, 0}
}

// Event that triggered the notification.
type NotificationStage_Event int32

const (
	// Default value.
	NotificationStage_EVENT_UNSPECIFIED NotificationStage_Event = 0
	// When a health status has been changed.
	NotificationStage_HEALTH_STATUS_CHANGE NotificationStage_Event = 1
)

// Enum value maps for NotificationStage_Event.
var (
	NotificationStage_Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "HEALTH_STATUS_CHANGE",
	}
	NotificationStage_Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":    0,
		"HEALTH_STATUS_CHANGE": 1,
	}
)

func (x NotificationStage_Event) Enum() *NotificationStage_Event {
	p := new(NotificationStage_Event)
	*p = x
	return p
}

func (x NotificationStage_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationStage_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_enumTypes[1].Descriptor()
}

func (NotificationStage_Event) Type() protoreflect.EnumType {
	return &file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_enumTypes[1]
}

func (x NotificationStage_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationStage_Event.Descriptor instead.
func (NotificationStage_Event) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescGZIP(), []int{0, 1}
}

// Payload proto for Notification logs.
type NotificationStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the Notification Service event.
	Stage NotificationStage_Stage `protobuf:"varint,1,opt,name=stage,proto3,enum=google.cloud.saasaccelerator.management.logs.v1.NotificationStage_Stage" json:"stage,omitempty"`
	// Time of the NotificationServiceEvent.
	EventTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The id of the notification.
	NotificationId string `protobuf:"bytes,3,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	// The event that triggered the notification.
	Event NotificationStage_Event `protobuf:"varint,4,opt,name=event,proto3,enum=google.cloud.saasaccelerator.management.logs.v1.NotificationStage_Event" json:"event,omitempty"`
	// Message to denote the error related to the event if applicable.
	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationStage) Reset() {
	*x = NotificationStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationStage) ProtoMessage() {}

func (x *NotificationStage) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationStage.ProtoReflect.Descriptor instead.
func (*NotificationStage) Descriptor() ([]byte, []int) {
	return file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationStage) GetStage() NotificationStage_Stage {
	if x != nil {
		return x.Stage
	}
	return NotificationStage_STAGE_UNSPECIFIED
}

func (x *NotificationStage) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *NotificationStage) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *NotificationStage) GetEvent() NotificationStage_Event {
	if x != nil {
		return x.Event
	}
	return NotificationStage_EVENT_UNSPECIFIED
}

func (x *NotificationStage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto protoreflect.FileDescriptor

var file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDesc = []byte{
	0x0a, 0x52, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x03, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x5e, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x61,
	0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x5e, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x52, 0x4f, 0x50, 0x50, 0x45,
	0x44, 0x10, 0x03, 0x22, 0x38, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x01, 0x42, 0xad, 0x01,
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f,
	0x67, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x1f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x61, 0x61, 0x73, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescOnce sync.Once
	file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescData = file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDesc
)

func file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescGZIP() []byte {
	file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescOnce.Do(func() {
		file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescData)
	})
	return file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDescData
}

var file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_goTypes = []interface{}{
	(NotificationStage_Stage)(0),  // 0: google.cloud.saasaccelerator.management.logs.v1.NotificationStage.Stage
	(NotificationStage_Event)(0),  // 1: google.cloud.saasaccelerator.management.logs.v1.NotificationStage.Event
	(*NotificationStage)(nil),     // 2: google.cloud.saasaccelerator.management.logs.v1.NotificationStage
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_depIdxs = []int32{
	0, // 0: google.cloud.saasaccelerator.management.logs.v1.NotificationStage.stage:type_name -> google.cloud.saasaccelerator.management.logs.v1.NotificationStage.Stage
	3, // 1: google.cloud.saasaccelerator.management.logs.v1.NotificationStage.event_time:type_name -> google.protobuf.Timestamp
	1, // 2: google.cloud.saasaccelerator.management.logs.v1.NotificationStage.event:type_name -> google.cloud.saasaccelerator.management.logs.v1.NotificationStage.Event
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_init()
}
func file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_init() {
	if File_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationStage); i {
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
			RawDescriptor: file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_goTypes,
		DependencyIndexes: file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_depIdxs,
		EnumInfos:         file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_enumTypes,
		MessageInfos:      file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_msgTypes,
	}.Build()
	File_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto = out.File
	file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_rawDesc = nil
	file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_goTypes = nil
	file_google_cloud_saasaccelerator_management_logs_v1_notification_service_payload_proto_depIdxs = nil
}
