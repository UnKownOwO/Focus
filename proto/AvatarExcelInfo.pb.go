// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AvatarExcelInfo.proto

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

type AvatarExcelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrefabPathHash           uint64 `protobuf:"varint,1,opt,name=prefab_path_hash,json=prefabPathHash,proto3" json:"prefab_path_hash,omitempty"`
	PrefabPathRemoteHash     uint64 `protobuf:"varint,2,opt,name=prefab_path_remote_hash,json=prefabPathRemoteHash,proto3" json:"prefab_path_remote_hash,omitempty"`
	ControllerPathHash       uint64 `protobuf:"varint,3,opt,name=controller_path_hash,json=controllerPathHash,proto3" json:"controller_path_hash,omitempty"`
	ControllerPathRemoteHash uint64 `protobuf:"varint,4,opt,name=controller_path_remote_hash,json=controllerPathRemoteHash,proto3" json:"controller_path_remote_hash,omitempty"`
	CombatConfigHash         uint64 `protobuf:"varint,5,opt,name=combat_config_hash,json=combatConfigHash,proto3" json:"combat_config_hash,omitempty"`
}

func (x *AvatarExcelInfo) Reset() {
	*x = AvatarExcelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AvatarExcelInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvatarExcelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvatarExcelInfo) ProtoMessage() {}

func (x *AvatarExcelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AvatarExcelInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvatarExcelInfo.ProtoReflect.Descriptor instead.
func (*AvatarExcelInfo) Descriptor() ([]byte, []int) {
	return file_AvatarExcelInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AvatarExcelInfo) GetPrefabPathHash() uint64 {
	if x != nil {
		return x.PrefabPathHash
	}
	return 0
}

func (x *AvatarExcelInfo) GetPrefabPathRemoteHash() uint64 {
	if x != nil {
		return x.PrefabPathRemoteHash
	}
	return 0
}

func (x *AvatarExcelInfo) GetControllerPathHash() uint64 {
	if x != nil {
		return x.ControllerPathHash
	}
	return 0
}

func (x *AvatarExcelInfo) GetControllerPathRemoteHash() uint64 {
	if x != nil {
		return x.ControllerPathRemoteHash
	}
	return 0
}

func (x *AvatarExcelInfo) GetCombatConfigHash() uint64 {
	if x != nil {
		return x.CombatConfigHash
	}
	return 0
}

var File_AvatarExcelInfo_proto protoreflect.FileDescriptor

var file_AvatarExcelInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x02, 0x0a, 0x0f, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x45, 0x78, 0x63, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x70,
	0x72, 0x65, 0x66, 0x61, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61, 0x74,
	0x68, 0x48, 0x61, 0x73, 0x68, 0x12, 0x35, 0x0a, 0x17, 0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x30, 0x0a, 0x14,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x48, 0x61, 0x73, 0x68, 0x12, 0x3d,
	0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x61, 0x73, 0x68, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AvatarExcelInfo_proto_rawDescOnce sync.Once
	file_AvatarExcelInfo_proto_rawDescData = file_AvatarExcelInfo_proto_rawDesc
)

func file_AvatarExcelInfo_proto_rawDescGZIP() []byte {
	file_AvatarExcelInfo_proto_rawDescOnce.Do(func() {
		file_AvatarExcelInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AvatarExcelInfo_proto_rawDescData)
	})
	return file_AvatarExcelInfo_proto_rawDescData
}

var file_AvatarExcelInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AvatarExcelInfo_proto_goTypes = []interface{}{
	(*AvatarExcelInfo)(nil), // 0: AvatarExcelInfo
}
var file_AvatarExcelInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AvatarExcelInfo_proto_init() }
func file_AvatarExcelInfo_proto_init() {
	if File_AvatarExcelInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_AvatarExcelInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvatarExcelInfo); i {
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
			RawDescriptor: file_AvatarExcelInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AvatarExcelInfo_proto_goTypes,
		DependencyIndexes: file_AvatarExcelInfo_proto_depIdxs,
		MessageInfos:      file_AvatarExcelInfo_proto_msgTypes,
	}.Build()
	File_AvatarExcelInfo_proto = out.File
	file_AvatarExcelInfo_proto_rawDesc = nil
	file_AvatarExcelInfo_proto_goTypes = nil
	file_AvatarExcelInfo_proto_depIdxs = nil
}
