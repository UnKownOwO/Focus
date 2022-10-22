// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: RacingGalleryInfo.proto

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

type RacingGalleryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordList []*Unk2700_OJJNGIHDJEH `protobuf:"bytes,7,rep,name=record_list,json=recordList,proto3" json:"record_list,omitempty"`
}

func (x *RacingGalleryInfo) Reset() {
	*x = RacingGalleryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RacingGalleryInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RacingGalleryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RacingGalleryInfo) ProtoMessage() {}

func (x *RacingGalleryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RacingGalleryInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RacingGalleryInfo.ProtoReflect.Descriptor instead.
func (*RacingGalleryInfo) Descriptor() ([]byte, []int) {
	return file_RacingGalleryInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RacingGalleryInfo) GetRecordList() []*Unk2700_OJJNGIHDJEH {
	if x != nil {
		return x.RecordList
	}
	return nil
}

var File_RacingGalleryInfo_proto protoreflect.FileDescriptor

var file_RacingGalleryInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x4f, 0x4a, 0x4a, 0x4e, 0x47, 0x49, 0x48, 0x44, 0x4a, 0x45, 0x48, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x11, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x47, 0x61,
	0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x4a, 0x4a, 0x4e, 0x47, 0x49, 0x48,
	0x44, 0x4a, 0x45, 0x48, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_RacingGalleryInfo_proto_rawDescOnce sync.Once
	file_RacingGalleryInfo_proto_rawDescData = file_RacingGalleryInfo_proto_rawDesc
)

func file_RacingGalleryInfo_proto_rawDescGZIP() []byte {
	file_RacingGalleryInfo_proto_rawDescOnce.Do(func() {
		file_RacingGalleryInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RacingGalleryInfo_proto_rawDescData)
	})
	return file_RacingGalleryInfo_proto_rawDescData
}

var file_RacingGalleryInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RacingGalleryInfo_proto_goTypes = []interface{}{
	(*RacingGalleryInfo)(nil),   // 0: RacingGalleryInfo
	(*Unk2700_OJJNGIHDJEH)(nil), // 1: Unk2700_OJJNGIHDJEH
}
var file_RacingGalleryInfo_proto_depIdxs = []int32{
	1, // 0: RacingGalleryInfo.record_list:type_name -> Unk2700_OJJNGIHDJEH
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RacingGalleryInfo_proto_init() }
func file_RacingGalleryInfo_proto_init() {
	if File_RacingGalleryInfo_proto != nil {
		return
	}
	file_Unk2700_OJJNGIHDJEH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RacingGalleryInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RacingGalleryInfo); i {
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
			RawDescriptor: file_RacingGalleryInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RacingGalleryInfo_proto_goTypes,
		DependencyIndexes: file_RacingGalleryInfo_proto_depIdxs,
		MessageInfos:      file_RacingGalleryInfo_proto_msgTypes,
	}.Build()
	File_RacingGalleryInfo_proto = out.File
	file_RacingGalleryInfo_proto_rawDesc = nil
	file_RacingGalleryInfo_proto_goTypes = nil
	file_RacingGalleryInfo_proto_depIdxs = nil
}
