// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HuntingFailNotify.proto

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

// CmdId: 4320
// EnetChannelId: 0
// EnetIsReliable: true
type HuntingFailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HuntingPair *HuntingPair `protobuf:"bytes,12,opt,name=hunting_pair,json=huntingPair,proto3" json:"hunting_pair,omitempty"`
}

func (x *HuntingFailNotify) Reset() {
	*x = HuntingFailNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HuntingFailNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuntingFailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuntingFailNotify) ProtoMessage() {}

func (x *HuntingFailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HuntingFailNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuntingFailNotify.ProtoReflect.Descriptor instead.
func (*HuntingFailNotify) Descriptor() ([]byte, []int) {
	return file_HuntingFailNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HuntingFailNotify) GetHuntingPair() *HuntingPair {
	if x != nil {
		return x.HuntingPair
	}
	return nil
}

var File_HuntingFailNotify_proto protoreflect.FileDescriptor

var file_HuntingFailNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x11,
	0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x2f, 0x0a, 0x0c, 0x68, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61, 0x69,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0b, 0x68, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x61,
	0x69, 0x72, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HuntingFailNotify_proto_rawDescOnce sync.Once
	file_HuntingFailNotify_proto_rawDescData = file_HuntingFailNotify_proto_rawDesc
)

func file_HuntingFailNotify_proto_rawDescGZIP() []byte {
	file_HuntingFailNotify_proto_rawDescOnce.Do(func() {
		file_HuntingFailNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HuntingFailNotify_proto_rawDescData)
	})
	return file_HuntingFailNotify_proto_rawDescData
}

var file_HuntingFailNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HuntingFailNotify_proto_goTypes = []interface{}{
	(*HuntingFailNotify)(nil), // 0: HuntingFailNotify
	(*HuntingPair)(nil),       // 1: HuntingPair
}
var file_HuntingFailNotify_proto_depIdxs = []int32{
	1, // 0: HuntingFailNotify.hunting_pair:type_name -> HuntingPair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HuntingFailNotify_proto_init() }
func file_HuntingFailNotify_proto_init() {
	if File_HuntingFailNotify_proto != nil {
		return
	}
	file_HuntingPair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HuntingFailNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuntingFailNotify); i {
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
			RawDescriptor: file_HuntingFailNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HuntingFailNotify_proto_goTypes,
		DependencyIndexes: file_HuntingFailNotify_proto_depIdxs,
		MessageInfos:      file_HuntingFailNotify_proto_msgTypes,
	}.Build()
	File_HuntingFailNotify_proto = out.File
	file_HuntingFailNotify_proto_rawDesc = nil
	file_HuntingFailNotify_proto_goTypes = nil
	file_HuntingFailNotify_proto_depIdxs = nil
}
