// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GetFurnitureCurModuleArrangeCountReq.proto

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

// CmdId: 4711
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type GetFurnitureCurModuleArrangeCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFurnitureCurModuleArrangeCountReq) Reset() {
	*x = GetFurnitureCurModuleArrangeCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetFurnitureCurModuleArrangeCountReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFurnitureCurModuleArrangeCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFurnitureCurModuleArrangeCountReq) ProtoMessage() {}

func (x *GetFurnitureCurModuleArrangeCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetFurnitureCurModuleArrangeCountReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFurnitureCurModuleArrangeCountReq.ProtoReflect.Descriptor instead.
func (*GetFurnitureCurModuleArrangeCountReq) Descriptor() ([]byte, []int) {
	return file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescGZIP(), []int{0}
}

var File_GetFurnitureCurModuleArrangeCountReq_proto protoreflect.FileDescriptor

var file_GetFurnitureCurModuleArrangeCountReq_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x47, 0x65, 0x74, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x43, 0x75,
	0x72, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x24,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x43, 0x75, 0x72, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescOnce sync.Once
	file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescData = file_GetFurnitureCurModuleArrangeCountReq_proto_rawDesc
)

func file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescGZIP() []byte {
	file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescOnce.Do(func() {
		file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescData)
	})
	return file_GetFurnitureCurModuleArrangeCountReq_proto_rawDescData
}

var file_GetFurnitureCurModuleArrangeCountReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetFurnitureCurModuleArrangeCountReq_proto_goTypes = []interface{}{
	(*GetFurnitureCurModuleArrangeCountReq)(nil), // 0: GetFurnitureCurModuleArrangeCountReq
}
var file_GetFurnitureCurModuleArrangeCountReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetFurnitureCurModuleArrangeCountReq_proto_init() }
func file_GetFurnitureCurModuleArrangeCountReq_proto_init() {
	if File_GetFurnitureCurModuleArrangeCountReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetFurnitureCurModuleArrangeCountReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFurnitureCurModuleArrangeCountReq); i {
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
			RawDescriptor: file_GetFurnitureCurModuleArrangeCountReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetFurnitureCurModuleArrangeCountReq_proto_goTypes,
		DependencyIndexes: file_GetFurnitureCurModuleArrangeCountReq_proto_depIdxs,
		MessageInfos:      file_GetFurnitureCurModuleArrangeCountReq_proto_msgTypes,
	}.Build()
	File_GetFurnitureCurModuleArrangeCountReq_proto = out.File
	file_GetFurnitureCurModuleArrangeCountReq_proto_rawDesc = nil
	file_GetFurnitureCurModuleArrangeCountReq_proto_goTypes = nil
	file_GetFurnitureCurModuleArrangeCountReq_proto_depIdxs = nil
}
