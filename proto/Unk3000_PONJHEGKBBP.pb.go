// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk3000_PONJHEGKBBP.proto

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

type Unk3000_PONJHEGKBBP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3000_MKNODEKEGJF map[uint32]Unk3000_AHNHHIOAHBC `protobuf:"bytes,6,rep,name=Unk3000_MKNODEKEGJF,json=Unk3000MKNODEKEGJF,proto3" json:"Unk3000_MKNODEKEGJF,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=Unk3000_AHNHHIOAHBC"`
	Unk3000_JPONGJJLGKF uint32                         `protobuf:"varint,12,opt,name=Unk3000_JPONGJJLGKF,json=Unk3000JPONGJJLGKF,proto3" json:"Unk3000_JPONGJJLGKF,omitempty"`
}

func (x *Unk3000_PONJHEGKBBP) Reset() {
	*x = Unk3000_PONJHEGKBBP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_PONJHEGKBBP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_PONJHEGKBBP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_PONJHEGKBBP) ProtoMessage() {}

func (x *Unk3000_PONJHEGKBBP) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_PONJHEGKBBP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_PONJHEGKBBP.ProtoReflect.Descriptor instead.
func (*Unk3000_PONJHEGKBBP) Descriptor() ([]byte, []int) {
	return file_Unk3000_PONJHEGKBBP_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_PONJHEGKBBP) GetUnk3000_MKNODEKEGJF() map[uint32]Unk3000_AHNHHIOAHBC {
	if x != nil {
		return x.Unk3000_MKNODEKEGJF
	}
	return nil
}

func (x *Unk3000_PONJHEGKBBP) GetUnk3000_JPONGJJLGKF() uint32 {
	if x != nil {
		return x.Unk3000_JPONGJJLGKF
	}
	return 0
}

var File_Unk3000_PONJHEGKBBP_proto protoreflect.FileDescriptor

var file_Unk3000_PONJHEGKBBP_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x4f, 0x4e, 0x4a, 0x48, 0x45,
	0x47, 0x4b, 0x42, 0x42, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x48, 0x4e, 0x48, 0x48, 0x49, 0x4f, 0x41, 0x48, 0x42, 0x43,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x5f, 0x50, 0x4f, 0x4e, 0x4a, 0x48, 0x45, 0x47, 0x4b, 0x42, 0x42, 0x50, 0x12, 0x5d,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4d, 0x4b, 0x4e, 0x4f, 0x44, 0x45,
	0x4b, 0x45, 0x47, 0x4a, 0x46, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x55, 0x6e,
	0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x4f, 0x4e, 0x4a, 0x48, 0x45, 0x47, 0x4b, 0x42, 0x42,
	0x50, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x4d, 0x4b, 0x4e, 0x4f, 0x44, 0x45, 0x4b,
	0x45, 0x47, 0x4a, 0x46, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x4d, 0x4b, 0x4e, 0x4f, 0x44, 0x45, 0x4b, 0x45, 0x47, 0x4a, 0x46, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4a, 0x50, 0x4f, 0x4e, 0x47, 0x4a, 0x4a,
	0x4c, 0x47, 0x4b, 0x46, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x30, 0x30, 0x30, 0x4a, 0x50, 0x4f, 0x4e, 0x47, 0x4a, 0x4a, 0x4c, 0x47, 0x4b, 0x46, 0x1a, 0x5b,
	0x0a, 0x17, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x4d, 0x4b, 0x4e, 0x4f, 0x44, 0x45, 0x4b,
	0x45, 0x47, 0x4a, 0x46, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x48, 0x4e, 0x48, 0x48, 0x49, 0x4f, 0x41, 0x48, 0x42, 0x43,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_PONJHEGKBBP_proto_rawDescOnce sync.Once
	file_Unk3000_PONJHEGKBBP_proto_rawDescData = file_Unk3000_PONJHEGKBBP_proto_rawDesc
)

func file_Unk3000_PONJHEGKBBP_proto_rawDescGZIP() []byte {
	file_Unk3000_PONJHEGKBBP_proto_rawDescOnce.Do(func() {
		file_Unk3000_PONJHEGKBBP_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_PONJHEGKBBP_proto_rawDescData)
	})
	return file_Unk3000_PONJHEGKBBP_proto_rawDescData
}

var file_Unk3000_PONJHEGKBBP_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Unk3000_PONJHEGKBBP_proto_goTypes = []interface{}{
	(*Unk3000_PONJHEGKBBP)(nil), // 0: Unk3000_PONJHEGKBBP
	nil,                         // 1: Unk3000_PONJHEGKBBP.Unk3000MKNODEKEGJFEntry
	(Unk3000_AHNHHIOAHBC)(0),    // 2: Unk3000_AHNHHIOAHBC
}
var file_Unk3000_PONJHEGKBBP_proto_depIdxs = []int32{
	1, // 0: Unk3000_PONJHEGKBBP.Unk3000_MKNODEKEGJF:type_name -> Unk3000_PONJHEGKBBP.Unk3000MKNODEKEGJFEntry
	2, // 1: Unk3000_PONJHEGKBBP.Unk3000MKNODEKEGJFEntry.value:type_name -> Unk3000_AHNHHIOAHBC
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk3000_PONJHEGKBBP_proto_init() }
func file_Unk3000_PONJHEGKBBP_proto_init() {
	if File_Unk3000_PONJHEGKBBP_proto != nil {
		return
	}
	file_Unk3000_AHNHHIOAHBC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_PONJHEGKBBP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_PONJHEGKBBP); i {
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
			RawDescriptor: file_Unk3000_PONJHEGKBBP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_PONJHEGKBBP_proto_goTypes,
		DependencyIndexes: file_Unk3000_PONJHEGKBBP_proto_depIdxs,
		MessageInfos:      file_Unk3000_PONJHEGKBBP_proto_msgTypes,
	}.Build()
	File_Unk3000_PONJHEGKBBP_proto = out.File
	file_Unk3000_PONJHEGKBBP_proto_rawDesc = nil
	file_Unk3000_PONJHEGKBBP_proto_goTypes = nil
	file_Unk3000_PONJHEGKBBP_proto_depIdxs = nil
}
