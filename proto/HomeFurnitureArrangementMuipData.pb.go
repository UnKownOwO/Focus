// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HomeFurnitureArrangementMuipData.proto

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

type HomeFurnitureArrangementMuipData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FurnitureId uint32  `protobuf:"varint,1,opt,name=furniture_id,json=furnitureId,proto3" json:"furniture_id,omitempty"`
	SpawnPos    *Vector `protobuf:"bytes,2,opt,name=spawn_pos,json=spawnPos,proto3" json:"spawn_pos,omitempty"`
	SpawnRot    *Vector `protobuf:"bytes,3,opt,name=spawn_rot,json=spawnRot,proto3" json:"spawn_rot,omitempty"`
}

func (x *HomeFurnitureArrangementMuipData) Reset() {
	*x = HomeFurnitureArrangementMuipData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeFurnitureArrangementMuipData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeFurnitureArrangementMuipData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeFurnitureArrangementMuipData) ProtoMessage() {}

func (x *HomeFurnitureArrangementMuipData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeFurnitureArrangementMuipData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeFurnitureArrangementMuipData.ProtoReflect.Descriptor instead.
func (*HomeFurnitureArrangementMuipData) Descriptor() ([]byte, []int) {
	return file_HomeFurnitureArrangementMuipData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeFurnitureArrangementMuipData) GetFurnitureId() uint32 {
	if x != nil {
		return x.FurnitureId
	}
	return 0
}

func (x *HomeFurnitureArrangementMuipData) GetSpawnPos() *Vector {
	if x != nil {
		return x.SpawnPos
	}
	return nil
}

func (x *HomeFurnitureArrangementMuipData) GetSpawnRot() *Vector {
	if x != nil {
		return x.SpawnRot
	}
	return nil
}

var File_HomeFurnitureArrangementMuipData_proto protoreflect.FileDescriptor

var file_HomeFurnitureArrangementMuipData_proto_rawDesc = []byte{
	0x0a, 0x26, 0x48, 0x6f, 0x6d, 0x65, 0x46, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x41,
	0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x20, 0x48, 0x6f, 0x6d, 0x65, 0x46,
	0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x66, 0x75, 0x72, 0x6e, 0x69, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x09, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73, 0x70, 0x61, 0x77,
	0x6e, 0x50, 0x6f, 0x73, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x5f, 0x72, 0x6f,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x08, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x52, 0x6f, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeFurnitureArrangementMuipData_proto_rawDescOnce sync.Once
	file_HomeFurnitureArrangementMuipData_proto_rawDescData = file_HomeFurnitureArrangementMuipData_proto_rawDesc
)

func file_HomeFurnitureArrangementMuipData_proto_rawDescGZIP() []byte {
	file_HomeFurnitureArrangementMuipData_proto_rawDescOnce.Do(func() {
		file_HomeFurnitureArrangementMuipData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeFurnitureArrangementMuipData_proto_rawDescData)
	})
	return file_HomeFurnitureArrangementMuipData_proto_rawDescData
}

var file_HomeFurnitureArrangementMuipData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeFurnitureArrangementMuipData_proto_goTypes = []interface{}{
	(*HomeFurnitureArrangementMuipData)(nil), // 0: HomeFurnitureArrangementMuipData
	(*Vector)(nil),                           // 1: Vector
}
var file_HomeFurnitureArrangementMuipData_proto_depIdxs = []int32{
	1, // 0: HomeFurnitureArrangementMuipData.spawn_pos:type_name -> Vector
	1, // 1: HomeFurnitureArrangementMuipData.spawn_rot:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HomeFurnitureArrangementMuipData_proto_init() }
func file_HomeFurnitureArrangementMuipData_proto_init() {
	if File_HomeFurnitureArrangementMuipData_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeFurnitureArrangementMuipData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeFurnitureArrangementMuipData); i {
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
			RawDescriptor: file_HomeFurnitureArrangementMuipData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeFurnitureArrangementMuipData_proto_goTypes,
		DependencyIndexes: file_HomeFurnitureArrangementMuipData_proto_depIdxs,
		MessageInfos:      file_HomeFurnitureArrangementMuipData_proto_msgTypes,
	}.Build()
	File_HomeFurnitureArrangementMuipData_proto = out.File
	file_HomeFurnitureArrangementMuipData_proto_rawDesc = nil
	file_HomeFurnitureArrangementMuipData_proto_goTypes = nil
	file_HomeFurnitureArrangementMuipData_proto_depIdxs = nil
}
