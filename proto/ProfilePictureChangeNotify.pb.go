// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ProfilePictureChangeNotify.proto

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

// CmdId: 4016
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ProfilePictureChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfilePicture *ProfilePicture `protobuf:"bytes,12,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
}

func (x *ProfilePictureChangeNotify) Reset() {
	*x = ProfilePictureChangeNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProfilePictureChangeNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfilePictureChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfilePictureChangeNotify) ProtoMessage() {}

func (x *ProfilePictureChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ProfilePictureChangeNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfilePictureChangeNotify.ProtoReflect.Descriptor instead.
func (*ProfilePictureChangeNotify) Descriptor() ([]byte, []int) {
	return file_ProfilePictureChangeNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ProfilePictureChangeNotify) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

var File_ProfilePictureChangeNotify_proto protoreflect.FileDescriptor

var file_ProfilePictureChangeNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x38, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ProfilePictureChangeNotify_proto_rawDescOnce sync.Once
	file_ProfilePictureChangeNotify_proto_rawDescData = file_ProfilePictureChangeNotify_proto_rawDesc
)

func file_ProfilePictureChangeNotify_proto_rawDescGZIP() []byte {
	file_ProfilePictureChangeNotify_proto_rawDescOnce.Do(func() {
		file_ProfilePictureChangeNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProfilePictureChangeNotify_proto_rawDescData)
	})
	return file_ProfilePictureChangeNotify_proto_rawDescData
}

var file_ProfilePictureChangeNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ProfilePictureChangeNotify_proto_goTypes = []interface{}{
	(*ProfilePictureChangeNotify)(nil), // 0: ProfilePictureChangeNotify
	(*ProfilePicture)(nil),             // 1: ProfilePicture
}
var file_ProfilePictureChangeNotify_proto_depIdxs = []int32{
	1, // 0: ProfilePictureChangeNotify.profile_picture:type_name -> ProfilePicture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ProfilePictureChangeNotify_proto_init() }
func file_ProfilePictureChangeNotify_proto_init() {
	if File_ProfilePictureChangeNotify_proto != nil {
		return
	}
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ProfilePictureChangeNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfilePictureChangeNotify); i {
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
			RawDescriptor: file_ProfilePictureChangeNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ProfilePictureChangeNotify_proto_goTypes,
		DependencyIndexes: file_ProfilePictureChangeNotify_proto_depIdxs,
		MessageInfos:      file_ProfilePictureChangeNotify_proto_msgTypes,
	}.Build()
	File_ProfilePictureChangeNotify_proto = out.File
	file_ProfilePictureChangeNotify_proto_rawDesc = nil
	file_ProfilePictureChangeNotify_proto_goTypes = nil
	file_ProfilePictureChangeNotify_proto_depIdxs = nil
}
