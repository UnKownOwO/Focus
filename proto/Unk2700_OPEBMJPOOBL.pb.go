// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_OPEBMJPOOBL.proto

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

type Unk2700_OPEBMJPOOBL int32

const (
	Unk2700_OPEBMJPOOBL_Unk2700_OPEBMJPOOBL_NONE                Unk2700_OPEBMJPOOBL = 0
	Unk2700_OPEBMJPOOBL_Unk2700_OPEBMJPOOBL_Unk2700_HONBFAOIDKK Unk2700_OPEBMJPOOBL = 1
)

// Enum value maps for Unk2700_OPEBMJPOOBL.
var (
	Unk2700_OPEBMJPOOBL_name = map[int32]string{
		0: "Unk2700_OPEBMJPOOBL_NONE",
		1: "Unk2700_OPEBMJPOOBL_Unk2700_HONBFAOIDKK",
	}
	Unk2700_OPEBMJPOOBL_value = map[string]int32{
		"Unk2700_OPEBMJPOOBL_NONE":                0,
		"Unk2700_OPEBMJPOOBL_Unk2700_HONBFAOIDKK": 1,
	}
)

func (x Unk2700_OPEBMJPOOBL) Enum() *Unk2700_OPEBMJPOOBL {
	p := new(Unk2700_OPEBMJPOOBL)
	*p = x
	return p
}

func (x Unk2700_OPEBMJPOOBL) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unk2700_OPEBMJPOOBL) Descriptor() protoreflect.EnumDescriptor {
	return file_Unk2700_OPEBMJPOOBL_proto_enumTypes[0].Descriptor()
}

func (Unk2700_OPEBMJPOOBL) Type() protoreflect.EnumType {
	return &file_Unk2700_OPEBMJPOOBL_proto_enumTypes[0]
}

func (x Unk2700_OPEBMJPOOBL) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unk2700_OPEBMJPOOBL.Descriptor instead.
func (Unk2700_OPEBMJPOOBL) EnumDescriptor() ([]byte, []int) {
	return file_Unk2700_OPEBMJPOOBL_proto_rawDescGZIP(), []int{0}
}

var File_Unk2700_OPEBMJPOOBL_proto protoreflect.FileDescriptor

var file_Unk2700_OPEBMJPOOBL_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x50, 0x45, 0x42, 0x4d, 0x4a,
	0x50, 0x4f, 0x4f, 0x42, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x60, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x50, 0x45, 0x42, 0x4d, 0x4a, 0x50, 0x4f, 0x4f,
	0x42, 0x4c, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x50,
	0x45, 0x42, 0x4d, 0x4a, 0x50, 0x4f, 0x4f, 0x42, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x50, 0x45, 0x42,
	0x4d, 0x4a, 0x50, 0x4f, 0x4f, 0x42, 0x4c, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x48, 0x4f, 0x4e, 0x42, 0x46, 0x41, 0x4f, 0x49, 0x44, 0x4b, 0x4b, 0x10, 0x01, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_OPEBMJPOOBL_proto_rawDescOnce sync.Once
	file_Unk2700_OPEBMJPOOBL_proto_rawDescData = file_Unk2700_OPEBMJPOOBL_proto_rawDesc
)

func file_Unk2700_OPEBMJPOOBL_proto_rawDescGZIP() []byte {
	file_Unk2700_OPEBMJPOOBL_proto_rawDescOnce.Do(func() {
		file_Unk2700_OPEBMJPOOBL_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_OPEBMJPOOBL_proto_rawDescData)
	})
	return file_Unk2700_OPEBMJPOOBL_proto_rawDescData
}

var file_Unk2700_OPEBMJPOOBL_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Unk2700_OPEBMJPOOBL_proto_goTypes = []interface{}{
	(Unk2700_OPEBMJPOOBL)(0), // 0: Unk2700_OPEBMJPOOBL
}
var file_Unk2700_OPEBMJPOOBL_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2700_OPEBMJPOOBL_proto_init() }
func file_Unk2700_OPEBMJPOOBL_proto_init() {
	if File_Unk2700_OPEBMJPOOBL_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Unk2700_OPEBMJPOOBL_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_OPEBMJPOOBL_proto_goTypes,
		DependencyIndexes: file_Unk2700_OPEBMJPOOBL_proto_depIdxs,
		EnumInfos:         file_Unk2700_OPEBMJPOOBL_proto_enumTypes,
	}.Build()
	File_Unk2700_OPEBMJPOOBL_proto = out.File
	file_Unk2700_OPEBMJPOOBL_proto_rawDesc = nil
	file_Unk2700_OPEBMJPOOBL_proto_goTypes = nil
	file_Unk2700_OPEBMJPOOBL_proto_depIdxs = nil
}
