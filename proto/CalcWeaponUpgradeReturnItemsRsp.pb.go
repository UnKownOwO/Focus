// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: CalcWeaponUpgradeReturnItemsRsp.proto

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

// CmdId: 684
// EnetChannelId: 0
// EnetIsReliable: true
type CalcWeaponUpgradeReturnItemsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemParamList    []*ItemParam `protobuf:"bytes,4,rep,name=item_param_list,json=itemParamList,proto3" json:"item_param_list,omitempty"`
	Retcode          int32        `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	TargetWeaponGuid uint64       `protobuf:"varint,8,opt,name=target_weapon_guid,json=targetWeaponGuid,proto3" json:"target_weapon_guid,omitempty"`
}

func (x *CalcWeaponUpgradeReturnItemsRsp) Reset() {
	*x = CalcWeaponUpgradeReturnItemsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CalcWeaponUpgradeReturnItemsRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcWeaponUpgradeReturnItemsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcWeaponUpgradeReturnItemsRsp) ProtoMessage() {}

func (x *CalcWeaponUpgradeReturnItemsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CalcWeaponUpgradeReturnItemsRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcWeaponUpgradeReturnItemsRsp.ProtoReflect.Descriptor instead.
func (*CalcWeaponUpgradeReturnItemsRsp) Descriptor() ([]byte, []int) {
	return file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescGZIP(), []int{0}
}

func (x *CalcWeaponUpgradeReturnItemsRsp) GetItemParamList() []*ItemParam {
	if x != nil {
		return x.ItemParamList
	}
	return nil
}

func (x *CalcWeaponUpgradeReturnItemsRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *CalcWeaponUpgradeReturnItemsRsp) GetTargetWeaponGuid() uint64 {
	if x != nil {
		return x.TargetWeaponGuid
	}
	return 0
}

var File_CalcWeaponUpgradeReturnItemsRsp_proto protoreflect.FileDescriptor

var file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDesc = []byte{
	0x0a, 0x25, 0x43, 0x61, 0x6c, 0x63, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x73,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x1f, 0x43, 0x61, 0x6c,
	0x63, 0x57, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x57, 0x65,
	0x61, 0x70, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescOnce sync.Once
	file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescData = file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDesc
)

func file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescGZIP() []byte {
	file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescOnce.Do(func() {
		file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescData)
	})
	return file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDescData
}

var file_CalcWeaponUpgradeReturnItemsRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CalcWeaponUpgradeReturnItemsRsp_proto_goTypes = []interface{}{
	(*CalcWeaponUpgradeReturnItemsRsp)(nil), // 0: CalcWeaponUpgradeReturnItemsRsp
	(*ItemParam)(nil),                       // 1: ItemParam
}
var file_CalcWeaponUpgradeReturnItemsRsp_proto_depIdxs = []int32{
	1, // 0: CalcWeaponUpgradeReturnItemsRsp.item_param_list:type_name -> ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CalcWeaponUpgradeReturnItemsRsp_proto_init() }
func file_CalcWeaponUpgradeReturnItemsRsp_proto_init() {
	if File_CalcWeaponUpgradeReturnItemsRsp_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CalcWeaponUpgradeReturnItemsRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcWeaponUpgradeReturnItemsRsp); i {
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
			RawDescriptor: file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CalcWeaponUpgradeReturnItemsRsp_proto_goTypes,
		DependencyIndexes: file_CalcWeaponUpgradeReturnItemsRsp_proto_depIdxs,
		MessageInfos:      file_CalcWeaponUpgradeReturnItemsRsp_proto_msgTypes,
	}.Build()
	File_CalcWeaponUpgradeReturnItemsRsp_proto = out.File
	file_CalcWeaponUpgradeReturnItemsRsp_proto_rawDesc = nil
	file_CalcWeaponUpgradeReturnItemsRsp_proto_goTypes = nil
	file_CalcWeaponUpgradeReturnItemsRsp_proto_depIdxs = nil
}
