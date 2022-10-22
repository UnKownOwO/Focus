// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: PlayerStoreNotify.proto

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

// CmdId: 672
// EnetChannelId: 0
// EnetIsReliable: true
type PlayerStoreNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList    []*Item   `protobuf:"bytes,15,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	WeightLimit uint32    `protobuf:"varint,8,opt,name=weight_limit,json=weightLimit,proto3" json:"weight_limit,omitempty"`
	StoreType   StoreType `protobuf:"varint,2,opt,name=store_type,json=storeType,proto3,enum=StoreType" json:"store_type,omitempty"`
}

func (x *PlayerStoreNotify) Reset() {
	*x = PlayerStoreNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerStoreNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerStoreNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStoreNotify) ProtoMessage() {}

func (x *PlayerStoreNotify) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerStoreNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStoreNotify.ProtoReflect.Descriptor instead.
func (*PlayerStoreNotify) Descriptor() ([]byte, []int) {
	return file_PlayerStoreNotify_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerStoreNotify) GetItemList() []*Item {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *PlayerStoreNotify) GetWeightLimit() uint32 {
	if x != nil {
		return x.WeightLimit
	}
	return 0
}

func (x *PlayerStoreNotify) GetStoreType() StoreType {
	if x != nil {
		return x.StoreType
	}
	return StoreType_STORE_TYPE_NONE
}

var File_PlayerStoreNotify_proto protoreflect.FileDescriptor

var file_PlayerStoreNotify_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x22, 0x0a, 0x09,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x29, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_PlayerStoreNotify_proto_rawDescOnce sync.Once
	file_PlayerStoreNotify_proto_rawDescData = file_PlayerStoreNotify_proto_rawDesc
)

func file_PlayerStoreNotify_proto_rawDescGZIP() []byte {
	file_PlayerStoreNotify_proto_rawDescOnce.Do(func() {
		file_PlayerStoreNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerStoreNotify_proto_rawDescData)
	})
	return file_PlayerStoreNotify_proto_rawDescData
}

var file_PlayerStoreNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerStoreNotify_proto_goTypes = []interface{}{
	(*PlayerStoreNotify)(nil), // 0: PlayerStoreNotify
	(*Item)(nil),              // 1: Item
	(StoreType)(0),            // 2: StoreType
}
var file_PlayerStoreNotify_proto_depIdxs = []int32{
	1, // 0: PlayerStoreNotify.item_list:type_name -> Item
	2, // 1: PlayerStoreNotify.store_type:type_name -> StoreType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerStoreNotify_proto_init() }
func file_PlayerStoreNotify_proto_init() {
	if File_PlayerStoreNotify_proto != nil {
		return
	}
	file_Item_proto_init()
	file_StoreType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerStoreNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerStoreNotify); i {
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
			RawDescriptor: file_PlayerStoreNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerStoreNotify_proto_goTypes,
		DependencyIndexes: file_PlayerStoreNotify_proto_depIdxs,
		MessageInfos:      file_PlayerStoreNotify_proto_msgTypes,
	}.Build()
	File_PlayerStoreNotify_proto = out.File
	file_PlayerStoreNotify_proto_rawDesc = nil
	file_PlayerStoreNotify_proto_goTypes = nil
	file_PlayerStoreNotify_proto_depIdxs = nil
}
