// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/scg/shared/v1/events.proto

package sharedv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type EventEnvelope struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventId       string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventType     string                 `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	ActorUuid     string                 `protobuf:"bytes,3,opt,name=actor_uuid,json=actorUuid,proto3" json:"actor_uuid,omitempty"`
	TraceId       string                 `protobuf:"bytes,4,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	OccurredAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	SchemaVersion int32                  `protobuf:"varint,6,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventEnvelope) Reset() {
	*x = EventEnvelope{}
	mi := &file_proto_scg_shared_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventEnvelope) ProtoMessage() {}

func (x *EventEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scg_shared_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventEnvelope.ProtoReflect.Descriptor instead.
func (*EventEnvelope) Descriptor() ([]byte, []int) {
	return file_proto_scg_shared_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *EventEnvelope) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventEnvelope) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *EventEnvelope) GetActorUuid() string {
	if x != nil {
		return x.ActorUuid
	}
	return ""
}

func (x *EventEnvelope) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *EventEnvelope) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

func (x *EventEnvelope) GetSchemaVersion() int32 {
	if x != nil {
		return x.SchemaVersion
	}
	return 0
}

var File_proto_scg_shared_v1_events_proto protoreflect.FileDescriptor

const file_proto_scg_shared_v1_events_proto_rawDesc = "" +
	"\n" +
	" proto/scg/shared/v1/events.proto\x12\x13proto.scg.shared.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe7\x01\n" +
	"\rEventEnvelope\x12\x19\n" +
	"\bevent_id\x18\x01 \x01(\tR\aeventId\x12\x1d\n" +
	"\n" +
	"event_type\x18\x02 \x01(\tR\teventType\x12\x1d\n" +
	"\n" +
	"actor_uuid\x18\x03 \x01(\tR\tactorUuid\x12\x19\n" +
	"\btrace_id\x18\x04 \x01(\tR\atraceId\x12;\n" +
	"\voccurred_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"occurredAt\x12%\n" +
	"\x0eschema_version\x18\x06 \x01(\x05R\rschemaVersionBHZFgithub.com/hbttundar/scg-contracts/gen/go/proto/scg/shared/v1;sharedv1b\x06proto3"

var (
	file_proto_scg_shared_v1_events_proto_rawDescOnce sync.Once
	file_proto_scg_shared_v1_events_proto_rawDescData []byte
)

func file_proto_scg_shared_v1_events_proto_rawDescGZIP() []byte {
	file_proto_scg_shared_v1_events_proto_rawDescOnce.Do(func() {
		file_proto_scg_shared_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_events_proto_rawDesc), len(file_proto_scg_shared_v1_events_proto_rawDesc)))
	})
	return file_proto_scg_shared_v1_events_proto_rawDescData
}

var file_proto_scg_shared_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_scg_shared_v1_events_proto_goTypes = []any{
	(*EventEnvelope)(nil),         // 0: proto.scg.shared.v1.EventEnvelope
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_proto_scg_shared_v1_events_proto_depIdxs = []int32{
	1, // 0: proto.scg.shared.v1.EventEnvelope.occurred_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_scg_shared_v1_events_proto_init() }
func file_proto_scg_shared_v1_events_proto_init() {
	if File_proto_scg_shared_v1_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_scg_shared_v1_events_proto_rawDesc), len(file_proto_scg_shared_v1_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_scg_shared_v1_events_proto_goTypes,
		DependencyIndexes: file_proto_scg_shared_v1_events_proto_depIdxs,
		MessageInfos:      file_proto_scg_shared_v1_events_proto_msgTypes,
	}.Build()
	File_proto_scg_shared_v1_events_proto = out.File
	file_proto_scg_shared_v1_events_proto_goTypes = nil
	file_proto_scg_shared_v1_events_proto_depIdxs = nil
}
