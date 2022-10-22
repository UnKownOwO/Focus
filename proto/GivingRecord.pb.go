// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GivingRecord.proto

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

type GivingRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFinished          bool              `protobuf:"varint,9,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	GroupId             uint32            `protobuf:"varint,5,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Unk2800_JBPPNEHPACC bool              `protobuf:"varint,8,opt,name=Unk2800_JBPPNEHPACC,json=Unk2800JBPPNEHPACC,proto3" json:"Unk2800_JBPPNEHPACC,omitempty"`
	GivingId            uint32            `protobuf:"varint,3,opt,name=giving_id,json=givingId,proto3" json:"giving_id,omitempty"`
	LastGroupId         uint32            `protobuf:"varint,6,opt,name=last_group_id,json=lastGroupId,proto3" json:"last_group_id,omitempty"`
	ConfigId            uint32            `protobuf:"varint,2,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	Unk2800_BDKKENPEEGD map[uint32]uint32 `protobuf:"bytes,15,rep,name=Unk2800_BDKKENPEEGD,json=Unk2800BDKKENPEEGD,proto3" json:"Unk2800_BDKKENPEEGD,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GivingRecord) Reset() {
	*x = GivingRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GivingRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GivingRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GivingRecord) ProtoMessage() {}

func (x *GivingRecord) ProtoReflect() protoreflect.Message {
	mi := &file_GivingRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GivingRecord.ProtoReflect.Descriptor instead.
func (*GivingRecord) Descriptor() ([]byte, []int) {
	return file_GivingRecord_proto_rawDescGZIP(), []int{0}
}

func (x *GivingRecord) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *GivingRecord) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GivingRecord) GetUnk2800_JBPPNEHPACC() bool {
	if x != nil {
		return x.Unk2800_JBPPNEHPACC
	}
	return false
}

func (x *GivingRecord) GetGivingId() uint32 {
	if x != nil {
		return x.GivingId
	}
	return 0
}

func (x *GivingRecord) GetLastGroupId() uint32 {
	if x != nil {
		return x.LastGroupId
	}
	return 0
}

func (x *GivingRecord) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *GivingRecord) GetUnk2800_BDKKENPEEGD() map[uint32]uint32 {
	if x != nil {
		return x.Unk2800_BDKKENPEEGD
	}
	return nil
}

var File_GivingRecord_proto protoreflect.FileDescriptor

var file_GivingRecord_proto_rawDesc = []byte{
	0x0a, 0x12, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a, 0x0c, 0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4a, 0x42, 0x50,
	0x50, 0x4e, 0x45, 0x48, 0x50, 0x41, 0x43, 0x43, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x4a, 0x42, 0x50, 0x50, 0x4e, 0x45, 0x48, 0x50, 0x41,
	0x43, 0x43, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64,
	0x12, 0x56, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x42, 0x44, 0x4b, 0x4b,
	0x45, 0x4e, 0x50, 0x45, 0x45, 0x47, 0x44, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x47, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x55, 0x6e, 0x6b,
	0x32, 0x38, 0x30, 0x30, 0x42, 0x44, 0x4b, 0x4b, 0x45, 0x4e, 0x50, 0x45, 0x45, 0x47, 0x44, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x42, 0x44, 0x4b,
	0x4b, 0x45, 0x4e, 0x50, 0x45, 0x45, 0x47, 0x44, 0x1a, 0x45, 0x0a, 0x17, 0x55, 0x6e, 0x6b, 0x32,
	0x38, 0x30, 0x30, 0x42, 0x44, 0x4b, 0x4b, 0x45, 0x4e, 0x50, 0x45, 0x45, 0x47, 0x44, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GivingRecord_proto_rawDescOnce sync.Once
	file_GivingRecord_proto_rawDescData = file_GivingRecord_proto_rawDesc
)

func file_GivingRecord_proto_rawDescGZIP() []byte {
	file_GivingRecord_proto_rawDescOnce.Do(func() {
		file_GivingRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_GivingRecord_proto_rawDescData)
	})
	return file_GivingRecord_proto_rawDescData
}

var file_GivingRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GivingRecord_proto_goTypes = []interface{}{
	(*GivingRecord)(nil), // 0: GivingRecord
	nil,                  // 1: GivingRecord.Unk2800BDKKENPEEGDEntry
}
var file_GivingRecord_proto_depIdxs = []int32{
	1, // 0: GivingRecord.Unk2800_BDKKENPEEGD:type_name -> GivingRecord.Unk2800BDKKENPEEGDEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GivingRecord_proto_init() }
func file_GivingRecord_proto_init() {
	if File_GivingRecord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GivingRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GivingRecord); i {
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
			RawDescriptor: file_GivingRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GivingRecord_proto_goTypes,
		DependencyIndexes: file_GivingRecord_proto_depIdxs,
		MessageInfos:      file_GivingRecord_proto_msgTypes,
	}.Build()
	File_GivingRecord_proto = out.File
	file_GivingRecord_proto_rawDesc = nil
	file_GivingRecord_proto_goTypes = nil
	file_GivingRecord_proto_depIdxs = nil
}
