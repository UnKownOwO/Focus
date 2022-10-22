// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: MechanicusOpenNotify.proto

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

// CmdId: 3907
// EnetChannelId: 0
// EnetIsReliable: true
type MechanicusOpenNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MechanicusId uint32 `protobuf:"varint,2,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
}

func (x *MechanicusOpenNotify) Reset() {
	*x = MechanicusOpenNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MechanicusOpenNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MechanicusOpenNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusOpenNotify) ProtoMessage() {}

func (x *MechanicusOpenNotify) ProtoReflect() protoreflect.Message {
	mi := &file_MechanicusOpenNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusOpenNotify.ProtoReflect.Descriptor instead.
func (*MechanicusOpenNotify) Descriptor() ([]byte, []int) {
	return file_MechanicusOpenNotify_proto_rawDescGZIP(), []int{0}
}

func (x *MechanicusOpenNotify) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

var File_MechanicusOpenNotify_proto protoreflect.FileDescriptor

var file_MechanicusOpenNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x6e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x14,
	0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63,
	0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MechanicusOpenNotify_proto_rawDescOnce sync.Once
	file_MechanicusOpenNotify_proto_rawDescData = file_MechanicusOpenNotify_proto_rawDesc
)

func file_MechanicusOpenNotify_proto_rawDescGZIP() []byte {
	file_MechanicusOpenNotify_proto_rawDescOnce.Do(func() {
		file_MechanicusOpenNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_MechanicusOpenNotify_proto_rawDescData)
	})
	return file_MechanicusOpenNotify_proto_rawDescData
}

var file_MechanicusOpenNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MechanicusOpenNotify_proto_goTypes = []interface{}{
	(*MechanicusOpenNotify)(nil), // 0: MechanicusOpenNotify
}
var file_MechanicusOpenNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MechanicusOpenNotify_proto_init() }
func file_MechanicusOpenNotify_proto_init() {
	if File_MechanicusOpenNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MechanicusOpenNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MechanicusOpenNotify); i {
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
			RawDescriptor: file_MechanicusOpenNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MechanicusOpenNotify_proto_goTypes,
		DependencyIndexes: file_MechanicusOpenNotify_proto_depIdxs,
		MessageInfos:      file_MechanicusOpenNotify_proto_msgTypes,
	}.Build()
	File_MechanicusOpenNotify_proto = out.File
	file_MechanicusOpenNotify_proto_rawDesc = nil
	file_MechanicusOpenNotify_proto_goTypes = nil
	file_MechanicusOpenNotify_proto_depIdxs = nil
}
