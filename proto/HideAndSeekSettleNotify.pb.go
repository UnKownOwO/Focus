// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HideAndSeekSettleNotify.proto

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

type HideAndSeekSettleNotify_SettleReason int32

const (
	HideAndSeekSettleNotify_SETTLE_REASON_TIME_OUT    HideAndSeekSettleNotify_SettleReason = 0
	HideAndSeekSettleNotify_SETTLE_REASON_PLAY_END    HideAndSeekSettleNotify_SettleReason = 1
	HideAndSeekSettleNotify_SETTLE_REASON_PLAYER_QUIT HideAndSeekSettleNotify_SettleReason = 2
)

// Enum value maps for HideAndSeekSettleNotify_SettleReason.
var (
	HideAndSeekSettleNotify_SettleReason_name = map[int32]string{
		0: "SETTLE_REASON_TIME_OUT",
		1: "SETTLE_REASON_PLAY_END",
		2: "SETTLE_REASON_PLAYER_QUIT",
	}
	HideAndSeekSettleNotify_SettleReason_value = map[string]int32{
		"SETTLE_REASON_TIME_OUT":    0,
		"SETTLE_REASON_PLAY_END":    1,
		"SETTLE_REASON_PLAYER_QUIT": 2,
	}
)

func (x HideAndSeekSettleNotify_SettleReason) Enum() *HideAndSeekSettleNotify_SettleReason {
	p := new(HideAndSeekSettleNotify_SettleReason)
	*p = x
	return p
}

func (x HideAndSeekSettleNotify_SettleReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HideAndSeekSettleNotify_SettleReason) Descriptor() protoreflect.EnumDescriptor {
	return file_HideAndSeekSettleNotify_proto_enumTypes[0].Descriptor()
}

func (HideAndSeekSettleNotify_SettleReason) Type() protoreflect.EnumType {
	return &file_HideAndSeekSettleNotify_proto_enumTypes[0]
}

func (x HideAndSeekSettleNotify_SettleReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HideAndSeekSettleNotify_SettleReason.Descriptor instead.
func (HideAndSeekSettleNotify_SettleReason) EnumDescriptor() ([]byte, []int) {
	return file_HideAndSeekSettleNotify_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 5317
// EnetChannelId: 0
// EnetIsReliable: true
type HideAndSeekSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostTime       uint32                               `protobuf:"varint,2,opt,name=cost_time,json=costTime,proto3" json:"cost_time,omitempty"`
	SettleInfoList []*HideAndSeekSettleInfo             `protobuf:"bytes,8,rep,name=settle_info_list,json=settleInfoList,proto3" json:"settle_info_list,omitempty"`
	WinnerList     []uint32                             `protobuf:"varint,15,rep,packed,name=winner_list,json=winnerList,proto3" json:"winner_list,omitempty"`
	Reason         HideAndSeekSettleNotify_SettleReason `protobuf:"varint,4,opt,name=reason,proto3,enum=HideAndSeekSettleNotify_SettleReason" json:"reason,omitempty"`
	PlayIndex      uint32                               `protobuf:"varint,13,opt,name=play_index,json=playIndex,proto3" json:"play_index,omitempty"`
	IsRecordScore  bool                                 `protobuf:"varint,6,opt,name=is_record_score,json=isRecordScore,proto3" json:"is_record_score,omitempty"`
	ScoreList      []*ExhibitionDisplayInfo             `protobuf:"bytes,9,rep,name=score_list,json=scoreList,proto3" json:"score_list,omitempty"`
	StageType      uint32                               `protobuf:"varint,14,opt,name=stage_type,json=stageType,proto3" json:"stage_type,omitempty"`
}

func (x *HideAndSeekSettleNotify) Reset() {
	*x = HideAndSeekSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HideAndSeekSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HideAndSeekSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HideAndSeekSettleNotify) ProtoMessage() {}

func (x *HideAndSeekSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HideAndSeekSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HideAndSeekSettleNotify.ProtoReflect.Descriptor instead.
func (*HideAndSeekSettleNotify) Descriptor() ([]byte, []int) {
	return file_HideAndSeekSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HideAndSeekSettleNotify) GetCostTime() uint32 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *HideAndSeekSettleNotify) GetSettleInfoList() []*HideAndSeekSettleInfo {
	if x != nil {
		return x.SettleInfoList
	}
	return nil
}

func (x *HideAndSeekSettleNotify) GetWinnerList() []uint32 {
	if x != nil {
		return x.WinnerList
	}
	return nil
}

func (x *HideAndSeekSettleNotify) GetReason() HideAndSeekSettleNotify_SettleReason {
	if x != nil {
		return x.Reason
	}
	return HideAndSeekSettleNotify_SETTLE_REASON_TIME_OUT
}

func (x *HideAndSeekSettleNotify) GetPlayIndex() uint32 {
	if x != nil {
		return x.PlayIndex
	}
	return 0
}

func (x *HideAndSeekSettleNotify) GetIsRecordScore() bool {
	if x != nil {
		return x.IsRecordScore
	}
	return false
}

func (x *HideAndSeekSettleNotify) GetScoreList() []*ExhibitionDisplayInfo {
	if x != nil {
		return x.ScoreList
	}
	return nil
}

func (x *HideAndSeekSettleNotify) GetStageType() uint32 {
	if x != nil {
		return x.StageType
	}
	return 0
}

var File_HideAndSeekSettleNotify_proto protoreflect.FileDescriptor

var file_HideAndSeekSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x48, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x48, 0x69,
	0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x17, 0x48, 0x69,
	0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x10, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x48,
	0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x48, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53,
	0x65, 0x65, 0x6b, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x45, 0x78, 0x68, 0x69, 0x62, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x65, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x16, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x45, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x5f, 0x51, 0x55, 0x49, 0x54, 0x10, 0x02, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HideAndSeekSettleNotify_proto_rawDescOnce sync.Once
	file_HideAndSeekSettleNotify_proto_rawDescData = file_HideAndSeekSettleNotify_proto_rawDesc
)

func file_HideAndSeekSettleNotify_proto_rawDescGZIP() []byte {
	file_HideAndSeekSettleNotify_proto_rawDescOnce.Do(func() {
		file_HideAndSeekSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HideAndSeekSettleNotify_proto_rawDescData)
	})
	return file_HideAndSeekSettleNotify_proto_rawDescData
}

var file_HideAndSeekSettleNotify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HideAndSeekSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HideAndSeekSettleNotify_proto_goTypes = []interface{}{
	(HideAndSeekSettleNotify_SettleReason)(0), // 0: HideAndSeekSettleNotify.SettleReason
	(*HideAndSeekSettleNotify)(nil),           // 1: HideAndSeekSettleNotify
	(*HideAndSeekSettleInfo)(nil),             // 2: HideAndSeekSettleInfo
	(*ExhibitionDisplayInfo)(nil),             // 3: ExhibitionDisplayInfo
}
var file_HideAndSeekSettleNotify_proto_depIdxs = []int32{
	2, // 0: HideAndSeekSettleNotify.settle_info_list:type_name -> HideAndSeekSettleInfo
	0, // 1: HideAndSeekSettleNotify.reason:type_name -> HideAndSeekSettleNotify.SettleReason
	3, // 2: HideAndSeekSettleNotify.score_list:type_name -> ExhibitionDisplayInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_HideAndSeekSettleNotify_proto_init() }
func file_HideAndSeekSettleNotify_proto_init() {
	if File_HideAndSeekSettleNotify_proto != nil {
		return
	}
	file_ExhibitionDisplayInfo_proto_init()
	file_HideAndSeekSettleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HideAndSeekSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HideAndSeekSettleNotify); i {
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
			RawDescriptor: file_HideAndSeekSettleNotify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HideAndSeekSettleNotify_proto_goTypes,
		DependencyIndexes: file_HideAndSeekSettleNotify_proto_depIdxs,
		EnumInfos:         file_HideAndSeekSettleNotify_proto_enumTypes,
		MessageInfos:      file_HideAndSeekSettleNotify_proto_msgTypes,
	}.Build()
	File_HideAndSeekSettleNotify_proto = out.File
	file_HideAndSeekSettleNotify_proto_rawDesc = nil
	file_HideAndSeekSettleNotify_proto_goTypes = nil
	file_HideAndSeekSettleNotify_proto_depIdxs = nil
}
