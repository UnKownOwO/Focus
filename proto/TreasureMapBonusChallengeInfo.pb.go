// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: TreasureMapBonusChallengeInfo.proto

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

type TreasureMapBonusChallengeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsDone      bool            `protobuf:"varint,5,opt,name=is_done,json=isDone,proto3" json:"is_done,omitempty"`
	ConfigId    uint32          `protobuf:"varint,10,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	IsActive    bool            `protobuf:"varint,1,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	FragmentMap map[uint32]bool `protobuf:"bytes,12,rep,name=fragment_map,json=fragmentMap,proto3" json:"fragment_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	SolutionId  uint32          `protobuf:"varint,8,opt,name=solution_id,json=solutionId,proto3" json:"solution_id,omitempty"`
}

func (x *TreasureMapBonusChallengeInfo) Reset() {
	*x = TreasureMapBonusChallengeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureMapBonusChallengeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureMapBonusChallengeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureMapBonusChallengeInfo) ProtoMessage() {}

func (x *TreasureMapBonusChallengeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureMapBonusChallengeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureMapBonusChallengeInfo.ProtoReflect.Descriptor instead.
func (*TreasureMapBonusChallengeInfo) Descriptor() ([]byte, []int) {
	return file_TreasureMapBonusChallengeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureMapBonusChallengeInfo) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *TreasureMapBonusChallengeInfo) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *TreasureMapBonusChallengeInfo) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *TreasureMapBonusChallengeInfo) GetFragmentMap() map[uint32]bool {
	if x != nil {
		return x.FragmentMap
	}
	return nil
}

func (x *TreasureMapBonusChallengeInfo) GetSolutionId() uint32 {
	if x != nil {
		return x.SolutionId
	}
	return 0
}

var File_TreasureMapBonusChallengeInfo_proto protoreflect.FileDescriptor

var file_TreasureMapBonusChallengeInfo_proto_rawDesc = []byte{
	0x0a, 0x23, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6e,
	0x75, 0x73, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x1d, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x64, 0x6f,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x52, 0x0a, 0x0c, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x42, 0x6f,
	0x6e, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a,
	0x3e, 0x0a, 0x10, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_TreasureMapBonusChallengeInfo_proto_rawDescOnce sync.Once
	file_TreasureMapBonusChallengeInfo_proto_rawDescData = file_TreasureMapBonusChallengeInfo_proto_rawDesc
)

func file_TreasureMapBonusChallengeInfo_proto_rawDescGZIP() []byte {
	file_TreasureMapBonusChallengeInfo_proto_rawDescOnce.Do(func() {
		file_TreasureMapBonusChallengeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureMapBonusChallengeInfo_proto_rawDescData)
	})
	return file_TreasureMapBonusChallengeInfo_proto_rawDescData
}

var file_TreasureMapBonusChallengeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_TreasureMapBonusChallengeInfo_proto_goTypes = []interface{}{
	(*TreasureMapBonusChallengeInfo)(nil), // 0: TreasureMapBonusChallengeInfo
	nil,                                   // 1: TreasureMapBonusChallengeInfo.FragmentMapEntry
}
var file_TreasureMapBonusChallengeInfo_proto_depIdxs = []int32{
	1, // 0: TreasureMapBonusChallengeInfo.fragment_map:type_name -> TreasureMapBonusChallengeInfo.FragmentMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TreasureMapBonusChallengeInfo_proto_init() }
func file_TreasureMapBonusChallengeInfo_proto_init() {
	if File_TreasureMapBonusChallengeInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TreasureMapBonusChallengeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureMapBonusChallengeInfo); i {
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
			RawDescriptor: file_TreasureMapBonusChallengeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureMapBonusChallengeInfo_proto_goTypes,
		DependencyIndexes: file_TreasureMapBonusChallengeInfo_proto_depIdxs,
		MessageInfos:      file_TreasureMapBonusChallengeInfo_proto_msgTypes,
	}.Build()
	File_TreasureMapBonusChallengeInfo_proto = out.File
	file_TreasureMapBonusChallengeInfo_proto_rawDesc = nil
	file_TreasureMapBonusChallengeInfo_proto_goTypes = nil
	file_TreasureMapBonusChallengeInfo_proto_depIdxs = nil
}
