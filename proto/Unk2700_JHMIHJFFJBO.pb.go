// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_JHMIHJFFJBO.proto

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

// CmdId: 8862
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_JHMIHJFFJBO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_DMJOJPGLFHE []*Unk2700_IMGLPJNBHCH `protobuf:"bytes,15,rep,name=Unk2700_DMJOJPGLFHE,json=Unk2700DMJOJPGLFHE,proto3" json:"Unk2700_DMJOJPGLFHE,omitempty"`
	Unk2700_AEHOPMMMHAP uint32                 `protobuf:"varint,13,opt,name=Unk2700_AEHOPMMMHAP,json=Unk2700AEHOPMMMHAP,proto3" json:"Unk2700_AEHOPMMMHAP,omitempty"`
	Unk2700_HMIBIIPHBAN uint32                 `protobuf:"varint,2,opt,name=Unk2700_HMIBIIPHBAN,json=Unk2700HMIBIIPHBAN,proto3" json:"Unk2700_HMIBIIPHBAN,omitempty"`
	Unk2700_FLMLLJIHOAI []*Unk2700_FEAENJPINFJ `protobuf:"bytes,8,rep,name=Unk2700_FLMLLJIHOAI,json=Unk2700FLMLLJIHOAI,proto3" json:"Unk2700_FLMLLJIHOAI,omitempty"`
	Unk2700_LOIMAGFKMOJ uint32                 `protobuf:"varint,6,opt,name=Unk2700_LOIMAGFKMOJ,json=Unk2700LOIMAGFKMOJ,proto3" json:"Unk2700_LOIMAGFKMOJ,omitempty"`
	StageId             uint32                 `protobuf:"varint,12,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	ChallengeId         uint32                 `protobuf:"varint,11,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	Unk2700_FGIIBJADECI uint32                 `protobuf:"varint,14,opt,name=Unk2700_FGIIBJADECI,json=Unk2700FGIIBJADECI,proto3" json:"Unk2700_FGIIBJADECI,omitempty"`
	Retcode             int32                  `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *Unk2700_JHMIHJFFJBO) Reset() {
	*x = Unk2700_JHMIHJFFJBO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_JHMIHJFFJBO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_JHMIHJFFJBO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_JHMIHJFFJBO) ProtoMessage() {}

func (x *Unk2700_JHMIHJFFJBO) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_JHMIHJFFJBO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_JHMIHJFFJBO.ProtoReflect.Descriptor instead.
func (*Unk2700_JHMIHJFFJBO) Descriptor() ([]byte, []int) {
	return file_Unk2700_JHMIHJFFJBO_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_JHMIHJFFJBO) GetUnk2700_DMJOJPGLFHE() []*Unk2700_IMGLPJNBHCH {
	if x != nil {
		return x.Unk2700_DMJOJPGLFHE
	}
	return nil
}

func (x *Unk2700_JHMIHJFFJBO) GetUnk2700_AEHOPMMMHAP() uint32 {
	if x != nil {
		return x.Unk2700_AEHOPMMMHAP
	}
	return 0
}

func (x *Unk2700_JHMIHJFFJBO) GetUnk2700_HMIBIIPHBAN() uint32 {
	if x != nil {
		return x.Unk2700_HMIBIIPHBAN
	}
	return 0
}

func (x *Unk2700_JHMIHJFFJBO) GetUnk2700_FLMLLJIHOAI() []*Unk2700_FEAENJPINFJ {
	if x != nil {
		return x.Unk2700_FLMLLJIHOAI
	}
	return nil
}

func (x *Unk2700_JHMIHJFFJBO) GetUnk2700_LOIMAGFKMOJ() uint32 {
	if x != nil {
		return x.Unk2700_LOIMAGFKMOJ
	}
	return 0
}

func (x *Unk2700_JHMIHJFFJBO) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *Unk2700_JHMIHJFFJBO) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *Unk2700_JHMIHJFFJBO) GetUnk2700_FGIIBJADECI() uint32 {
	if x != nil {
		return x.Unk2700_FGIIBJADECI
	}
	return 0
}

func (x *Unk2700_JHMIHJFFJBO) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_Unk2700_JHMIHJFFJBO_proto protoreflect.FileDescriptor

var file_Unk2700_JHMIHJFFJBO_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x48, 0x4d, 0x49, 0x48, 0x4a,
	0x46, 0x46, 0x4a, 0x42, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x45, 0x41, 0x45, 0x4e, 0x4a, 0x50, 0x49, 0x4e, 0x46, 0x4a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x49, 0x4d, 0x47, 0x4c, 0x50, 0x4a, 0x4e, 0x42, 0x48, 0x43, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbf, 0x03, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x48,
	0x4d, 0x49, 0x48, 0x4a, 0x46, 0x46, 0x4a, 0x42, 0x4f, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x4d, 0x4a, 0x4f, 0x4a, 0x50, 0x47, 0x4c, 0x46, 0x48, 0x45,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x49, 0x4d, 0x47, 0x4c, 0x50, 0x4a, 0x4e, 0x42, 0x48, 0x43, 0x48, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x4d, 0x4a, 0x4f, 0x4a, 0x50, 0x47, 0x4c, 0x46, 0x48, 0x45,
	0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x45, 0x48, 0x4f,
	0x50, 0x4d, 0x4d, 0x4d, 0x48, 0x41, 0x50, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x41, 0x45, 0x48, 0x4f, 0x50, 0x4d, 0x4d, 0x4d, 0x48, 0x41,
	0x50, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x4d, 0x49,
	0x42, 0x49, 0x49, 0x50, 0x48, 0x42, 0x41, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x48, 0x4d, 0x49, 0x42, 0x49, 0x49, 0x50, 0x48, 0x42,
	0x41, 0x4e, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x4c,
	0x4d, 0x4c, 0x4c, 0x4a, 0x49, 0x48, 0x4f, 0x41, 0x49, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x45, 0x41, 0x45, 0x4e, 0x4a,
	0x50, 0x49, 0x4e, 0x46, 0x4a, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x46, 0x4c,
	0x4d, 0x4c, 0x4c, 0x4a, 0x49, 0x48, 0x4f, 0x41, 0x49, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x4c, 0x4f, 0x49, 0x4d, 0x41, 0x47, 0x46, 0x4b, 0x4d, 0x4f, 0x4a,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4c,
	0x4f, 0x49, 0x4d, 0x41, 0x47, 0x46, 0x4b, 0x4d, 0x4f, 0x4a, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x46, 0x47, 0x49, 0x49, 0x42, 0x4a, 0x41, 0x44, 0x45, 0x43, 0x49, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x46, 0x47,
	0x49, 0x49, 0x42, 0x4a, 0x41, 0x44, 0x45, 0x43, 0x49, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_JHMIHJFFJBO_proto_rawDescOnce sync.Once
	file_Unk2700_JHMIHJFFJBO_proto_rawDescData = file_Unk2700_JHMIHJFFJBO_proto_rawDesc
)

func file_Unk2700_JHMIHJFFJBO_proto_rawDescGZIP() []byte {
	file_Unk2700_JHMIHJFFJBO_proto_rawDescOnce.Do(func() {
		file_Unk2700_JHMIHJFFJBO_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_JHMIHJFFJBO_proto_rawDescData)
	})
	return file_Unk2700_JHMIHJFFJBO_proto_rawDescData
}

var file_Unk2700_JHMIHJFFJBO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_JHMIHJFFJBO_proto_goTypes = []interface{}{
	(*Unk2700_JHMIHJFFJBO)(nil), // 0: Unk2700_JHMIHJFFJBO
	(*Unk2700_IMGLPJNBHCH)(nil), // 1: Unk2700_IMGLPJNBHCH
	(*Unk2700_FEAENJPINFJ)(nil), // 2: Unk2700_FEAENJPINFJ
}
var file_Unk2700_JHMIHJFFJBO_proto_depIdxs = []int32{
	1, // 0: Unk2700_JHMIHJFFJBO.Unk2700_DMJOJPGLFHE:type_name -> Unk2700_IMGLPJNBHCH
	2, // 1: Unk2700_JHMIHJFFJBO.Unk2700_FLMLLJIHOAI:type_name -> Unk2700_FEAENJPINFJ
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk2700_JHMIHJFFJBO_proto_init() }
func file_Unk2700_JHMIHJFFJBO_proto_init() {
	if File_Unk2700_JHMIHJFFJBO_proto != nil {
		return
	}
	file_Unk2700_FEAENJPINFJ_proto_init()
	file_Unk2700_IMGLPJNBHCH_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_JHMIHJFFJBO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_JHMIHJFFJBO); i {
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
			RawDescriptor: file_Unk2700_JHMIHJFFJBO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_JHMIHJFFJBO_proto_goTypes,
		DependencyIndexes: file_Unk2700_JHMIHJFFJBO_proto_depIdxs,
		MessageInfos:      file_Unk2700_JHMIHJFFJBO_proto_msgTypes,
	}.Build()
	File_Unk2700_JHMIHJFFJBO_proto = out.File
	file_Unk2700_JHMIHJFFJBO_proto_rawDesc = nil
	file_Unk2700_JHMIHJFFJBO_proto_goTypes = nil
	file_Unk2700_JHMIHJFFJBO_proto_depIdxs = nil
}
