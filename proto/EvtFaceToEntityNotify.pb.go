// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: EvtFaceToEntityNotify.proto

package proto

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

// CmdId: 303
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type EvtFaceToEntityNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceEntityId uint32      `protobuf:"varint,5,opt,name=face_entity_id,json=faceEntityId,proto3" json:"face_entity_id,omitempty"`
	ForwardType  ForwardType `protobuf:"varint,9,opt,name=forward_type,json=forwardType,proto3,enum=ForwardType" json:"forward_type,omitempty"`
	EntityId     uint32      `protobuf:"varint,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *EvtFaceToEntityNotify) Reset() {
	*x = EvtFaceToEntityNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvtFaceToEntityNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvtFaceToEntityNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvtFaceToEntityNotify) ProtoMessage() {}

func (x *EvtFaceToEntityNotify) ProtoReflect() protoreflect.Message {
	mi := &file_EvtFaceToEntityNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvtFaceToEntityNotify.ProtoReflect.Descriptor instead.
func (*EvtFaceToEntityNotify) Descriptor() ([]byte, []int) {
	return file_EvtFaceToEntityNotify_proto_rawDescGZIP(), []int{0}
}

func (x *EvtFaceToEntityNotify) GetFaceEntityId() uint32 {
	if x != nil {
		return x.FaceEntityId
	}
	return 0
}

func (x *EvtFaceToEntityNotify) GetForwardType() ForwardType {
	if x != nil {
		return x.ForwardType
	}
	return ForwardType_FORWARD_TYPE_LOCAL
}

func (x *EvtFaceToEntityNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_EvtFaceToEntityNotify_proto protoreflect.FileDescriptor

var file_EvtFaceToEntityNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x45, 0x76, 0x74, 0x46, 0x61, 0x63, 0x65, 0x54, 0x6f, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8b, 0x01, 0x0a, 0x15, 0x45, 0x76, 0x74, 0x46, 0x61, 0x63, 0x65, 0x54, 0x6f, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x66, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_EvtFaceToEntityNotify_proto_rawDescOnce sync.Once
	file_EvtFaceToEntityNotify_proto_rawDescData = file_EvtFaceToEntityNotify_proto_rawDesc
)

func file_EvtFaceToEntityNotify_proto_rawDescGZIP() []byte {
	file_EvtFaceToEntityNotify_proto_rawDescOnce.Do(func() {
		file_EvtFaceToEntityNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvtFaceToEntityNotify_proto_rawDescData)
	})
	return file_EvtFaceToEntityNotify_proto_rawDescData
}

var file_EvtFaceToEntityNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvtFaceToEntityNotify_proto_goTypes = []interface{}{
	(*EvtFaceToEntityNotify)(nil), // 0: EvtFaceToEntityNotify
	(ForwardType)(0),              // 1: ForwardType
}
var file_EvtFaceToEntityNotify_proto_depIdxs = []int32{
	1, // 0: EvtFaceToEntityNotify.forward_type:type_name -> ForwardType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EvtFaceToEntityNotify_proto_init() }
func file_EvtFaceToEntityNotify_proto_init() {
	if File_EvtFaceToEntityNotify_proto != nil {
		return
	}
	file_ForwardType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvtFaceToEntityNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvtFaceToEntityNotify); i {
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
			RawDescriptor: file_EvtFaceToEntityNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvtFaceToEntityNotify_proto_goTypes,
		DependencyIndexes: file_EvtFaceToEntityNotify_proto_depIdxs,
		MessageInfos:      file_EvtFaceToEntityNotify_proto_msgTypes,
	}.Build()
	File_EvtFaceToEntityNotify_proto = out.File
	file_EvtFaceToEntityNotify_proto_rawDesc = nil
	file_EvtFaceToEntityNotify_proto_goTypes = nil
	file_EvtFaceToEntityNotify_proto_depIdxs = nil
}
