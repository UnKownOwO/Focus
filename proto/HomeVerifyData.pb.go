// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: HomeVerifyData.proto

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

type HomeVerifyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_OAKBDKKBFHP string                        `protobuf:"bytes,7,opt,name=Unk2700_OAKBDKKBFHP,json=Unk2700OAKBDKKBFHP,proto3" json:"Unk2700_OAKBDKKBFHP,omitempty"`
	Timestamp           uint32                        `protobuf:"fixed32,15,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Uid                 uint32                        `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	Unk2700_CDELDBLKLDO *HomeSceneArrangementMuipData `protobuf:"bytes,9,opt,name=Unk2700_CDELDBLKLDO,json=Unk2700CDELDBLKLDO,proto3" json:"Unk2700_CDELDBLKLDO,omitempty"`
	Region              string                        `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Token               string                        `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	HomeInfo            *HomeVerifySceneData          `protobuf:"bytes,6,opt,name=home_info,json=homeInfo,proto3" json:"home_info,omitempty"`
	Lang                LanguageType                  `protobuf:"varint,8,opt,name=lang,proto3,enum=LanguageType" json:"lang,omitempty"`
}

func (x *HomeVerifyData) Reset() {
	*x = HomeVerifyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeVerifyData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeVerifyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeVerifyData) ProtoMessage() {}

func (x *HomeVerifyData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeVerifyData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeVerifyData.ProtoReflect.Descriptor instead.
func (*HomeVerifyData) Descriptor() ([]byte, []int) {
	return file_HomeVerifyData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeVerifyData) GetUnk2700_OAKBDKKBFHP() string {
	if x != nil {
		return x.Unk2700_OAKBDKKBFHP
	}
	return ""
}

func (x *HomeVerifyData) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HomeVerifyData) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *HomeVerifyData) GetUnk2700_CDELDBLKLDO() *HomeSceneArrangementMuipData {
	if x != nil {
		return x.Unk2700_CDELDBLKLDO
	}
	return nil
}

func (x *HomeVerifyData) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *HomeVerifyData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *HomeVerifyData) GetHomeInfo() *HomeVerifySceneData {
	if x != nil {
		return x.HomeInfo
	}
	return nil
}

func (x *HomeVerifyData) GetLang() LanguageType {
	if x != nil {
		return x.Lang
	}
	return LanguageType_LANGUAGE_TYPE_NONE
}

var File_HomeVerifyData_proto protoreflect.FileDescriptor

var file_HomeVerifyData_proto_rawDesc = []byte{
	0x0a, 0x14, 0x48, 0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x69, 0x70,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x48, 0x6f, 0x6d, 0x65,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x0e, 0x48, 0x6f,
	0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f, 0x41, 0x4b, 0x42, 0x44, 0x4b, 0x4b, 0x42,
	0x46, 0x48, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x4f, 0x41, 0x4b, 0x42, 0x44, 0x4b, 0x4b, 0x42, 0x46, 0x48, 0x50, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x4e, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x44, 0x45, 0x4c, 0x44, 0x42, 0x4c,
	0x4b, 0x4c, 0x44, 0x4f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x48, 0x6f, 0x6d,
	0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x75, 0x69, 0x70, 0x44, 0x61, 0x74, 0x61, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x43, 0x44, 0x45, 0x4c, 0x44, 0x42, 0x4c, 0x4b, 0x4c, 0x44, 0x4f, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x09, 0x68,
	0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x68, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21,
	0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeVerifyData_proto_rawDescOnce sync.Once
	file_HomeVerifyData_proto_rawDescData = file_HomeVerifyData_proto_rawDesc
)

func file_HomeVerifyData_proto_rawDescGZIP() []byte {
	file_HomeVerifyData_proto_rawDescOnce.Do(func() {
		file_HomeVerifyData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeVerifyData_proto_rawDescData)
	})
	return file_HomeVerifyData_proto_rawDescData
}

var file_HomeVerifyData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeVerifyData_proto_goTypes = []interface{}{
	(*HomeVerifyData)(nil),               // 0: HomeVerifyData
	(*HomeSceneArrangementMuipData)(nil), // 1: HomeSceneArrangementMuipData
	(*HomeVerifySceneData)(nil),          // 2: HomeVerifySceneData
	(LanguageType)(0),                    // 3: LanguageType
}
var file_HomeVerifyData_proto_depIdxs = []int32{
	1, // 0: HomeVerifyData.Unk2700_CDELDBLKLDO:type_name -> HomeSceneArrangementMuipData
	2, // 1: HomeVerifyData.home_info:type_name -> HomeVerifySceneData
	3, // 2: HomeVerifyData.lang:type_name -> LanguageType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_HomeVerifyData_proto_init() }
func file_HomeVerifyData_proto_init() {
	if File_HomeVerifyData_proto != nil {
		return
	}
	file_HomeSceneArrangementMuipData_proto_init()
	file_HomeVerifySceneData_proto_init()
	file_LanguageType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeVerifyData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeVerifyData); i {
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
			RawDescriptor: file_HomeVerifyData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeVerifyData_proto_goTypes,
		DependencyIndexes: file_HomeVerifyData_proto_depIdxs,
		MessageInfos:      file_HomeVerifyData_proto_msgTypes,
	}.Build()
	File_HomeVerifyData_proto = out.File
	file_HomeVerifyData_proto_rawDesc = nil
	file_HomeVerifyData_proto_goTypes = nil
	file_HomeVerifyData_proto_depIdxs = nil
}
