// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_PIEJMALFKIF.proto

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

// CmdId: 8531
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_PIEJMALFKIF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId             uint32                 `protobuf:"varint,13,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	Unk2700_FHEHGDABALE uint32                 `protobuf:"varint,7,opt,name=Unk2700_FHEHGDABALE,json=Unk2700FHEHGDABALE,proto3" json:"Unk2700_FHEHGDABALE,omitempty"`
	DungeonAvatarList   []*Unk2700_KHDMDKKDOCD `protobuf:"bytes,6,rep,name=dungeon_avatar_list,json=dungeonAvatarList,proto3" json:"dungeon_avatar_list,omitempty"`
	LevelId             uint32                 `protobuf:"varint,8,opt,name=level_id,json=levelId,proto3" json:"level_id,omitempty"`
	Unk2700_HKFEBBCMBHL uint32                 `protobuf:"varint,5,opt,name=Unk2700_HKFEBBCMBHL,json=Unk2700HKFEBBCMBHL,proto3" json:"Unk2700_HKFEBBCMBHL,omitempty"`
}

func (x *Unk2700_PIEJMALFKIF) Reset() {
	*x = Unk2700_PIEJMALFKIF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_PIEJMALFKIF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_PIEJMALFKIF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_PIEJMALFKIF) ProtoMessage() {}

func (x *Unk2700_PIEJMALFKIF) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_PIEJMALFKIF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_PIEJMALFKIF.ProtoReflect.Descriptor instead.
func (*Unk2700_PIEJMALFKIF) Descriptor() ([]byte, []int) {
	return file_Unk2700_PIEJMALFKIF_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_PIEJMALFKIF) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *Unk2700_PIEJMALFKIF) GetUnk2700_FHEHGDABALE() uint32 {
	if x != nil {
		return x.Unk2700_FHEHGDABALE
	}
	return 0
}

func (x *Unk2700_PIEJMALFKIF) GetDungeonAvatarList() []*Unk2700_KHDMDKKDOCD {
	if x != nil {
		return x.DungeonAvatarList
	}
	return nil
}

func (x *Unk2700_PIEJMALFKIF) GetLevelId() uint32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *Unk2700_PIEJMALFKIF) GetUnk2700_HKFEBBCMBHL() uint32 {
	if x != nil {
		return x.Unk2700_HKFEBBCMBHL
	}
	return 0
}

var File_Unk2700_PIEJMALFKIF_proto protoreflect.FileDescriptor

var file_Unk2700_PIEJMALFKIF_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x49, 0x45, 0x4a, 0x4d, 0x41,
	0x4c, 0x46, 0x4b, 0x49, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x48, 0x44, 0x4d, 0x44, 0x4b, 0x4b, 0x44, 0x4f, 0x43, 0x44,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x50, 0x49, 0x45, 0x4a, 0x4d, 0x41, 0x4c, 0x46, 0x4b, 0x49, 0x46, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x48, 0x45, 0x48, 0x47, 0x44, 0x41, 0x42, 0x41, 0x4c, 0x45,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x46,
	0x48, 0x45, 0x48, 0x47, 0x44, 0x41, 0x42, 0x41, 0x4c, 0x45, 0x12, 0x44, 0x0a, 0x13, 0x64, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x5f, 0x4b, 0x48, 0x44, 0x4d, 0x44, 0x4b, 0x4b, 0x44, 0x4f, 0x43, 0x44, 0x52, 0x11, 0x64,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x4b, 0x46, 0x45, 0x42, 0x42, 0x43, 0x4d, 0x42,
	0x48, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x48, 0x4b, 0x46, 0x45, 0x42, 0x42, 0x43, 0x4d, 0x42, 0x48, 0x4c, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_PIEJMALFKIF_proto_rawDescOnce sync.Once
	file_Unk2700_PIEJMALFKIF_proto_rawDescData = file_Unk2700_PIEJMALFKIF_proto_rawDesc
)

func file_Unk2700_PIEJMALFKIF_proto_rawDescGZIP() []byte {
	file_Unk2700_PIEJMALFKIF_proto_rawDescOnce.Do(func() {
		file_Unk2700_PIEJMALFKIF_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_PIEJMALFKIF_proto_rawDescData)
	})
	return file_Unk2700_PIEJMALFKIF_proto_rawDescData
}

var file_Unk2700_PIEJMALFKIF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_PIEJMALFKIF_proto_goTypes = []interface{}{
	(*Unk2700_PIEJMALFKIF)(nil), // 0: Unk2700_PIEJMALFKIF
	(*Unk2700_KHDMDKKDOCD)(nil), // 1: Unk2700_KHDMDKKDOCD
}
var file_Unk2700_PIEJMALFKIF_proto_depIdxs = []int32{
	1, // 0: Unk2700_PIEJMALFKIF.dungeon_avatar_list:type_name -> Unk2700_KHDMDKKDOCD
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk2700_PIEJMALFKIF_proto_init() }
func file_Unk2700_PIEJMALFKIF_proto_init() {
	if File_Unk2700_PIEJMALFKIF_proto != nil {
		return
	}
	file_Unk2700_KHDMDKKDOCD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_PIEJMALFKIF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_PIEJMALFKIF); i {
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
			RawDescriptor: file_Unk2700_PIEJMALFKIF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_PIEJMALFKIF_proto_goTypes,
		DependencyIndexes: file_Unk2700_PIEJMALFKIF_proto_depIdxs,
		MessageInfos:      file_Unk2700_PIEJMALFKIF_proto_msgTypes,
	}.Build()
	File_Unk2700_PIEJMALFKIF_proto = out.File
	file_Unk2700_PIEJMALFKIF_proto_rawDesc = nil
	file_Unk2700_PIEJMALFKIF_proto_goTypes = nil
	file_Unk2700_PIEJMALFKIF_proto_depIdxs = nil
}