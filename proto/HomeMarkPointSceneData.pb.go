// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HomeMarkPointSceneData.proto

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

type HomeMarkPointSceneData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FurnitureList       []*HomeMarkPointFurnitureData `protobuf:"bytes,6,rep,name=furniture_list,json=furnitureList,proto3" json:"furniture_list,omitempty"`
	TeapotSpiritPos     *Vector                       `protobuf:"bytes,4,opt,name=teapot_spirit_pos,json=teapotSpiritPos,proto3" json:"teapot_spirit_pos,omitempty"`
	SceneId             uint32                        `protobuf:"varint,2,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	ModuleId            uint32                        `protobuf:"varint,5,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	Unk3100_ABBFBELGECB *Vector                       `protobuf:"bytes,11,opt,name=Unk3100_ABBFBELGECB,json=Unk3100ABBFBELGECB,proto3" json:"Unk3100_ABBFBELGECB,omitempty"`
}

func (x *HomeMarkPointSceneData) Reset() {
	*x = HomeMarkPointSceneData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeMarkPointSceneData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeMarkPointSceneData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeMarkPointSceneData) ProtoMessage() {}

func (x *HomeMarkPointSceneData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeMarkPointSceneData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeMarkPointSceneData.ProtoReflect.Descriptor instead.
func (*HomeMarkPointSceneData) Descriptor() ([]byte, []int) {
	return file_HomeMarkPointSceneData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeMarkPointSceneData) GetFurnitureList() []*HomeMarkPointFurnitureData {
	if x != nil {
		return x.FurnitureList
	}
	return nil
}

func (x *HomeMarkPointSceneData) GetTeapotSpiritPos() *Vector {
	if x != nil {
		return x.TeapotSpiritPos
	}
	return nil
}

func (x *HomeMarkPointSceneData) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *HomeMarkPointSceneData) GetModuleId() uint32 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *HomeMarkPointSceneData) GetUnk3100_ABBFBELGECB() *Vector {
	if x != nil {
		return x.Unk3100_ABBFBELGECB
	}
	return nil
}

var File_HomeMarkPointSceneData_proto protoreflect.FileDescriptor

var file_HomeMarkPointSceneData_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x75, 0x72,
	0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83,
	0x02, 0x0a, 0x16, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x0e, 0x66, 0x75, 0x72,
	0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0d,
	0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a,
	0x11, 0x74, 0x65, 0x61, 0x70, 0x6f, 0x74, 0x5f, 0x73, 0x70, 0x69, 0x72, 0x69, 0x74, 0x5f, 0x70,
	0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x0f, 0x74, 0x65, 0x61, 0x70, 0x6f, 0x74, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x50,
	0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x41, 0x42, 0x42, 0x46, 0x42, 0x45, 0x4c, 0x47, 0x45, 0x43,
	0x42, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x41, 0x42, 0x42, 0x46, 0x42, 0x45, 0x4c,
	0x47, 0x45, 0x43, 0x42, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeMarkPointSceneData_proto_rawDescOnce sync.Once
	file_HomeMarkPointSceneData_proto_rawDescData = file_HomeMarkPointSceneData_proto_rawDesc
)

func file_HomeMarkPointSceneData_proto_rawDescGZIP() []byte {
	file_HomeMarkPointSceneData_proto_rawDescOnce.Do(func() {
		file_HomeMarkPointSceneData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeMarkPointSceneData_proto_rawDescData)
	})
	return file_HomeMarkPointSceneData_proto_rawDescData
}

var file_HomeMarkPointSceneData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeMarkPointSceneData_proto_goTypes = []interface{}{
	(*HomeMarkPointSceneData)(nil),     // 0: HomeMarkPointSceneData
	(*HomeMarkPointFurnitureData)(nil), // 1: HomeMarkPointFurnitureData
	(*Vector)(nil),                     // 2: Vector
}
var file_HomeMarkPointSceneData_proto_depIdxs = []int32{
	1, // 0: HomeMarkPointSceneData.furniture_list:type_name -> HomeMarkPointFurnitureData
	2, // 1: HomeMarkPointSceneData.teapot_spirit_pos:type_name -> Vector
	2, // 2: HomeMarkPointSceneData.Unk3100_ABBFBELGECB:type_name -> Vector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_HomeMarkPointSceneData_proto_init() }
func file_HomeMarkPointSceneData_proto_init() {
	if File_HomeMarkPointSceneData_proto != nil {
		return
	}
	file_HomeMarkPointFurnitureData_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeMarkPointSceneData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeMarkPointSceneData); i {
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
			RawDescriptor: file_HomeMarkPointSceneData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeMarkPointSceneData_proto_goTypes,
		DependencyIndexes: file_HomeMarkPointSceneData_proto_depIdxs,
		MessageInfos:      file_HomeMarkPointSceneData_proto_msgTypes,
	}.Build()
	File_HomeMarkPointSceneData_proto = out.File
	file_HomeMarkPointSceneData_proto_rawDesc = nil
	file_HomeMarkPointSceneData_proto_goTypes = nil
	file_HomeMarkPointSceneData_proto_depIdxs = nil
}
