// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk3000_NPPMPMGBBLM.proto

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

// CmdId: 6368
// EnetChannelId: 0
// EnetIsReliable: true
type Unk3000_NPPMPMGBBLM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk3000_JPONGJJLGKF uint32              `protobuf:"varint,7,opt,name=Unk3000_JPONGJJLGKF,json=Unk3000JPONGJJLGKF,proto3" json:"Unk3000_JPONGJJLGKF,omitempty"`
	Unk3000_HPKDIOBGGHN Unk3000_AHNHHIOAHBC `protobuf:"varint,12,opt,name=Unk3000_HPKDIOBGGHN,json=Unk3000HPKDIOBGGHN,proto3,enum=Unk3000_AHNHHIOAHBC" json:"Unk3000_HPKDIOBGGHN,omitempty"`
	Unk3000_OAFAKPMJCEN Unk3000_AHNHHIOAHBC `protobuf:"varint,15,opt,name=Unk3000_OAFAKPMJCEN,json=Unk3000OAFAKPMJCEN,proto3,enum=Unk3000_AHNHHIOAHBC" json:"Unk3000_OAFAKPMJCEN,omitempty"`
	Unk3000_BIACMOKGHKF uint32              `protobuf:"varint,8,opt,name=Unk3000_BIACMOKGHKF,json=Unk3000BIACMOKGHKF,proto3" json:"Unk3000_BIACMOKGHKF,omitempty"`
}

func (x *Unk3000_NPPMPMGBBLM) Reset() {
	*x = Unk3000_NPPMPMGBBLM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_NPPMPMGBBLM_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_NPPMPMGBBLM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_NPPMPMGBBLM) ProtoMessage() {}

func (x *Unk3000_NPPMPMGBBLM) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_NPPMPMGBBLM_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_NPPMPMGBBLM.ProtoReflect.Descriptor instead.
func (*Unk3000_NPPMPMGBBLM) Descriptor() ([]byte, []int) {
	return file_Unk3000_NPPMPMGBBLM_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_NPPMPMGBBLM) GetUnk3000_JPONGJJLGKF() uint32 {
	if x != nil {
		return x.Unk3000_JPONGJJLGKF
	}
	return 0
}

func (x *Unk3000_NPPMPMGBBLM) GetUnk3000_HPKDIOBGGHN() Unk3000_AHNHHIOAHBC {
	if x != nil {
		return x.Unk3000_HPKDIOBGGHN
	}
	return Unk3000_AHNHHIOAHBC_Unk3000_AHNHHIOAHBC_NONE
}

func (x *Unk3000_NPPMPMGBBLM) GetUnk3000_OAFAKPMJCEN() Unk3000_AHNHHIOAHBC {
	if x != nil {
		return x.Unk3000_OAFAKPMJCEN
	}
	return Unk3000_AHNHHIOAHBC_Unk3000_AHNHHIOAHBC_NONE
}

func (x *Unk3000_NPPMPMGBBLM) GetUnk3000_BIACMOKGHKF() uint32 {
	if x != nil {
		return x.Unk3000_BIACMOKGHKF
	}
	return 0
}

var File_Unk3000_NPPMPMGBBLM_proto protoreflect.FileDescriptor

var file_Unk3000_NPPMPMGBBLM_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4e, 0x50, 0x50, 0x4d, 0x50, 0x4d,
	0x47, 0x42, 0x42, 0x4c, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x48, 0x4e, 0x48, 0x48, 0x49, 0x4f, 0x41, 0x48, 0x42, 0x43,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x5f, 0x4e, 0x50, 0x50, 0x4d, 0x50, 0x4d, 0x47, 0x42, 0x42, 0x4c, 0x4d, 0x12, 0x2f,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4a, 0x50, 0x4f, 0x4e, 0x47, 0x4a,
	0x4a, 0x4c, 0x47, 0x4b, 0x46, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x4a, 0x50, 0x4f, 0x4e, 0x47, 0x4a, 0x4a, 0x4c, 0x47, 0x4b, 0x46, 0x12,
	0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x48, 0x50, 0x4b, 0x44, 0x49,
	0x4f, 0x42, 0x47, 0x47, 0x48, 0x4e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55,
	0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x48, 0x4e, 0x48, 0x48, 0x49, 0x4f, 0x41, 0x48,
	0x42, 0x43, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x48, 0x50, 0x4b, 0x44, 0x49,
	0x4f, 0x42, 0x47, 0x47, 0x48, 0x4e, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30,
	0x30, 0x5f, 0x4f, 0x41, 0x46, 0x41, 0x4b, 0x50, 0x4d, 0x4a, 0x43, 0x45, 0x4e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x48,
	0x4e, 0x48, 0x48, 0x49, 0x4f, 0x41, 0x48, 0x42, 0x43, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x4f, 0x41, 0x46, 0x41, 0x4b, 0x50, 0x4d, 0x4a, 0x43, 0x45, 0x4e, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x42, 0x49, 0x41, 0x43, 0x4d, 0x4f, 0x4b,
	0x47, 0x48, 0x4b, 0x46, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x30, 0x30, 0x30, 0x42, 0x49, 0x41, 0x43, 0x4d, 0x4f, 0x4b, 0x47, 0x48, 0x4b, 0x46, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Unk3000_NPPMPMGBBLM_proto_rawDescOnce sync.Once
	file_Unk3000_NPPMPMGBBLM_proto_rawDescData = file_Unk3000_NPPMPMGBBLM_proto_rawDesc
)

func file_Unk3000_NPPMPMGBBLM_proto_rawDescGZIP() []byte {
	file_Unk3000_NPPMPMGBBLM_proto_rawDescOnce.Do(func() {
		file_Unk3000_NPPMPMGBBLM_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_NPPMPMGBBLM_proto_rawDescData)
	})
	return file_Unk3000_NPPMPMGBBLM_proto_rawDescData
}

var file_Unk3000_NPPMPMGBBLM_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3000_NPPMPMGBBLM_proto_goTypes = []interface{}{
	(*Unk3000_NPPMPMGBBLM)(nil), // 0: Unk3000_NPPMPMGBBLM
	(Unk3000_AHNHHIOAHBC)(0),    // 1: Unk3000_AHNHHIOAHBC
}
var file_Unk3000_NPPMPMGBBLM_proto_depIdxs = []int32{
	1, // 0: Unk3000_NPPMPMGBBLM.Unk3000_HPKDIOBGGHN:type_name -> Unk3000_AHNHHIOAHBC
	1, // 1: Unk3000_NPPMPMGBBLM.Unk3000_OAFAKPMJCEN:type_name -> Unk3000_AHNHHIOAHBC
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk3000_NPPMPMGBBLM_proto_init() }
func file_Unk3000_NPPMPMGBBLM_proto_init() {
	if File_Unk3000_NPPMPMGBBLM_proto != nil {
		return
	}
	file_Unk3000_AHNHHIOAHBC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_NPPMPMGBBLM_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_NPPMPMGBBLM); i {
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
			RawDescriptor: file_Unk3000_NPPMPMGBBLM_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_NPPMPMGBBLM_proto_goTypes,
		DependencyIndexes: file_Unk3000_NPPMPMGBBLM_proto_depIdxs,
		MessageInfos:      file_Unk3000_NPPMPMGBBLM_proto_msgTypes,
	}.Build()
	File_Unk3000_NPPMPMGBBLM_proto = out.File
	file_Unk3000_NPPMPMGBBLM_proto_rawDesc = nil
	file_Unk3000_NPPMPMGBBLM_proto_goTypes = nil
	file_Unk3000_NPPMPMGBBLM_proto_depIdxs = nil
}
