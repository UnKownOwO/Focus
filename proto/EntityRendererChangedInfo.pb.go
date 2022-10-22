// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: EntityRendererChangedInfo.proto

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

type EntityRendererChangedInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangedRenderers map[string]uint32 `protobuf:"bytes,1,rep,name=changed_renderers,json=changedRenderers,proto3" json:"changed_renderers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	VisibilityCount  uint32            `protobuf:"varint,2,opt,name=visibility_count,json=visibilityCount,proto3" json:"visibility_count,omitempty"`
	IsCached         bool              `protobuf:"varint,3,opt,name=is_cached,json=isCached,proto3" json:"is_cached,omitempty"`
}

func (x *EntityRendererChangedInfo) Reset() {
	*x = EntityRendererChangedInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityRendererChangedInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityRendererChangedInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityRendererChangedInfo) ProtoMessage() {}

func (x *EntityRendererChangedInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EntityRendererChangedInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityRendererChangedInfo.ProtoReflect.Descriptor instead.
func (*EntityRendererChangedInfo) Descriptor() ([]byte, []int) {
	return file_EntityRendererChangedInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EntityRendererChangedInfo) GetChangedRenderers() map[string]uint32 {
	if x != nil {
		return x.ChangedRenderers
	}
	return nil
}

func (x *EntityRendererChangedInfo) GetVisibilityCount() uint32 {
	if x != nil {
		return x.VisibilityCount
	}
	return 0
}

func (x *EntityRendererChangedInfo) GetIsCached() bool {
	if x != nil {
		return x.IsCached
	}
	return false
}

var File_EntityRendererChangedInfo_proto protoreflect.FileDescriptor

var file_EntityRendererChangedInfo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x87, 0x02, 0x0a, 0x19, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x5d, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x1a, 0x43, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityRendererChangedInfo_proto_rawDescOnce sync.Once
	file_EntityRendererChangedInfo_proto_rawDescData = file_EntityRendererChangedInfo_proto_rawDesc
)

func file_EntityRendererChangedInfo_proto_rawDescGZIP() []byte {
	file_EntityRendererChangedInfo_proto_rawDescOnce.Do(func() {
		file_EntityRendererChangedInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityRendererChangedInfo_proto_rawDescData)
	})
	return file_EntityRendererChangedInfo_proto_rawDescData
}

var file_EntityRendererChangedInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_EntityRendererChangedInfo_proto_goTypes = []interface{}{
	(*EntityRendererChangedInfo)(nil), // 0: EntityRendererChangedInfo
	nil,                               // 1: EntityRendererChangedInfo.ChangedRenderersEntry
}
var file_EntityRendererChangedInfo_proto_depIdxs = []int32{
	1, // 0: EntityRendererChangedInfo.changed_renderers:type_name -> EntityRendererChangedInfo.ChangedRenderersEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EntityRendererChangedInfo_proto_init() }
func file_EntityRendererChangedInfo_proto_init() {
	if File_EntityRendererChangedInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EntityRendererChangedInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityRendererChangedInfo); i {
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
			RawDescriptor: file_EntityRendererChangedInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityRendererChangedInfo_proto_goTypes,
		DependencyIndexes: file_EntityRendererChangedInfo_proto_depIdxs,
		MessageInfos:      file_EntityRendererChangedInfo_proto_msgTypes,
	}.Build()
	File_EntityRendererChangedInfo_proto = out.File
	file_EntityRendererChangedInfo_proto_rawDesc = nil
	file_EntityRendererChangedInfo_proto_goTypes = nil
	file_EntityRendererChangedInfo_proto_depIdxs = nil
}
