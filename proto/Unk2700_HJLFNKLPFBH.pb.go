// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_HJLFNKLPFBH.proto

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

type Unk2700_HJLFNKLPFBH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar              *Unk2700_FMGGGEDNGGN  `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Level               uint32                `protobuf:"varint,14,opt,name=level,proto3" json:"level,omitempty"`
	Unk2700_EGKOIPOHCHG uint32                `protobuf:"varint,13,opt,name=Unk2700_EGKOIPOHCHG,json=Unk2700EGKOIPOHCHG,proto3" json:"Unk2700_EGKOIPOHCHG,omitempty"`
	Unk2700_JCKLLFKOFCG []Unk2700_AGIDJODJNEA `protobuf:"varint,9,rep,packed,name=Unk2700_JCKLLFKOFCG,json=Unk2700JCKLLFKOFCG,proto3,enum=Unk2700_AGIDJODJNEA" json:"Unk2700_JCKLLFKOFCG,omitempty"`
}

func (x *Unk2700_HJLFNKLPFBH) Reset() {
	*x = Unk2700_HJLFNKLPFBH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_HJLFNKLPFBH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_HJLFNKLPFBH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_HJLFNKLPFBH) ProtoMessage() {}

func (x *Unk2700_HJLFNKLPFBH) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_HJLFNKLPFBH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_HJLFNKLPFBH.ProtoReflect.Descriptor instead.
func (*Unk2700_HJLFNKLPFBH) Descriptor() ([]byte, []int) {
	return file_Unk2700_HJLFNKLPFBH_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_HJLFNKLPFBH) GetAvatar() *Unk2700_FMGGGEDNGGN {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *Unk2700_HJLFNKLPFBH) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Unk2700_HJLFNKLPFBH) GetUnk2700_EGKOIPOHCHG() uint32 {
	if x != nil {
		return x.Unk2700_EGKOIPOHCHG
	}
	return 0
}

func (x *Unk2700_HJLFNKLPFBH) GetUnk2700_JCKLLFKOFCG() []Unk2700_AGIDJODJNEA {
	if x != nil {
		return x.Unk2700_JCKLLFKOFCG
	}
	return nil
}

var File_Unk2700_HJLFNKLPFBH_proto protoreflect.FileDescriptor

var file_Unk2700_HJLFNKLPFBH_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x4a, 0x4c, 0x46, 0x4e, 0x4b,
	0x4c, 0x50, 0x46, 0x42, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x47, 0x49, 0x44, 0x4a, 0x4f, 0x44, 0x4a, 0x4e, 0x45, 0x41,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x46, 0x4d, 0x47, 0x47, 0x47, 0x45, 0x44, 0x4e, 0x47, 0x47, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x4a,
	0x4c, 0x46, 0x4e, 0x4b, 0x4c, 0x50, 0x46, 0x42, 0x48, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x46, 0x4d, 0x47, 0x47, 0x47, 0x45, 0x44, 0x4e, 0x47, 0x47, 0x4e, 0x52,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x47, 0x4b, 0x4f, 0x49, 0x50, 0x4f,
	0x48, 0x43, 0x48, 0x47, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x45, 0x47, 0x4b, 0x4f, 0x49, 0x50, 0x4f, 0x48, 0x43, 0x48, 0x47, 0x12, 0x45,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x43, 0x4b, 0x4c, 0x4c, 0x46,
	0x4b, 0x4f, 0x46, 0x43, 0x47, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x47, 0x49, 0x44, 0x4a, 0x4f, 0x44, 0x4a, 0x4e, 0x45,
	0x41, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4a, 0x43, 0x4b, 0x4c, 0x4c, 0x46,
	0x4b, 0x4f, 0x46, 0x43, 0x47, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_HJLFNKLPFBH_proto_rawDescOnce sync.Once
	file_Unk2700_HJLFNKLPFBH_proto_rawDescData = file_Unk2700_HJLFNKLPFBH_proto_rawDesc
)

func file_Unk2700_HJLFNKLPFBH_proto_rawDescGZIP() []byte {
	file_Unk2700_HJLFNKLPFBH_proto_rawDescOnce.Do(func() {
		file_Unk2700_HJLFNKLPFBH_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_HJLFNKLPFBH_proto_rawDescData)
	})
	return file_Unk2700_HJLFNKLPFBH_proto_rawDescData
}

var file_Unk2700_HJLFNKLPFBH_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_HJLFNKLPFBH_proto_goTypes = []interface{}{
	(*Unk2700_HJLFNKLPFBH)(nil), // 0: Unk2700_HJLFNKLPFBH
	(*Unk2700_FMGGGEDNGGN)(nil), // 1: Unk2700_FMGGGEDNGGN
	(Unk2700_AGIDJODJNEA)(0),    // 2: Unk2700_AGIDJODJNEA
}
var file_Unk2700_HJLFNKLPFBH_proto_depIdxs = []int32{
	1, // 0: Unk2700_HJLFNKLPFBH.avatar:type_name -> Unk2700_FMGGGEDNGGN
	2, // 1: Unk2700_HJLFNKLPFBH.Unk2700_JCKLLFKOFCG:type_name -> Unk2700_AGIDJODJNEA
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk2700_HJLFNKLPFBH_proto_init() }
func file_Unk2700_HJLFNKLPFBH_proto_init() {
	if File_Unk2700_HJLFNKLPFBH_proto != nil {
		return
	}
	file_Unk2700_AGIDJODJNEA_proto_init()
	file_Unk2700_FMGGGEDNGGN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_HJLFNKLPFBH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_HJLFNKLPFBH); i {
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
			RawDescriptor: file_Unk2700_HJLFNKLPFBH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_HJLFNKLPFBH_proto_goTypes,
		DependencyIndexes: file_Unk2700_HJLFNKLPFBH_proto_depIdxs,
		MessageInfos:      file_Unk2700_HJLFNKLPFBH_proto_msgTypes,
	}.Build()
	File_Unk2700_HJLFNKLPFBH_proto = out.File
	file_Unk2700_HJLFNKLPFBH_proto_rawDesc = nil
	file_Unk2700_HJLFNKLPFBH_proto_goTypes = nil
	file_Unk2700_HJLFNKLPFBH_proto_depIdxs = nil
}
