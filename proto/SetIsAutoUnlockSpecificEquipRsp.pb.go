// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SetIsAutoUnlockSpecificEquipRsp.proto

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

// CmdId: 664
// EnetChannelId: 0
// EnetIsReliable: true
type SetIsAutoUnlockSpecificEquipRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SetIsAutoUnlockSpecificEquipRsp) Reset() {
	*x = SetIsAutoUnlockSpecificEquipRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetIsAutoUnlockSpecificEquipRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetIsAutoUnlockSpecificEquipRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIsAutoUnlockSpecificEquipRsp) ProtoMessage() {}

func (x *SetIsAutoUnlockSpecificEquipRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetIsAutoUnlockSpecificEquipRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIsAutoUnlockSpecificEquipRsp.ProtoReflect.Descriptor instead.
func (*SetIsAutoUnlockSpecificEquipRsp) Descriptor() ([]byte, []int) {
	return file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetIsAutoUnlockSpecificEquipRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SetIsAutoUnlockSpecificEquipRsp_proto protoreflect.FileDescriptor

var file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDesc = []byte{
	0x0a, 0x25, 0x53, 0x65, 0x74, 0x49, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x1f, 0x53, 0x65, 0x74, 0x49, 0x73,
	0x41, 0x75, 0x74, 0x6f, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescOnce sync.Once
	file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescData = file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDesc
)

func file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescGZIP() []byte {
	file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescOnce.Do(func() {
		file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescData)
	})
	return file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDescData
}

var file_SetIsAutoUnlockSpecificEquipRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetIsAutoUnlockSpecificEquipRsp_proto_goTypes = []interface{}{
	(*SetIsAutoUnlockSpecificEquipRsp)(nil), // 0: SetIsAutoUnlockSpecificEquipRsp
}
var file_SetIsAutoUnlockSpecificEquipRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetIsAutoUnlockSpecificEquipRsp_proto_init() }
func file_SetIsAutoUnlockSpecificEquipRsp_proto_init() {
	if File_SetIsAutoUnlockSpecificEquipRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetIsAutoUnlockSpecificEquipRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetIsAutoUnlockSpecificEquipRsp); i {
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
			RawDescriptor: file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetIsAutoUnlockSpecificEquipRsp_proto_goTypes,
		DependencyIndexes: file_SetIsAutoUnlockSpecificEquipRsp_proto_depIdxs,
		MessageInfos:      file_SetIsAutoUnlockSpecificEquipRsp_proto_msgTypes,
	}.Build()
	File_SetIsAutoUnlockSpecificEquipRsp_proto = out.File
	file_SetIsAutoUnlockSpecificEquipRsp_proto_rawDesc = nil
	file_SetIsAutoUnlockSpecificEquipRsp_proto_goTypes = nil
	file_SetIsAutoUnlockSpecificEquipRsp_proto_depIdxs = nil
}