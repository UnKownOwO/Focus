// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HomeSceneArrangementMuipData.proto

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

type HomeSceneArrangementMuipData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleId      uint32                          `protobuf:"varint,1,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	SceneId       uint32                          `protobuf:"varint,2,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	IsRoom        bool                            `protobuf:"varint,3,opt,name=is_room,json=isRoom,proto3" json:"is_room,omitempty"`
	BlockDataList []*HomeBlockArrangementMuipData `protobuf:"bytes,4,rep,name=block_data_list,json=blockDataList,proto3" json:"block_data_list,omitempty"`
}

func (x *HomeSceneArrangementMuipData) Reset() {
	*x = HomeSceneArrangementMuipData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeSceneArrangementMuipData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeSceneArrangementMuipData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeSceneArrangementMuipData) ProtoMessage() {}

func (x *HomeSceneArrangementMuipData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeSceneArrangementMuipData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeSceneArrangementMuipData.ProtoReflect.Descriptor instead.
func (*HomeSceneArrangementMuipData) Descriptor() ([]byte, []int) {
	return file_HomeSceneArrangementMuipData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeSceneArrangementMuipData) GetModuleId() uint32 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *HomeSceneArrangementMuipData) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *HomeSceneArrangementMuipData) GetIsRoom() bool {
	if x != nil {
		return x.IsRoom
	}
	return false
}

func (x *HomeSceneArrangementMuipData) GetBlockDataList() []*HomeBlockArrangementMuipData {
	if x != nil {
		return x.BlockDataList
	}
	return nil
}

var File_HomeSceneArrangementMuipData_proto protoreflect.FileDescriptor

var file_HomeSceneArrangementMuipData_proto_rawDesc = []byte{
	0x0a, 0x22, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41,
	0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x1c, 0x48, 0x6f, 0x6d,
	0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x45, 0x0a, 0x0f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41,
	0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeSceneArrangementMuipData_proto_rawDescOnce sync.Once
	file_HomeSceneArrangementMuipData_proto_rawDescData = file_HomeSceneArrangementMuipData_proto_rawDesc
)

func file_HomeSceneArrangementMuipData_proto_rawDescGZIP() []byte {
	file_HomeSceneArrangementMuipData_proto_rawDescOnce.Do(func() {
		file_HomeSceneArrangementMuipData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeSceneArrangementMuipData_proto_rawDescData)
	})
	return file_HomeSceneArrangementMuipData_proto_rawDescData
}

var file_HomeSceneArrangementMuipData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeSceneArrangementMuipData_proto_goTypes = []interface{}{
	(*HomeSceneArrangementMuipData)(nil), // 0: HomeSceneArrangementMuipData
	(*HomeBlockArrangementMuipData)(nil), // 1: HomeBlockArrangementMuipData
}
var file_HomeSceneArrangementMuipData_proto_depIdxs = []int32{
	1, // 0: HomeSceneArrangementMuipData.block_data_list:type_name -> HomeBlockArrangementMuipData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeSceneArrangementMuipData_proto_init() }
func file_HomeSceneArrangementMuipData_proto_init() {
	if File_HomeSceneArrangementMuipData_proto != nil {
		return
	}
	file_HomeBlockArrangementMuipData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeSceneArrangementMuipData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeSceneArrangementMuipData); i {
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
			RawDescriptor: file_HomeSceneArrangementMuipData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeSceneArrangementMuipData_proto_goTypes,
		DependencyIndexes: file_HomeSceneArrangementMuipData_proto_depIdxs,
		MessageInfos:      file_HomeSceneArrangementMuipData_proto_msgTypes,
	}.Build()
	File_HomeSceneArrangementMuipData_proto = out.File
	file_HomeSceneArrangementMuipData_proto_rawDesc = nil
	file_HomeSceneArrangementMuipData_proto_goTypes = nil
	file_HomeSceneArrangementMuipData_proto_depIdxs = nil
}
