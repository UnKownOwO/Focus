// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ChangeMpTeamAvatarRsp.proto

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

// CmdId: 1753
// EnetChannelId: 0
// EnetIsReliable: true
type ChangeMpTeamAvatarRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode        int32    `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AvatarGuidList []uint64 `protobuf:"varint,3,rep,packed,name=avatar_guid_list,json=avatarGuidList,proto3" json:"avatar_guid_list,omitempty"`
	CurAvatarGuid  uint64   `protobuf:"varint,13,opt,name=cur_avatar_guid,json=curAvatarGuid,proto3" json:"cur_avatar_guid,omitempty"`
}

func (x *ChangeMpTeamAvatarRsp) Reset() {
	*x = ChangeMpTeamAvatarRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChangeMpTeamAvatarRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeMpTeamAvatarRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeMpTeamAvatarRsp) ProtoMessage() {}

func (x *ChangeMpTeamAvatarRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChangeMpTeamAvatarRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeMpTeamAvatarRsp.ProtoReflect.Descriptor instead.
func (*ChangeMpTeamAvatarRsp) Descriptor() ([]byte, []int) {
	return file_ChangeMpTeamAvatarRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeMpTeamAvatarRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChangeMpTeamAvatarRsp) GetAvatarGuidList() []uint64 {
	if x != nil {
		return x.AvatarGuidList
	}
	return nil
}

func (x *ChangeMpTeamAvatarRsp) GetCurAvatarGuid() uint64 {
	if x != nil {
		return x.CurAvatarGuid
	}
	return 0
}

var File_ChangeMpTeamAvatarRsp_proto protoreflect.FileDescriptor

var file_ChangeMpTeamAvatarRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x70, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01,
	0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x70, 0x54, 0x65, 0x61, 0x6d, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0e, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x75, 0x72, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x47,
	0x75, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChangeMpTeamAvatarRsp_proto_rawDescOnce sync.Once
	file_ChangeMpTeamAvatarRsp_proto_rawDescData = file_ChangeMpTeamAvatarRsp_proto_rawDesc
)

func file_ChangeMpTeamAvatarRsp_proto_rawDescGZIP() []byte {
	file_ChangeMpTeamAvatarRsp_proto_rawDescOnce.Do(func() {
		file_ChangeMpTeamAvatarRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChangeMpTeamAvatarRsp_proto_rawDescData)
	})
	return file_ChangeMpTeamAvatarRsp_proto_rawDescData
}

var file_ChangeMpTeamAvatarRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChangeMpTeamAvatarRsp_proto_goTypes = []interface{}{
	(*ChangeMpTeamAvatarRsp)(nil), // 0: ChangeMpTeamAvatarRsp
}
var file_ChangeMpTeamAvatarRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChangeMpTeamAvatarRsp_proto_init() }
func file_ChangeMpTeamAvatarRsp_proto_init() {
	if File_ChangeMpTeamAvatarRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChangeMpTeamAvatarRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeMpTeamAvatarRsp); i {
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
			RawDescriptor: file_ChangeMpTeamAvatarRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChangeMpTeamAvatarRsp_proto_goTypes,
		DependencyIndexes: file_ChangeMpTeamAvatarRsp_proto_depIdxs,
		MessageInfos:      file_ChangeMpTeamAvatarRsp_proto_msgTypes,
	}.Build()
	File_ChangeMpTeamAvatarRsp_proto = out.File
	file_ChangeMpTeamAvatarRsp_proto_rawDesc = nil
	file_ChangeMpTeamAvatarRsp_proto_goTypes = nil
	file_ChangeMpTeamAvatarRsp_proto_depIdxs = nil
}