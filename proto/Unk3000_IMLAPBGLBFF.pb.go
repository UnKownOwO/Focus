// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk3000_IMLAPBGLBFF.proto

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

// CmdId: 1687
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type Unk3000_IMLAPBGLBFF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Unk3000_IMLAPBGLBFF) Reset() {
	*x = Unk3000_IMLAPBGLBFF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_IMLAPBGLBFF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_IMLAPBGLBFF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_IMLAPBGLBFF) ProtoMessage() {}

func (x *Unk3000_IMLAPBGLBFF) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_IMLAPBGLBFF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_IMLAPBGLBFF.ProtoReflect.Descriptor instead.
func (*Unk3000_IMLAPBGLBFF) Descriptor() ([]byte, []int) {
	return file_Unk3000_IMLAPBGLBFF_proto_rawDescGZIP(), []int{0}
}

var File_Unk3000_IMLAPBGLBFF_proto protoreflect.FileDescriptor

var file_Unk3000_IMLAPBGLBFF_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x49, 0x4d, 0x4c, 0x41, 0x50, 0x42,
	0x47, 0x4c, 0x42, 0x46, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x49, 0x4d, 0x4c, 0x41, 0x50, 0x42, 0x47, 0x4c, 0x42,
	0x46, 0x46, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_IMLAPBGLBFF_proto_rawDescOnce sync.Once
	file_Unk3000_IMLAPBGLBFF_proto_rawDescData = file_Unk3000_IMLAPBGLBFF_proto_rawDesc
)

func file_Unk3000_IMLAPBGLBFF_proto_rawDescGZIP() []byte {
	file_Unk3000_IMLAPBGLBFF_proto_rawDescOnce.Do(func() {
		file_Unk3000_IMLAPBGLBFF_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_IMLAPBGLBFF_proto_rawDescData)
	})
	return file_Unk3000_IMLAPBGLBFF_proto_rawDescData
}

var file_Unk3000_IMLAPBGLBFF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3000_IMLAPBGLBFF_proto_goTypes = []interface{}{
	(*Unk3000_IMLAPBGLBFF)(nil), // 0: Unk3000_IMLAPBGLBFF
}
var file_Unk3000_IMLAPBGLBFF_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk3000_IMLAPBGLBFF_proto_init() }
func file_Unk3000_IMLAPBGLBFF_proto_init() {
	if File_Unk3000_IMLAPBGLBFF_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_IMLAPBGLBFF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_IMLAPBGLBFF); i {
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
			RawDescriptor: file_Unk3000_IMLAPBGLBFF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_IMLAPBGLBFF_proto_goTypes,
		DependencyIndexes: file_Unk3000_IMLAPBGLBFF_proto_depIdxs,
		MessageInfos:      file_Unk3000_IMLAPBGLBFF_proto_msgTypes,
	}.Build()
	File_Unk3000_IMLAPBGLBFF_proto = out.File
	file_Unk3000_IMLAPBGLBFF_proto_rawDesc = nil
	file_Unk3000_IMLAPBGLBFF_proto_goTypes = nil
	file_Unk3000_IMLAPBGLBFF_proto_depIdxs = nil
}