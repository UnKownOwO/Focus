// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: FurnitureMakeBeHelpedNotify.proto

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

// CmdId: 4578
// EnetChannelId: 0
// EnetIsReliable: true
type FurnitureMakeBeHelpedNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FurnitureMakeSlot       *FurnitureMakeSlot         `protobuf:"bytes,7,opt,name=furniture_make_slot,json=furnitureMakeSlot,proto3" json:"furniture_make_slot,omitempty"`
	FurnitureMakeHelpedData *FurnitureMakeBeHelpedData `protobuf:"bytes,2,opt,name=furniture_make_helped_data,json=furnitureMakeHelpedData,proto3" json:"furniture_make_helped_data,omitempty"`
}

func (x *FurnitureMakeBeHelpedNotify) Reset() {
	*x = FurnitureMakeBeHelpedNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FurnitureMakeBeHelpedNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FurnitureMakeBeHelpedNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FurnitureMakeBeHelpedNotify) ProtoMessage() {}

func (x *FurnitureMakeBeHelpedNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FurnitureMakeBeHelpedNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FurnitureMakeBeHelpedNotify.ProtoReflect.Descriptor instead.
func (*FurnitureMakeBeHelpedNotify) Descriptor() ([]byte, []int) {
	return file_FurnitureMakeBeHelpedNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FurnitureMakeBeHelpedNotify) GetFurnitureMakeSlot() *FurnitureMakeSlot {
	if x != nil {
		return x.FurnitureMakeSlot
	}
	return nil
}

func (x *FurnitureMakeBeHelpedNotify) GetFurnitureMakeHelpedData() *FurnitureMakeBeHelpedData {
	if x != nil {
		return x.FurnitureMakeHelpedData
	}
	return nil
}

var File_FurnitureMakeBeHelpedNotify_proto protoreflect.FileDescriptor

var file_FurnitureMakeBeHelpedNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x42,
	0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61,
	0x6b, 0x65, 0x42, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d,
	0x61, 0x6b, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01,
	0x0a, 0x1b, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x42,
	0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x42, 0x0a,
	0x13, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x61, 0x6b, 0x65, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46, 0x75, 0x72,
	0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x11,
	0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65, 0x53, 0x6c, 0x6f,
	0x74, 0x12, 0x57, 0x0a, 0x1a, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6d,
	0x61, 0x6b, 0x65, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72,
	0x65, 0x4d, 0x61, 0x6b, 0x65, 0x42, 0x65, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x17, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x6b, 0x65,
	0x48, 0x65, 0x6c, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FurnitureMakeBeHelpedNotify_proto_rawDescOnce sync.Once
	file_FurnitureMakeBeHelpedNotify_proto_rawDescData = file_FurnitureMakeBeHelpedNotify_proto_rawDesc
)

func file_FurnitureMakeBeHelpedNotify_proto_rawDescGZIP() []byte {
	file_FurnitureMakeBeHelpedNotify_proto_rawDescOnce.Do(func() {
		file_FurnitureMakeBeHelpedNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FurnitureMakeBeHelpedNotify_proto_rawDescData)
	})
	return file_FurnitureMakeBeHelpedNotify_proto_rawDescData
}

var file_FurnitureMakeBeHelpedNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FurnitureMakeBeHelpedNotify_proto_goTypes = []interface{}{
	(*FurnitureMakeBeHelpedNotify)(nil), // 0: FurnitureMakeBeHelpedNotify
	(*FurnitureMakeSlot)(nil),           // 1: FurnitureMakeSlot
	(*FurnitureMakeBeHelpedData)(nil),   // 2: FurnitureMakeBeHelpedData
}
var file_FurnitureMakeBeHelpedNotify_proto_depIdxs = []int32{
	1, // 0: FurnitureMakeBeHelpedNotify.furniture_make_slot:type_name -> FurnitureMakeSlot
	2, // 1: FurnitureMakeBeHelpedNotify.furniture_make_helped_data:type_name -> FurnitureMakeBeHelpedData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FurnitureMakeBeHelpedNotify_proto_init() }
func file_FurnitureMakeBeHelpedNotify_proto_init() {
	if File_FurnitureMakeBeHelpedNotify_proto != nil {
		return
	}
	file_FurnitureMakeBeHelpedData_proto_init()
	file_FurnitureMakeSlot_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FurnitureMakeBeHelpedNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FurnitureMakeBeHelpedNotify); i {
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
			RawDescriptor: file_FurnitureMakeBeHelpedNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FurnitureMakeBeHelpedNotify_proto_goTypes,
		DependencyIndexes: file_FurnitureMakeBeHelpedNotify_proto_depIdxs,
		MessageInfos:      file_FurnitureMakeBeHelpedNotify_proto_msgTypes,
	}.Build()
	File_FurnitureMakeBeHelpedNotify_proto = out.File
	file_FurnitureMakeBeHelpedNotify_proto_rawDesc = nil
	file_FurnitureMakeBeHelpedNotify_proto_goTypes = nil
	file_FurnitureMakeBeHelpedNotify_proto_depIdxs = nil
}