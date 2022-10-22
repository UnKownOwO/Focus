// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AbilityApplyLevelModifier.proto

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

type AbilityApplyLevelModifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyEntityId uint32 `protobuf:"varint,6,opt,name=apply_entity_id,json=applyEntityId,proto3" json:"apply_entity_id,omitempty"`
}

func (x *AbilityApplyLevelModifier) Reset() {
	*x = AbilityApplyLevelModifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityApplyLevelModifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityApplyLevelModifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityApplyLevelModifier) ProtoMessage() {}

func (x *AbilityApplyLevelModifier) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityApplyLevelModifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityApplyLevelModifier.ProtoReflect.Descriptor instead.
func (*AbilityApplyLevelModifier) Descriptor() ([]byte, []int) {
	return file_AbilityApplyLevelModifier_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityApplyLevelModifier) GetApplyEntityId() uint32 {
	if x != nil {
		return x.ApplyEntityId
	}
	return 0
}

var File_AbilityApplyLevelModifier_proto protoreflect.FileDescriptor

var file_AbilityApplyLevelModifier_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x43, 0x0a, 0x19, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityApplyLevelModifier_proto_rawDescOnce sync.Once
	file_AbilityApplyLevelModifier_proto_rawDescData = file_AbilityApplyLevelModifier_proto_rawDesc
)

func file_AbilityApplyLevelModifier_proto_rawDescGZIP() []byte {
	file_AbilityApplyLevelModifier_proto_rawDescOnce.Do(func() {
		file_AbilityApplyLevelModifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityApplyLevelModifier_proto_rawDescData)
	})
	return file_AbilityApplyLevelModifier_proto_rawDescData
}

var file_AbilityApplyLevelModifier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityApplyLevelModifier_proto_goTypes = []interface{}{
	(*AbilityApplyLevelModifier)(nil), // 0: AbilityApplyLevelModifier
}
var file_AbilityApplyLevelModifier_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AbilityApplyLevelModifier_proto_init() }
func file_AbilityApplyLevelModifier_proto_init() {
	if File_AbilityApplyLevelModifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AbilityApplyLevelModifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityApplyLevelModifier); i {
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
			RawDescriptor: file_AbilityApplyLevelModifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityApplyLevelModifier_proto_goTypes,
		DependencyIndexes: file_AbilityApplyLevelModifier_proto_depIdxs,
		MessageInfos:      file_AbilityApplyLevelModifier_proto_msgTypes,
	}.Build()
	File_AbilityApplyLevelModifier_proto = out.File
	file_AbilityApplyLevelModifier_proto_rawDesc = nil
	file_AbilityApplyLevelModifier_proto_goTypes = nil
	file_AbilityApplyLevelModifier_proto_depIdxs = nil
}
