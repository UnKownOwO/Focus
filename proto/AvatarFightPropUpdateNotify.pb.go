// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AvatarFightPropUpdateNotify.proto

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

// CmdId: 1221
// EnetChannelId: 0
// EnetIsReliable: true
type AvatarFightPropUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FightPropMap map[uint32]float32 `protobuf:"bytes,15,rep,name=fight_prop_map,json=fightPropMap,proto3" json:"fight_prop_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	AvatarGuid   uint64             `protobuf:"varint,13,opt,name=avatar_guid,json=avatarGuid,proto3" json:"avatar_guid,omitempty"`
}

func (x *AvatarFightPropUpdateNotify) Reset() {
	*x = AvatarFightPropUpdateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarFightPropUpdateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarFightPropUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarFightPropUpdateNotify) ProtoMessage() {}

func (x *AvatarFightPropUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarFightPropUpdateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarFightPropUpdateNotify.ProtoReflect.Descriptor instead.
func (*AvatarFightPropUpdateNotify) Descriptor() ([]byte, []int) {
	return file_AvatarFightPropUpdateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarFightPropUpdateNotify) GetFightPropMap() map[uint32]float32 {
	if x != nil {
		return x.FightPropMap
	}
	return nil
}

func (x *AvatarFightPropUpdateNotify) GetAvatarGuid() uint64 {
	if x != nil {
		return x.AvatarGuid
	}
	return 0
}

var File_AvatarFightPropUpdateNotify_proto protoreflect.FileDescriptor

var file_AvatarFightPropUpdateNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f,
	0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x1b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x69,
	0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x54, 0x0a, 0x0e, 0x66, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x70, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x46, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x46, 0x69, 0x67, 0x68, 0x74, 0x50,
	0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x66, 0x69, 0x67,
	0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x1a, 0x3f, 0x0a, 0x11, 0x46, 0x69,
	0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarFightPropUpdateNotify_proto_rawDescOnce sync.Once
	file_AvatarFightPropUpdateNotify_proto_rawDescData = file_AvatarFightPropUpdateNotify_proto_rawDesc
)

func file_AvatarFightPropUpdateNotify_proto_rawDescGZIP() []byte {
	file_AvatarFightPropUpdateNotify_proto_rawDescOnce.Do(func() {
		file_AvatarFightPropUpdateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarFightPropUpdateNotify_proto_rawDescData)
	})
	return file_AvatarFightPropUpdateNotify_proto_rawDescData
}

var file_AvatarFightPropUpdateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_AvatarFightPropUpdateNotify_proto_goTypes = []interface{}{
	(*AvatarFightPropUpdateNotify)(nil), // 0: AvatarFightPropUpdateNotify
	nil,                                 // 1: AvatarFightPropUpdateNotify.FightPropMapEntry
}
var file_AvatarFightPropUpdateNotify_proto_depIdxs = []int32{
	1, // 0: AvatarFightPropUpdateNotify.fight_prop_map:type_name -> AvatarFightPropUpdateNotify.FightPropMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AvatarFightPropUpdateNotify_proto_init() }
func file_AvatarFightPropUpdateNotify_proto_init() {
	if File_AvatarFightPropUpdateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarFightPropUpdateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarFightPropUpdateNotify); i {
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
			RawDescriptor: file_AvatarFightPropUpdateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarFightPropUpdateNotify_proto_goTypes,
		DependencyIndexes: file_AvatarFightPropUpdateNotify_proto_depIdxs,
		MessageInfos:      file_AvatarFightPropUpdateNotify_proto_msgTypes,
	}.Build()
	File_AvatarFightPropUpdateNotify_proto = out.File
	file_AvatarFightPropUpdateNotify_proto_rawDesc = nil
	file_AvatarFightPropUpdateNotify_proto_goTypes = nil
	file_AvatarFightPropUpdateNotify_proto_depIdxs = nil
}
