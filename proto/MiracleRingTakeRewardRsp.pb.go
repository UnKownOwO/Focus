// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: MiracleRingTakeRewardRsp.proto

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

// CmdId: 5202
// EnetChannelId: 0
// EnetIsReliable: true
type MiracleRingTakeRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *MiracleRingTakeRewardRsp) Reset() {
	*x = MiracleRingTakeRewardRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MiracleRingTakeRewardRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiracleRingTakeRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingTakeRewardRsp) ProtoMessage() {}

func (x *MiracleRingTakeRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MiracleRingTakeRewardRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingTakeRewardRsp.ProtoReflect.Descriptor instead.
func (*MiracleRingTakeRewardRsp) Descriptor() ([]byte, []int) {
	return file_MiracleRingTakeRewardRsp_proto_rawDescGZIP(), []int{0}
}

func (x *MiracleRingTakeRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_MiracleRingTakeRewardRsp_proto protoreflect.FileDescriptor

var file_MiracleRingTakeRewardRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6b,
	0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x34, 0x0a, 0x18, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MiracleRingTakeRewardRsp_proto_rawDescOnce sync.Once
	file_MiracleRingTakeRewardRsp_proto_rawDescData = file_MiracleRingTakeRewardRsp_proto_rawDesc
)

func file_MiracleRingTakeRewardRsp_proto_rawDescGZIP() []byte {
	file_MiracleRingTakeRewardRsp_proto_rawDescOnce.Do(func() {
		file_MiracleRingTakeRewardRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MiracleRingTakeRewardRsp_proto_rawDescData)
	})
	return file_MiracleRingTakeRewardRsp_proto_rawDescData
}

var file_MiracleRingTakeRewardRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MiracleRingTakeRewardRsp_proto_goTypes = []interface{}{
	(*MiracleRingTakeRewardRsp)(nil), // 0: MiracleRingTakeRewardRsp
}
var file_MiracleRingTakeRewardRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MiracleRingTakeRewardRsp_proto_init() }
func file_MiracleRingTakeRewardRsp_proto_init() {
	if File_MiracleRingTakeRewardRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MiracleRingTakeRewardRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiracleRingTakeRewardRsp); i {
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
			RawDescriptor: file_MiracleRingTakeRewardRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MiracleRingTakeRewardRsp_proto_goTypes,
		DependencyIndexes: file_MiracleRingTakeRewardRsp_proto_depIdxs,
		MessageInfos:      file_MiracleRingTakeRewardRsp_proto_msgTypes,
	}.Build()
	File_MiracleRingTakeRewardRsp_proto = out.File
	file_MiracleRingTakeRewardRsp_proto_rawDesc = nil
	file_MiracleRingTakeRewardRsp_proto_goTypes = nil
	file_MiracleRingTakeRewardRsp_proto_depIdxs = nil
}
