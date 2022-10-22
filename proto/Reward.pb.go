// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Reward.proto

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

type Reward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RewardId uint32       `protobuf:"varint,1,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	ItemList []*ItemParam `protobuf:"bytes,2,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *Reward) Reset() {
	*x = Reward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reward) ProtoMessage() {}

func (x *Reward) ProtoReflect() protoreflect.Message {
	mi := &file_Reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reward.ProtoReflect.Descriptor instead.
func (*Reward) Descriptor() ([]byte, []int) {
	return file_Reward_proto_rawDescGZIP(), []int{0}
}

func (x *Reward) GetRewardId() uint32 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *Reward) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_Reward_proto protoreflect.FileDescriptor

var file_Reward_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4e, 0x0a, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Reward_proto_rawDescOnce sync.Once
	file_Reward_proto_rawDescData = file_Reward_proto_rawDesc
)

func file_Reward_proto_rawDescGZIP() []byte {
	file_Reward_proto_rawDescOnce.Do(func() {
		file_Reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_Reward_proto_rawDescData)
	})
	return file_Reward_proto_rawDescData
}

var file_Reward_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Reward_proto_goTypes = []interface{}{
	(*Reward)(nil),    // 0: Reward
	(*ItemParam)(nil), // 1: ItemParam
}
var file_Reward_proto_depIdxs = []int32{
	1, // 0: Reward.item_list:type_name -> ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Reward_proto_init() }
func file_Reward_proto_init() {
	if File_Reward_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reward); i {
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
			RawDescriptor: file_Reward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Reward_proto_goTypes,
		DependencyIndexes: file_Reward_proto_depIdxs,
		MessageInfos:      file_Reward_proto_msgTypes,
	}.Build()
	File_Reward_proto = out.File
	file_Reward_proto_rawDesc = nil
	file_Reward_proto_goTypes = nil
	file_Reward_proto_depIdxs = nil
}
