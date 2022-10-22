// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: InBattleMechanicusSettleInfo.proto

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

type InBattleMechanicusSettleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneTimeMs          uint64                         `protobuf:"varint,15,opt,name=scene_time_ms,json=sceneTimeMs,proto3" json:"scene_time_ms,omitempty"`
	TotalToken           uint32                         `protobuf:"varint,4,opt,name=total_token,json=totalToken,proto3" json:"total_token,omitempty"`
	RealToken            uint32                         `protobuf:"varint,8,opt,name=real_token,json=realToken,proto3" json:"real_token,omitempty"`
	WatcherList          []*MultistageSettleWatcherInfo `protobuf:"bytes,7,rep,name=watcher_list,json=watcherList,proto3" json:"watcher_list,omitempty"`
	IsSuccess            bool                           `protobuf:"varint,6,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	PlayIndex            uint32                         `protobuf:"varint,3,opt,name=play_index,json=playIndex,proto3" json:"play_index,omitempty"`
	DifficultyPercentage uint32                         `protobuf:"varint,10,opt,name=difficulty_percentage,json=difficultyPercentage,proto3" json:"difficulty_percentage,omitempty"`
	GroupId              uint32                         `protobuf:"varint,13,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *InBattleMechanicusSettleInfo) Reset() {
	*x = InBattleMechanicusSettleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleMechanicusSettleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleMechanicusSettleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleMechanicusSettleInfo) ProtoMessage() {}

func (x *InBattleMechanicusSettleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleMechanicusSettleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleMechanicusSettleInfo.ProtoReflect.Descriptor instead.
func (*InBattleMechanicusSettleInfo) Descriptor() ([]byte, []int) {
	return file_InBattleMechanicusSettleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleMechanicusSettleInfo) GetSceneTimeMs() uint64 {
	if x != nil {
		return x.SceneTimeMs
	}
	return 0
}

func (x *InBattleMechanicusSettleInfo) GetTotalToken() uint32 {
	if x != nil {
		return x.TotalToken
	}
	return 0
}

func (x *InBattleMechanicusSettleInfo) GetRealToken() uint32 {
	if x != nil {
		return x.RealToken
	}
	return 0
}

func (x *InBattleMechanicusSettleInfo) GetWatcherList() []*MultistageSettleWatcherInfo {
	if x != nil {
		return x.WatcherList
	}
	return nil
}

func (x *InBattleMechanicusSettleInfo) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *InBattleMechanicusSettleInfo) GetPlayIndex() uint32 {
	if x != nil {
		return x.PlayIndex
	}
	return 0
}

func (x *InBattleMechanicusSettleInfo) GetDifficultyPercentage() uint32 {
	if x != nil {
		return x.DifficultyPercentage
	}
	return 0
}

func (x *InBattleMechanicusSettleInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_InBattleMechanicusSettleInfo_proto protoreflect.FileDescriptor

var file_InBattleMechanicusSettleInfo_proto_rawDesc = []byte{
	0x0a, 0x22, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x1c, 0x49, 0x6e, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x63, 0x65, 0x6e,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3f, 0x0a, 0x0c,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0b, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x33, 0x0a, 0x15, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleMechanicusSettleInfo_proto_rawDescOnce sync.Once
	file_InBattleMechanicusSettleInfo_proto_rawDescData = file_InBattleMechanicusSettleInfo_proto_rawDesc
)

func file_InBattleMechanicusSettleInfo_proto_rawDescGZIP() []byte {
	file_InBattleMechanicusSettleInfo_proto_rawDescOnce.Do(func() {
		file_InBattleMechanicusSettleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleMechanicusSettleInfo_proto_rawDescData)
	})
	return file_InBattleMechanicusSettleInfo_proto_rawDescData
}

var file_InBattleMechanicusSettleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InBattleMechanicusSettleInfo_proto_goTypes = []interface{}{
	(*InBattleMechanicusSettleInfo)(nil), // 0: InBattleMechanicusSettleInfo
	(*MultistageSettleWatcherInfo)(nil),  // 1: MultistageSettleWatcherInfo
}
var file_InBattleMechanicusSettleInfo_proto_depIdxs = []int32{
	1, // 0: InBattleMechanicusSettleInfo.watcher_list:type_name -> MultistageSettleWatcherInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_InBattleMechanicusSettleInfo_proto_init() }
func file_InBattleMechanicusSettleInfo_proto_init() {
	if File_InBattleMechanicusSettleInfo_proto != nil {
		return
	}
	file_MultistageSettleWatcherInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InBattleMechanicusSettleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleMechanicusSettleInfo); i {
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
			RawDescriptor: file_InBattleMechanicusSettleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleMechanicusSettleInfo_proto_goTypes,
		DependencyIndexes: file_InBattleMechanicusSettleInfo_proto_depIdxs,
		MessageInfos:      file_InBattleMechanicusSettleInfo_proto_msgTypes,
	}.Build()
	File_InBattleMechanicusSettleInfo_proto = out.File
	file_InBattleMechanicusSettleInfo_proto_rawDesc = nil
	file_InBattleMechanicusSettleInfo_proto_goTypes = nil
	file_InBattleMechanicusSettleInfo_proto_depIdxs = nil
}
