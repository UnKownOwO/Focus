// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ForceUpdateInfo.proto

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

type ForceUpdateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ForceUpdateUrl string `protobuf:"bytes,1,opt,name=force_update_url,json=forceUpdateUrl,proto3" json:"force_update_url,omitempty"`
}

func (x *ForceUpdateInfo) Reset() {
	*x = ForceUpdateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ForceUpdateInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceUpdateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceUpdateInfo) ProtoMessage() {}

func (x *ForceUpdateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ForceUpdateInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceUpdateInfo.ProtoReflect.Descriptor instead.
func (*ForceUpdateInfo) Descriptor() ([]byte, []int) {
	return file_ForceUpdateInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ForceUpdateInfo) GetForceUpdateUrl() string {
	if x != nil {
		return x.ForceUpdateUrl
	}
	return ""
}

var File_ForceUpdateInfo_proto protoreflect.FileDescriptor

var file_ForceUpdateInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x72, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ForceUpdateInfo_proto_rawDescOnce sync.Once
	file_ForceUpdateInfo_proto_rawDescData = file_ForceUpdateInfo_proto_rawDesc
)

func file_ForceUpdateInfo_proto_rawDescGZIP() []byte {
	file_ForceUpdateInfo_proto_rawDescOnce.Do(func() {
		file_ForceUpdateInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ForceUpdateInfo_proto_rawDescData)
	})
	return file_ForceUpdateInfo_proto_rawDescData
}

var file_ForceUpdateInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ForceUpdateInfo_proto_goTypes = []interface{}{
	(*ForceUpdateInfo)(nil), // 0: ForceUpdateInfo
}
var file_ForceUpdateInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ForceUpdateInfo_proto_init() }
func file_ForceUpdateInfo_proto_init() {
	if File_ForceUpdateInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ForceUpdateInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceUpdateInfo); i {
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
			RawDescriptor: file_ForceUpdateInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ForceUpdateInfo_proto_goTypes,
		DependencyIndexes: file_ForceUpdateInfo_proto_depIdxs,
		MessageInfos:      file_ForceUpdateInfo_proto_msgTypes,
	}.Build()
	File_ForceUpdateInfo_proto = out.File
	file_ForceUpdateInfo_proto_rawDesc = nil
	file_ForceUpdateInfo_proto_goTypes = nil
	file_ForceUpdateInfo_proto_depIdxs = nil
}
