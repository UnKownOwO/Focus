// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AsterActivityDetailInfo.proto

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

type AsterActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AsterLittle          *AsterLittleDetailInfo   `protobuf:"bytes,7,opt,name=aster_little,json=asterLittle,proto3" json:"aster_little,omitempty"`
	AsterCredit          uint32                   `protobuf:"varint,14,opt,name=aster_credit,json=asterCredit,proto3" json:"aster_credit,omitempty"`
	AsterLarge           *AsterLargeDetailInfo    `protobuf:"bytes,9,opt,name=aster_large,json=asterLarge,proto3" json:"aster_large,omitempty"`
	IsSpecialRewardTaken bool                     `protobuf:"varint,1,opt,name=is_special_reward_taken,json=isSpecialRewardTaken,proto3" json:"is_special_reward_taken,omitempty"`
	IsContentClosed      bool                     `protobuf:"varint,13,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	ContentCloseTime     uint32                   `protobuf:"varint,8,opt,name=content_close_time,json=contentCloseTime,proto3" json:"content_close_time,omitempty"`
	AsterToken           uint32                   `protobuf:"varint,5,opt,name=aster_token,json=asterToken,proto3" json:"aster_token,omitempty"`
	AsterMid             *AsterMidDetailInfo      `protobuf:"bytes,6,opt,name=aster_mid,json=asterMid,proto3" json:"aster_mid,omitempty"`
	AsterProgress        *AsterProgressDetailInfo `protobuf:"bytes,2,opt,name=aster_progress,json=asterProgress,proto3" json:"aster_progress,omitempty"`
}

func (x *AsterActivityDetailInfo) Reset() {
	*x = AsterActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AsterActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsterActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsterActivityDetailInfo) ProtoMessage() {}

func (x *AsterActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_AsterActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsterActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*AsterActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_AsterActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *AsterActivityDetailInfo) GetAsterLittle() *AsterLittleDetailInfo {
	if x != nil {
		return x.AsterLittle
	}
	return nil
}

func (x *AsterActivityDetailInfo) GetAsterCredit() uint32 {
	if x != nil {
		return x.AsterCredit
	}
	return 0
}

func (x *AsterActivityDetailInfo) GetAsterLarge() *AsterLargeDetailInfo {
	if x != nil {
		return x.AsterLarge
	}
	return nil
}

func (x *AsterActivityDetailInfo) GetIsSpecialRewardTaken() bool {
	if x != nil {
		return x.IsSpecialRewardTaken
	}
	return false
}

func (x *AsterActivityDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *AsterActivityDetailInfo) GetContentCloseTime() uint32 {
	if x != nil {
		return x.ContentCloseTime
	}
	return 0
}

func (x *AsterActivityDetailInfo) GetAsterToken() uint32 {
	if x != nil {
		return x.AsterToken
	}
	return 0
}

func (x *AsterActivityDetailInfo) GetAsterMid() *AsterMidDetailInfo {
	if x != nil {
		return x.AsterMid
	}
	return nil
}

func (x *AsterActivityDetailInfo) GetAsterProgress() *AsterProgressDetailInfo {
	if x != nil {
		return x.AsterProgress
	}
	return nil
}

var File_AsterActivityDetailInfo_proto protoreflect.FileDescriptor

var file_AsterActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x41, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x41, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x41, 0x73, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x73, 0x74, 0x65, 0x72, 0x4d,
	0x69, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x41, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd4, 0x03, 0x0a, 0x17, 0x41, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a,
	0x0c, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x41, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x74, 0x74, 0x6c,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x41, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x61,
	0x72, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x73, 0x74, 0x65, 0x72,
	0x4d, 0x69, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x4d, 0x69, 0x64, 0x12, 0x3f, 0x0a, 0x0e, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x41, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AsterActivityDetailInfo_proto_rawDescOnce sync.Once
	file_AsterActivityDetailInfo_proto_rawDescData = file_AsterActivityDetailInfo_proto_rawDesc
)

func file_AsterActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_AsterActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_AsterActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_AsterActivityDetailInfo_proto_rawDescData)
	})
	return file_AsterActivityDetailInfo_proto_rawDescData
}

var file_AsterActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AsterActivityDetailInfo_proto_goTypes = []interface{}{
	(*AsterActivityDetailInfo)(nil), // 0: AsterActivityDetailInfo
	(*AsterLittleDetailInfo)(nil),   // 1: AsterLittleDetailInfo
	(*AsterLargeDetailInfo)(nil),    // 2: AsterLargeDetailInfo
	(*AsterMidDetailInfo)(nil),      // 3: AsterMidDetailInfo
	(*AsterProgressDetailInfo)(nil), // 4: AsterProgressDetailInfo
}
var file_AsterActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: AsterActivityDetailInfo.aster_little:type_name -> AsterLittleDetailInfo
	2, // 1: AsterActivityDetailInfo.aster_large:type_name -> AsterLargeDetailInfo
	3, // 2: AsterActivityDetailInfo.aster_mid:type_name -> AsterMidDetailInfo
	4, // 3: AsterActivityDetailInfo.aster_progress:type_name -> AsterProgressDetailInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_AsterActivityDetailInfo_proto_init() }
func file_AsterActivityDetailInfo_proto_init() {
	if File_AsterActivityDetailInfo_proto != nil {
		return
	}
	file_AsterLargeDetailInfo_proto_init()
	file_AsterLittleDetailInfo_proto_init()
	file_AsterMidDetailInfo_proto_init()
	file_AsterProgressDetailInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AsterActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsterActivityDetailInfo); i {
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
			RawDescriptor: file_AsterActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AsterActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_AsterActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_AsterActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_AsterActivityDetailInfo_proto = out.File
	file_AsterActivityDetailInfo_proto_rawDesc = nil
	file_AsterActivityDetailInfo_proto_goTypes = nil
	file_AsterActivityDetailInfo_proto_depIdxs = nil
}
