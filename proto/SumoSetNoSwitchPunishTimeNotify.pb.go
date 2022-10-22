// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SumoSetNoSwitchPunishTimeNotify.proto

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

// CmdId: 8935
// EnetChannelId: 0
// EnetIsReliable: true
type SumoSetNoSwitchPunishTimeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurTeamIndex        uint32             `protobuf:"varint,15,opt,name=cur_team_index,json=curTeamIndex,proto3" json:"cur_team_index,omitempty"`
	StageId             uint32             `protobuf:"varint,13,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	DungeonTeamList     []*SumoDungeonTeam `protobuf:"bytes,11,rep,name=dungeon_team_list,json=dungeonTeamList,proto3" json:"dungeon_team_list,omitempty"`
	NoSwitchPunishTime  uint32             `protobuf:"varint,2,opt,name=no_switch_punish_time,json=noSwitchPunishTime,proto3" json:"no_switch_punish_time,omitempty"`
	NextValidSwitchTime uint32             `protobuf:"varint,14,opt,name=next_valid_switch_time,json=nextValidSwitchTime,proto3" json:"next_valid_switch_time,omitempty"`
	ActivityId          uint32             `protobuf:"varint,9,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *SumoSetNoSwitchPunishTimeNotify) Reset() {
	*x = SumoSetNoSwitchPunishTimeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SumoSetNoSwitchPunishTimeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumoSetNoSwitchPunishTimeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumoSetNoSwitchPunishTimeNotify) ProtoMessage() {}

func (x *SumoSetNoSwitchPunishTimeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SumoSetNoSwitchPunishTimeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumoSetNoSwitchPunishTimeNotify.ProtoReflect.Descriptor instead.
func (*SumoSetNoSwitchPunishTimeNotify) Descriptor() ([]byte, []int) {
	return file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SumoSetNoSwitchPunishTimeNotify) GetCurTeamIndex() uint32 {
	if x != nil {
		return x.CurTeamIndex
	}
	return 0
}

func (x *SumoSetNoSwitchPunishTimeNotify) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *SumoSetNoSwitchPunishTimeNotify) GetDungeonTeamList() []*SumoDungeonTeam {
	if x != nil {
		return x.DungeonTeamList
	}
	return nil
}

func (x *SumoSetNoSwitchPunishTimeNotify) GetNoSwitchPunishTime() uint32 {
	if x != nil {
		return x.NoSwitchPunishTime
	}
	return 0
}

func (x *SumoSetNoSwitchPunishTimeNotify) GetNextValidSwitchTime() uint32 {
	if x != nil {
		return x.NextValidSwitchTime
	}
	return 0
}

func (x *SumoSetNoSwitchPunishTimeNotify) GetActivityId() uint32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

var File_SumoSetNoSwitchPunishTimeNotify_proto protoreflect.FileDescriptor

var file_SumoSetNoSwitchPunishTimeNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x53, 0x75, 0x6d, 0x6f, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x50, 0x75, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x75, 0x6d, 0x6f, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9,
	0x02, 0x0a, 0x1f, 0x53, 0x75, 0x6d, 0x6f, 0x53, 0x65, 0x74, 0x4e, 0x6f, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x50, 0x75, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x5f, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x54,
	0x65, 0x61, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x11, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x53, 0x75, 0x6d, 0x6f, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x0f, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x54, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x15, 0x6e, 0x6f, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x70,
	0x75, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x6e, 0x6f, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x50, 0x75, 0x6e, 0x69, 0x73, 0x68,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6e, 0x65, 0x78, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x53,
	0x77, 0x69, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescOnce sync.Once
	file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescData = file_SumoSetNoSwitchPunishTimeNotify_proto_rawDesc
)

func file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescGZIP() []byte {
	file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescOnce.Do(func() {
		file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescData)
	})
	return file_SumoSetNoSwitchPunishTimeNotify_proto_rawDescData
}

var file_SumoSetNoSwitchPunishTimeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SumoSetNoSwitchPunishTimeNotify_proto_goTypes = []interface{}{
	(*SumoSetNoSwitchPunishTimeNotify)(nil), // 0: SumoSetNoSwitchPunishTimeNotify
	(*SumoDungeonTeam)(nil),                 // 1: SumoDungeonTeam
}
var file_SumoSetNoSwitchPunishTimeNotify_proto_depIdxs = []int32{
	1, // 0: SumoSetNoSwitchPunishTimeNotify.dungeon_team_list:type_name -> SumoDungeonTeam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SumoSetNoSwitchPunishTimeNotify_proto_init() }
func file_SumoSetNoSwitchPunishTimeNotify_proto_init() {
	if File_SumoSetNoSwitchPunishTimeNotify_proto != nil {
		return
	}
	file_SumoDungeonTeam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SumoSetNoSwitchPunishTimeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumoSetNoSwitchPunishTimeNotify); i {
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
			RawDescriptor: file_SumoSetNoSwitchPunishTimeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SumoSetNoSwitchPunishTimeNotify_proto_goTypes,
		DependencyIndexes: file_SumoSetNoSwitchPunishTimeNotify_proto_depIdxs,
		MessageInfos:      file_SumoSetNoSwitchPunishTimeNotify_proto_msgTypes,
	}.Build()
	File_SumoSetNoSwitchPunishTimeNotify_proto = out.File
	file_SumoSetNoSwitchPunishTimeNotify_proto_rawDesc = nil
	file_SumoSetNoSwitchPunishTimeNotify_proto_goTypes = nil
	file_SumoSetNoSwitchPunishTimeNotify_proto_depIdxs = nil
}
