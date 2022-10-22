// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: VehicleLocationInfo.proto

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

type VehicleLocationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rot      *Vector  `protobuf:"bytes,14,opt,name=rot,proto3" json:"rot,omitempty"`
	EntityId uint32   `protobuf:"varint,15,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	CurHp    float32  `protobuf:"fixed32,11,opt,name=cur_hp,json=curHp,proto3" json:"cur_hp,omitempty"`
	OwnerUid uint32   `protobuf:"varint,5,opt,name=owner_uid,json=ownerUid,proto3" json:"owner_uid,omitempty"`
	Pos      *Vector  `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	UidList  []uint32 `protobuf:"varint,3,rep,packed,name=uid_list,json=uidList,proto3" json:"uid_list,omitempty"`
	GadgetId uint32   `protobuf:"varint,13,opt,name=gadget_id,json=gadgetId,proto3" json:"gadget_id,omitempty"`
	MaxHp    float32  `protobuf:"fixed32,6,opt,name=max_hp,json=maxHp,proto3" json:"max_hp,omitempty"`
}

func (x *VehicleLocationInfo) Reset() {
	*x = VehicleLocationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_VehicleLocationInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleLocationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleLocationInfo) ProtoMessage() {}

func (x *VehicleLocationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_VehicleLocationInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleLocationInfo.ProtoReflect.Descriptor instead.
func (*VehicleLocationInfo) Descriptor() ([]byte, []int) {
	return file_VehicleLocationInfo_proto_rawDescGZIP(), []int{0}
}

func (x *VehicleLocationInfo) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

func (x *VehicleLocationInfo) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *VehicleLocationInfo) GetCurHp() float32 {
	if x != nil {
		return x.CurHp
	}
	return 0
}

func (x *VehicleLocationInfo) GetOwnerUid() uint32 {
	if x != nil {
		return x.OwnerUid
	}
	return 0
}

func (x *VehicleLocationInfo) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *VehicleLocationInfo) GetUidList() []uint32 {
	if x != nil {
		return x.UidList
	}
	return nil
}

func (x *VehicleLocationInfo) GetGadgetId() uint32 {
	if x != nil {
		return x.GadgetId
	}
	return 0
}

func (x *VehicleLocationInfo) GetMaxHp() float32 {
	if x != nil {
		return x.MaxHp
	}
	return 0
}

var File_VehicleLocationInfo_proto protoreflect.FileDescriptor

var file_VehicleLocationInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x01, 0x0a, 0x13, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x19, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x75, 0x72,
	0x5f, 0x68, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x63, 0x75, 0x72, 0x48, 0x70,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x03, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x69, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x69, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x68, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x6d, 0x61, 0x78, 0x48, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VehicleLocationInfo_proto_rawDescOnce sync.Once
	file_VehicleLocationInfo_proto_rawDescData = file_VehicleLocationInfo_proto_rawDesc
)

func file_VehicleLocationInfo_proto_rawDescGZIP() []byte {
	file_VehicleLocationInfo_proto_rawDescOnce.Do(func() {
		file_VehicleLocationInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_VehicleLocationInfo_proto_rawDescData)
	})
	return file_VehicleLocationInfo_proto_rawDescData
}

var file_VehicleLocationInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_VehicleLocationInfo_proto_goTypes = []interface{}{
	(*VehicleLocationInfo)(nil), // 0: VehicleLocationInfo
	(*Vector)(nil),              // 1: Vector
}
var file_VehicleLocationInfo_proto_depIdxs = []int32{
	1, // 0: VehicleLocationInfo.rot:type_name -> Vector
	1, // 1: VehicleLocationInfo.pos:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_VehicleLocationInfo_proto_init() }
func file_VehicleLocationInfo_proto_init() {
	if File_VehicleLocationInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_VehicleLocationInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleLocationInfo); i {
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
			RawDescriptor: file_VehicleLocationInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VehicleLocationInfo_proto_goTypes,
		DependencyIndexes: file_VehicleLocationInfo_proto_depIdxs,
		MessageInfos:      file_VehicleLocationInfo_proto_msgTypes,
	}.Build()
	File_VehicleLocationInfo_proto = out.File
	file_VehicleLocationInfo_proto_rawDesc = nil
	file_VehicleLocationInfo_proto_goTypes = nil
	file_VehicleLocationInfo_proto_depIdxs = nil
}