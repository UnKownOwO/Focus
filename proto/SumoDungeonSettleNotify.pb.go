// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SumoDungeonSettleNotify.proto

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

// CmdId: 8291
// EnetChannelId: 0
// EnetIsReliable: true
type SumoDungeonSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinalScore          uint32 `protobuf:"varint,7,opt,name=final_score,json=finalScore,proto3" json:"final_score,omitempty"`
	DifficultyId        uint32 `protobuf:"varint,14,opt,name=difficulty_id,json=difficultyId,proto3" json:"difficulty_id,omitempty"`
	KillEliteMonsterNum uint32 `protobuf:"varint,15,opt,name=kill_elite_monster_num,json=killEliteMonsterNum,proto3" json:"kill_elite_monster_num,omitempty"`
	StageId             uint32 `protobuf:"varint,12,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	KillMonsterNum      uint32 `protobuf:"varint,4,opt,name=kill_monster_num,json=killMonsterNum,proto3" json:"kill_monster_num,omitempty"`
	IsNewRecord         bool   `protobuf:"varint,5,opt,name=is_new_record,json=isNewRecord,proto3" json:"is_new_record,omitempty"`
}

func (x *SumoDungeonSettleNotify) Reset() {
	*x = SumoDungeonSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SumoDungeonSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumoDungeonSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumoDungeonSettleNotify) ProtoMessage() {}

func (x *SumoDungeonSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SumoDungeonSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumoDungeonSettleNotify.ProtoReflect.Descriptor instead.
func (*SumoDungeonSettleNotify) Descriptor() ([]byte, []int) {
	return file_SumoDungeonSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SumoDungeonSettleNotify) GetFinalScore() uint32 {
	if x != nil {
		return x.FinalScore
	}
	return 0
}

func (x *SumoDungeonSettleNotify) GetDifficultyId() uint32 {
	if x != nil {
		return x.DifficultyId
	}
	return 0
}

func (x *SumoDungeonSettleNotify) GetKillEliteMonsterNum() uint32 {
	if x != nil {
		return x.KillEliteMonsterNum
	}
	return 0
}

func (x *SumoDungeonSettleNotify) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *SumoDungeonSettleNotify) GetKillMonsterNum() uint32 {
	if x != nil {
		return x.KillMonsterNum
	}
	return 0
}

func (x *SumoDungeonSettleNotify) GetIsNewRecord() bool {
	if x != nil {
		return x.IsNewRecord
	}
	return false
}

var File_SumoDungeonSettleNotify_proto protoreflect.FileDescriptor

var file_SumoDungeonSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x75, 0x6d, 0x6f, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfd, 0x01, 0x0a, 0x17, 0x53, 0x75, 0x6d, 0x6f, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x16, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x6c, 0x69, 0x74, 0x65, 0x5f,
	0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x13, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x6c, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x10, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6b, 0x69, 0x6c,
	0x6c, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x69,
	0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SumoDungeonSettleNotify_proto_rawDescOnce sync.Once
	file_SumoDungeonSettleNotify_proto_rawDescData = file_SumoDungeonSettleNotify_proto_rawDesc
)

func file_SumoDungeonSettleNotify_proto_rawDescGZIP() []byte {
	file_SumoDungeonSettleNotify_proto_rawDescOnce.Do(func() {
		file_SumoDungeonSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SumoDungeonSettleNotify_proto_rawDescData)
	})
	return file_SumoDungeonSettleNotify_proto_rawDescData
}

var file_SumoDungeonSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SumoDungeonSettleNotify_proto_goTypes = []interface{}{
	(*SumoDungeonSettleNotify)(nil), // 0: SumoDungeonSettleNotify
}
var file_SumoDungeonSettleNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SumoDungeonSettleNotify_proto_init() }
func file_SumoDungeonSettleNotify_proto_init() {
	if File_SumoDungeonSettleNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SumoDungeonSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumoDungeonSettleNotify); i {
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
			RawDescriptor: file_SumoDungeonSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SumoDungeonSettleNotify_proto_goTypes,
		DependencyIndexes: file_SumoDungeonSettleNotify_proto_depIdxs,
		MessageInfos:      file_SumoDungeonSettleNotify_proto_msgTypes,
	}.Build()
	File_SumoDungeonSettleNotify_proto = out.File
	file_SumoDungeonSettleNotify_proto_rawDesc = nil
	file_SumoDungeonSettleNotify_proto_goTypes = nil
	file_SumoDungeonSettleNotify_proto_depIdxs = nil
}