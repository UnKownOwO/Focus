// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: TowerMonthlyBrief.proto

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

type TowerMonthlyBrief struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TowerScheduleId uint32 `protobuf:"varint,15,opt,name=tower_schedule_id,json=towerScheduleId,proto3" json:"tower_schedule_id,omitempty"`
	BestFloorIndex  uint32 `protobuf:"varint,6,opt,name=best_floor_index,json=bestFloorIndex,proto3" json:"best_floor_index,omitempty"`
	BestLevelIndex  uint32 `protobuf:"varint,3,opt,name=best_level_index,json=bestLevelIndex,proto3" json:"best_level_index,omitempty"`
	TotalStarCount  uint32 `protobuf:"varint,12,opt,name=total_star_count,json=totalStarCount,proto3" json:"total_star_count,omitempty"`
}

func (x *TowerMonthlyBrief) Reset() {
	*x = TowerMonthlyBrief{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TowerMonthlyBrief_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerMonthlyBrief) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerMonthlyBrief) ProtoMessage() {}

func (x *TowerMonthlyBrief) ProtoReflect() protoreflect.Message {
	mi := &file_TowerMonthlyBrief_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerMonthlyBrief.ProtoReflect.Descriptor instead.
func (*TowerMonthlyBrief) Descriptor() ([]byte, []int) {
	return file_TowerMonthlyBrief_proto_rawDescGZIP(), []int{0}
}

func (x *TowerMonthlyBrief) GetTowerScheduleId() uint32 {
	if x != nil {
		return x.TowerScheduleId
	}
	return 0
}

func (x *TowerMonthlyBrief) GetBestFloorIndex() uint32 {
	if x != nil {
		return x.BestFloorIndex
	}
	return 0
}

func (x *TowerMonthlyBrief) GetBestLevelIndex() uint32 {
	if x != nil {
		return x.BestLevelIndex
	}
	return 0
}

func (x *TowerMonthlyBrief) GetTotalStarCount() uint32 {
	if x != nil {
		return x.TotalStarCount
	}
	return 0
}

var File_TowerMonthlyBrief_proto protoreflect.FileDescriptor

var file_TowerMonthlyBrief_proto_rawDesc = []byte{
	0x0a, 0x17, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x72,
	0x69, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x11, 0x54, 0x6f,
	0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x42, 0x72, 0x69, 0x65, 0x66, 0x12,
	0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x74, 0x6f, 0x77, 0x65,
	0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x62,
	0x65, 0x73, 0x74, 0x5f, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x65, 0x73, 0x74, 0x46, 0x6c, 0x6f, 0x6f, 0x72,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x62, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TowerMonthlyBrief_proto_rawDescOnce sync.Once
	file_TowerMonthlyBrief_proto_rawDescData = file_TowerMonthlyBrief_proto_rawDesc
)

func file_TowerMonthlyBrief_proto_rawDescGZIP() []byte {
	file_TowerMonthlyBrief_proto_rawDescOnce.Do(func() {
		file_TowerMonthlyBrief_proto_rawDescData = protoimpl.X.CompressGZIP(file_TowerMonthlyBrief_proto_rawDescData)
	})
	return file_TowerMonthlyBrief_proto_rawDescData
}

var file_TowerMonthlyBrief_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TowerMonthlyBrief_proto_goTypes = []interface{}{
	(*TowerMonthlyBrief)(nil), // 0: TowerMonthlyBrief
}
var file_TowerMonthlyBrief_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TowerMonthlyBrief_proto_init() }
func file_TowerMonthlyBrief_proto_init() {
	if File_TowerMonthlyBrief_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TowerMonthlyBrief_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerMonthlyBrief); i {
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
			RawDescriptor: file_TowerMonthlyBrief_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TowerMonthlyBrief_proto_goTypes,
		DependencyIndexes: file_TowerMonthlyBrief_proto_depIdxs,
		MessageInfos:      file_TowerMonthlyBrief_proto_msgTypes,
	}.Build()
	File_TowerMonthlyBrief_proto = out.File
	file_TowerMonthlyBrief_proto_rawDesc = nil
	file_TowerMonthlyBrief_proto_goTypes = nil
	file_TowerMonthlyBrief_proto_depIdxs = nil
}
