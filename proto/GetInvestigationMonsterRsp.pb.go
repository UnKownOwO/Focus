// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GetInvestigationMonsterRsp.proto

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

// CmdId: 1910
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type GetInvestigationMonsterRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonsterList         []*InvestigationMonster `protobuf:"bytes,10,rep,name=monster_list,json=monsterList,proto3" json:"monster_list,omitempty"`
	Retcode             int32                   `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Unk2700_DEMFDHNFBBJ bool                    `protobuf:"varint,2,opt,name=Unk2700_DEMFDHNFBBJ,json=Unk2700DEMFDHNFBBJ,proto3" json:"Unk2700_DEMFDHNFBBJ,omitempty"`
}

func (x *GetInvestigationMonsterRsp) Reset() {
	*x = GetInvestigationMonsterRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetInvestigationMonsterRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvestigationMonsterRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvestigationMonsterRsp) ProtoMessage() {}

func (x *GetInvestigationMonsterRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetInvestigationMonsterRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvestigationMonsterRsp.ProtoReflect.Descriptor instead.
func (*GetInvestigationMonsterRsp) Descriptor() ([]byte, []int) {
	return file_GetInvestigationMonsterRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetInvestigationMonsterRsp) GetMonsterList() []*InvestigationMonster {
	if x != nil {
		return x.MonsterList
	}
	return nil
}

func (x *GetInvestigationMonsterRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetInvestigationMonsterRsp) GetUnk2700_DEMFDHNFBBJ() bool {
	if x != nil {
		return x.Unk2700_DEMFDHNFBBJ
	}
	return false
}

var File_GetInvestigationMonsterRsp_proto protoreflect.FileDescriptor

var file_GetInvestigationMonsterRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1,
	0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a,
	0x0c, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x45, 0x4d,
	0x46, 0x44, 0x48, 0x4e, 0x46, 0x42, 0x42, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x45, 0x4d, 0x46, 0x44, 0x48, 0x4e, 0x46, 0x42,
	0x42, 0x4a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetInvestigationMonsterRsp_proto_rawDescOnce sync.Once
	file_GetInvestigationMonsterRsp_proto_rawDescData = file_GetInvestigationMonsterRsp_proto_rawDesc
)

func file_GetInvestigationMonsterRsp_proto_rawDescGZIP() []byte {
	file_GetInvestigationMonsterRsp_proto_rawDescOnce.Do(func() {
		file_GetInvestigationMonsterRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetInvestigationMonsterRsp_proto_rawDescData)
	})
	return file_GetInvestigationMonsterRsp_proto_rawDescData
}

var file_GetInvestigationMonsterRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetInvestigationMonsterRsp_proto_goTypes = []interface{}{
	(*GetInvestigationMonsterRsp)(nil), // 0: GetInvestigationMonsterRsp
	(*InvestigationMonster)(nil),       // 1: InvestigationMonster
}
var file_GetInvestigationMonsterRsp_proto_depIdxs = []int32{
	1, // 0: GetInvestigationMonsterRsp.monster_list:type_name -> InvestigationMonster
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetInvestigationMonsterRsp_proto_init() }
func file_GetInvestigationMonsterRsp_proto_init() {
	if File_GetInvestigationMonsterRsp_proto != nil {
		return
	}
	file_InvestigationMonster_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetInvestigationMonsterRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvestigationMonsterRsp); i {
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
			RawDescriptor: file_GetInvestigationMonsterRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetInvestigationMonsterRsp_proto_goTypes,
		DependencyIndexes: file_GetInvestigationMonsterRsp_proto_depIdxs,
		MessageInfos:      file_GetInvestigationMonsterRsp_proto_msgTypes,
	}.Build()
	File_GetInvestigationMonsterRsp_proto = out.File
	file_GetInvestigationMonsterRsp_proto_rawDesc = nil
	file_GetInvestigationMonsterRsp_proto_goTypes = nil
	file_GetInvestigationMonsterRsp_proto_depIdxs = nil
}
