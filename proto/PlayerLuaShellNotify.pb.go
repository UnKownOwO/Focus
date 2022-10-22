// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: PlayerLuaShellNotify.proto

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

// CmdId: 133
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerLuaShellNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_JJMHFFHNJJO Unk2700_JOEPIGNPDGH `protobuf:"varint,7,opt,name=Unk2700_JJMHFFHNJJO,json=Unk2700JJMHFFHNJJO,proto3,enum=Unk2700_JOEPIGNPDGH" json:"Unk2700_JJMHFFHNJJO,omitempty"`
	Id                  uint32              `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	LuaShell            []byte              `protobuf:"bytes,12,opt,name=lua_shell,json=luaShell,proto3" json:"lua_shell,omitempty"`
	UseType             uint32              `protobuf:"varint,10,opt,name=use_type,json=useType,proto3" json:"use_type,omitempty"`
}

func (x *PlayerLuaShellNotify) Reset() {
	*x = PlayerLuaShellNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLuaShellNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLuaShellNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLuaShellNotify) ProtoMessage() {}

func (x *PlayerLuaShellNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLuaShellNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLuaShellNotify.ProtoReflect.Descriptor instead.
func (*PlayerLuaShellNotify) Descriptor() ([]byte, []int) {
	return file_PlayerLuaShellNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLuaShellNotify) GetUnk2700_JJMHFFHNJJO() Unk2700_JOEPIGNPDGH {
	if x != nil {
		return x.Unk2700_JJMHFFHNJJO
	}
	return Unk2700_JOEPIGNPDGH_Unk2700_JOEPIGNPDGH_Unk2700_GIGONJIGKBM
}

func (x *PlayerLuaShellNotify) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlayerLuaShellNotify) GetLuaShell() []byte {
	if x != nil {
		return x.LuaShell
	}
	return nil
}

func (x *PlayerLuaShellNotify) GetUseType() uint32 {
	if x != nil {
		return x.UseType
	}
	return 0
}

var File_PlayerLuaShellNotify_proto protoreflect.FileDescriptor

var file_PlayerLuaShellNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x75, 0x61, 0x53, 0x68, 0x65, 0x6c, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4f, 0x45, 0x50, 0x49, 0x47, 0x4e, 0x50, 0x44, 0x47,
	0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x4c, 0x75, 0x61, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4a, 0x4d, 0x48,
	0x46, 0x46, 0x48, 0x4e, 0x4a, 0x4a, 0x4f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4f, 0x45, 0x50, 0x49, 0x47, 0x4e, 0x50,
	0x44, 0x47, 0x48, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4a, 0x4a, 0x4d, 0x48,
	0x46, 0x46, 0x48, 0x4e, 0x4a, 0x4a, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x75, 0x61, 0x5f, 0x73,
	0x68, 0x65, 0x6c, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x75, 0x61, 0x53,
	0x68, 0x65, 0x6c, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PlayerLuaShellNotify_proto_rawDescOnce sync.Once
	file_PlayerLuaShellNotify_proto_rawDescData = file_PlayerLuaShellNotify_proto_rawDesc
)

func file_PlayerLuaShellNotify_proto_rawDescGZIP() []byte {
	file_PlayerLuaShellNotify_proto_rawDescOnce.Do(func() {
		file_PlayerLuaShellNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLuaShellNotify_proto_rawDescData)
	})
	return file_PlayerLuaShellNotify_proto_rawDescData
}

var file_PlayerLuaShellNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerLuaShellNotify_proto_goTypes = []interface{}{
	(*PlayerLuaShellNotify)(nil), // 0: PlayerLuaShellNotify
	(Unk2700_JOEPIGNPDGH)(0),     // 1: Unk2700_JOEPIGNPDGH
}
var file_PlayerLuaShellNotify_proto_depIdxs = []int32{
	1, // 0: PlayerLuaShellNotify.Unk2700_JJMHFFHNJJO:type_name -> Unk2700_JOEPIGNPDGH
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerLuaShellNotify_proto_init() }
func file_PlayerLuaShellNotify_proto_init() {
	if File_PlayerLuaShellNotify_proto != nil {
		return
	}
	file_Unk2700_JOEPIGNPDGH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerLuaShellNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLuaShellNotify); i {
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
			RawDescriptor: file_PlayerLuaShellNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLuaShellNotify_proto_goTypes,
		DependencyIndexes: file_PlayerLuaShellNotify_proto_depIdxs,
		MessageInfos:      file_PlayerLuaShellNotify_proto_msgTypes,
	}.Build()
	File_PlayerLuaShellNotify_proto = out.File
	file_PlayerLuaShellNotify_proto_rawDesc = nil
	file_PlayerLuaShellNotify_proto_goTypes = nil
	file_PlayerLuaShellNotify_proto_depIdxs = nil
}
