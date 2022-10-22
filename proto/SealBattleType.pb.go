// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SealBattleType.proto

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

type SealBattleType int32

const (
	SealBattleType_SEAL_BATTLE_TYPE_KEEP_ALIVE    SealBattleType = 0
	SealBattleType_SEAL_BATTLE_TYPE_KILL_MONSTER  SealBattleType = 1
	SealBattleType_SEAL_BATTLE_TYPE_ENERGY_CHARGE SealBattleType = 2
)

// Enum value maps for SealBattleType.
var (
	SealBattleType_name = map[int32]string{
		0: "SEAL_BATTLE_TYPE_KEEP_ALIVE",
		1: "SEAL_BATTLE_TYPE_KILL_MONSTER",
		2: "SEAL_BATTLE_TYPE_ENERGY_CHARGE",
	}
	SealBattleType_value = map[string]int32{
		"SEAL_BATTLE_TYPE_KEEP_ALIVE":    0,
		"SEAL_BATTLE_TYPE_KILL_MONSTER":  1,
		"SEAL_BATTLE_TYPE_ENERGY_CHARGE": 2,
	}
)

func (x SealBattleType) Enum() *SealBattleType {
	p := new(SealBattleType)
	*p = x
	return p
}

func (x SealBattleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SealBattleType) Descriptor() protoreflect.EnumDescriptor {
	return file_SealBattleType_proto_enumTypes[0].Descriptor()
}

func (SealBattleType) Type() protoreflect.EnumType {
	return &file_SealBattleType_proto_enumTypes[0]
}

func (x SealBattleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SealBattleType.Descriptor instead.
func (SealBattleType) EnumDescriptor() ([]byte, []int) {
	return file_SealBattleType_proto_rawDescGZIP(), []int{0}
}

var File_SealBattleType_proto protoreflect.FileDescriptor

var file_SealBattleType_proto_rawDesc = []byte{
	0x0a, 0x14, 0x53, 0x65, 0x61, 0x6c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x78, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x6c, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x45, 0x41, 0x4c,
	0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x45, 0x45,
	0x50, 0x5f, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x45, 0x41,
	0x4c, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x49,
	0x4c, 0x4c, 0x5f, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e,
	0x53, 0x45, 0x41, 0x4c, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x4e, 0x45, 0x52, 0x47, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x10, 0x02,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SealBattleType_proto_rawDescOnce sync.Once
	file_SealBattleType_proto_rawDescData = file_SealBattleType_proto_rawDesc
)

func file_SealBattleType_proto_rawDescGZIP() []byte {
	file_SealBattleType_proto_rawDescOnce.Do(func() {
		file_SealBattleType_proto_rawDescData = protoimpl.X.CompressGZIP(file_SealBattleType_proto_rawDescData)
	})
	return file_SealBattleType_proto_rawDescData
}

var file_SealBattleType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SealBattleType_proto_goTypes = []interface{}{
	(SealBattleType)(0), // 0: SealBattleType
}
var file_SealBattleType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SealBattleType_proto_init() }
func file_SealBattleType_proto_init() {
	if File_SealBattleType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SealBattleType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SealBattleType_proto_goTypes,
		DependencyIndexes: file_SealBattleType_proto_depIdxs,
		EnumInfos:         file_SealBattleType_proto_enumTypes,
	}.Build()
	File_SealBattleType_proto = out.File
	file_SealBattleType_proto_rawDesc = nil
	file_SealBattleType_proto_goTypes = nil
	file_SealBattleType_proto_depIdxs = nil
}
