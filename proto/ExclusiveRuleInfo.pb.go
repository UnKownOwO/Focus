// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ExclusiveRuleInfo.proto

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

type ExclusiveRuleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectIdList []uint32 `protobuf:"varint,1,rep,packed,name=object_id_list,json=objectIdList,proto3" json:"object_id_list,omitempty"`
	RuleType     uint32   `protobuf:"varint,10,opt,name=rule_type,json=ruleType,proto3" json:"rule_type,omitempty"`
}

func (x *ExclusiveRuleInfo) Reset() {
	*x = ExclusiveRuleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExclusiveRuleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExclusiveRuleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExclusiveRuleInfo) ProtoMessage() {}

func (x *ExclusiveRuleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ExclusiveRuleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExclusiveRuleInfo.ProtoReflect.Descriptor instead.
func (*ExclusiveRuleInfo) Descriptor() ([]byte, []int) {
	return file_ExclusiveRuleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ExclusiveRuleInfo) GetObjectIdList() []uint32 {
	if x != nil {
		return x.ObjectIdList
	}
	return nil
}

func (x *ExclusiveRuleInfo) GetRuleType() uint32 {
	if x != nil {
		return x.RuleType
	}
	return 0
}

var File_ExclusiveRuleInfo_proto protoreflect.FileDescriptor

var file_ExclusiveRuleInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x11, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24,
	0x0a, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x75, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExclusiveRuleInfo_proto_rawDescOnce sync.Once
	file_ExclusiveRuleInfo_proto_rawDescData = file_ExclusiveRuleInfo_proto_rawDesc
)

func file_ExclusiveRuleInfo_proto_rawDescGZIP() []byte {
	file_ExclusiveRuleInfo_proto_rawDescOnce.Do(func() {
		file_ExclusiveRuleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExclusiveRuleInfo_proto_rawDescData)
	})
	return file_ExclusiveRuleInfo_proto_rawDescData
}

var file_ExclusiveRuleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExclusiveRuleInfo_proto_goTypes = []interface{}{
	(*ExclusiveRuleInfo)(nil), // 0: ExclusiveRuleInfo
}
var file_ExclusiveRuleInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ExclusiveRuleInfo_proto_init() }
func file_ExclusiveRuleInfo_proto_init() {
	if File_ExclusiveRuleInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ExclusiveRuleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExclusiveRuleInfo); i {
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
			RawDescriptor: file_ExclusiveRuleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExclusiveRuleInfo_proto_goTypes,
		DependencyIndexes: file_ExclusiveRuleInfo_proto_depIdxs,
		MessageInfos:      file_ExclusiveRuleInfo_proto_msgTypes,
	}.Build()
	File_ExclusiveRuleInfo_proto = out.File
	file_ExclusiveRuleInfo_proto_rawDesc = nil
	file_ExclusiveRuleInfo_proto_goTypes = nil
	file_ExclusiveRuleInfo_proto_depIdxs = nil
}
