// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GetFriendShowNameCardInfoRsp.proto

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

// CmdId: 4029
// EnetChannelId: 0
// EnetIsReliable: true
type GetFriendShowNameCardInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode            int32    `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Uid                uint32   `protobuf:"varint,7,opt,name=uid,proto3" json:"uid,omitempty"`
	ShowNameCardIdList []uint32 `protobuf:"varint,10,rep,packed,name=show_name_card_id_list,json=showNameCardIdList,proto3" json:"show_name_card_id_list,omitempty"`
}

func (x *GetFriendShowNameCardInfoRsp) Reset() {
	*x = GetFriendShowNameCardInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetFriendShowNameCardInfoRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendShowNameCardInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendShowNameCardInfoRsp) ProtoMessage() {}

func (x *GetFriendShowNameCardInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetFriendShowNameCardInfoRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendShowNameCardInfoRsp.ProtoReflect.Descriptor instead.
func (*GetFriendShowNameCardInfoRsp) Descriptor() ([]byte, []int) {
	return file_GetFriendShowNameCardInfoRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetFriendShowNameCardInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetFriendShowNameCardInfoRsp) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetFriendShowNameCardInfoRsp) GetShowNameCardIdList() []uint32 {
	if x != nil {
		return x.ShowNameCardIdList
	}
	return nil
}

var File_GetFriendShowNameCardInfoRsp_proto protoreflect.FileDescriptor

var file_GetFriendShowNameCardInfoRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x77, 0x4e,
	0x61, 0x6d, 0x65, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x53, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x12, 0x32, 0x0a, 0x16, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x12, 0x73, 0x68, 0x6f, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x61, 0x72, 0x64, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetFriendShowNameCardInfoRsp_proto_rawDescOnce sync.Once
	file_GetFriendShowNameCardInfoRsp_proto_rawDescData = file_GetFriendShowNameCardInfoRsp_proto_rawDesc
)

func file_GetFriendShowNameCardInfoRsp_proto_rawDescGZIP() []byte {
	file_GetFriendShowNameCardInfoRsp_proto_rawDescOnce.Do(func() {
		file_GetFriendShowNameCardInfoRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetFriendShowNameCardInfoRsp_proto_rawDescData)
	})
	return file_GetFriendShowNameCardInfoRsp_proto_rawDescData
}

var file_GetFriendShowNameCardInfoRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetFriendShowNameCardInfoRsp_proto_goTypes = []interface{}{
	(*GetFriendShowNameCardInfoRsp)(nil), // 0: GetFriendShowNameCardInfoRsp
}
var file_GetFriendShowNameCardInfoRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetFriendShowNameCardInfoRsp_proto_init() }
func file_GetFriendShowNameCardInfoRsp_proto_init() {
	if File_GetFriendShowNameCardInfoRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetFriendShowNameCardInfoRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendShowNameCardInfoRsp); i {
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
			RawDescriptor: file_GetFriendShowNameCardInfoRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetFriendShowNameCardInfoRsp_proto_goTypes,
		DependencyIndexes: file_GetFriendShowNameCardInfoRsp_proto_depIdxs,
		MessageInfos:      file_GetFriendShowNameCardInfoRsp_proto_msgTypes,
	}.Build()
	File_GetFriendShowNameCardInfoRsp_proto = out.File
	file_GetFriendShowNameCardInfoRsp_proto_rawDesc = nil
	file_GetFriendShowNameCardInfoRsp_proto_goTypes = nil
	file_GetFriendShowNameCardInfoRsp_proto_depIdxs = nil
}
