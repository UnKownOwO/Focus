// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GetPlayerTokenReq.proto

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

// CmdId: 172
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type GetPlayerTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountToken        string `protobuf:"bytes,10,opt,name=account_token,json=accountToken,proto3" json:"account_token,omitempty"`
	AccountUid          string `protobuf:"bytes,11,opt,name=account_uid,json=accountUid,proto3" json:"account_uid,omitempty"`
	PsnRegion           string `protobuf:"bytes,4,opt,name=psn_region,json=psnRegion,proto3" json:"psn_region,omitempty"`
	OnlineId            string `protobuf:"bytes,7,opt,name=online_id,json=onlineId,proto3" json:"online_id,omitempty"`
	ChannelId           uint32 `protobuf:"varint,15,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	AccountExt          string `protobuf:"bytes,9,opt,name=account_ext,json=accountExt,proto3" json:"account_ext,omitempty"`
	CountryCode         string `protobuf:"bytes,5,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	ClientSeed          string `protobuf:"bytes,760,opt,name=client_seed,json=clientSeed,proto3" json:"client_seed,omitempty"`
	IsGuest             bool   `protobuf:"varint,6,opt,name=is_guest,json=isGuest,proto3" json:"is_guest,omitempty"`
	Birthday            string `protobuf:"bytes,1718,opt,name=birthday,proto3" json:"birthday,omitempty"`
	SubChannelId        uint32 `protobuf:"varint,8,opt,name=sub_channel_id,json=subChannelId,proto3" json:"sub_channel_id,omitempty"`
	PlatformType        uint32 `protobuf:"varint,12,opt,name=platform_type,json=platformType,proto3" json:"platform_type,omitempty"`
	ClientIpStr         string `protobuf:"bytes,3,opt,name=client_ip_str,json=clientIpStr,proto3" json:"client_ip_str,omitempty"`
	PsnId               string `protobuf:"bytes,13,opt,name=psn_id,json=psnId,proto3" json:"psn_id,omitempty"`
	AccountType         uint32 `protobuf:"varint,1,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
	Unk2700_NOJPEHIBDJH uint32 `protobuf:"varint,995,opt,name=Unk2700_NOJPEHIBDJH,json=Unk2700NOJPEHIBDJH,proto3" json:"Unk2700_NOJPEHIBDJH,omitempty"`
	CloudClientIp       uint32 `protobuf:"varint,14,opt,name=cloud_client_ip,json=cloudClientIp,proto3" json:"cloud_client_ip,omitempty"`
	KeyId               uint32 `protobuf:"varint,1787,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Uid                 uint32 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetPlayerTokenReq) Reset() {
	*x = GetPlayerTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPlayerTokenReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerTokenReq) ProtoMessage() {}

func (x *GetPlayerTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetPlayerTokenReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerTokenReq.ProtoReflect.Descriptor instead.
func (*GetPlayerTokenReq) Descriptor() ([]byte, []int) {
	return file_GetPlayerTokenReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerTokenReq) GetAccountToken() string {
	if x != nil {
		return x.AccountToken
	}
	return ""
}

func (x *GetPlayerTokenReq) GetAccountUid() string {
	if x != nil {
		return x.AccountUid
	}
	return ""
}

func (x *GetPlayerTokenReq) GetPsnRegion() string {
	if x != nil {
		return x.PsnRegion
	}
	return ""
}

func (x *GetPlayerTokenReq) GetOnlineId() string {
	if x != nil {
		return x.OnlineId
	}
	return ""
}

func (x *GetPlayerTokenReq) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetPlayerTokenReq) GetAccountExt() string {
	if x != nil {
		return x.AccountExt
	}
	return ""
}

func (x *GetPlayerTokenReq) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GetPlayerTokenReq) GetClientSeed() string {
	if x != nil {
		return x.ClientSeed
	}
	return ""
}

func (x *GetPlayerTokenReq) GetIsGuest() bool {
	if x != nil {
		return x.IsGuest
	}
	return false
}

func (x *GetPlayerTokenReq) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetPlayerTokenReq) GetSubChannelId() uint32 {
	if x != nil {
		return x.SubChannelId
	}
	return 0
}

func (x *GetPlayerTokenReq) GetPlatformType() uint32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *GetPlayerTokenReq) GetClientIpStr() string {
	if x != nil {
		return x.ClientIpStr
	}
	return ""
}

func (x *GetPlayerTokenReq) GetPsnId() string {
	if x != nil {
		return x.PsnId
	}
	return ""
}

func (x *GetPlayerTokenReq) GetAccountType() uint32 {
	if x != nil {
		return x.AccountType
	}
	return 0
}

func (x *GetPlayerTokenReq) GetUnk2700_NOJPEHIBDJH() uint32 {
	if x != nil {
		return x.Unk2700_NOJPEHIBDJH
	}
	return 0
}

func (x *GetPlayerTokenReq) GetCloudClientIp() uint32 {
	if x != nil {
		return x.CloudClientIp
	}
	return 0
}

func (x *GetPlayerTokenReq) GetKeyId() uint32 {
	if x != nil {
		return x.KeyId
	}
	return 0
}

func (x *GetPlayerTokenReq) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_GetPlayerTokenReq_proto protoreflect.FileDescriptor

var file_GetPlayerTokenReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x04, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x73, 0x6e, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x73, 0x6e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x78,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x65, 0x64, 0x18, 0xf8, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0xb6, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x53, 0x74, 0x72, 0x12, 0x15, 0x0a,
	0x06, 0x70, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x73, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x4e, 0x4f, 0x4a, 0x50, 0x45, 0x48, 0x49, 0x42, 0x44, 0x4a, 0x48, 0x18, 0xe3,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4e, 0x4f,
	0x4a, 0x50, 0x45, 0x48, 0x49, 0x42, 0x44, 0x4a, 0x48, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0xfb, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetPlayerTokenReq_proto_rawDescOnce sync.Once
	file_GetPlayerTokenReq_proto_rawDescData = file_GetPlayerTokenReq_proto_rawDesc
)

func file_GetPlayerTokenReq_proto_rawDescGZIP() []byte {
	file_GetPlayerTokenReq_proto_rawDescOnce.Do(func() {
		file_GetPlayerTokenReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPlayerTokenReq_proto_rawDescData)
	})
	return file_GetPlayerTokenReq_proto_rawDescData
}

var file_GetPlayerTokenReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPlayerTokenReq_proto_goTypes = []interface{}{
	(*GetPlayerTokenReq)(nil), // 0: GetPlayerTokenReq
}
var file_GetPlayerTokenReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetPlayerTokenReq_proto_init() }
func file_GetPlayerTokenReq_proto_init() {
	if File_GetPlayerTokenReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetPlayerTokenReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerTokenReq); i {
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
			RawDescriptor: file_GetPlayerTokenReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPlayerTokenReq_proto_goTypes,
		DependencyIndexes: file_GetPlayerTokenReq_proto_depIdxs,
		MessageInfos:      file_GetPlayerTokenReq_proto_msgTypes,
	}.Build()
	File_GetPlayerTokenReq_proto = out.File
	file_GetPlayerTokenReq_proto_rawDesc = nil
	file_GetPlayerTokenReq_proto_goTypes = nil
	file_GetPlayerTokenReq_proto_depIdxs = nil
}
