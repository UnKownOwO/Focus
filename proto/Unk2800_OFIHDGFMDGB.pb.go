// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2800_OFIHDGFMDGB.proto

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

// CmdId: 171
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type Unk2800_OFIHDGFMDGB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GivingId uint32 `protobuf:"varint,4,opt,name=giving_id,json=givingId,proto3" json:"giving_id,omitempty"`
}

func (x *Unk2800_OFIHDGFMDGB) Reset() {
	*x = Unk2800_OFIHDGFMDGB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2800_OFIHDGFMDGB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2800_OFIHDGFMDGB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2800_OFIHDGFMDGB) ProtoMessage() {}

func (x *Unk2800_OFIHDGFMDGB) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2800_OFIHDGFMDGB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2800_OFIHDGFMDGB.ProtoReflect.Descriptor instead.
func (*Unk2800_OFIHDGFMDGB) Descriptor() ([]byte, []int) {
	return file_Unk2800_OFIHDGFMDGB_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2800_OFIHDGFMDGB) GetGivingId() uint32 {
	if x != nil {
		return x.GivingId
	}
	return 0
}

var File_Unk2800_OFIHDGFMDGB_proto protoreflect.FileDescriptor

var file_Unk2800_OFIHDGFMDGB_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4f, 0x46, 0x49, 0x48, 0x44, 0x47,
	0x46, 0x4d, 0x44, 0x47, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x13, 0x55,
	0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4f, 0x46, 0x49, 0x48, 0x44, 0x47, 0x46, 0x4d, 0x44,
	0x47, 0x42, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Unk2800_OFIHDGFMDGB_proto_rawDescOnce sync.Once
	file_Unk2800_OFIHDGFMDGB_proto_rawDescData = file_Unk2800_OFIHDGFMDGB_proto_rawDesc
)

func file_Unk2800_OFIHDGFMDGB_proto_rawDescGZIP() []byte {
	file_Unk2800_OFIHDGFMDGB_proto_rawDescOnce.Do(func() {
		file_Unk2800_OFIHDGFMDGB_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2800_OFIHDGFMDGB_proto_rawDescData)
	})
	return file_Unk2800_OFIHDGFMDGB_proto_rawDescData
}

var file_Unk2800_OFIHDGFMDGB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2800_OFIHDGFMDGB_proto_goTypes = []interface{}{
	(*Unk2800_OFIHDGFMDGB)(nil), // 0: Unk2800_OFIHDGFMDGB
}
var file_Unk2800_OFIHDGFMDGB_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2800_OFIHDGFMDGB_proto_init() }
func file_Unk2800_OFIHDGFMDGB_proto_init() {
	if File_Unk2800_OFIHDGFMDGB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk2800_OFIHDGFMDGB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2800_OFIHDGFMDGB); i {
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
			RawDescriptor: file_Unk2800_OFIHDGFMDGB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2800_OFIHDGFMDGB_proto_goTypes,
		DependencyIndexes: file_Unk2800_OFIHDGFMDGB_proto_depIdxs,
		MessageInfos:      file_Unk2800_OFIHDGFMDGB_proto_msgTypes,
	}.Build()
	File_Unk2800_OFIHDGFMDGB_proto = out.File
	file_Unk2800_OFIHDGFMDGB_proto_rawDesc = nil
	file_Unk2800_OFIHDGFMDGB_proto_goTypes = nil
	file_Unk2800_OFIHDGFMDGB_proto_depIdxs = nil
}
