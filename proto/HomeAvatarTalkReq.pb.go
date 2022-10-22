// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HomeAvatarTalkReq.proto

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

// CmdId: 4688
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type HomeAvatarTalkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TalkId   uint32 `protobuf:"varint,12,opt,name=talk_id,json=talkId,proto3" json:"talk_id,omitempty"`
	AvatarId uint32 `protobuf:"varint,15,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
}

func (x *HomeAvatarTalkReq) Reset() {
	*x = HomeAvatarTalkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeAvatarTalkReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeAvatarTalkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeAvatarTalkReq) ProtoMessage() {}

func (x *HomeAvatarTalkReq) ProtoReflect() protoreflect.Message {
	mi := &file_HomeAvatarTalkReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeAvatarTalkReq.ProtoReflect.Descriptor instead.
func (*HomeAvatarTalkReq) Descriptor() ([]byte, []int) {
	return file_HomeAvatarTalkReq_proto_rawDescGZIP(), []int{0}
}

func (x *HomeAvatarTalkReq) GetTalkId() uint32 {
	if x != nil {
		return x.TalkId
	}
	return 0
}

func (x *HomeAvatarTalkReq) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

var File_HomeAvatarTalkReq_proto protoreflect.FileDescriptor

var file_HomeAvatarTalkReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x61, 0x6c, 0x6b,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x11, 0x48, 0x6f, 0x6d,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x61, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x6c, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x74, 0x61, 0x6c, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeAvatarTalkReq_proto_rawDescOnce sync.Once
	file_HomeAvatarTalkReq_proto_rawDescData = file_HomeAvatarTalkReq_proto_rawDesc
)

func file_HomeAvatarTalkReq_proto_rawDescGZIP() []byte {
	file_HomeAvatarTalkReq_proto_rawDescOnce.Do(func() {
		file_HomeAvatarTalkReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeAvatarTalkReq_proto_rawDescData)
	})
	return file_HomeAvatarTalkReq_proto_rawDescData
}

var file_HomeAvatarTalkReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeAvatarTalkReq_proto_goTypes = []interface{}{
	(*HomeAvatarTalkReq)(nil), // 0: HomeAvatarTalkReq
}
var file_HomeAvatarTalkReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HomeAvatarTalkReq_proto_init() }
func file_HomeAvatarTalkReq_proto_init() {
	if File_HomeAvatarTalkReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HomeAvatarTalkReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeAvatarTalkReq); i {
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
			RawDescriptor: file_HomeAvatarTalkReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeAvatarTalkReq_proto_goTypes,
		DependencyIndexes: file_HomeAvatarTalkReq_proto_depIdxs,
		MessageInfos:      file_HomeAvatarTalkReq_proto_msgTypes,
	}.Build()
	File_HomeAvatarTalkReq_proto = out.File
	file_HomeAvatarTalkReq_proto_rawDesc = nil
	file_HomeAvatarTalkReq_proto_goTypes = nil
	file_HomeAvatarTalkReq_proto_depIdxs = nil
}
