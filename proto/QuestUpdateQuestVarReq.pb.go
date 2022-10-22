// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: QuestUpdateQuestVarReq.proto

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

// CmdId: 447
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type QuestUpdateQuestVarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentQuestId     uint32        `protobuf:"varint,9,opt,name=parent_quest_id,json=parentQuestId,proto3" json:"parent_quest_id,omitempty"`
	QuestVarOpList    []*QuestVarOp `protobuf:"bytes,4,rep,name=quest_var_op_list,json=questVarOpList,proto3" json:"quest_var_op_list,omitempty"`
	QuestId           uint32        `protobuf:"varint,11,opt,name=quest_id,json=questId,proto3" json:"quest_id,omitempty"`
	ParentQuestVarSeq uint32        `protobuf:"varint,1,opt,name=parent_quest_var_seq,json=parentQuestVarSeq,proto3" json:"parent_quest_var_seq,omitempty"`
}

func (x *QuestUpdateQuestVarReq) Reset() {
	*x = QuestUpdateQuestVarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_QuestUpdateQuestVarReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestUpdateQuestVarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestUpdateQuestVarReq) ProtoMessage() {}

func (x *QuestUpdateQuestVarReq) ProtoReflect() protoreflect.Message {
	mi := &file_QuestUpdateQuestVarReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestUpdateQuestVarReq.ProtoReflect.Descriptor instead.
func (*QuestUpdateQuestVarReq) Descriptor() ([]byte, []int) {
	return file_QuestUpdateQuestVarReq_proto_rawDescGZIP(), []int{0}
}

func (x *QuestUpdateQuestVarReq) GetParentQuestId() uint32 {
	if x != nil {
		return x.ParentQuestId
	}
	return 0
}

func (x *QuestUpdateQuestVarReq) GetQuestVarOpList() []*QuestVarOp {
	if x != nil {
		return x.QuestVarOpList
	}
	return nil
}

func (x *QuestUpdateQuestVarReq) GetQuestId() uint32 {
	if x != nil {
		return x.QuestId
	}
	return 0
}

func (x *QuestUpdateQuestVarReq) GetParentQuestVarSeq() uint32 {
	if x != nil {
		return x.ParentQuestVarSeq
	}
	return 0
}

var File_QuestUpdateQuestVarReq_proto protoreflect.FileDescriptor

var file_QuestUpdateQuestVarReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72, 0x4f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc4, 0x01, 0x0a, 0x16, 0x51, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0f, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x11, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x72,
	0x5f, 0x6f, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x56, 0x61, 0x72, 0x4f, 0x70, 0x52, 0x0e, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x61, 0x72, 0x4f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x56, 0x61, 0x72, 0x53, 0x65, 0x71, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_QuestUpdateQuestVarReq_proto_rawDescOnce sync.Once
	file_QuestUpdateQuestVarReq_proto_rawDescData = file_QuestUpdateQuestVarReq_proto_rawDesc
)

func file_QuestUpdateQuestVarReq_proto_rawDescGZIP() []byte {
	file_QuestUpdateQuestVarReq_proto_rawDescOnce.Do(func() {
		file_QuestUpdateQuestVarReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_QuestUpdateQuestVarReq_proto_rawDescData)
	})
	return file_QuestUpdateQuestVarReq_proto_rawDescData
}

var file_QuestUpdateQuestVarReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_QuestUpdateQuestVarReq_proto_goTypes = []interface{}{
	(*QuestUpdateQuestVarReq)(nil), // 0: QuestUpdateQuestVarReq
	(*QuestVarOp)(nil),             // 1: QuestVarOp
}
var file_QuestUpdateQuestVarReq_proto_depIdxs = []int32{
	1, // 0: QuestUpdateQuestVarReq.quest_var_op_list:type_name -> QuestVarOp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_QuestUpdateQuestVarReq_proto_init() }
func file_QuestUpdateQuestVarReq_proto_init() {
	if File_QuestUpdateQuestVarReq_proto != nil {
		return
	}
	file_QuestVarOp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_QuestUpdateQuestVarReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestUpdateQuestVarReq); i {
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
			RawDescriptor: file_QuestUpdateQuestVarReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_QuestUpdateQuestVarReq_proto_goTypes,
		DependencyIndexes: file_QuestUpdateQuestVarReq_proto_depIdxs,
		MessageInfos:      file_QuestUpdateQuestVarReq_proto_msgTypes,
	}.Build()
	File_QuestUpdateQuestVarReq_proto = out.File
	file_QuestUpdateQuestVarReq_proto_rawDesc = nil
	file_QuestUpdateQuestVarReq_proto_goTypes = nil
	file_QuestUpdateQuestVarReq_proto_depIdxs = nil
}
