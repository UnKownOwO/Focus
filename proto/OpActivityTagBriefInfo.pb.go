// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: OpActivityTagBriefInfo.proto

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

type OpActivityTagBriefInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId       uint32 `protobuf:"varint,2,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	HasReward      bool   `protobuf:"varint,3,opt,name=has_reward,json=hasReward,proto3" json:"has_reward,omitempty"`
	OpActivityType uint32 `protobuf:"varint,11,opt,name=op_activity_type,json=opActivityType,proto3" json:"op_activity_type,omitempty"`
}

func (x *OpActivityTagBriefInfo) Reset() {
	*x = OpActivityTagBriefInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_OpActivityTagBriefInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpActivityTagBriefInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpActivityTagBriefInfo) ProtoMessage() {}

func (x *OpActivityTagBriefInfo) ProtoReflect() protoreflect.Message {
	mi := &file_OpActivityTagBriefInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpActivityTagBriefInfo.ProtoReflect.Descriptor instead.
func (*OpActivityTagBriefInfo) Descriptor() ([]byte, []int) {
	return file_OpActivityTagBriefInfo_proto_rawDescGZIP(), []int{0}
}

func (x *OpActivityTagBriefInfo) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *OpActivityTagBriefInfo) GetHasReward() bool {
	if x != nil {
		return x.HasReward
	}
	return false
}

func (x *OpActivityTagBriefInfo) GetOpActivityType() uint32 {
	if x != nil {
		return x.OpActivityType
	}
	return 0
}

var File_OpActivityTagBriefInfo_proto protoreflect.FileDescriptor

var file_OpActivityTagBriefInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x61, 0x67, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e,
	0x0a, 0x16, 0x4f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x61, 0x67, 0x42,
	0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x70, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e,
	0x6f, 0x70, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_OpActivityTagBriefInfo_proto_rawDescOnce sync.Once
	file_OpActivityTagBriefInfo_proto_rawDescData = file_OpActivityTagBriefInfo_proto_rawDesc
)

func file_OpActivityTagBriefInfo_proto_rawDescGZIP() []byte {
	file_OpActivityTagBriefInfo_proto_rawDescOnce.Do(func() {
		file_OpActivityTagBriefInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_OpActivityTagBriefInfo_proto_rawDescData)
	})
	return file_OpActivityTagBriefInfo_proto_rawDescData
}

var file_OpActivityTagBriefInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_OpActivityTagBriefInfo_proto_goTypes = []interface{}{
	(*OpActivityTagBriefInfo)(nil), // 0: OpActivityTagBriefInfo
}
var file_OpActivityTagBriefInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_OpActivityTagBriefInfo_proto_init() }
func file_OpActivityTagBriefInfo_proto_init() {
	if File_OpActivityTagBriefInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_OpActivityTagBriefInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpActivityTagBriefInfo); i {
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
			RawDescriptor: file_OpActivityTagBriefInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_OpActivityTagBriefInfo_proto_goTypes,
		DependencyIndexes: file_OpActivityTagBriefInfo_proto_depIdxs,
		MessageInfos:      file_OpActivityTagBriefInfo_proto_msgTypes,
	}.Build()
	File_OpActivityTagBriefInfo_proto = out.File
	file_OpActivityTagBriefInfo_proto_rawDesc = nil
	file_OpActivityTagBriefInfo_proto_goTypes = nil
	file_OpActivityTagBriefInfo_proto_depIdxs = nil
}