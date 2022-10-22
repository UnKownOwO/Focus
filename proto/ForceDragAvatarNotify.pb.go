// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ForceDragAvatarNotify.proto

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

// CmdId: 3235
// EnetChannelId: 0
// EnetIsReliable: true
type ForceDragAvatarNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneTime      uint32      `protobuf:"varint,3,opt,name=scene_time,json=sceneTime,proto3" json:"scene_time,omitempty"`
	DeltaTimeMs    uint64      `protobuf:"varint,1,opt,name=delta_time_ms,json=deltaTimeMs,proto3" json:"delta_time_ms,omitempty"`
	EntityId       uint32      `protobuf:"varint,2,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	MotionInfo     *MotionInfo `protobuf:"bytes,10,opt,name=motion_info,json=motionInfo,proto3" json:"motion_info,omitempty"`
	IsFirstValid   bool        `protobuf:"varint,8,opt,name=is_first_valid,json=isFirstValid,proto3" json:"is_first_valid,omitempty"`
	LastMoveTimeMs uint64      `protobuf:"varint,12,opt,name=last_move_time_ms,json=lastMoveTimeMs,proto3" json:"last_move_time_ms,omitempty"`
}

func (x *ForceDragAvatarNotify) Reset() {
	*x = ForceDragAvatarNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ForceDragAvatarNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceDragAvatarNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceDragAvatarNotify) ProtoMessage() {}

func (x *ForceDragAvatarNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ForceDragAvatarNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceDragAvatarNotify.ProtoReflect.Descriptor instead.
func (*ForceDragAvatarNotify) Descriptor() ([]byte, []int) {
	return file_ForceDragAvatarNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ForceDragAvatarNotify) GetSceneTime() uint32 {
	if x != nil {
		return x.SceneTime
	}
	return 0
}

func (x *ForceDragAvatarNotify) GetDeltaTimeMs() uint64 {
	if x != nil {
		return x.DeltaTimeMs
	}
	return 0
}

func (x *ForceDragAvatarNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *ForceDragAvatarNotify) GetMotionInfo() *MotionInfo {
	if x != nil {
		return x.MotionInfo
	}
	return nil
}

func (x *ForceDragAvatarNotify) GetIsFirstValid() bool {
	if x != nil {
		return x.IsFirstValid
	}
	return false
}

func (x *ForceDragAvatarNotify) GetLastMoveTimeMs() uint64 {
	if x != nil {
		return x.LastMoveTimeMs
	}
	return 0
}

var File_ForceDragAvatarNotify_proto protoreflect.FileDescriptor

var file_ForceDragAvatarNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x44, 0x72, 0x61, 0x67, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf6, 0x01, 0x0a, 0x15, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x44, 0x72, 0x61, 0x67, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x6d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x29, 0x0a,
	0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6d, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ForceDragAvatarNotify_proto_rawDescOnce sync.Once
	file_ForceDragAvatarNotify_proto_rawDescData = file_ForceDragAvatarNotify_proto_rawDesc
)

func file_ForceDragAvatarNotify_proto_rawDescGZIP() []byte {
	file_ForceDragAvatarNotify_proto_rawDescOnce.Do(func() {
		file_ForceDragAvatarNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ForceDragAvatarNotify_proto_rawDescData)
	})
	return file_ForceDragAvatarNotify_proto_rawDescData
}

var file_ForceDragAvatarNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ForceDragAvatarNotify_proto_goTypes = []interface{}{
	(*ForceDragAvatarNotify)(nil), // 0: ForceDragAvatarNotify
	(*MotionInfo)(nil),            // 1: MotionInfo
}
var file_ForceDragAvatarNotify_proto_depIdxs = []int32{
	1, // 0: ForceDragAvatarNotify.motion_info:type_name -> MotionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ForceDragAvatarNotify_proto_init() }
func file_ForceDragAvatarNotify_proto_init() {
	if File_ForceDragAvatarNotify_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ForceDragAvatarNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceDragAvatarNotify); i {
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
			RawDescriptor: file_ForceDragAvatarNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ForceDragAvatarNotify_proto_goTypes,
		DependencyIndexes: file_ForceDragAvatarNotify_proto_depIdxs,
		MessageInfos:      file_ForceDragAvatarNotify_proto_msgTypes,
	}.Build()
	File_ForceDragAvatarNotify_proto = out.File
	file_ForceDragAvatarNotify_proto_rawDesc = nil
	file_ForceDragAvatarNotify_proto_goTypes = nil
	file_ForceDragAvatarNotify_proto_depIdxs = nil
}