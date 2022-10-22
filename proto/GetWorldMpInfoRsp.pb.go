// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GetWorldMpInfoRsp.proto

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

// CmdId: 3320
// EnetChannelId: 0
// EnetIsReliable: true
type GetWorldMpInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode         int32  `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IsInMpMode      bool   `protobuf:"varint,1,opt,name=is_in_mp_mode,json=isInMpMode,proto3" json:"is_in_mp_mode,omitempty"`
	QuitMpValidTime uint32 `protobuf:"varint,9,opt,name=quit_mp_valid_time,json=quitMpValidTime,proto3" json:"quit_mp_valid_time,omitempty"`
}

func (x *GetWorldMpInfoRsp) Reset() {
	*x = GetWorldMpInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetWorldMpInfoRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorldMpInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorldMpInfoRsp) ProtoMessage() {}

func (x *GetWorldMpInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetWorldMpInfoRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorldMpInfoRsp.ProtoReflect.Descriptor instead.
func (*GetWorldMpInfoRsp) Descriptor() ([]byte, []int) {
	return file_GetWorldMpInfoRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetWorldMpInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetWorldMpInfoRsp) GetIsInMpMode() bool {
	if x != nil {
		return x.IsInMpMode
	}
	return false
}

func (x *GetWorldMpInfoRsp) GetQuitMpValidTime() uint32 {
	if x != nil {
		return x.QuitMpValidTime
	}
	return 0
}

var File_GetWorldMpInfoRsp_proto protoreflect.FileDescriptor

var file_GetWorldMpInfoRsp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4d, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4d, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x69,
	0x6e, 0x5f, 0x6d, 0x70, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x49, 0x6e, 0x4d, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x12, 0x71,
	0x75, 0x69, 0x74, 0x5f, 0x6d, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x71, 0x75, 0x69, 0x74, 0x4d, 0x70, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetWorldMpInfoRsp_proto_rawDescOnce sync.Once
	file_GetWorldMpInfoRsp_proto_rawDescData = file_GetWorldMpInfoRsp_proto_rawDesc
)

func file_GetWorldMpInfoRsp_proto_rawDescGZIP() []byte {
	file_GetWorldMpInfoRsp_proto_rawDescOnce.Do(func() {
		file_GetWorldMpInfoRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetWorldMpInfoRsp_proto_rawDescData)
	})
	return file_GetWorldMpInfoRsp_proto_rawDescData
}

var file_GetWorldMpInfoRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetWorldMpInfoRsp_proto_goTypes = []interface{}{
	(*GetWorldMpInfoRsp)(nil), // 0: GetWorldMpInfoRsp
}
var file_GetWorldMpInfoRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetWorldMpInfoRsp_proto_init() }
func file_GetWorldMpInfoRsp_proto_init() {
	if File_GetWorldMpInfoRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetWorldMpInfoRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorldMpInfoRsp); i {
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
			RawDescriptor: file_GetWorldMpInfoRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetWorldMpInfoRsp_proto_goTypes,
		DependencyIndexes: file_GetWorldMpInfoRsp_proto_depIdxs,
		MessageInfos:      file_GetWorldMpInfoRsp_proto_msgTypes,
	}.Build()
	File_GetWorldMpInfoRsp_proto = out.File
	file_GetWorldMpInfoRsp_proto_rawDesc = nil
	file_GetWorldMpInfoRsp_proto_goTypes = nil
	file_GetWorldMpInfoRsp_proto_depIdxs = nil
}
