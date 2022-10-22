// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SelectEffigyChallengeConditionRsp.proto

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

// CmdId: 2039
// EnetChannelId: 0
// EnetIsReliable: true
type SelectEffigyChallengeConditionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConditionIdList []uint32 `protobuf:"varint,12,rep,packed,name=condition_id_list,json=conditionIdList,proto3" json:"condition_id_list,omitempty"`
	Retcode         int32    `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DifficultyId    uint32   `protobuf:"varint,7,opt,name=difficulty_id,json=difficultyId,proto3" json:"difficulty_id,omitempty"`
	ChallengeId     uint32   `protobuf:"varint,2,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
}

func (x *SelectEffigyChallengeConditionRsp) Reset() {
	*x = SelectEffigyChallengeConditionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SelectEffigyChallengeConditionRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectEffigyChallengeConditionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectEffigyChallengeConditionRsp) ProtoMessage() {}

func (x *SelectEffigyChallengeConditionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SelectEffigyChallengeConditionRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectEffigyChallengeConditionRsp.ProtoReflect.Descriptor instead.
func (*SelectEffigyChallengeConditionRsp) Descriptor() ([]byte, []int) {
	return file_SelectEffigyChallengeConditionRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SelectEffigyChallengeConditionRsp) GetConditionIdList() []uint32 {
	if x != nil {
		return x.ConditionIdList
	}
	return nil
}

func (x *SelectEffigyChallengeConditionRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SelectEffigyChallengeConditionRsp) GetDifficultyId() uint32 {
	if x != nil {
		return x.DifficultyId
	}
	return 0
}

func (x *SelectEffigyChallengeConditionRsp) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

var File_SelectEffigyChallengeConditionRsp_proto protoreflect.FileDescriptor

var file_SelectEffigyChallengeConditionRsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x45, 0x66, 0x66, 0x69, 0x67, 0x79, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x21, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x45, 0x66, 0x66, 0x69, 0x67, 0x79, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12,
	0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SelectEffigyChallengeConditionRsp_proto_rawDescOnce sync.Once
	file_SelectEffigyChallengeConditionRsp_proto_rawDescData = file_SelectEffigyChallengeConditionRsp_proto_rawDesc
)

func file_SelectEffigyChallengeConditionRsp_proto_rawDescGZIP() []byte {
	file_SelectEffigyChallengeConditionRsp_proto_rawDescOnce.Do(func() {
		file_SelectEffigyChallengeConditionRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SelectEffigyChallengeConditionRsp_proto_rawDescData)
	})
	return file_SelectEffigyChallengeConditionRsp_proto_rawDescData
}

var file_SelectEffigyChallengeConditionRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SelectEffigyChallengeConditionRsp_proto_goTypes = []interface{}{
	(*SelectEffigyChallengeConditionRsp)(nil), // 0: SelectEffigyChallengeConditionRsp
}
var file_SelectEffigyChallengeConditionRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SelectEffigyChallengeConditionRsp_proto_init() }
func file_SelectEffigyChallengeConditionRsp_proto_init() {
	if File_SelectEffigyChallengeConditionRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SelectEffigyChallengeConditionRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectEffigyChallengeConditionRsp); i {
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
			RawDescriptor: file_SelectEffigyChallengeConditionRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SelectEffigyChallengeConditionRsp_proto_goTypes,
		DependencyIndexes: file_SelectEffigyChallengeConditionRsp_proto_depIdxs,
		MessageInfos:      file_SelectEffigyChallengeConditionRsp_proto_msgTypes,
	}.Build()
	File_SelectEffigyChallengeConditionRsp_proto = out.File
	file_SelectEffigyChallengeConditionRsp_proto_rawDesc = nil
	file_SelectEffigyChallengeConditionRsp_proto_goTypes = nil
	file_SelectEffigyChallengeConditionRsp_proto_depIdxs = nil
}
