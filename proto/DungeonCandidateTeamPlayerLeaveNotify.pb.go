// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: DungeonCandidateTeamPlayerLeaveNotify.proto

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

// CmdId: 926
// EnetChannelId: 0
// EnetIsReliable: true
type DungeonCandidateTeamPlayerLeaveNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason    DungeonCandidateTeamPlayerLeaveReason `protobuf:"varint,3,opt,name=reason,proto3,enum=DungeonCandidateTeamPlayerLeaveReason" json:"reason,omitempty"`
	PlayerUid uint32                                `protobuf:"varint,13,opt,name=player_uid,json=playerUid,proto3" json:"player_uid,omitempty"`
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) Reset() {
	*x = DungeonCandidateTeamPlayerLeaveNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonCandidateTeamPlayerLeaveNotify) ProtoMessage() {}

func (x *DungeonCandidateTeamPlayerLeaveNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonCandidateTeamPlayerLeaveNotify.ProtoReflect.Descriptor instead.
func (*DungeonCandidateTeamPlayerLeaveNotify) Descriptor() ([]byte, []int) {
	return file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) GetReason() DungeonCandidateTeamPlayerLeaveReason {
	if x != nil {
		return x.Reason
	}
	return DungeonCandidateTeamPlayerLeaveReason_DUNGEON_CANDIDATE_TEAM_PLAYER_LEAVE_REASON_TPLR_NORMAL
}

func (x *DungeonCandidateTeamPlayerLeaveNotify) GetPlayerUid() uint32 {
	if x != nil {
		return x.PlayerUid
	}
	return 0
}

var File_DungeonCandidateTeamPlayerLeaveNotify_proto protoreflect.FileDescriptor

var file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x25, 0x44,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x75,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescOnce sync.Once
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData = file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc
)

func file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescGZIP() []byte {
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescOnce.Do(func() {
		file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData)
	})
	return file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDescData
}

var file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DungeonCandidateTeamPlayerLeaveNotify_proto_goTypes = []interface{}{
	(*DungeonCandidateTeamPlayerLeaveNotify)(nil), // 0: DungeonCandidateTeamPlayerLeaveNotify
	(DungeonCandidateTeamPlayerLeaveReason)(0),    // 1: DungeonCandidateTeamPlayerLeaveReason
}
var file_DungeonCandidateTeamPlayerLeaveNotify_proto_depIdxs = []int32{
	1, // 0: DungeonCandidateTeamPlayerLeaveNotify.reason:type_name -> DungeonCandidateTeamPlayerLeaveReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DungeonCandidateTeamPlayerLeaveNotify_proto_init() }
func file_DungeonCandidateTeamPlayerLeaveNotify_proto_init() {
	if File_DungeonCandidateTeamPlayerLeaveNotify_proto != nil {
		return
	}
	file_DungeonCandidateTeamPlayerLeaveReason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonCandidateTeamPlayerLeaveNotify); i {
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
			RawDescriptor: file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonCandidateTeamPlayerLeaveNotify_proto_goTypes,
		DependencyIndexes: file_DungeonCandidateTeamPlayerLeaveNotify_proto_depIdxs,
		MessageInfos:      file_DungeonCandidateTeamPlayerLeaveNotify_proto_msgTypes,
	}.Build()
	File_DungeonCandidateTeamPlayerLeaveNotify_proto = out.File
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_rawDesc = nil
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_goTypes = nil
	file_DungeonCandidateTeamPlayerLeaveNotify_proto_depIdxs = nil
}