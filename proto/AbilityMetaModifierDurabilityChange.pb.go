// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AbilityMetaModifierDurabilityChange.proto

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

type AbilityMetaModifierDurabilityChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReduceDurability float32 `protobuf:"fixed32,6,opt,name=reduce_durability,json=reduceDurability,proto3" json:"reduce_durability,omitempty"`
	RemainDurability float32 `protobuf:"fixed32,15,opt,name=remain_durability,json=remainDurability,proto3" json:"remain_durability,omitempty"`
}

func (x *AbilityMetaModifierDurabilityChange) Reset() {
	*x = AbilityMetaModifierDurabilityChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityMetaModifierDurabilityChange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMetaModifierDurabilityChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMetaModifierDurabilityChange) ProtoMessage() {}

func (x *AbilityMetaModifierDurabilityChange) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityMetaModifierDurabilityChange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMetaModifierDurabilityChange.ProtoReflect.Descriptor instead.
func (*AbilityMetaModifierDurabilityChange) Descriptor() ([]byte, []int) {
	return file_AbilityMetaModifierDurabilityChange_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityMetaModifierDurabilityChange) GetReduceDurability() float32 {
	if x != nil {
		return x.ReduceDurability
	}
	return 0
}

func (x *AbilityMetaModifierDurabilityChange) GetRemainDurability() float32 {
	if x != nil {
		return x.RemainDurability
	}
	return 0
}

var File_AbilityMetaModifierDurabilityChange_proto protoreflect.FileDescriptor

var file_AbilityMetaModifierDurabilityChange_proto_rawDesc = []byte{
	0x0a, 0x29, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x23, 0x41,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x64, 0x75, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x72,
	0x65, 0x64, 0x75, 0x63, 0x65, 0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x2b, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x72, 0x65, 0x6d, 0x61,
	0x69, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityMetaModifierDurabilityChange_proto_rawDescOnce sync.Once
	file_AbilityMetaModifierDurabilityChange_proto_rawDescData = file_AbilityMetaModifierDurabilityChange_proto_rawDesc
)

func file_AbilityMetaModifierDurabilityChange_proto_rawDescGZIP() []byte {
	file_AbilityMetaModifierDurabilityChange_proto_rawDescOnce.Do(func() {
		file_AbilityMetaModifierDurabilityChange_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityMetaModifierDurabilityChange_proto_rawDescData)
	})
	return file_AbilityMetaModifierDurabilityChange_proto_rawDescData
}

var file_AbilityMetaModifierDurabilityChange_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityMetaModifierDurabilityChange_proto_goTypes = []interface{}{
	(*AbilityMetaModifierDurabilityChange)(nil), // 0: AbilityMetaModifierDurabilityChange
}
var file_AbilityMetaModifierDurabilityChange_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AbilityMetaModifierDurabilityChange_proto_init() }
func file_AbilityMetaModifierDurabilityChange_proto_init() {
	if File_AbilityMetaModifierDurabilityChange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AbilityMetaModifierDurabilityChange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMetaModifierDurabilityChange); i {
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
			RawDescriptor: file_AbilityMetaModifierDurabilityChange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityMetaModifierDurabilityChange_proto_goTypes,
		DependencyIndexes: file_AbilityMetaModifierDurabilityChange_proto_depIdxs,
		MessageInfos:      file_AbilityMetaModifierDurabilityChange_proto_msgTypes,
	}.Build()
	File_AbilityMetaModifierDurabilityChange_proto = out.File
	file_AbilityMetaModifierDurabilityChange_proto_rawDesc = nil
	file_AbilityMetaModifierDurabilityChange_proto_goTypes = nil
	file_AbilityMetaModifierDurabilityChange_proto_depIdxs = nil
}