// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: LunaRiteHintPointRsp.proto

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

// CmdId: 8765
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type LunaRiteHintPointRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HintStatus LunaRiteHintStatusType `protobuf:"varint,4,opt,name=hint_status,json=hintStatus,proto3,enum=LunaRiteHintStatusType" json:"hint_status,omitempty"`
	AreaId     uint32                 `protobuf:"varint,5,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Retcode    int32                  `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	HintPoint  []*LunaRiteHintPoint   `protobuf:"bytes,9,rep,name=hint_point,json=hintPoint,proto3" json:"hint_point,omitempty"`
}

func (x *LunaRiteHintPointRsp) Reset() {
	*x = LunaRiteHintPointRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LunaRiteHintPointRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LunaRiteHintPointRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LunaRiteHintPointRsp) ProtoMessage() {}

func (x *LunaRiteHintPointRsp) ProtoReflect() protoreflect.Message {
	mi := &file_LunaRiteHintPointRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LunaRiteHintPointRsp.ProtoReflect.Descriptor instead.
func (*LunaRiteHintPointRsp) Descriptor() ([]byte, []int) {
	return file_LunaRiteHintPointRsp_proto_rawDescGZIP(), []int{0}
}

func (x *LunaRiteHintPointRsp) GetHintStatus() LunaRiteHintStatusType {
	if x != nil {
		return x.HintStatus
	}
	return LunaRiteHintStatusType_LUNA_RITE_HINT_STATUS_TYPE_DEFAULT
}

func (x *LunaRiteHintPointRsp) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *LunaRiteHintPointRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *LunaRiteHintPointRsp) GetHintPoint() []*LunaRiteHintPoint {
	if x != nil {
		return x.HintPoint
	}
	return nil
}

var File_LunaRiteHintPointRsp_proto protoreflect.FileDescriptor

var file_LunaRiteHintPointRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x4c, 0x75, 0x6e, 0x61, 0x52, 0x69, 0x74, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x4c, 0x75,
	0x6e, 0x61, 0x52, 0x69, 0x74, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x4c, 0x75, 0x6e, 0x61, 0x52, 0x69, 0x74, 0x65, 0x48,
	0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x14, 0x4c, 0x75, 0x6e, 0x61, 0x52, 0x69, 0x74, 0x65,
	0x48, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0b,
	0x68, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x4c, 0x75, 0x6e, 0x61, 0x52, 0x69, 0x74, 0x65, 0x48, 0x69, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x68, 0x69, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x68, 0x69, 0x6e,
	0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x4c, 0x75, 0x6e, 0x61, 0x52, 0x69, 0x74, 0x65, 0x48, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x09, 0x68, 0x69, 0x6e, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LunaRiteHintPointRsp_proto_rawDescOnce sync.Once
	file_LunaRiteHintPointRsp_proto_rawDescData = file_LunaRiteHintPointRsp_proto_rawDesc
)

func file_LunaRiteHintPointRsp_proto_rawDescGZIP() []byte {
	file_LunaRiteHintPointRsp_proto_rawDescOnce.Do(func() {
		file_LunaRiteHintPointRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_LunaRiteHintPointRsp_proto_rawDescData)
	})
	return file_LunaRiteHintPointRsp_proto_rawDescData
}

var file_LunaRiteHintPointRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LunaRiteHintPointRsp_proto_goTypes = []interface{}{
	(*LunaRiteHintPointRsp)(nil), // 0: LunaRiteHintPointRsp
	(LunaRiteHintStatusType)(0),  // 1: LunaRiteHintStatusType
	(*LunaRiteHintPoint)(nil),    // 2: LunaRiteHintPoint
}
var file_LunaRiteHintPointRsp_proto_depIdxs = []int32{
	1, // 0: LunaRiteHintPointRsp.hint_status:type_name -> LunaRiteHintStatusType
	2, // 1: LunaRiteHintPointRsp.hint_point:type_name -> LunaRiteHintPoint
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_LunaRiteHintPointRsp_proto_init() }
func file_LunaRiteHintPointRsp_proto_init() {
	if File_LunaRiteHintPointRsp_proto != nil {
		return
	}
	file_LunaRiteHintPoint_proto_init()
	file_LunaRiteHintStatusType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LunaRiteHintPointRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LunaRiteHintPointRsp); i {
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
			RawDescriptor: file_LunaRiteHintPointRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LunaRiteHintPointRsp_proto_goTypes,
		DependencyIndexes: file_LunaRiteHintPointRsp_proto_depIdxs,
		MessageInfos:      file_LunaRiteHintPointRsp_proto_msgTypes,
	}.Build()
	File_LunaRiteHintPointRsp_proto = out.File
	file_LunaRiteHintPointRsp_proto_rawDesc = nil
	file_LunaRiteHintPointRsp_proto_goTypes = nil
	file_LunaRiteHintPointRsp_proto_depIdxs = nil
}
