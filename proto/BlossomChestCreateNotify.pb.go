// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: BlossomChestCreateNotify.proto

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

// CmdId: 2721
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type BlossomChestCreateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshId    uint32 `protobuf:"varint,1,opt,name=refresh_id,json=refreshId,proto3" json:"refresh_id,omitempty"`
	CircleCampId uint32 `protobuf:"varint,10,opt,name=circle_camp_id,json=circleCampId,proto3" json:"circle_camp_id,omitempty"`
}

func (x *BlossomChestCreateNotify) Reset() {
	*x = BlossomChestCreateNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlossomChestCreateNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlossomChestCreateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlossomChestCreateNotify) ProtoMessage() {}

func (x *BlossomChestCreateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BlossomChestCreateNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlossomChestCreateNotify.ProtoReflect.Descriptor instead.
func (*BlossomChestCreateNotify) Descriptor() ([]byte, []int) {
	return file_BlossomChestCreateNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BlossomChestCreateNotify) GetRefreshId() uint32 {
	if x != nil {
		return x.RefreshId
	}
	return 0
}

func (x *BlossomChestCreateNotify) GetCircleCampId() uint32 {
	if x != nil {
		return x.CircleCampId
	}
	return 0
}

var File_BlossomChestCreateNotify_proto protoreflect.FileDescriptor

var file_BlossomChestCreateNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5f, 0x0a, 0x18, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x43, 0x68, 0x65, 0x73, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x63,
	0x69, 0x72, 0x63, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x49,
	0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BlossomChestCreateNotify_proto_rawDescOnce sync.Once
	file_BlossomChestCreateNotify_proto_rawDescData = file_BlossomChestCreateNotify_proto_rawDesc
)

func file_BlossomChestCreateNotify_proto_rawDescGZIP() []byte {
	file_BlossomChestCreateNotify_proto_rawDescOnce.Do(func() {
		file_BlossomChestCreateNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BlossomChestCreateNotify_proto_rawDescData)
	})
	return file_BlossomChestCreateNotify_proto_rawDescData
}

var file_BlossomChestCreateNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BlossomChestCreateNotify_proto_goTypes = []interface{}{
	(*BlossomChestCreateNotify)(nil), // 0: BlossomChestCreateNotify
}
var file_BlossomChestCreateNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BlossomChestCreateNotify_proto_init() }
func file_BlossomChestCreateNotify_proto_init() {
	if File_BlossomChestCreateNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BlossomChestCreateNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlossomChestCreateNotify); i {
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
			RawDescriptor: file_BlossomChestCreateNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BlossomChestCreateNotify_proto_goTypes,
		DependencyIndexes: file_BlossomChestCreateNotify_proto_depIdxs,
		MessageInfos:      file_BlossomChestCreateNotify_proto_msgTypes,
	}.Build()
	File_BlossomChestCreateNotify_proto = out.File
	file_BlossomChestCreateNotify_proto_rawDesc = nil
	file_BlossomChestCreateNotify_proto_goTypes = nil
	file_BlossomChestCreateNotify_proto_depIdxs = nil
}