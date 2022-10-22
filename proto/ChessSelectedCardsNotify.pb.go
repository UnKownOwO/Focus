// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ChessSelectedCardsNotify.proto

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

// CmdId: 5392
// EnetChannelId: 0
// EnetIsReliable: true
type ChessSelectedCardsNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelectedCardInfoList []*ChessCardInfo `protobuf:"bytes,4,rep,name=selected_card_info_list,json=selectedCardInfoList,proto3" json:"selected_card_info_list,omitempty"`
}

func (x *ChessSelectedCardsNotify) Reset() {
	*x = ChessSelectedCardsNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessSelectedCardsNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessSelectedCardsNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessSelectedCardsNotify) ProtoMessage() {}

func (x *ChessSelectedCardsNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChessSelectedCardsNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessSelectedCardsNotify.ProtoReflect.Descriptor instead.
func (*ChessSelectedCardsNotify) Descriptor() ([]byte, []int) {
	return file_ChessSelectedCardsNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChessSelectedCardsNotify) GetSelectedCardInfoList() []*ChessCardInfo {
	if x != nil {
		return x.SelectedCardInfoList
	}
	return nil
}

var File_ChessSelectedCardsNotify_proto protoreflect.FileDescriptor

var file_ChessSelectedCardsNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43,
	0x61, 0x72, 0x64, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x43, 0x68, 0x65, 0x73, 0x73, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x45, 0x0a, 0x17, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x14, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessSelectedCardsNotify_proto_rawDescOnce sync.Once
	file_ChessSelectedCardsNotify_proto_rawDescData = file_ChessSelectedCardsNotify_proto_rawDesc
)

func file_ChessSelectedCardsNotify_proto_rawDescGZIP() []byte {
	file_ChessSelectedCardsNotify_proto_rawDescOnce.Do(func() {
		file_ChessSelectedCardsNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessSelectedCardsNotify_proto_rawDescData)
	})
	return file_ChessSelectedCardsNotify_proto_rawDescData
}

var file_ChessSelectedCardsNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessSelectedCardsNotify_proto_goTypes = []interface{}{
	(*ChessSelectedCardsNotify)(nil), // 0: ChessSelectedCardsNotify
	(*ChessCardInfo)(nil),            // 1: ChessCardInfo
}
var file_ChessSelectedCardsNotify_proto_depIdxs = []int32{
	1, // 0: ChessSelectedCardsNotify.selected_card_info_list:type_name -> ChessCardInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChessSelectedCardsNotify_proto_init() }
func file_ChessSelectedCardsNotify_proto_init() {
	if File_ChessSelectedCardsNotify_proto != nil {
		return
	}
	file_ChessCardInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessSelectedCardsNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessSelectedCardsNotify); i {
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
			RawDescriptor: file_ChessSelectedCardsNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessSelectedCardsNotify_proto_goTypes,
		DependencyIndexes: file_ChessSelectedCardsNotify_proto_depIdxs,
		MessageInfos:      file_ChessSelectedCardsNotify_proto_msgTypes,
	}.Build()
	File_ChessSelectedCardsNotify_proto = out.File
	file_ChessSelectedCardsNotify_proto_rawDesc = nil
	file_ChessSelectedCardsNotify_proto_goTypes = nil
	file_ChessSelectedCardsNotify_proto_depIdxs = nil
}
