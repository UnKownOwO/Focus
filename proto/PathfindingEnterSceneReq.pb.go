// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: PathfindingEnterSceneReq.proto

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

// CmdId: 2307
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type PathfindingEnterSceneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId             uint32          `protobuf:"varint,12,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	ActivityId          []uint32        `protobuf:"varint,14,rep,packed,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	Unk2800_NCDFAFMGMIG uint32          `protobuf:"varint,15,opt,name=Unk2800_NCDFAFMGMIG,json=Unk2800NCDFAFMGMIG,proto3" json:"Unk2800_NCDFAFMGMIG,omitempty"`
	Version             uint32          `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	IsEditor            bool            `protobuf:"varint,11,opt,name=is_editor,json=isEditor,proto3" json:"is_editor,omitempty"`
	Obstacles           []*ObstacleInfo `protobuf:"bytes,13,rep,name=obstacles,proto3" json:"obstacles,omitempty"`
	Unk2800_MBDFGODMPFN uint32          `protobuf:"varint,4,opt,name=Unk2800_MBDFGODMPFN,json=Unk2800MBDFGODMPFN,proto3" json:"Unk2800_MBDFGODMPFN,omitempty"`
}

func (x *PathfindingEnterSceneReq) Reset() {
	*x = PathfindingEnterSceneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PathfindingEnterSceneReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathfindingEnterSceneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathfindingEnterSceneReq) ProtoMessage() {}

func (x *PathfindingEnterSceneReq) ProtoReflect() protoreflect.Message {
	mi := &file_PathfindingEnterSceneReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathfindingEnterSceneReq.ProtoReflect.Descriptor instead.
func (*PathfindingEnterSceneReq) Descriptor() ([]byte, []int) {
	return file_PathfindingEnterSceneReq_proto_rawDescGZIP(), []int{0}
}

func (x *PathfindingEnterSceneReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *PathfindingEnterSceneReq) GetActivityId() []uint32 {
	if x != nil {
		return x.ActivityId
	}
	return nil
}

func (x *PathfindingEnterSceneReq) GetUnk2800_NCDFAFMGMIG() uint32 {
	if x != nil {
		return x.Unk2800_NCDFAFMGMIG
	}
	return 0
}

func (x *PathfindingEnterSceneReq) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *PathfindingEnterSceneReq) GetIsEditor() bool {
	if x != nil {
		return x.IsEditor
	}
	return false
}

func (x *PathfindingEnterSceneReq) GetObstacles() []*ObstacleInfo {
	if x != nil {
		return x.Obstacles
	}
	return nil
}

func (x *PathfindingEnterSceneReq) GetUnk2800_MBDFGODMPFN() uint32 {
	if x != nil {
		return x.Unk2800_MBDFGODMPFN
	}
	return 0
}

var File_PathfindingEnterSceneReq_proto protoreflect.FileDescriptor

var file_PathfindingEnterSceneReq_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x50, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x4f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x18, 0x50, 0x61, 0x74, 0x68, 0x66, 0x69, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4e, 0x43, 0x44, 0x46, 0x41, 0x46, 0x4d,
	0x47, 0x4d, 0x49, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32,
	0x38, 0x30, 0x30, 0x4e, 0x43, 0x44, 0x46, 0x41, 0x46, 0x4d, 0x47, 0x4d, 0x49, 0x47, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x45,
	0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x09, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c,
	0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4f, 0x62, 0x73, 0x74, 0x61,
	0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6f, 0x62, 0x73, 0x74, 0x61, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4d, 0x42,
	0x44, 0x46, 0x47, 0x4f, 0x44, 0x4d, 0x50, 0x46, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x4d, 0x42, 0x44, 0x46, 0x47, 0x4f, 0x44, 0x4d,
	0x50, 0x46, 0x4e, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PathfindingEnterSceneReq_proto_rawDescOnce sync.Once
	file_PathfindingEnterSceneReq_proto_rawDescData = file_PathfindingEnterSceneReq_proto_rawDesc
)

func file_PathfindingEnterSceneReq_proto_rawDescGZIP() []byte {
	file_PathfindingEnterSceneReq_proto_rawDescOnce.Do(func() {
		file_PathfindingEnterSceneReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PathfindingEnterSceneReq_proto_rawDescData)
	})
	return file_PathfindingEnterSceneReq_proto_rawDescData
}

var file_PathfindingEnterSceneReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PathfindingEnterSceneReq_proto_goTypes = []interface{}{
	(*PathfindingEnterSceneReq)(nil), // 0: PathfindingEnterSceneReq
	(*ObstacleInfo)(nil),             // 1: ObstacleInfo
}
var file_PathfindingEnterSceneReq_proto_depIdxs = []int32{
	1, // 0: PathfindingEnterSceneReq.obstacles:type_name -> ObstacleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PathfindingEnterSceneReq_proto_init() }
func file_PathfindingEnterSceneReq_proto_init() {
	if File_PathfindingEnterSceneReq_proto != nil {
		return
	}
	file_ObstacleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PathfindingEnterSceneReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathfindingEnterSceneReq); i {
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
			RawDescriptor: file_PathfindingEnterSceneReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PathfindingEnterSceneReq_proto_goTypes,
		DependencyIndexes: file_PathfindingEnterSceneReq_proto_depIdxs,
		MessageInfos:      file_PathfindingEnterSceneReq_proto_msgTypes,
	}.Build()
	File_PathfindingEnterSceneReq_proto = out.File
	file_PathfindingEnterSceneReq_proto_rawDesc = nil
	file_PathfindingEnterSceneReq_proto_goTypes = nil
	file_PathfindingEnterSceneReq_proto_depIdxs = nil
}
