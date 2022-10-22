// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: PlayerQuitFromMpNotify.proto

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

type PlayerQuitFromMpNotify_QuitReason int32

const (
	PlayerQuitFromMpNotify_QUIT_REASON_INVALID                 PlayerQuitFromMpNotify_QuitReason = 0
	PlayerQuitFromMpNotify_QUIT_REASON_HOST_NO_OTHER_PLAYER    PlayerQuitFromMpNotify_QuitReason = 1
	PlayerQuitFromMpNotify_QUIT_REASON_KICK_BY_HOST            PlayerQuitFromMpNotify_QuitReason = 2
	PlayerQuitFromMpNotify_QUIT_REASON_BACK_TO_MY_WORLD        PlayerQuitFromMpNotify_QuitReason = 3
	PlayerQuitFromMpNotify_QUIT_REASON_KICK_BY_HOST_LOGOUT     PlayerQuitFromMpNotify_QuitReason = 4
	PlayerQuitFromMpNotify_QUIT_REASON_KICK_BY_HOST_BLOCK      PlayerQuitFromMpNotify_QuitReason = 5
	PlayerQuitFromMpNotify_QUIT_REASON_BE_BLOCKED              PlayerQuitFromMpNotify_QuitReason = 6
	PlayerQuitFromMpNotify_QUIT_REASON_KICK_BY_HOST_ENTER_HOME PlayerQuitFromMpNotify_QuitReason = 7
	PlayerQuitFromMpNotify_QUIT_REASON_HOST_SCENE_INVALID      PlayerQuitFromMpNotify_QuitReason = 8
	PlayerQuitFromMpNotify_QUIT_REASON_KICK_BY_PLAY            PlayerQuitFromMpNotify_QuitReason = 9
	PlayerQuitFromMpNotify_QUIT_REASON_Unk2800_FDECHAHJFDA     PlayerQuitFromMpNotify_QuitReason = 10
)

// Enum value maps for PlayerQuitFromMpNotify_QuitReason.
var (
	PlayerQuitFromMpNotify_QuitReason_name = map[int32]string{
		0:  "QUIT_REASON_INVALID",
		1:  "QUIT_REASON_HOST_NO_OTHER_PLAYER",
		2:  "QUIT_REASON_KICK_BY_HOST",
		3:  "QUIT_REASON_BACK_TO_MY_WORLD",
		4:  "QUIT_REASON_KICK_BY_HOST_LOGOUT",
		5:  "QUIT_REASON_KICK_BY_HOST_BLOCK",
		6:  "QUIT_REASON_BE_BLOCKED",
		7:  "QUIT_REASON_KICK_BY_HOST_ENTER_HOME",
		8:  "QUIT_REASON_HOST_SCENE_INVALID",
		9:  "QUIT_REASON_KICK_BY_PLAY",
		10: "QUIT_REASON_Unk2800_FDECHAHJFDA",
	}
	PlayerQuitFromMpNotify_QuitReason_value = map[string]int32{
		"QUIT_REASON_INVALID":                 0,
		"QUIT_REASON_HOST_NO_OTHER_PLAYER":    1,
		"QUIT_REASON_KICK_BY_HOST":            2,
		"QUIT_REASON_BACK_TO_MY_WORLD":        3,
		"QUIT_REASON_KICK_BY_HOST_LOGOUT":     4,
		"QUIT_REASON_KICK_BY_HOST_BLOCK":      5,
		"QUIT_REASON_BE_BLOCKED":              6,
		"QUIT_REASON_KICK_BY_HOST_ENTER_HOME": 7,
		"QUIT_REASON_HOST_SCENE_INVALID":      8,
		"QUIT_REASON_KICK_BY_PLAY":            9,
		"QUIT_REASON_Unk2800_FDECHAHJFDA":     10,
	}
)

func (x PlayerQuitFromMpNotify_QuitReason) Enum() *PlayerQuitFromMpNotify_QuitReason {
	p := new(PlayerQuitFromMpNotify_QuitReason)
	*p = x
	return p
}

func (x PlayerQuitFromMpNotify_QuitReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerQuitFromMpNotify_QuitReason) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerQuitFromMpNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerQuitFromMpNotify_QuitReason) Type() protoreflect.EnumType {
	return &file_PlayerQuitFromMpNotify_proto_enumTypes[0]
}

func (x PlayerQuitFromMpNotify_QuitReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerQuitFromMpNotify_QuitReason.Descriptor instead.
func (PlayerQuitFromMpNotify_QuitReason) EnumDescriptor() ([]byte, []int) {
	return file_PlayerQuitFromMpNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 1829
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerQuitFromMpNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason PlayerQuitFromMpNotify_QuitReason `protobuf:"varint,11,opt,name=reason,proto3,enum=PlayerQuitFromMpNotify_QuitReason" json:"reason,omitempty"`
}

func (x *PlayerQuitFromMpNotify) Reset() {
	*x = PlayerQuitFromMpNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerQuitFromMpNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerQuitFromMpNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerQuitFromMpNotify) ProtoMessage() {}

func (x *PlayerQuitFromMpNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerQuitFromMpNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerQuitFromMpNotify.ProtoReflect.Descriptor instead.
func (*PlayerQuitFromMpNotify) Descriptor() ([]byte, []int) {
	return file_PlayerQuitFromMpNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerQuitFromMpNotify) GetReason() PlayerQuitFromMpNotify_QuitReason {
	if x != nil {
		return x.Reason
	}
	return PlayerQuitFromMpNotify_QUIT_REASON_INVALID
}

var File_PlayerQuitFromMpNotify_proto protoreflect.FileDescriptor

var file_PlayerQuitFromMpNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x69, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7,
	0x03, 0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x51, 0x75, 0x69, 0x74, 0x46, 0x72, 0x6f,
	0x6d, 0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x51, 0x75, 0x69, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x80, 0x03, 0x0a, 0x0a, 0x51, 0x75, 0x69, 0x74, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41,
	0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a,
	0x20, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x48, 0x4f, 0x53,
	0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x59, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x10,
	0x02, 0x12, 0x20, 0x0a, 0x1c, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x54, 0x4f, 0x5f, 0x4d, 0x59, 0x5f, 0x57, 0x4f, 0x52, 0x4c,
	0x44, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x59, 0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x4c, 0x4f, 0x47, 0x4f, 0x55, 0x54, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x51, 0x55, 0x49, 0x54,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x59, 0x5f,
	0x48, 0x4f, 0x53, 0x54, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16,
	0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x42, 0x45, 0x5f, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x06, 0x12, 0x27, 0x0a, 0x23, 0x51, 0x55, 0x49, 0x54,
	0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x59, 0x5f,
	0x48, 0x4f, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x10,
	0x07, 0x12, 0x22, 0x0a, 0x1e, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x42, 0x59, 0x5f, 0x50, 0x4c, 0x41,
	0x59, 0x10, 0x09, 0x12, 0x23, 0x0a, 0x1f, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x46, 0x44, 0x45, 0x43, 0x48,
	0x41, 0x48, 0x4a, 0x46, 0x44, 0x41, 0x10, 0x0a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerQuitFromMpNotify_proto_rawDescOnce sync.Once
	file_PlayerQuitFromMpNotify_proto_rawDescData = file_PlayerQuitFromMpNotify_proto_rawDesc
)

func file_PlayerQuitFromMpNotify_proto_rawDescGZIP() []byte {
	file_PlayerQuitFromMpNotify_proto_rawDescOnce.Do(func() {
		file_PlayerQuitFromMpNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerQuitFromMpNotify_proto_rawDescData)
	})
	return file_PlayerQuitFromMpNotify_proto_rawDescData
}

var file_PlayerQuitFromMpNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerQuitFromMpNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerQuitFromMpNotify_proto_goTypes = []interface{}{
	(PlayerQuitFromMpNotify_QuitReason)(0), // 0: PlayerQuitFromMpNotify.QuitReason
	(*PlayerQuitFromMpNotify)(nil),         // 1: PlayerQuitFromMpNotify
}
var file_PlayerQuitFromMpNotify_proto_depIdxs = []int32{
	0, // 0: PlayerQuitFromMpNotify.reason:type_name -> PlayerQuitFromMpNotify.QuitReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerQuitFromMpNotify_proto_init() }
func file_PlayerQuitFromMpNotify_proto_init() {
	if File_PlayerQuitFromMpNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerQuitFromMpNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerQuitFromMpNotify); i {
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
			RawDescriptor: file_PlayerQuitFromMpNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerQuitFromMpNotify_proto_goTypes,
		DependencyIndexes: file_PlayerQuitFromMpNotify_proto_depIdxs,
		EnumInfos:         file_PlayerQuitFromMpNotify_proto_enumTypes,
		MessageInfos:      file_PlayerQuitFromMpNotify_proto_msgTypes,
	}.Build()
	File_PlayerQuitFromMpNotify_proto = out.File
	file_PlayerQuitFromMpNotify_proto_rawDesc = nil
	file_PlayerQuitFromMpNotify_proto_goTypes = nil
	file_PlayerQuitFromMpNotify_proto_depIdxs = nil
}
