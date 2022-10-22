// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AvatarAddNotify.proto

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

// CmdId: 1769
// EnetChannelId: 0
// EnetIsReliable: true
type AvatarAddNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avatar   *AvatarInfo `protobuf:"bytes,13,opt,name=avatar,proto3" json:"avatar,omitempty"`
	IsInTeam bool        `protobuf:"varint,12,opt,name=is_in_team,json=isInTeam,proto3" json:"is_in_team,omitempty"`
}

func (x *AvatarAddNotify) Reset() {
	*x = AvatarAddNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarAddNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarAddNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarAddNotify) ProtoMessage() {}

func (x *AvatarAddNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarAddNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarAddNotify.ProtoReflect.Descriptor instead.
func (*AvatarAddNotify) Descriptor() ([]byte, []int) {
	return file_AvatarAddNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarAddNotify) GetAvatar() *AvatarInfo {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *AvatarAddNotify) GetIsInTeam() bool {
	if x != nil {
		return x.IsInTeam
	}
	return false
}

var File_AvatarAddNotify_proto protoreflect.FileDescriptor

var file_AvatarAddNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0f, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x1c, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x49, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_AvatarAddNotify_proto_rawDescOnce sync.Once
	file_AvatarAddNotify_proto_rawDescData = file_AvatarAddNotify_proto_rawDesc
)

func file_AvatarAddNotify_proto_rawDescGZIP() []byte {
	file_AvatarAddNotify_proto_rawDescOnce.Do(func() {
		file_AvatarAddNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarAddNotify_proto_rawDescData)
	})
	return file_AvatarAddNotify_proto_rawDescData
}

var file_AvatarAddNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarAddNotify_proto_goTypes = []interface{}{
	(*AvatarAddNotify)(nil), // 0: AvatarAddNotify
	(*AvatarInfo)(nil),      // 1: AvatarInfo
}
var file_AvatarAddNotify_proto_depIdxs = []int32{
	1, // 0: AvatarAddNotify.avatar:type_name -> AvatarInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AvatarAddNotify_proto_init() }
func file_AvatarAddNotify_proto_init() {
	if File_AvatarAddNotify_proto != nil {
		return
	}
	file_AvatarInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AvatarAddNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarAddNotify); i {
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
			RawDescriptor: file_AvatarAddNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarAddNotify_proto_goTypes,
		DependencyIndexes: file_AvatarAddNotify_proto_depIdxs,
		MessageInfos:      file_AvatarAddNotify_proto_msgTypes,
	}.Build()
	File_AvatarAddNotify_proto = out.File
	file_AvatarAddNotify_proto_rawDesc = nil
	file_AvatarAddNotify_proto_goTypes = nil
	file_AvatarAddNotify_proto_depIdxs = nil
}
