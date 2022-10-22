// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: WorldRoutineTypeInfo.proto

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

type WorldRoutineTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoutineType          uint32              `protobuf:"varint,13,opt,name=routine_type,json=routineType,proto3" json:"routine_type,omitempty"`
	NextRefreshTime      uint32              `protobuf:"varint,12,opt,name=next_refresh_time,json=nextRefreshTime,proto3" json:"next_refresh_time,omitempty"`
	WorldRoutineInfoList []*WorldRoutineInfo `protobuf:"bytes,3,rep,name=world_routine_info_list,json=worldRoutineInfoList,proto3" json:"world_routine_info_list,omitempty"`
}

func (x *WorldRoutineTypeInfo) Reset() {
	*x = WorldRoutineTypeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WorldRoutineTypeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRoutineTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRoutineTypeInfo) ProtoMessage() {}

func (x *WorldRoutineTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_WorldRoutineTypeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRoutineTypeInfo.ProtoReflect.Descriptor instead.
func (*WorldRoutineTypeInfo) Descriptor() ([]byte, []int) {
	return file_WorldRoutineTypeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *WorldRoutineTypeInfo) GetRoutineType() uint32 {
	if x != nil {
		return x.RoutineType
	}
	return 0
}

func (x *WorldRoutineTypeInfo) GetNextRefreshTime() uint32 {
	if x != nil {
		return x.NextRefreshTime
	}
	return 0
}

func (x *WorldRoutineTypeInfo) GetWorldRoutineInfoList() []*WorldRoutineInfo {
	if x != nil {
		return x.WorldRoutineInfoList
	}
	return nil
}

var File_WorldRoutineTypeInfo_proto protoreflect.FileDescriptor

var file_WorldRoutineTypeInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x17,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x14, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_WorldRoutineTypeInfo_proto_rawDescOnce sync.Once
	file_WorldRoutineTypeInfo_proto_rawDescData = file_WorldRoutineTypeInfo_proto_rawDesc
)

func file_WorldRoutineTypeInfo_proto_rawDescGZIP() []byte {
	file_WorldRoutineTypeInfo_proto_rawDescOnce.Do(func() {
		file_WorldRoutineTypeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_WorldRoutineTypeInfo_proto_rawDescData)
	})
	return file_WorldRoutineTypeInfo_proto_rawDescData
}

var file_WorldRoutineTypeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WorldRoutineTypeInfo_proto_goTypes = []interface{}{
	(*WorldRoutineTypeInfo)(nil), // 0: WorldRoutineTypeInfo
	(*WorldRoutineInfo)(nil),     // 1: WorldRoutineInfo
}
var file_WorldRoutineTypeInfo_proto_depIdxs = []int32{
	1, // 0: WorldRoutineTypeInfo.world_routine_info_list:type_name -> WorldRoutineInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_WorldRoutineTypeInfo_proto_init() }
func file_WorldRoutineTypeInfo_proto_init() {
	if File_WorldRoutineTypeInfo_proto != nil {
		return
	}
	file_WorldRoutineInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WorldRoutineTypeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRoutineTypeInfo); i {
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
			RawDescriptor: file_WorldRoutineTypeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WorldRoutineTypeInfo_proto_goTypes,
		DependencyIndexes: file_WorldRoutineTypeInfo_proto_depIdxs,
		MessageInfos:      file_WorldRoutineTypeInfo_proto_msgTypes,
	}.Build()
	File_WorldRoutineTypeInfo_proto = out.File
	file_WorldRoutineTypeInfo_proto_rawDesc = nil
	file_WorldRoutineTypeInfo_proto_goTypes = nil
	file_WorldRoutineTypeInfo_proto_depIdxs = nil
}
