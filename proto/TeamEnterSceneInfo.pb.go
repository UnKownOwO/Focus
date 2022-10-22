// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: TeamEnterSceneInfo.proto

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

type TeamEnterSceneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbilityControlBlock *AbilityControlBlock  `protobuf:"bytes,7,opt,name=ability_control_block,json=abilityControlBlock,proto3" json:"ability_control_block,omitempty"`
	TeamAbilityInfo     *AbilitySyncStateInfo `protobuf:"bytes,10,opt,name=team_ability_info,json=teamAbilityInfo,proto3" json:"team_ability_info,omitempty"`
	TeamEntityId        uint32                `protobuf:"varint,15,opt,name=team_entity_id,json=teamEntityId,proto3" json:"team_entity_id,omitempty"`
}

func (x *TeamEnterSceneInfo) Reset() {
	*x = TeamEnterSceneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TeamEnterSceneInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamEnterSceneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamEnterSceneInfo) ProtoMessage() {}

func (x *TeamEnterSceneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TeamEnterSceneInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamEnterSceneInfo.ProtoReflect.Descriptor instead.
func (*TeamEnterSceneInfo) Descriptor() ([]byte, []int) {
	return file_TeamEnterSceneInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TeamEnterSceneInfo) GetAbilityControlBlock() *AbilityControlBlock {
	if x != nil {
		return x.AbilityControlBlock
	}
	return nil
}

func (x *TeamEnterSceneInfo) GetTeamAbilityInfo() *AbilitySyncStateInfo {
	if x != nil {
		return x.TeamAbilityInfo
	}
	return nil
}

func (x *TeamEnterSceneInfo) GetTeamEntityId() uint32 {
	if x != nil {
		return x.TeamEntityId
	}
	return 0
}

var File_TeamEnterSceneInfo_proto protoreflect.FileDescriptor

var file_TeamEnterSceneInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x41, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x12, 0x54, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x48, 0x0a, 0x15, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x13, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x41, 0x0a, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x74, 0x65, 0x61, 0x6d, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74,
	0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TeamEnterSceneInfo_proto_rawDescOnce sync.Once
	file_TeamEnterSceneInfo_proto_rawDescData = file_TeamEnterSceneInfo_proto_rawDesc
)

func file_TeamEnterSceneInfo_proto_rawDescGZIP() []byte {
	file_TeamEnterSceneInfo_proto_rawDescOnce.Do(func() {
		file_TeamEnterSceneInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TeamEnterSceneInfo_proto_rawDescData)
	})
	return file_TeamEnterSceneInfo_proto_rawDescData
}

var file_TeamEnterSceneInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TeamEnterSceneInfo_proto_goTypes = []interface{}{
	(*TeamEnterSceneInfo)(nil),   // 0: TeamEnterSceneInfo
	(*AbilityControlBlock)(nil),  // 1: AbilityControlBlock
	(*AbilitySyncStateInfo)(nil), // 2: AbilitySyncStateInfo
}
var file_TeamEnterSceneInfo_proto_depIdxs = []int32{
	1, // 0: TeamEnterSceneInfo.ability_control_block:type_name -> AbilityControlBlock
	2, // 1: TeamEnterSceneInfo.team_ability_info:type_name -> AbilitySyncStateInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TeamEnterSceneInfo_proto_init() }
func file_TeamEnterSceneInfo_proto_init() {
	if File_TeamEnterSceneInfo_proto != nil {
		return
	}
	file_AbilityControlBlock_proto_init()
	file_AbilitySyncStateInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TeamEnterSceneInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamEnterSceneInfo); i {
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
			RawDescriptor: file_TeamEnterSceneInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TeamEnterSceneInfo_proto_goTypes,
		DependencyIndexes: file_TeamEnterSceneInfo_proto_depIdxs,
		MessageInfos:      file_TeamEnterSceneInfo_proto_msgTypes,
	}.Build()
	File_TeamEnterSceneInfo_proto = out.File
	file_TeamEnterSceneInfo_proto_rawDesc = nil
	file_TeamEnterSceneInfo_proto_goTypes = nil
	file_TeamEnterSceneInfo_proto_depIdxs = nil
}
