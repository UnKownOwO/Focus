// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: InstableSprayDetailInfo.proto

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

type InstableSprayDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_PHKHIPLDOOA []*Unk3000_ICLKJJNGOHN `protobuf:"bytes,9,rep,name=Unk2700_PHKHIPLDOOA,json=Unk2700PHKHIPLDOOA,proto3" json:"Unk2700_PHKHIPLDOOA,omitempty"`
}

func (x *InstableSprayDetailInfo) Reset() {
	*x = InstableSprayDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InstableSprayDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstableSprayDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstableSprayDetailInfo) ProtoMessage() {}

func (x *InstableSprayDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InstableSprayDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstableSprayDetailInfo.ProtoReflect.Descriptor instead.
func (*InstableSprayDetailInfo) Descriptor() ([]byte, []int) {
	return file_InstableSprayDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InstableSprayDetailInfo) GetUnk2700_PHKHIPLDOOA() []*Unk3000_ICLKJJNGOHN {
	if x != nil {
		return x.Unk2700_PHKHIPLDOOA
	}
	return nil
}

var File_InstableSprayDetailInfo_proto protoreflect.FileDescriptor

var file_InstableSprayDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x72, 0x61, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x49, 0x43, 0x4c, 0x4b, 0x4a, 0x4a, 0x4e,
	0x47, 0x4f, 0x48, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x17, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x72, 0x61, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x50, 0x48, 0x4b, 0x48, 0x49, 0x50, 0x4c, 0x44, 0x4f, 0x4f, 0x41, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x49, 0x43, 0x4c,
	0x4b, 0x4a, 0x4a, 0x4e, 0x47, 0x4f, 0x48, 0x4e, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x50, 0x48, 0x4b, 0x48, 0x49, 0x50, 0x4c, 0x44, 0x4f, 0x4f, 0x41, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InstableSprayDetailInfo_proto_rawDescOnce sync.Once
	file_InstableSprayDetailInfo_proto_rawDescData = file_InstableSprayDetailInfo_proto_rawDesc
)

func file_InstableSprayDetailInfo_proto_rawDescGZIP() []byte {
	file_InstableSprayDetailInfo_proto_rawDescOnce.Do(func() {
		file_InstableSprayDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InstableSprayDetailInfo_proto_rawDescData)
	})
	return file_InstableSprayDetailInfo_proto_rawDescData
}

var file_InstableSprayDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InstableSprayDetailInfo_proto_goTypes = []interface{}{
	(*InstableSprayDetailInfo)(nil), // 0: InstableSprayDetailInfo
	(*Unk3000_ICLKJJNGOHN)(nil),     // 1: Unk3000_ICLKJJNGOHN
}
var file_InstableSprayDetailInfo_proto_depIdxs = []int32{
	1, // 0: InstableSprayDetailInfo.Unk2700_PHKHIPLDOOA:type_name -> Unk3000_ICLKJJNGOHN
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InstableSprayDetailInfo_proto_init() }
func file_InstableSprayDetailInfo_proto_init() {
	if File_InstableSprayDetailInfo_proto != nil {
		return
	}
	file_Unk3000_ICLKJJNGOHN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InstableSprayDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstableSprayDetailInfo); i {
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
			RawDescriptor: file_InstableSprayDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InstableSprayDetailInfo_proto_goTypes,
		DependencyIndexes: file_InstableSprayDetailInfo_proto_depIdxs,
		MessageInfos:      file_InstableSprayDetailInfo_proto_msgTypes,
	}.Build()
	File_InstableSprayDetailInfo_proto = out.File
	file_InstableSprayDetailInfo_proto_rawDesc = nil
	file_InstableSprayDetailInfo_proto_goTypes = nil
	file_InstableSprayDetailInfo_proto_depIdxs = nil
}