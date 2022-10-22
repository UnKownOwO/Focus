// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SetPlayerSignatureRsp.proto

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

// CmdId: 4005
// EnetChannelId: 0
// EnetIsReliable: true
type SetPlayerSignatureRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Retcode   int32  `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *SetPlayerSignatureRsp) Reset() {
	*x = SetPlayerSignatureRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetPlayerSignatureRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPlayerSignatureRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPlayerSignatureRsp) ProtoMessage() {}

func (x *SetPlayerSignatureRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetPlayerSignatureRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPlayerSignatureRsp.ProtoReflect.Descriptor instead.
func (*SetPlayerSignatureRsp) Descriptor() ([]byte, []int) {
	return file_SetPlayerSignatureRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetPlayerSignatureRsp) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SetPlayerSignatureRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_SetPlayerSignatureRsp_proto protoreflect.FileDescriptor

var file_SetPlayerSignatureRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a,
	0x15, 0x53, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_SetPlayerSignatureRsp_proto_rawDescOnce sync.Once
	file_SetPlayerSignatureRsp_proto_rawDescData = file_SetPlayerSignatureRsp_proto_rawDesc
)

func file_SetPlayerSignatureRsp_proto_rawDescGZIP() []byte {
	file_SetPlayerSignatureRsp_proto_rawDescOnce.Do(func() {
		file_SetPlayerSignatureRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetPlayerSignatureRsp_proto_rawDescData)
	})
	return file_SetPlayerSignatureRsp_proto_rawDescData
}

var file_SetPlayerSignatureRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetPlayerSignatureRsp_proto_goTypes = []interface{}{
	(*SetPlayerSignatureRsp)(nil), // 0: SetPlayerSignatureRsp
}
var file_SetPlayerSignatureRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetPlayerSignatureRsp_proto_init() }
func file_SetPlayerSignatureRsp_proto_init() {
	if File_SetPlayerSignatureRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetPlayerSignatureRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPlayerSignatureRsp); i {
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
			RawDescriptor: file_SetPlayerSignatureRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetPlayerSignatureRsp_proto_goTypes,
		DependencyIndexes: file_SetPlayerSignatureRsp_proto_depIdxs,
		MessageInfos:      file_SetPlayerSignatureRsp_proto_msgTypes,
	}.Build()
	File_SetPlayerSignatureRsp_proto = out.File
	file_SetPlayerSignatureRsp_proto_rawDesc = nil
	file_SetPlayerSignatureRsp_proto_goTypes = nil
	file_SetPlayerSignatureRsp_proto_depIdxs = nil
}
