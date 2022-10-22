// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: FightPropPair.proto

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

type FightPropPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropType  uint32  `protobuf:"varint,1,opt,name=prop_type,json=propType,proto3" json:"prop_type,omitempty"`
	PropValue float32 `protobuf:"fixed32,2,opt,name=prop_value,json=propValue,proto3" json:"prop_value,omitempty"`
}

func (x *FightPropPair) Reset() {
	*x = FightPropPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FightPropPair_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FightPropPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FightPropPair) ProtoMessage() {}

func (x *FightPropPair) ProtoReflect() protoreflect.Message {
	mi := &file_FightPropPair_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FightPropPair.ProtoReflect.Descriptor instead.
func (*FightPropPair) Descriptor() ([]byte, []int) {
	return file_FightPropPair_proto_rawDescGZIP(), []int{0}
}

func (x *FightPropPair) GetPropType() uint32 {
	if x != nil {
		return x.PropType
	}
	return 0
}

func (x *FightPropPair) GetPropValue() float32 {
	if x != nil {
		return x.PropValue
	}
	return 0
}

var File_FightPropPair_proto protoreflect.FileDescriptor

var file_FightPropPair_proto_rawDesc = []byte{
	0x0a, 0x13, 0x46, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x69, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0d, 0x46, 0x69, 0x67, 0x68, 0x74, 0x50, 0x72,
	0x6f, 0x70, 0x50, 0x61, 0x69, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FightPropPair_proto_rawDescOnce sync.Once
	file_FightPropPair_proto_rawDescData = file_FightPropPair_proto_rawDesc
)

func file_FightPropPair_proto_rawDescGZIP() []byte {
	file_FightPropPair_proto_rawDescOnce.Do(func() {
		file_FightPropPair_proto_rawDescData = protoimpl.X.CompressGZIP(file_FightPropPair_proto_rawDescData)
	})
	return file_FightPropPair_proto_rawDescData
}

var file_FightPropPair_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FightPropPair_proto_goTypes = []interface{}{
	(*FightPropPair)(nil), // 0: FightPropPair
}
var file_FightPropPair_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FightPropPair_proto_init() }
func file_FightPropPair_proto_init() {
	if File_FightPropPair_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FightPropPair_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FightPropPair); i {
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
			RawDescriptor: file_FightPropPair_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FightPropPair_proto_goTypes,
		DependencyIndexes: file_FightPropPair_proto_depIdxs,
		MessageInfos:      file_FightPropPair_proto_msgTypes,
	}.Build()
	File_FightPropPair_proto = out.File
	file_FightPropPair_proto_rawDesc = nil
	file_FightPropPair_proto_goTypes = nil
	file_FightPropPair_proto_depIdxs = nil
}