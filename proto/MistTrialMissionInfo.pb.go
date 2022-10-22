// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: MistTrialMissionInfo.proto

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

type MistTrialMissionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param         uint32 `protobuf:"varint,9,opt,name=param,proto3" json:"param,omitempty"`
	WatcherListId uint32 `protobuf:"varint,13,opt,name=watcher_list_id,json=watcherListId,proto3" json:"watcher_list_id,omitempty"`
}

func (x *MistTrialMissionInfo) Reset() {
	*x = MistTrialMissionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MistTrialMissionInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MistTrialMissionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MistTrialMissionInfo) ProtoMessage() {}

func (x *MistTrialMissionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MistTrialMissionInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MistTrialMissionInfo.ProtoReflect.Descriptor instead.
func (*MistTrialMissionInfo) Descriptor() ([]byte, []int) {
	return file_MistTrialMissionInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MistTrialMissionInfo) GetParam() uint32 {
	if x != nil {
		return x.Param
	}
	return 0
}

func (x *MistTrialMissionInfo) GetWatcherListId() uint32 {
	if x != nil {
		return x.WatcherListId
	}
	return 0
}

var File_MistTrialMissionInfo_proto protoreflect.FileDescriptor

var file_MistTrialMissionInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x14,
	0x4d, 0x69, 0x73, 0x74, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MistTrialMissionInfo_proto_rawDescOnce sync.Once
	file_MistTrialMissionInfo_proto_rawDescData = file_MistTrialMissionInfo_proto_rawDesc
)

func file_MistTrialMissionInfo_proto_rawDescGZIP() []byte {
	file_MistTrialMissionInfo_proto_rawDescOnce.Do(func() {
		file_MistTrialMissionInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MistTrialMissionInfo_proto_rawDescData)
	})
	return file_MistTrialMissionInfo_proto_rawDescData
}

var file_MistTrialMissionInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MistTrialMissionInfo_proto_goTypes = []interface{}{
	(*MistTrialMissionInfo)(nil), // 0: MistTrialMissionInfo
}
var file_MistTrialMissionInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MistTrialMissionInfo_proto_init() }
func file_MistTrialMissionInfo_proto_init() {
	if File_MistTrialMissionInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MistTrialMissionInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MistTrialMissionInfo); i {
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
			RawDescriptor: file_MistTrialMissionInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MistTrialMissionInfo_proto_goTypes,
		DependencyIndexes: file_MistTrialMissionInfo_proto_depIdxs,
		MessageInfos:      file_MistTrialMissionInfo_proto_msgTypes,
	}.Build()
	File_MistTrialMissionInfo_proto = out.File
	file_MistTrialMissionInfo_proto_rawDesc = nil
	file_MistTrialMissionInfo_proto_goTypes = nil
	file_MistTrialMissionInfo_proto_depIdxs = nil
}
