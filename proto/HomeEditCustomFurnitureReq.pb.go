// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HomeEditCustomFurnitureReq.proto

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

// CmdId: 4724
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type HomeEditCustomFurnitureReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomFurnitureInfo *HomeCustomFurnitureInfo `protobuf:"bytes,15,opt,name=custom_furniture_info,json=customFurnitureInfo,proto3" json:"custom_furniture_info,omitempty"`
}

func (x *HomeEditCustomFurnitureReq) Reset() {
	*x = HomeEditCustomFurnitureReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeEditCustomFurnitureReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeEditCustomFurnitureReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeEditCustomFurnitureReq) ProtoMessage() {}

func (x *HomeEditCustomFurnitureReq) ProtoReflect() protoreflect.Message {
	mi := &file_HomeEditCustomFurnitureReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeEditCustomFurnitureReq.ProtoReflect.Descriptor instead.
func (*HomeEditCustomFurnitureReq) Descriptor() ([]byte, []int) {
	return file_HomeEditCustomFurnitureReq_proto_rawDescGZIP(), []int{0}
}

func (x *HomeEditCustomFurnitureReq) GetCustomFurnitureInfo() *HomeCustomFurnitureInfo {
	if x != nil {
		return x.CustomFurnitureInfo
	}
	return nil
}

var File_HomeEditCustomFurnitureReq_proto protoreflect.FileDescriptor

var file_HomeEditCustomFurnitureReq_proto_rawDesc = []byte{
	0x0a, 0x20, 0x48, 0x6f, 0x6d, 0x65, 0x45, 0x64, 0x69, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x75,
	0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6a, 0x0a, 0x1a, 0x48, 0x6f, 0x6d, 0x65, 0x45, 0x64, 0x69, 0x74, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x4c, 0x0a, 0x15, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x75, 0x72, 0x6e, 0x69,
	0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeEditCustomFurnitureReq_proto_rawDescOnce sync.Once
	file_HomeEditCustomFurnitureReq_proto_rawDescData = file_HomeEditCustomFurnitureReq_proto_rawDesc
)

func file_HomeEditCustomFurnitureReq_proto_rawDescGZIP() []byte {
	file_HomeEditCustomFurnitureReq_proto_rawDescOnce.Do(func() {
		file_HomeEditCustomFurnitureReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeEditCustomFurnitureReq_proto_rawDescData)
	})
	return file_HomeEditCustomFurnitureReq_proto_rawDescData
}

var file_HomeEditCustomFurnitureReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeEditCustomFurnitureReq_proto_goTypes = []interface{}{
	(*HomeEditCustomFurnitureReq)(nil), // 0: HomeEditCustomFurnitureReq
	(*HomeCustomFurnitureInfo)(nil),    // 1: HomeCustomFurnitureInfo
}
var file_HomeEditCustomFurnitureReq_proto_depIdxs = []int32{
	1, // 0: HomeEditCustomFurnitureReq.custom_furniture_info:type_name -> HomeCustomFurnitureInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeEditCustomFurnitureReq_proto_init() }
func file_HomeEditCustomFurnitureReq_proto_init() {
	if File_HomeEditCustomFurnitureReq_proto != nil {
		return
	}
	file_HomeCustomFurnitureInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeEditCustomFurnitureReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeEditCustomFurnitureReq); i {
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
			RawDescriptor: file_HomeEditCustomFurnitureReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeEditCustomFurnitureReq_proto_goTypes,
		DependencyIndexes: file_HomeEditCustomFurnitureReq_proto_depIdxs,
		MessageInfos:      file_HomeEditCustomFurnitureReq_proto_msgTypes,
	}.Build()
	File_HomeEditCustomFurnitureReq_proto = out.File
	file_HomeEditCustomFurnitureReq_proto_rawDesc = nil
	file_HomeEditCustomFurnitureReq_proto_goTypes = nil
	file_HomeEditCustomFurnitureReq_proto_depIdxs = nil
}
