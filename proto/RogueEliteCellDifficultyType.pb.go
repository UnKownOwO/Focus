// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: RogueEliteCellDifficultyType.proto

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

type RogueEliteCellDifficultyType int32

const (
	RogueEliteCellDifficultyType_ROGUE_ELITE_CELL_DIFFICULTY_TYPE_NORMAL RogueEliteCellDifficultyType = 0
	RogueEliteCellDifficultyType_ROGUE_ELITE_CELL_DIFFICULTY_TYPE_HARD   RogueEliteCellDifficultyType = 1
)

// Enum value maps for RogueEliteCellDifficultyType.
var (
	RogueEliteCellDifficultyType_name = map[int32]string{
		0: "ROGUE_ELITE_CELL_DIFFICULTY_TYPE_NORMAL",
		1: "ROGUE_ELITE_CELL_DIFFICULTY_TYPE_HARD",
	}
	RogueEliteCellDifficultyType_value = map[string]int32{
		"ROGUE_ELITE_CELL_DIFFICULTY_TYPE_NORMAL": 0,
		"ROGUE_ELITE_CELL_DIFFICULTY_TYPE_HARD":   1,
	}
)

func (x RogueEliteCellDifficultyType) Enum() *RogueEliteCellDifficultyType {
	p := new(RogueEliteCellDifficultyType)
	*p = x
	return p
}

func (x RogueEliteCellDifficultyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueEliteCellDifficultyType) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueEliteCellDifficultyType_proto_enumTypes[0].Descriptor()
}

func (RogueEliteCellDifficultyType) Type() protoreflect.EnumType {
	return &file_RogueEliteCellDifficultyType_proto_enumTypes[0]
}

func (x RogueEliteCellDifficultyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueEliteCellDifficultyType.Descriptor instead.
func (RogueEliteCellDifficultyType) EnumDescriptor() ([]byte, []int) {
	return file_RogueEliteCellDifficultyType_proto_rawDescGZIP(), []int{0}
}

var File_RogueEliteCellDifficultyType_proto protoreflect.FileDescriptor

var file_RogueEliteCellDifficultyType_proto_rawDesc = []byte{
	0x0a, 0x22, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6c, 0x69, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c,
	0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x76, 0x0a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6c, 0x69,
	0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x27, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x45, 0x4c,
	0x49, 0x54, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x44, 0x49, 0x46, 0x46, 0x49, 0x43, 0x55,
	0x4c, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x00, 0x12, 0x29, 0x0a, 0x25, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x45, 0x4c, 0x49, 0x54, 0x45,
	0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x5f, 0x44, 0x49, 0x46, 0x46, 0x49, 0x43, 0x55, 0x4c, 0x54, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x41, 0x52, 0x44, 0x10, 0x01, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueEliteCellDifficultyType_proto_rawDescOnce sync.Once
	file_RogueEliteCellDifficultyType_proto_rawDescData = file_RogueEliteCellDifficultyType_proto_rawDesc
)

func file_RogueEliteCellDifficultyType_proto_rawDescGZIP() []byte {
	file_RogueEliteCellDifficultyType_proto_rawDescOnce.Do(func() {
		file_RogueEliteCellDifficultyType_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueEliteCellDifficultyType_proto_rawDescData)
	})
	return file_RogueEliteCellDifficultyType_proto_rawDescData
}

var file_RogueEliteCellDifficultyType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueEliteCellDifficultyType_proto_goTypes = []interface{}{
	(RogueEliteCellDifficultyType)(0), // 0: RogueEliteCellDifficultyType
}
var file_RogueEliteCellDifficultyType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueEliteCellDifficultyType_proto_init() }
func file_RogueEliteCellDifficultyType_proto_init() {
	if File_RogueEliteCellDifficultyType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueEliteCellDifficultyType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueEliteCellDifficultyType_proto_goTypes,
		DependencyIndexes: file_RogueEliteCellDifficultyType_proto_depIdxs,
		EnumInfos:         file_RogueEliteCellDifficultyType_proto_enumTypes,
	}.Build()
	File_RogueEliteCellDifficultyType_proto = out.File
	file_RogueEliteCellDifficultyType_proto_rawDesc = nil
	file_RogueEliteCellDifficultyType_proto_goTypes = nil
	file_RogueEliteCellDifficultyType_proto_depIdxs = nil
}
