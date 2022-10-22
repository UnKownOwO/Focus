// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: TrialAvatarInfo.proto

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

type TrialAvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrialAvatarId  uint32                  `protobuf:"varint,1,opt,name=trial_avatar_id,json=trialAvatarId,proto3" json:"trial_avatar_id,omitempty"`
	TrialEquipList []*Item                 `protobuf:"bytes,2,rep,name=trial_equip_list,json=trialEquipList,proto3" json:"trial_equip_list,omitempty"`
	GrantRecord    *TrialAvatarGrantRecord `protobuf:"bytes,3,opt,name=grant_record,json=grantRecord,proto3" json:"grant_record,omitempty"`
}

func (x *TrialAvatarInfo) Reset() {
	*x = TrialAvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TrialAvatarInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrialAvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrialAvatarInfo) ProtoMessage() {}

func (x *TrialAvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TrialAvatarInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrialAvatarInfo.ProtoReflect.Descriptor instead.
func (*TrialAvatarInfo) Descriptor() ([]byte, []int) {
	return file_TrialAvatarInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TrialAvatarInfo) GetTrialAvatarId() uint32 {
	if x != nil {
		return x.TrialAvatarId
	}
	return 0
}

func (x *TrialAvatarInfo) GetTrialEquipList() []*Item {
	if x != nil {
		return x.TrialEquipList
	}
	return nil
}

func (x *TrialAvatarInfo) GetGrantRecord() *TrialAvatarGrantRecord {
	if x != nil {
		return x.GrantRecord
	}
	return nil
}

var File_TrialAvatarInfo_proto protoreflect.FileDescriptor

var file_TrialAvatarInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x10, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0e,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x71, 0x75, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x0c, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TrialAvatarInfo_proto_rawDescOnce sync.Once
	file_TrialAvatarInfo_proto_rawDescData = file_TrialAvatarInfo_proto_rawDesc
)

func file_TrialAvatarInfo_proto_rawDescGZIP() []byte {
	file_TrialAvatarInfo_proto_rawDescOnce.Do(func() {
		file_TrialAvatarInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TrialAvatarInfo_proto_rawDescData)
	})
	return file_TrialAvatarInfo_proto_rawDescData
}

var file_TrialAvatarInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TrialAvatarInfo_proto_goTypes = []interface{}{
	(*TrialAvatarInfo)(nil),        // 0: TrialAvatarInfo
	(*Item)(nil),                   // 1: Item
	(*TrialAvatarGrantRecord)(nil), // 2: TrialAvatarGrantRecord
}
var file_TrialAvatarInfo_proto_depIdxs = []int32{
	1, // 0: TrialAvatarInfo.trial_equip_list:type_name -> Item
	2, // 1: TrialAvatarInfo.grant_record:type_name -> TrialAvatarGrantRecord
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TrialAvatarInfo_proto_init() }
func file_TrialAvatarInfo_proto_init() {
	if File_TrialAvatarInfo_proto != nil {
		return
	}
	file_Item_proto_init()
	file_TrialAvatarGrantRecord_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TrialAvatarInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrialAvatarInfo); i {
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
			RawDescriptor: file_TrialAvatarInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TrialAvatarInfo_proto_goTypes,
		DependencyIndexes: file_TrialAvatarInfo_proto_depIdxs,
		MessageInfos:      file_TrialAvatarInfo_proto_msgTypes,
	}.Build()
	File_TrialAvatarInfo_proto = out.File
	file_TrialAvatarInfo_proto_rawDesc = nil
	file_TrialAvatarInfo_proto_goTypes = nil
	file_TrialAvatarInfo_proto_depIdxs = nil
}
