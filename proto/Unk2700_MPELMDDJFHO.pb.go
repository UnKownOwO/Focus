// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_MPELMDDJFHO.proto

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

type Unk2700_MPELMDDJFHO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_ONOOJBEABOE uint64               `protobuf:"varint,1,opt,name=Unk2700_ONOOJBEABOE,json=Unk2700ONOOJBEABOE,proto3" json:"Unk2700_ONOOJBEABOE,omitempty"`
	DungeonId           uint32               `protobuf:"varint,2,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	Unk2700_MONNIDCNDFI string               `protobuf:"bytes,3,opt,name=Unk2700_MONNIDCNDFI,json=Unk2700MONNIDCNDFI,proto3" json:"Unk2700_MONNIDCNDFI,omitempty"`
	TagList             []uint32             `protobuf:"varint,4,rep,packed,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	Unk2700_JGFDODPBGFL *Unk2700_GBHAPPCDCIL `protobuf:"bytes,5,opt,name=Unk2700_JGFDODPBGFL,json=Unk2700JGFDODPBGFL,proto3" json:"Unk2700_JGFDODPBGFL,omitempty"`
	Unk2700_PCFIKJEDEGN *Unk2700_IOONEPPHCJP `protobuf:"bytes,6,opt,name=Unk2700_PCFIKJEDEGN,json=Unk2700PCFIKJEDEGN,proto3" json:"Unk2700_PCFIKJEDEGN,omitempty"`
	Unk2700_IKGOMKLAJLH *Unk2700_PDGLEKKMCBD `protobuf:"bytes,7,opt,name=Unk2700_IKGOMKLAJLH,json=Unk2700IKGOMKLAJLH,proto3" json:"Unk2700_IKGOMKLAJLH,omitempty"`
}

func (x *Unk2700_MPELMDDJFHO) Reset() {
	*x = Unk2700_MPELMDDJFHO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_MPELMDDJFHO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_MPELMDDJFHO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_MPELMDDJFHO) ProtoMessage() {}

func (x *Unk2700_MPELMDDJFHO) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_MPELMDDJFHO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_MPELMDDJFHO.ProtoReflect.Descriptor instead.
func (*Unk2700_MPELMDDJFHO) Descriptor() ([]byte, []int) {
	return file_Unk2700_MPELMDDJFHO_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_MPELMDDJFHO) GetUnk2700_ONOOJBEABOE() uint64 {
	if x != nil {
		return x.Unk2700_ONOOJBEABOE
	}
	return 0
}

func (x *Unk2700_MPELMDDJFHO) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *Unk2700_MPELMDDJFHO) GetUnk2700_MONNIDCNDFI() string {
	if x != nil {
		return x.Unk2700_MONNIDCNDFI
	}
	return ""
}

func (x *Unk2700_MPELMDDJFHO) GetTagList() []uint32 {
	if x != nil {
		return x.TagList
	}
	return nil
}

func (x *Unk2700_MPELMDDJFHO) GetUnk2700_JGFDODPBGFL() *Unk2700_GBHAPPCDCIL {
	if x != nil {
		return x.Unk2700_JGFDODPBGFL
	}
	return nil
}

func (x *Unk2700_MPELMDDJFHO) GetUnk2700_PCFIKJEDEGN() *Unk2700_IOONEPPHCJP {
	if x != nil {
		return x.Unk2700_PCFIKJEDEGN
	}
	return nil
}

func (x *Unk2700_MPELMDDJFHO) GetUnk2700_IKGOMKLAJLH() *Unk2700_PDGLEKKMCBD {
	if x != nil {
		return x.Unk2700_IKGOMKLAJLH
	}
	return nil
}

var File_Unk2700_MPELMDDJFHO_proto protoreflect.FileDescriptor

var file_Unk2700_MPELMDDJFHO_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x50, 0x45, 0x4c, 0x4d, 0x44,
	0x44, 0x4a, 0x46, 0x48, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x47, 0x42, 0x48, 0x41, 0x50, 0x50, 0x43, 0x44, 0x43, 0x49, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x49, 0x4f, 0x4f, 0x4e, 0x45, 0x50, 0x50, 0x48, 0x43, 0x4a, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x44, 0x47, 0x4c, 0x45,
	0x4b, 0x4b, 0x4d, 0x43, 0x42, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x50, 0x45, 0x4c, 0x4d, 0x44, 0x44,
	0x4a, 0x46, 0x48, 0x4f, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x4f, 0x4e, 0x4f, 0x4f, 0x4a, 0x42, 0x45, 0x41, 0x42, 0x4f, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4f, 0x4e, 0x4f, 0x4f, 0x4a, 0x42,
	0x45, 0x41, 0x42, 0x4f, 0x45, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x4d, 0x4f, 0x4e, 0x4e, 0x49, 0x44, 0x43, 0x4e, 0x44, 0x46, 0x49, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4d, 0x4f, 0x4e, 0x4e, 0x49, 0x44,
	0x43, 0x4e, 0x44, 0x46, 0x49, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x47, 0x46, 0x44,
	0x4f, 0x44, 0x50, 0x42, 0x47, 0x46, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x47, 0x42, 0x48, 0x41, 0x50, 0x50, 0x43, 0x44,
	0x43, 0x49, 0x4c, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4a, 0x47, 0x46, 0x44,
	0x4f, 0x44, 0x50, 0x42, 0x47, 0x46, 0x4c, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x50, 0x43, 0x46, 0x49, 0x4b, 0x4a, 0x45, 0x44, 0x45, 0x47, 0x4e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x49,
	0x4f, 0x4f, 0x4e, 0x45, 0x50, 0x50, 0x48, 0x43, 0x4a, 0x50, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x50, 0x43, 0x46, 0x49, 0x4b, 0x4a, 0x45, 0x44, 0x45, 0x47, 0x4e, 0x12, 0x45,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x49, 0x4b, 0x47, 0x4f, 0x4d, 0x4b,
	0x4c, 0x41, 0x4a, 0x4c, 0x48, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x44, 0x47, 0x4c, 0x45, 0x4b, 0x4b, 0x4d, 0x43, 0x42,
	0x44, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x4b, 0x47, 0x4f, 0x4d, 0x4b,
	0x4c, 0x41, 0x4a, 0x4c, 0x48, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_MPELMDDJFHO_proto_rawDescOnce sync.Once
	file_Unk2700_MPELMDDJFHO_proto_rawDescData = file_Unk2700_MPELMDDJFHO_proto_rawDesc
)

func file_Unk2700_MPELMDDJFHO_proto_rawDescGZIP() []byte {
	file_Unk2700_MPELMDDJFHO_proto_rawDescOnce.Do(func() {
		file_Unk2700_MPELMDDJFHO_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_MPELMDDJFHO_proto_rawDescData)
	})
	return file_Unk2700_MPELMDDJFHO_proto_rawDescData
}

var file_Unk2700_MPELMDDJFHO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_MPELMDDJFHO_proto_goTypes = []interface{}{
	(*Unk2700_MPELMDDJFHO)(nil), // 0: Unk2700_MPELMDDJFHO
	(*Unk2700_GBHAPPCDCIL)(nil), // 1: Unk2700_GBHAPPCDCIL
	(*Unk2700_IOONEPPHCJP)(nil), // 2: Unk2700_IOONEPPHCJP
	(*Unk2700_PDGLEKKMCBD)(nil), // 3: Unk2700_PDGLEKKMCBD
}
var file_Unk2700_MPELMDDJFHO_proto_depIdxs = []int32{
	1, // 0: Unk2700_MPELMDDJFHO.Unk2700_JGFDODPBGFL:type_name -> Unk2700_GBHAPPCDCIL
	2, // 1: Unk2700_MPELMDDJFHO.Unk2700_PCFIKJEDEGN:type_name -> Unk2700_IOONEPPHCJP
	3, // 2: Unk2700_MPELMDDJFHO.Unk2700_IKGOMKLAJLH:type_name -> Unk2700_PDGLEKKMCBD
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Unk2700_MPELMDDJFHO_proto_init() }
func file_Unk2700_MPELMDDJFHO_proto_init() {
	if File_Unk2700_MPELMDDJFHO_proto != nil {
		return
	}
	file_Unk2700_GBHAPPCDCIL_proto_init()
	file_Unk2700_IOONEPPHCJP_proto_init()
	file_Unk2700_PDGLEKKMCBD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_MPELMDDJFHO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_MPELMDDJFHO); i {
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
			RawDescriptor: file_Unk2700_MPELMDDJFHO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_MPELMDDJFHO_proto_goTypes,
		DependencyIndexes: file_Unk2700_MPELMDDJFHO_proto_depIdxs,
		MessageInfos:      file_Unk2700_MPELMDDJFHO_proto_msgTypes,
	}.Build()
	File_Unk2700_MPELMDDJFHO_proto = out.File
	file_Unk2700_MPELMDDJFHO_proto_rawDesc = nil
	file_Unk2700_MPELMDDJFHO_proto_goTypes = nil
	file_Unk2700_MPELMDDJFHO_proto_depIdxs = nil
}
