// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: PlayerDieOption.proto

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

type PlayerDieOption int32

const (
	PlayerDieOption_PLAYER_DIE_OPTION_OPT_NONE   PlayerDieOption = 0
	PlayerDieOption_PLAYER_DIE_OPTION_OPT_REPLAY PlayerDieOption = 1
	PlayerDieOption_PLAYER_DIE_OPTION_OPT_CANCEL PlayerDieOption = 2
	PlayerDieOption_PLAYER_DIE_OPTION_OPT_REVIVE PlayerDieOption = 3
)

// Enum value maps for PlayerDieOption.
var (
	PlayerDieOption_name = map[int32]string{
		0: "PLAYER_DIE_OPTION_OPT_NONE",
		1: "PLAYER_DIE_OPTION_OPT_REPLAY",
		2: "PLAYER_DIE_OPTION_OPT_CANCEL",
		3: "PLAYER_DIE_OPTION_OPT_REVIVE",
	}
	PlayerDieOption_value = map[string]int32{
		"PLAYER_DIE_OPTION_OPT_NONE":   0,
		"PLAYER_DIE_OPTION_OPT_REPLAY": 1,
		"PLAYER_DIE_OPTION_OPT_CANCEL": 2,
		"PLAYER_DIE_OPTION_OPT_REVIVE": 3,
	}
)

func (x PlayerDieOption) Enum() *PlayerDieOption {
	p := new(PlayerDieOption)
	*p = x
	return p
}

func (x PlayerDieOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerDieOption) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerDieOption_proto_enumTypes[0].Descriptor()
}

func (PlayerDieOption) Type() protoreflect.EnumType {
	return &file_PlayerDieOption_proto_enumTypes[0]
}

func (x PlayerDieOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerDieOption.Descriptor instead.
func (PlayerDieOption) EnumDescriptor() ([]byte, []int) {
	return file_PlayerDieOption_proto_rawDescGZIP(), []int{0}
}

var File_PlayerDieOption_proto protoreflect.FileDescriptor

var file_PlayerDieOption_proto_rawDesc = []byte{
	0x0a, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x97, 0x01, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x69, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a, 0x50,
	0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x50,
	0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x01, 0x12, 0x20, 0x0a,
	0x1c, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x4f, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x02, 0x12,
	0x20, 0x0a, 0x1c, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x4f, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x56, 0x45, 0x10,
	0x03, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerDieOption_proto_rawDescOnce sync.Once
	file_PlayerDieOption_proto_rawDescData = file_PlayerDieOption_proto_rawDesc
)

func file_PlayerDieOption_proto_rawDescGZIP() []byte {
	file_PlayerDieOption_proto_rawDescOnce.Do(func() {
		file_PlayerDieOption_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerDieOption_proto_rawDescData)
	})
	return file_PlayerDieOption_proto_rawDescData
}

var file_PlayerDieOption_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerDieOption_proto_goTypes = []interface{}{
	(PlayerDieOption)(0), // 0: PlayerDieOption
}
var file_PlayerDieOption_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerDieOption_proto_init() }
func file_PlayerDieOption_proto_init() {
	if File_PlayerDieOption_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlayerDieOption_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerDieOption_proto_goTypes,
		DependencyIndexes: file_PlayerDieOption_proto_depIdxs,
		EnumInfos:         file_PlayerDieOption_proto_enumTypes,
	}.Build()
	File_PlayerDieOption_proto = out.File
	file_PlayerDieOption_proto_rawDesc = nil
	file_PlayerDieOption_proto_goTypes = nil
	file_PlayerDieOption_proto_depIdxs = nil
}
