// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_CPNDLPDOPGN.proto

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

type Unk2700_CPNDLPDOPGN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_ONOOJBEABOE uint64 `protobuf:"varint,3,opt,name=Unk2700_ONOOJBEABOE,json=Unk2700ONOOJBEABOE,proto3" json:"Unk2700_ONOOJBEABOE,omitempty"`
	Uid                 uint32 `protobuf:"varint,15,opt,name=uid,proto3" json:"uid,omitempty"`
	Timestamp           uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Region              string `protobuf:"bytes,11,opt,name=region,proto3" json:"region,omitempty"`
	Lang                uint32 `protobuf:"varint,13,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *Unk2700_CPNDLPDOPGN) Reset() {
	*x = Unk2700_CPNDLPDOPGN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_CPNDLPDOPGN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_CPNDLPDOPGN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_CPNDLPDOPGN) ProtoMessage() {}

func (x *Unk2700_CPNDLPDOPGN) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_CPNDLPDOPGN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_CPNDLPDOPGN.ProtoReflect.Descriptor instead.
func (*Unk2700_CPNDLPDOPGN) Descriptor() ([]byte, []int) {
	return file_Unk2700_CPNDLPDOPGN_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_CPNDLPDOPGN) GetUnk2700_ONOOJBEABOE() uint64 {
	if x != nil {
		return x.Unk2700_ONOOJBEABOE
	}
	return 0
}

func (x *Unk2700_CPNDLPDOPGN) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Unk2700_CPNDLPDOPGN) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Unk2700_CPNDLPDOPGN) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Unk2700_CPNDLPDOPGN) GetLang() uint32 {
	if x != nil {
		return x.Lang
	}
	return 0
}

var File_Unk2700_CPNDLPDOPGN_proto protoreflect.FileDescriptor

var file_Unk2700_CPNDLPDOPGN_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x50, 0x4e, 0x44, 0x4c, 0x50,
	0x44, 0x4f, 0x50, 0x47, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x50, 0x4e, 0x44, 0x4c, 0x50, 0x44, 0x4f,
	0x50, 0x47, 0x4e, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f,
	0x4e, 0x4f, 0x4f, 0x4a, 0x42, 0x45, 0x41, 0x42, 0x4f, 0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4f, 0x4e, 0x4f, 0x4f, 0x4a, 0x42, 0x45,
	0x41, 0x42, 0x4f, 0x45, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x61, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_CPNDLPDOPGN_proto_rawDescOnce sync.Once
	file_Unk2700_CPNDLPDOPGN_proto_rawDescData = file_Unk2700_CPNDLPDOPGN_proto_rawDesc
)

func file_Unk2700_CPNDLPDOPGN_proto_rawDescGZIP() []byte {
	file_Unk2700_CPNDLPDOPGN_proto_rawDescOnce.Do(func() {
		file_Unk2700_CPNDLPDOPGN_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_CPNDLPDOPGN_proto_rawDescData)
	})
	return file_Unk2700_CPNDLPDOPGN_proto_rawDescData
}

var file_Unk2700_CPNDLPDOPGN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_CPNDLPDOPGN_proto_goTypes = []interface{}{
	(*Unk2700_CPNDLPDOPGN)(nil), // 0: Unk2700_CPNDLPDOPGN
}
var file_Unk2700_CPNDLPDOPGN_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2700_CPNDLPDOPGN_proto_init() }
func file_Unk2700_CPNDLPDOPGN_proto_init() {
	if File_Unk2700_CPNDLPDOPGN_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_CPNDLPDOPGN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_CPNDLPDOPGN); i {
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
			RawDescriptor: file_Unk2700_CPNDLPDOPGN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_CPNDLPDOPGN_proto_goTypes,
		DependencyIndexes: file_Unk2700_CPNDLPDOPGN_proto_depIdxs,
		MessageInfos:      file_Unk2700_CPNDLPDOPGN_proto_msgTypes,
	}.Build()
	File_Unk2700_CPNDLPDOPGN_proto = out.File
	file_Unk2700_CPNDLPDOPGN_proto_rawDesc = nil
	file_Unk2700_CPNDLPDOPGN_proto_goTypes = nil
	file_Unk2700_CPNDLPDOPGN_proto_depIdxs = nil
}
