// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: StrengthenPointData.proto

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

type StrengthenPointData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasePoint uint32 `protobuf:"varint,10,opt,name=base_point,json=basePoint,proto3" json:"base_point,omitempty"`
	CurPoint  uint32 `protobuf:"varint,11,opt,name=cur_point,json=curPoint,proto3" json:"cur_point,omitempty"`
}

func (x *StrengthenPointData) Reset() {
	*x = StrengthenPointData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StrengthenPointData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrengthenPointData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrengthenPointData) ProtoMessage() {}

func (x *StrengthenPointData) ProtoReflect() protoreflect.Message {
	mi := &file_StrengthenPointData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrengthenPointData.ProtoReflect.Descriptor instead.
func (*StrengthenPointData) Descriptor() ([]byte, []int) {
	return file_StrengthenPointData_proto_rawDescGZIP(), []int{0}
}

func (x *StrengthenPointData) GetBasePoint() uint32 {
	if x != nil {
		return x.BasePoint
	}
	return 0
}

func (x *StrengthenPointData) GetCurPoint() uint32 {
	if x != nil {
		return x.CurPoint
	}
	return 0
}

var File_StrengthenPointData_proto protoreflect.FileDescriptor

var file_StrengthenPointData_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x13, 0x53,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_StrengthenPointData_proto_rawDescOnce sync.Once
	file_StrengthenPointData_proto_rawDescData = file_StrengthenPointData_proto_rawDesc
)

func file_StrengthenPointData_proto_rawDescGZIP() []byte {
	file_StrengthenPointData_proto_rawDescOnce.Do(func() {
		file_StrengthenPointData_proto_rawDescData = protoimpl.X.CompressGZIP(file_StrengthenPointData_proto_rawDescData)
	})
	return file_StrengthenPointData_proto_rawDescData
}

var file_StrengthenPointData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StrengthenPointData_proto_goTypes = []interface{}{
	(*StrengthenPointData)(nil), // 0: StrengthenPointData
}
var file_StrengthenPointData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_StrengthenPointData_proto_init() }
func file_StrengthenPointData_proto_init() {
	if File_StrengthenPointData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_StrengthenPointData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrengthenPointData); i {
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
			RawDescriptor: file_StrengthenPointData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StrengthenPointData_proto_goTypes,
		DependencyIndexes: file_StrengthenPointData_proto_depIdxs,
		MessageInfos:      file_StrengthenPointData_proto_msgTypes,
	}.Build()
	File_StrengthenPointData_proto = out.File
	file_StrengthenPointData_proto_rawDesc = nil
	file_StrengthenPointData_proto_goTypes = nil
	file_StrengthenPointData_proto_depIdxs = nil
}
