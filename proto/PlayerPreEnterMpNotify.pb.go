// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: PlayerPreEnterMpNotify.proto

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

type PlayerPreEnterMpNotify_State int32

const (
	PlayerPreEnterMpNotify_STATE_INVALID PlayerPreEnterMpNotify_State = 0
	PlayerPreEnterMpNotify_STATE_START   PlayerPreEnterMpNotify_State = 1
	PlayerPreEnterMpNotify_STATE_TIMEOUT PlayerPreEnterMpNotify_State = 2
)

// Enum value maps for PlayerPreEnterMpNotify_State.
var (
	PlayerPreEnterMpNotify_State_name = map[int32]string{
		0: "STATE_INVALID",
		1: "STATE_START",
		2: "STATE_TIMEOUT",
	}
	PlayerPreEnterMpNotify_State_value = map[string]int32{
		"STATE_INVALID": 0,
		"STATE_START":   1,
		"STATE_TIMEOUT": 2,
	}
)

func (x PlayerPreEnterMpNotify_State) Enum() *PlayerPreEnterMpNotify_State {
	p := new(PlayerPreEnterMpNotify_State)
	*p = x
	return p
}

func (x PlayerPreEnterMpNotify_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerPreEnterMpNotify_State) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerPreEnterMpNotify_proto_enumTypes[0].Descriptor()
}

func (PlayerPreEnterMpNotify_State) Type() protoreflect.EnumType {
	return &file_PlayerPreEnterMpNotify_proto_enumTypes[0]
}

func (x PlayerPreEnterMpNotify_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerPreEnterMpNotify_State.Descriptor instead.
func (PlayerPreEnterMpNotify_State) EnumDescriptor() ([]byte, []int) {
	return file_PlayerPreEnterMpNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 1822
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerPreEnterMpNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    PlayerPreEnterMpNotify_State `protobuf:"varint,2,opt,name=state,proto3,enum=PlayerPreEnterMpNotify_State" json:"state,omitempty"`
	Uid      uint32                       `protobuf:"varint,14,opt,name=uid,proto3" json:"uid,omitempty"`
	Nickname string                       `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *PlayerPreEnterMpNotify) Reset() {
	*x = PlayerPreEnterMpNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerPreEnterMpNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPreEnterMpNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPreEnterMpNotify) ProtoMessage() {}

func (x *PlayerPreEnterMpNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerPreEnterMpNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPreEnterMpNotify.ProtoReflect.Descriptor instead.
func (*PlayerPreEnterMpNotify) Descriptor() ([]byte, []int) {
	return file_PlayerPreEnterMpNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerPreEnterMpNotify) GetState() PlayerPreEnterMpNotify_State {
	if x != nil {
		return x.State
	}
	return PlayerPreEnterMpNotify_STATE_INVALID
}

func (x *PlayerPreEnterMpNotify) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlayerPreEnterMpNotify) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_PlayerPreEnterMpNotify_proto protoreflect.FileDescriptor

var file_PlayerPreEnterMpNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb,
	0x01, 0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x50, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerPreEnterMpNotify_proto_rawDescOnce sync.Once
	file_PlayerPreEnterMpNotify_proto_rawDescData = file_PlayerPreEnterMpNotify_proto_rawDesc
)

func file_PlayerPreEnterMpNotify_proto_rawDescGZIP() []byte {
	file_PlayerPreEnterMpNotify_proto_rawDescOnce.Do(func() {
		file_PlayerPreEnterMpNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerPreEnterMpNotify_proto_rawDescData)
	})
	return file_PlayerPreEnterMpNotify_proto_rawDescData
}

var file_PlayerPreEnterMpNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerPreEnterMpNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerPreEnterMpNotify_proto_goTypes = []interface{}{
	(PlayerPreEnterMpNotify_State)(0), // 0: PlayerPreEnterMpNotify.State
	(*PlayerPreEnterMpNotify)(nil),    // 1: PlayerPreEnterMpNotify
}
var file_PlayerPreEnterMpNotify_proto_depIdxs = []int32{
	0, // 0: PlayerPreEnterMpNotify.state:type_name -> PlayerPreEnterMpNotify.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PlayerPreEnterMpNotify_proto_init() }
func file_PlayerPreEnterMpNotify_proto_init() {
	if File_PlayerPreEnterMpNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerPreEnterMpNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerPreEnterMpNotify); i {
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
			RawDescriptor: file_PlayerPreEnterMpNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerPreEnterMpNotify_proto_goTypes,
		DependencyIndexes: file_PlayerPreEnterMpNotify_proto_depIdxs,
		EnumInfos:         file_PlayerPreEnterMpNotify_proto_enumTypes,
		MessageInfos:      file_PlayerPreEnterMpNotify_proto_msgTypes,
	}.Build()
	File_PlayerPreEnterMpNotify_proto = out.File
	file_PlayerPreEnterMpNotify_proto_rawDesc = nil
	file_PlayerPreEnterMpNotify_proto_goTypes = nil
	file_PlayerPreEnterMpNotify_proto_depIdxs = nil
}
