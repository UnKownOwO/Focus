// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ClearRoguelikeCurseNotify.proto

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

// CmdId: 8207
// EnetChannelId: 0
// EnetIsReliable: true
type ClearRoguelikeCurseNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClearCurseMap   map[uint32]uint32 `protobuf:"bytes,10,rep,name=clear_curse_map,json=clearCurseMap,proto3" json:"clear_curse_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	IsClearAll      bool              `protobuf:"varint,11,opt,name=is_clear_all,json=isClearAll,proto3" json:"is_clear_all,omitempty"`
	CardId          uint32            `protobuf:"varint,8,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	IsCurseAllClear bool              `protobuf:"varint,1,opt,name=is_curse_all_clear,json=isCurseAllClear,proto3" json:"is_curse_all_clear,omitempty"`
}

func (x *ClearRoguelikeCurseNotify) Reset() {
	*x = ClearRoguelikeCurseNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClearRoguelikeCurseNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearRoguelikeCurseNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearRoguelikeCurseNotify) ProtoMessage() {}

func (x *ClearRoguelikeCurseNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ClearRoguelikeCurseNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearRoguelikeCurseNotify.ProtoReflect.Descriptor instead.
func (*ClearRoguelikeCurseNotify) Descriptor() ([]byte, []int) {
	return file_ClearRoguelikeCurseNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ClearRoguelikeCurseNotify) GetClearCurseMap() map[uint32]uint32 {
	if x != nil {
		return x.ClearCurseMap
	}
	return nil
}

func (x *ClearRoguelikeCurseNotify) GetIsClearAll() bool {
	if x != nil {
		return x.IsClearAll
	}
	return false
}

func (x *ClearRoguelikeCurseNotify) GetCardId() uint32 {
	if x != nil {
		return x.CardId
	}
	return 0
}

func (x *ClearRoguelikeCurseNotify) GetIsCurseAllClear() bool {
	if x != nil {
		return x.IsCurseAllClear
	}
	return false
}

var File_ClearRoguelikeCurseNotify_proto protoreflect.FileDescriptor

var file_ClearRoguelikeCurseNotify_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65,
	0x43, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x19, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x6c, 0x69, 0x6b, 0x65, 0x43, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x55, 0x0a, 0x0f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d,
	0x61, 0x70, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x75, 0x72, 0x73, 0x65, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x75,
	0x72, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69,
	0x73, 0x43, 0x75, 0x72, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x1a, 0x40,
	0x0a, 0x12, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x43, 0x75, 0x72, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ClearRoguelikeCurseNotify_proto_rawDescOnce sync.Once
	file_ClearRoguelikeCurseNotify_proto_rawDescData = file_ClearRoguelikeCurseNotify_proto_rawDesc
)

func file_ClearRoguelikeCurseNotify_proto_rawDescGZIP() []byte {
	file_ClearRoguelikeCurseNotify_proto_rawDescOnce.Do(func() {
		file_ClearRoguelikeCurseNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClearRoguelikeCurseNotify_proto_rawDescData)
	})
	return file_ClearRoguelikeCurseNotify_proto_rawDescData
}

var file_ClearRoguelikeCurseNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ClearRoguelikeCurseNotify_proto_goTypes = []interface{}{
	(*ClearRoguelikeCurseNotify)(nil), // 0: ClearRoguelikeCurseNotify
	nil,                               // 1: ClearRoguelikeCurseNotify.ClearCurseMapEntry
}
var file_ClearRoguelikeCurseNotify_proto_depIdxs = []int32{
	1, // 0: ClearRoguelikeCurseNotify.clear_curse_map:type_name -> ClearRoguelikeCurseNotify.ClearCurseMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ClearRoguelikeCurseNotify_proto_init() }
func file_ClearRoguelikeCurseNotify_proto_init() {
	if File_ClearRoguelikeCurseNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClearRoguelikeCurseNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearRoguelikeCurseNotify); i {
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
			RawDescriptor: file_ClearRoguelikeCurseNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClearRoguelikeCurseNotify_proto_goTypes,
		DependencyIndexes: file_ClearRoguelikeCurseNotify_proto_depIdxs,
		MessageInfos:      file_ClearRoguelikeCurseNotify_proto_msgTypes,
	}.Build()
	File_ClearRoguelikeCurseNotify_proto = out.File
	file_ClearRoguelikeCurseNotify_proto_rawDesc = nil
	file_ClearRoguelikeCurseNotify_proto_goTypes = nil
	file_ClearRoguelikeCurseNotify_proto_depIdxs = nil
}
