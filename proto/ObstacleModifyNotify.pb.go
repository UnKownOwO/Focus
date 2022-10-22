// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ObstacleModifyNotify.proto

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

// CmdId: 2312
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ObstacleModifyNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoveObstacleIds []int32         `protobuf:"varint,9,rep,packed,name=remove_obstacle_ids,json=removeObstacleIds,proto3" json:"remove_obstacle_ids,omitempty"`
	AddObstacles      []*ObstacleInfo `protobuf:"bytes,6,rep,name=add_obstacles,json=addObstacles,proto3" json:"add_obstacles,omitempty"`
	SceneId           uint32          `protobuf:"varint,5,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *ObstacleModifyNotify) Reset() {
	*x = ObstacleModifyNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ObstacleModifyNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObstacleModifyNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObstacleModifyNotify) ProtoMessage() {}

func (x *ObstacleModifyNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ObstacleModifyNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObstacleModifyNotify.ProtoReflect.Descriptor instead.
func (*ObstacleModifyNotify) Descriptor() ([]byte, []int) {
	return file_ObstacleModifyNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ObstacleModifyNotify) GetRemoveObstacleIds() []int32 {
	if x != nil {
		return x.RemoveObstacleIds
	}
	return nil
}

func (x *ObstacleModifyNotify) GetAddObstacles() []*ObstacleInfo {
	if x != nil {
		return x.AddObstacles
	}
	return nil
}

func (x *ObstacleModifyNotify) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

var File_ObstacleModifyNotify_proto protoreflect.FileDescriptor

var file_ObstacleModifyNotify_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4f, 0x62,
	0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x95, 0x01, 0x0a, 0x14, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x5f, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x62,
	0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x0d, 0x61, 0x64, 0x64,
	0x5f, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x61, 0x64, 0x64, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ObstacleModifyNotify_proto_rawDescOnce sync.Once
	file_ObstacleModifyNotify_proto_rawDescData = file_ObstacleModifyNotify_proto_rawDesc
)

func file_ObstacleModifyNotify_proto_rawDescGZIP() []byte {
	file_ObstacleModifyNotify_proto_rawDescOnce.Do(func() {
		file_ObstacleModifyNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ObstacleModifyNotify_proto_rawDescData)
	})
	return file_ObstacleModifyNotify_proto_rawDescData
}

var file_ObstacleModifyNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ObstacleModifyNotify_proto_goTypes = []interface{}{
	(*ObstacleModifyNotify)(nil), // 0: ObstacleModifyNotify
	(*ObstacleInfo)(nil),         // 1: ObstacleInfo
}
var file_ObstacleModifyNotify_proto_depIdxs = []int32{
	1, // 0: ObstacleModifyNotify.add_obstacles:type_name -> ObstacleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ObstacleModifyNotify_proto_init() }
func file_ObstacleModifyNotify_proto_init() {
	if File_ObstacleModifyNotify_proto != nil {
		return
	}
	file_ObstacleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ObstacleModifyNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObstacleModifyNotify); i {
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
			RawDescriptor: file_ObstacleModifyNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ObstacleModifyNotify_proto_goTypes,
		DependencyIndexes: file_ObstacleModifyNotify_proto_depIdxs,
		MessageInfos:      file_ObstacleModifyNotify_proto_msgTypes,
	}.Build()
	File_ObstacleModifyNotify_proto = out.File
	file_ObstacleModifyNotify_proto_rawDesc = nil
	file_ObstacleModifyNotify_proto_goTypes = nil
	file_ObstacleModifyNotify_proto_depIdxs = nil
}
