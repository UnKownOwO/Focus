// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk3000_DHOFMKPKFMF.proto

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

// CmdId: 1749
// EnetChannelId: 0
// EnetIsReliable: true
type Unk3000_DHOFMKPKFMF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TempAvatarGuidList  []uint64               `protobuf:"varint,6,rep,packed,name=temp_avatar_guid_list,json=tempAvatarGuidList,proto3" json:"temp_avatar_guid_list,omitempty"`
	AvatarTeamMap       map[uint32]*AvatarTeam `protobuf:"bytes,3,rep,name=avatar_team_map,json=avatarTeamMap,proto3" json:"avatar_team_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Unk3000_NIGPICLBHMA []uint32               `protobuf:"varint,1,rep,packed,name=Unk3000_NIGPICLBHMA,json=Unk3000NIGPICLBHMA,proto3" json:"Unk3000_NIGPICLBHMA,omitempty"`
}

func (x *Unk3000_DHOFMKPKFMF) Reset() {
	*x = Unk3000_DHOFMKPKFMF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3000_DHOFMKPKFMF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3000_DHOFMKPKFMF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3000_DHOFMKPKFMF) ProtoMessage() {}

func (x *Unk3000_DHOFMKPKFMF) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3000_DHOFMKPKFMF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3000_DHOFMKPKFMF.ProtoReflect.Descriptor instead.
func (*Unk3000_DHOFMKPKFMF) Descriptor() ([]byte, []int) {
	return file_Unk3000_DHOFMKPKFMF_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3000_DHOFMKPKFMF) GetTempAvatarGuidList() []uint64 {
	if x != nil {
		return x.TempAvatarGuidList
	}
	return nil
}

func (x *Unk3000_DHOFMKPKFMF) GetAvatarTeamMap() map[uint32]*AvatarTeam {
	if x != nil {
		return x.AvatarTeamMap
	}
	return nil
}

func (x *Unk3000_DHOFMKPKFMF) GetUnk3000_NIGPICLBHMA() []uint32 {
	if x != nil {
		return x.Unk3000_NIGPICLBHMA
	}
	return nil
}

var File_Unk3000_DHOFMKPKFMF_proto protoreflect.FileDescriptor

var file_Unk3000_DHOFMKPKFMF_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x44, 0x48, 0x4f, 0x46, 0x4d, 0x4b,
	0x50, 0x4b, 0x46, 0x4d, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x44, 0x48, 0x4f, 0x46, 0x4d, 0x4b,
	0x50, 0x4b, 0x46, 0x4d, 0x46, 0x12, 0x31, 0x0a, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x44, 0x48, 0x4f, 0x46,
	0x4d, 0x4b, 0x50, 0x4b, 0x46, 0x4d, 0x46, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65,
	0x61, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x33, 0x30, 0x30, 0x30, 0x5f, 0x4e, 0x49, 0x47, 0x50, 0x49, 0x43, 0x4c, 0x42, 0x48, 0x4d, 0x41,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x4e,
	0x49, 0x47, 0x50, 0x49, 0x43, 0x4c, 0x42, 0x48, 0x4d, 0x41, 0x1a, 0x4d, 0x0a, 0x12, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_DHOFMKPKFMF_proto_rawDescOnce sync.Once
	file_Unk3000_DHOFMKPKFMF_proto_rawDescData = file_Unk3000_DHOFMKPKFMF_proto_rawDesc
)

func file_Unk3000_DHOFMKPKFMF_proto_rawDescGZIP() []byte {
	file_Unk3000_DHOFMKPKFMF_proto_rawDescOnce.Do(func() {
		file_Unk3000_DHOFMKPKFMF_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_DHOFMKPKFMF_proto_rawDescData)
	})
	return file_Unk3000_DHOFMKPKFMF_proto_rawDescData
}

var file_Unk3000_DHOFMKPKFMF_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Unk3000_DHOFMKPKFMF_proto_goTypes = []interface{}{
	(*Unk3000_DHOFMKPKFMF)(nil), // 0: Unk3000_DHOFMKPKFMF
	nil,                         // 1: Unk3000_DHOFMKPKFMF.AvatarTeamMapEntry
	(*AvatarTeam)(nil),          // 2: AvatarTeam
}
var file_Unk3000_DHOFMKPKFMF_proto_depIdxs = []int32{
	1, // 0: Unk3000_DHOFMKPKFMF.avatar_team_map:type_name -> Unk3000_DHOFMKPKFMF.AvatarTeamMapEntry
	2, // 1: Unk3000_DHOFMKPKFMF.AvatarTeamMapEntry.value:type_name -> AvatarTeam
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk3000_DHOFMKPKFMF_proto_init() }
func file_Unk3000_DHOFMKPKFMF_proto_init() {
	if File_Unk3000_DHOFMKPKFMF_proto != nil {
		return
	}
	file_AvatarTeam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3000_DHOFMKPKFMF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3000_DHOFMKPKFMF); i {
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
			RawDescriptor: file_Unk3000_DHOFMKPKFMF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_DHOFMKPKFMF_proto_goTypes,
		DependencyIndexes: file_Unk3000_DHOFMKPKFMF_proto_depIdxs,
		MessageInfos:      file_Unk3000_DHOFMKPKFMF_proto_msgTypes,
	}.Build()
	File_Unk3000_DHOFMKPKFMF_proto = out.File
	file_Unk3000_DHOFMKPKFMF_proto_rawDesc = nil
	file_Unk3000_DHOFMKPKFMF_proto_goTypes = nil
	file_Unk3000_DHOFMKPKFMF_proto_depIdxs = nil
}