// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_JJAFAJIKDDK_ServerRsp.proto

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

// CmdId: 6307
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_JJAFAJIKDDK_ServerRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_CEPGMKAHHCD uint64              `protobuf:"varint,3,opt,name=Unk2700_CEPGMKAHHCD,json=Unk2700CEPGMKAHHCD,proto3" json:"Unk2700_CEPGMKAHHCD,omitempty"`
	Unk2700_KHBDAPGDOJA Unk2700_OPEBMJPOOBL `protobuf:"varint,11,opt,name=Unk2700_KHBDAPGDOJA,json=Unk2700KHBDAPGDOJA,proto3,enum=Unk2700_OPEBMJPOOBL" json:"Unk2700_KHBDAPGDOJA,omitempty"`
	Retcode             int32               `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	// Types that are assignable to Unk2700_ILHNBMNOMHO:
	//
	//	*Unk2700_JJAFAJIKDDK_ServerRsp_MusicBriefInfo
	Unk2700_ILHNBMNOMHO isUnk2700_JJAFAJIKDDK_ServerRsp_Unk2700_ILHNBMNOMHO `protobuf_oneof:"Unk2700_ILHNBMNOMHO"`
}

func (x *Unk2700_JJAFAJIKDDK_ServerRsp) Reset() {
	*x = Unk2700_JJAFAJIKDDK_ServerRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_JJAFAJIKDDK_ServerRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_JJAFAJIKDDK_ServerRsp) ProtoMessage() {}

func (x *Unk2700_JJAFAJIKDDK_ServerRsp) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_JJAFAJIKDDK_ServerRsp.ProtoReflect.Descriptor instead.
func (*Unk2700_JJAFAJIKDDK_ServerRsp) Descriptor() ([]byte, []int) {
	return file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_JJAFAJIKDDK_ServerRsp) GetUnk2700_CEPGMKAHHCD() uint64 {
	if x != nil {
		return x.Unk2700_CEPGMKAHHCD
	}
	return 0
}

func (x *Unk2700_JJAFAJIKDDK_ServerRsp) GetUnk2700_KHBDAPGDOJA() Unk2700_OPEBMJPOOBL {
	if x != nil {
		return x.Unk2700_KHBDAPGDOJA
	}
	return Unk2700_OPEBMJPOOBL_Unk2700_OPEBMJPOOBL_NONE
}

func (x *Unk2700_JJAFAJIKDDK_ServerRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (m *Unk2700_JJAFAJIKDDK_ServerRsp) GetUnk2700_ILHNBMNOMHO() isUnk2700_JJAFAJIKDDK_ServerRsp_Unk2700_ILHNBMNOMHO {
	if m != nil {
		return m.Unk2700_ILHNBMNOMHO
	}
	return nil
}

func (x *Unk2700_JJAFAJIKDDK_ServerRsp) GetMusicBriefInfo() *MusicBriefInfo {
	if x, ok := x.GetUnk2700_ILHNBMNOMHO().(*Unk2700_JJAFAJIKDDK_ServerRsp_MusicBriefInfo); ok {
		return x.MusicBriefInfo
	}
	return nil
}

type isUnk2700_JJAFAJIKDDK_ServerRsp_Unk2700_ILHNBMNOMHO interface {
	isUnk2700_JJAFAJIKDDK_ServerRsp_Unk2700_ILHNBMNOMHO()
}

type Unk2700_JJAFAJIKDDK_ServerRsp_MusicBriefInfo struct {
	MusicBriefInfo *MusicBriefInfo `protobuf:"bytes,2,opt,name=music_brief_info,json=musicBriefInfo,proto3,oneof"`
}

func (*Unk2700_JJAFAJIKDDK_ServerRsp_MusicBriefInfo) isUnk2700_JJAFAJIKDDK_ServerRsp_Unk2700_ILHNBMNOMHO() {
}

var File_Unk2700_JJAFAJIKDDK_ServerRsp_proto protoreflect.FileDescriptor

var file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x4a, 0x41, 0x46, 0x41, 0x4a,
	0x49, 0x4b, 0x44, 0x44, 0x4b, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x72, 0x69, 0x65,
	0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x50, 0x45, 0x42, 0x4d, 0x4a, 0x50, 0x4f, 0x4f, 0x42, 0x4c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x1d, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x4a, 0x4a, 0x41, 0x46, 0x41, 0x4a, 0x49, 0x4b, 0x44, 0x44, 0x4b, 0x5f, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x43, 0x45, 0x50, 0x47, 0x4d, 0x4b, 0x41, 0x48, 0x48, 0x43, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x43, 0x45,
	0x50, 0x47, 0x4d, 0x4b, 0x41, 0x48, 0x48, 0x43, 0x44, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x48, 0x42, 0x44, 0x41, 0x50, 0x47, 0x44, 0x4f, 0x4a, 0x41,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x4f, 0x50, 0x45, 0x42, 0x4d, 0x4a, 0x50, 0x4f, 0x4f, 0x42, 0x4c, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x4b, 0x48, 0x42, 0x44, 0x41, 0x50, 0x47, 0x44, 0x4f, 0x4a, 0x41,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x10, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x5f, 0x62, 0x72, 0x69, 0x65, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x72, 0x69, 0x65,
	0x66, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x72,
	0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x15, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x49, 0x4c, 0x48, 0x4e, 0x42, 0x4d, 0x4e, 0x4f, 0x4d, 0x48, 0x4f, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescOnce sync.Once
	file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescData = file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDesc
)

func file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescGZIP() []byte {
	file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescOnce.Do(func() {
		file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescData)
	})
	return file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDescData
}

var file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_goTypes = []interface{}{
	(*Unk2700_JJAFAJIKDDK_ServerRsp)(nil), // 0: Unk2700_JJAFAJIKDDK_ServerRsp
	(Unk2700_OPEBMJPOOBL)(0),              // 1: Unk2700_OPEBMJPOOBL
	(*MusicBriefInfo)(nil),                // 2: MusicBriefInfo
}
var file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_depIdxs = []int32{
	1, // 0: Unk2700_JJAFAJIKDDK_ServerRsp.Unk2700_KHBDAPGDOJA:type_name -> Unk2700_OPEBMJPOOBL
	2, // 1: Unk2700_JJAFAJIKDDK_ServerRsp.music_brief_info:type_name -> MusicBriefInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_init() }
func file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_init() {
	if File_Unk2700_JJAFAJIKDDK_ServerRsp_proto != nil {
		return
	}
	file_MusicBriefInfo_proto_init()
	file_Unk2700_OPEBMJPOOBL_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_JJAFAJIKDDK_ServerRsp); i {
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
	file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Unk2700_JJAFAJIKDDK_ServerRsp_MusicBriefInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_goTypes,
		DependencyIndexes: file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_depIdxs,
		MessageInfos:      file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_msgTypes,
	}.Build()
	File_Unk2700_JJAFAJIKDDK_ServerRsp_proto = out.File
	file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_rawDesc = nil
	file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_goTypes = nil
	file_Unk2700_JJAFAJIKDDK_ServerRsp_proto_depIdxs = nil
}