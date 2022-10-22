// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GalleryBalloonScoreNotify.proto

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

// CmdId: 5512
// EnetChannelId: 0
// EnetIsReliable: true
type GalleryBalloonScoreNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GalleryId   uint32            `protobuf:"varint,9,opt,name=gallery_id,json=galleryId,proto3" json:"gallery_id,omitempty"`
	UidScoreMap map[uint32]uint32 `protobuf:"bytes,7,rep,name=uid_score_map,json=uidScoreMap,proto3" json:"uid_score_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GalleryBalloonScoreNotify) Reset() {
	*x = GalleryBalloonScoreNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GalleryBalloonScoreNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GalleryBalloonScoreNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GalleryBalloonScoreNotify) ProtoMessage() {}

func (x *GalleryBalloonScoreNotify) ProtoReflect() protoreflect.Message {
	mi := &file_GalleryBalloonScoreNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GalleryBalloonScoreNotify.ProtoReflect.Descriptor instead.
func (*GalleryBalloonScoreNotify) Descriptor() ([]byte, []int) {
	return file_GalleryBalloonScoreNotify_proto_rawDescGZIP(), []int{0}
}

func (x *GalleryBalloonScoreNotify) GetGalleryId() uint32 {
	if x != nil {
		return x.GalleryId
	}
	return 0
}

func (x *GalleryBalloonScoreNotify) GetUidScoreMap() map[uint32]uint32 {
	if x != nil {
		return x.UidScoreMap
	}
	return nil
}

var File_GalleryBalloonScoreNotify_proto protoreflect.FileDescriptor

var file_GalleryBalloonScoreNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x19, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x42, 0x61, 0x6c,
	0x6c, 0x6f, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x4f,
	0x0a, 0x0d, 0x75, 0x69, 0x64, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x42,
	0x61, 0x6c, 0x6c, 0x6f, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x55, 0x69, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x75, 0x69, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x1a,
	0x3e, 0x0a, 0x10, 0x55, 0x69, 0x64, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GalleryBalloonScoreNotify_proto_rawDescOnce sync.Once
	file_GalleryBalloonScoreNotify_proto_rawDescData = file_GalleryBalloonScoreNotify_proto_rawDesc
)

func file_GalleryBalloonScoreNotify_proto_rawDescGZIP() []byte {
	file_GalleryBalloonScoreNotify_proto_rawDescOnce.Do(func() {
		file_GalleryBalloonScoreNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_GalleryBalloonScoreNotify_proto_rawDescData)
	})
	return file_GalleryBalloonScoreNotify_proto_rawDescData
}

var file_GalleryBalloonScoreNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GalleryBalloonScoreNotify_proto_goTypes = []interface{}{
	(*GalleryBalloonScoreNotify)(nil), // 0: GalleryBalloonScoreNotify
	nil,                               // 1: GalleryBalloonScoreNotify.UidScoreMapEntry
}
var file_GalleryBalloonScoreNotify_proto_depIdxs = []int32{
	1, // 0: GalleryBalloonScoreNotify.uid_score_map:type_name -> GalleryBalloonScoreNotify.UidScoreMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GalleryBalloonScoreNotify_proto_init() }
func file_GalleryBalloonScoreNotify_proto_init() {
	if File_GalleryBalloonScoreNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GalleryBalloonScoreNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GalleryBalloonScoreNotify); i {
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
			RawDescriptor: file_GalleryBalloonScoreNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GalleryBalloonScoreNotify_proto_goTypes,
		DependencyIndexes: file_GalleryBalloonScoreNotify_proto_depIdxs,
		MessageInfos:      file_GalleryBalloonScoreNotify_proto_msgTypes,
	}.Build()
	File_GalleryBalloonScoreNotify_proto = out.File
	file_GalleryBalloonScoreNotify_proto_rawDesc = nil
	file_GalleryBalloonScoreNotify_proto_goTypes = nil
	file_GalleryBalloonScoreNotify_proto_depIdxs = nil
}
