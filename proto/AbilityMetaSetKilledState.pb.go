// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AbilityMetaSetKilledState.proto

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

type AbilityMetaSetKilledState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Killed bool `protobuf:"varint,2,opt,name=killed,proto3" json:"killed,omitempty"`
}

func (x *AbilityMetaSetKilledState) Reset() {
	*x = AbilityMetaSetKilledState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityMetaSetKilledState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityMetaSetKilledState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityMetaSetKilledState) ProtoMessage() {}

func (x *AbilityMetaSetKilledState) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityMetaSetKilledState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityMetaSetKilledState.ProtoReflect.Descriptor instead.
func (*AbilityMetaSetKilledState) Descriptor() ([]byte, []int) {
	return file_AbilityMetaSetKilledState_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityMetaSetKilledState) GetKilled() bool {
	if x != nil {
		return x.Killed
	}
	return false
}

var File_AbilityMetaSetKilledState_proto protoreflect.FileDescriptor

var file_AbilityMetaSetKilledState_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x65, 0x74,
	0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x33, 0x0a, 0x19, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x53, 0x65, 0x74, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x6b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AbilityMetaSetKilledState_proto_rawDescOnce sync.Once
	file_AbilityMetaSetKilledState_proto_rawDescData = file_AbilityMetaSetKilledState_proto_rawDesc
)

func file_AbilityMetaSetKilledState_proto_rawDescGZIP() []byte {
	file_AbilityMetaSetKilledState_proto_rawDescOnce.Do(func() {
		file_AbilityMetaSetKilledState_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityMetaSetKilledState_proto_rawDescData)
	})
	return file_AbilityMetaSetKilledState_proto_rawDescData
}

var file_AbilityMetaSetKilledState_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityMetaSetKilledState_proto_goTypes = []interface{}{
	(*AbilityMetaSetKilledState)(nil), // 0: AbilityMetaSetKilledState
}
var file_AbilityMetaSetKilledState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AbilityMetaSetKilledState_proto_init() }
func file_AbilityMetaSetKilledState_proto_init() {
	if File_AbilityMetaSetKilledState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AbilityMetaSetKilledState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityMetaSetKilledState); i {
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
			RawDescriptor: file_AbilityMetaSetKilledState_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityMetaSetKilledState_proto_goTypes,
		DependencyIndexes: file_AbilityMetaSetKilledState_proto_depIdxs,
		MessageInfos:      file_AbilityMetaSetKilledState_proto_msgTypes,
	}.Build()
	File_AbilityMetaSetKilledState_proto = out.File
	file_AbilityMetaSetKilledState_proto_rawDesc = nil
	file_AbilityMetaSetKilledState_proto_goTypes = nil
	file_AbilityMetaSetKilledState_proto_depIdxs = nil
}
