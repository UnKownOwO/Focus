// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SceneGalleryHomeSeekFurnitureInfo.proto

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

type SceneGalleryHomeSeekFurnitureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_KDBENBBODGP uint32            `protobuf:"varint,6,opt,name=Unk2700_KDBENBBODGP,json=Unk2700KDBENBBODGP,proto3" json:"Unk2700_KDBENBBODGP,omitempty"`
	Unk2700_DDHOJHOICBL map[uint32]uint32 `protobuf:"bytes,8,rep,name=Unk2700_DDHOJHOICBL,json=Unk2700DDHOJHOICBL,proto3" json:"Unk2700_DDHOJHOICBL,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Unk2700_LODFFCPFJLC uint32            `protobuf:"varint,12,opt,name=Unk2700_LODFFCPFJLC,json=Unk2700LODFFCPFJLC,proto3" json:"Unk2700_LODFFCPFJLC,omitempty"`
	Unk2700_HLCIHCCGFFC uint32            `protobuf:"varint,9,opt,name=Unk2700_HLCIHCCGFFC,json=Unk2700HLCIHCCGFFC,proto3" json:"Unk2700_HLCIHCCGFFC,omitempty"`
}

func (x *SceneGalleryHomeSeekFurnitureInfo) Reset() {
	*x = SceneGalleryHomeSeekFurnitureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGalleryHomeSeekFurnitureInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGalleryHomeSeekFurnitureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGalleryHomeSeekFurnitureInfo) ProtoMessage() {}

func (x *SceneGalleryHomeSeekFurnitureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGalleryHomeSeekFurnitureInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGalleryHomeSeekFurnitureInfo.ProtoReflect.Descriptor instead.
func (*SceneGalleryHomeSeekFurnitureInfo) Descriptor() ([]byte, []int) {
	return file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGalleryHomeSeekFurnitureInfo) GetUnk2700_KDBENBBODGP() uint32 {
	if x != nil {
		return x.Unk2700_KDBENBBODGP
	}
	return 0
}

func (x *SceneGalleryHomeSeekFurnitureInfo) GetUnk2700_DDHOJHOICBL() map[uint32]uint32 {
	if x != nil {
		return x.Unk2700_DDHOJHOICBL
	}
	return nil
}

func (x *SceneGalleryHomeSeekFurnitureInfo) GetUnk2700_LODFFCPFJLC() uint32 {
	if x != nil {
		return x.Unk2700_LODFFCPFJLC
	}
	return 0
}

func (x *SceneGalleryHomeSeekFurnitureInfo) GetUnk2700_HLCIHCCGFFC() uint32 {
	if x != nil {
		return x.Unk2700_HLCIHCCGFFC
	}
	return 0
}

var File_SceneGalleryHomeSeekFurnitureInfo_proto protoreflect.FileDescriptor

var file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDesc = []byte{
	0x0a, 0x27, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x48, 0x6f,
	0x6d, 0x65, 0x53, 0x65, 0x65, 0x6b, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x21, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x65,
	0x65, 0x6b, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x44, 0x42, 0x45, 0x4e,
	0x42, 0x42, 0x4f, 0x44, 0x47, 0x50, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x4b, 0x44, 0x42, 0x45, 0x4e, 0x42, 0x42, 0x4f, 0x44, 0x47, 0x50,
	0x12, 0x6b, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x44, 0x48, 0x4f,
	0x4a, 0x48, 0x4f, 0x49, 0x43, 0x42, 0x4c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x48, 0x6f, 0x6d, 0x65,
	0x53, 0x65, 0x65, 0x6b, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x44, 0x48, 0x4f, 0x4a, 0x48, 0x4f,
	0x49, 0x43, 0x42, 0x4c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x44, 0x44, 0x48, 0x4f, 0x4a, 0x48, 0x4f, 0x49, 0x43, 0x42, 0x4c, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4c, 0x4f, 0x44, 0x46, 0x46, 0x43, 0x50,
	0x46, 0x4a, 0x4c, 0x43, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x4c, 0x4f, 0x44, 0x46, 0x46, 0x43, 0x50, 0x46, 0x4a, 0x4c, 0x43, 0x12, 0x2f,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x4c, 0x43, 0x49, 0x48, 0x43,
	0x43, 0x47, 0x46, 0x46, 0x43, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x48, 0x4c, 0x43, 0x49, 0x48, 0x43, 0x43, 0x47, 0x46, 0x46, 0x43, 0x1a,
	0x45, 0x0a, 0x17, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x44, 0x48, 0x4f, 0x4a, 0x48,
	0x4f, 0x49, 0x43, 0x42, 0x4c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescOnce sync.Once
	file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescData = file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDesc
)

func file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescGZIP() []byte {
	file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescOnce.Do(func() {
		file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescData)
	})
	return file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDescData
}

var file_SceneGalleryHomeSeekFurnitureInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SceneGalleryHomeSeekFurnitureInfo_proto_goTypes = []interface{}{
	(*SceneGalleryHomeSeekFurnitureInfo)(nil), // 0: SceneGalleryHomeSeekFurnitureInfo
	nil, // 1: SceneGalleryHomeSeekFurnitureInfo.Unk2700DDHOJHOICBLEntry
}
var file_SceneGalleryHomeSeekFurnitureInfo_proto_depIdxs = []int32{
	1, // 0: SceneGalleryHomeSeekFurnitureInfo.Unk2700_DDHOJHOICBL:type_name -> SceneGalleryHomeSeekFurnitureInfo.Unk2700DDHOJHOICBLEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneGalleryHomeSeekFurnitureInfo_proto_init() }
func file_SceneGalleryHomeSeekFurnitureInfo_proto_init() {
	if File_SceneGalleryHomeSeekFurnitureInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneGalleryHomeSeekFurnitureInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGalleryHomeSeekFurnitureInfo); i {
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
			RawDescriptor: file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGalleryHomeSeekFurnitureInfo_proto_goTypes,
		DependencyIndexes: file_SceneGalleryHomeSeekFurnitureInfo_proto_depIdxs,
		MessageInfos:      file_SceneGalleryHomeSeekFurnitureInfo_proto_msgTypes,
	}.Build()
	File_SceneGalleryHomeSeekFurnitureInfo_proto = out.File
	file_SceneGalleryHomeSeekFurnitureInfo_proto_rawDesc = nil
	file_SceneGalleryHomeSeekFurnitureInfo_proto_goTypes = nil
	file_SceneGalleryHomeSeekFurnitureInfo_proto_depIdxs = nil
}
