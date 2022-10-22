// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: BuildingInfo.proto

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

type BuildingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildingId          uint32 `protobuf:"varint,1,opt,name=building_id,json=buildingId,proto3" json:"building_id,omitempty"`
	PointConfigId       uint32 `protobuf:"varint,2,opt,name=point_config_id,json=pointConfigId,proto3" json:"point_config_id,omitempty"`
	Cost                uint32 `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	Refund              uint32 `protobuf:"varint,5,opt,name=refund,proto3" json:"refund,omitempty"`
	OwnerUid            uint32 `protobuf:"varint,6,opt,name=owner_uid,json=ownerUid,proto3" json:"owner_uid,omitempty"`
	Unk2700_MDJOPHOHFDB uint32 `protobuf:"varint,7,opt,name=Unk2700_MDJOPHOHFDB,json=Unk2700MDJOPHOHFDB,proto3" json:"Unk2700_MDJOPHOHFDB,omitempty"`
	Unk2700_COFBIGLBNGP uint32 `protobuf:"varint,8,opt,name=Unk2700_COFBIGLBNGP,json=Unk2700COFBIGLBNGP,proto3" json:"Unk2700_COFBIGLBNGP,omitempty"`
}

func (x *BuildingInfo) Reset() {
	*x = BuildingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BuildingInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildingInfo) ProtoMessage() {}

func (x *BuildingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BuildingInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildingInfo.ProtoReflect.Descriptor instead.
func (*BuildingInfo) Descriptor() ([]byte, []int) {
	return file_BuildingInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BuildingInfo) GetBuildingId() uint32 {
	if x != nil {
		return x.BuildingId
	}
	return 0
}

func (x *BuildingInfo) GetPointConfigId() uint32 {
	if x != nil {
		return x.PointConfigId
	}
	return 0
}

func (x *BuildingInfo) GetCost() uint32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *BuildingInfo) GetRefund() uint32 {
	if x != nil {
		return x.Refund
	}
	return 0
}

func (x *BuildingInfo) GetOwnerUid() uint32 {
	if x != nil {
		return x.OwnerUid
	}
	return 0
}

func (x *BuildingInfo) GetUnk2700_MDJOPHOHFDB() uint32 {
	if x != nil {
		return x.Unk2700_MDJOPHOHFDB
	}
	return 0
}

func (x *BuildingInfo) GetUnk2700_COFBIGLBNGP() uint32 {
	if x != nil {
		return x.Unk2700_COFBIGLBNGP
	}
	return 0
}

var File_BuildingInfo_proto protoreflect.FileDescriptor

var file_BuildingInfo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x02, 0x0a, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x66, 0x75, 0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x4d, 0x44, 0x4a, 0x4f, 0x50, 0x48, 0x4f, 0x48, 0x46, 0x44, 0x42, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4d, 0x44, 0x4a,
	0x4f, 0x50, 0x48, 0x4f, 0x48, 0x46, 0x44, 0x42, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x43, 0x4f, 0x46, 0x42, 0x49, 0x47, 0x4c, 0x42, 0x4e, 0x47, 0x50, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x43, 0x4f,
	0x46, 0x42, 0x49, 0x47, 0x4c, 0x42, 0x4e, 0x47, 0x50, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BuildingInfo_proto_rawDescOnce sync.Once
	file_BuildingInfo_proto_rawDescData = file_BuildingInfo_proto_rawDesc
)

func file_BuildingInfo_proto_rawDescGZIP() []byte {
	file_BuildingInfo_proto_rawDescOnce.Do(func() {
		file_BuildingInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BuildingInfo_proto_rawDescData)
	})
	return file_BuildingInfo_proto_rawDescData
}

var file_BuildingInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BuildingInfo_proto_goTypes = []interface{}{
	(*BuildingInfo)(nil), // 0: BuildingInfo
}
var file_BuildingInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BuildingInfo_proto_init() }
func file_BuildingInfo_proto_init() {
	if File_BuildingInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BuildingInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildingInfo); i {
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
			RawDescriptor: file_BuildingInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BuildingInfo_proto_goTypes,
		DependencyIndexes: file_BuildingInfo_proto_depIdxs,
		MessageInfos:      file_BuildingInfo_proto_msgTypes,
	}.Build()
	File_BuildingInfo_proto = out.File
	file_BuildingInfo_proto_rawDesc = nil
	file_BuildingInfo_proto_goTypes = nil
	file_BuildingInfo_proto_depIdxs = nil
}
