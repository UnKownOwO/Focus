// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2700_HHGMCHANCBJ_ServerNotify.proto

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

// CmdId: 6217
// EnetChannelId: 0
// EnetIsReliable: true
type Unk2700_HHGMCHANCBJ_ServerNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_LGBODABIKLL Unk2700_KBBDJNLFAKD  `protobuf:"varint,14,opt,name=Unk2700_LGBODABIKLL,json=Unk2700LGBODABIKLL,proto3,enum=Unk2700_KBBDJNLFAKD" json:"Unk2700_LGBODABIKLL,omitempty"`
	Unk2700_NBAIINBBBPK Unk2700_ADGLMHECKKJ  `protobuf:"varint,3,opt,name=Unk2700_NBAIINBBBPK,json=Unk2700NBAIINBBBPK,proto3,enum=Unk2700_ADGLMHECKKJ" json:"Unk2700_NBAIINBBBPK,omitempty"`
	Unk2700_EJHNBDLLLFO *Unk2700_NLFDMMFNMIO `protobuf:"bytes,10,opt,name=Unk2700_EJHNBDLLLFO,json=Unk2700EJHNBDLLLFO,proto3" json:"Unk2700_EJHNBDLLLFO,omitempty"`
	Unk2700_EIOPOPABBNC []uint32             `protobuf:"varint,12,rep,packed,name=Unk2700_EIOPOPABBNC,json=Unk2700EIOPOPABBNC,proto3" json:"Unk2700_EIOPOPABBNC,omitempty"`
}

func (x *Unk2700_HHGMCHANCBJ_ServerNotify) Reset() {
	*x = Unk2700_HHGMCHANCBJ_ServerNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_HHGMCHANCBJ_ServerNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_HHGMCHANCBJ_ServerNotify) ProtoMessage() {}

func (x *Unk2700_HHGMCHANCBJ_ServerNotify) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_HHGMCHANCBJ_ServerNotify.ProtoReflect.Descriptor instead.
func (*Unk2700_HHGMCHANCBJ_ServerNotify) Descriptor() ([]byte, []int) {
	return file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_HHGMCHANCBJ_ServerNotify) GetUnk2700_LGBODABIKLL() Unk2700_KBBDJNLFAKD {
	if x != nil {
		return x.Unk2700_LGBODABIKLL
	}
	return Unk2700_KBBDJNLFAKD_Unk2700_KBBDJNLFAKD_Unk2700_FACJMMHAOLB
}

func (x *Unk2700_HHGMCHANCBJ_ServerNotify) GetUnk2700_NBAIINBBBPK() Unk2700_ADGLMHECKKJ {
	if x != nil {
		return x.Unk2700_NBAIINBBBPK
	}
	return Unk2700_ADGLMHECKKJ_Unk2700_ADGLMHECKKJ_Unk2700_KHKEKEIAPBP
}

func (x *Unk2700_HHGMCHANCBJ_ServerNotify) GetUnk2700_EJHNBDLLLFO() *Unk2700_NLFDMMFNMIO {
	if x != nil {
		return x.Unk2700_EJHNBDLLLFO
	}
	return nil
}

func (x *Unk2700_HHGMCHANCBJ_ServerNotify) GetUnk2700_EIOPOPABBNC() []uint32 {
	if x != nil {
		return x.Unk2700_EIOPOPABBNC
	}
	return nil
}

var File_Unk2700_HHGMCHANCBJ_ServerNotify_proto protoreflect.FileDescriptor

var file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x48, 0x47, 0x4d, 0x43, 0x48,
	0x41, 0x4e, 0x43, 0x42, 0x4a, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x5f, 0x41, 0x44, 0x47, 0x4c, 0x4d, 0x48, 0x45, 0x43, 0x4b, 0x4b, 0x4a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x42, 0x42,
	0x44, 0x4a, 0x4e, 0x4c, 0x46, 0x41, 0x4b, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4e, 0x4c, 0x46, 0x44, 0x4d, 0x4d, 0x46, 0x4e,
	0x4d, 0x49, 0x4f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x20, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x48, 0x48, 0x47, 0x4d, 0x43, 0x48, 0x41, 0x4e, 0x43, 0x42,
	0x4a, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x45,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4c, 0x47, 0x42, 0x4f, 0x44, 0x41,
	0x42, 0x49, 0x4b, 0x4c, 0x4c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4b, 0x42, 0x42, 0x44, 0x4a, 0x4e, 0x4c, 0x46, 0x41, 0x4b,
	0x44, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4c, 0x47, 0x42, 0x4f, 0x44, 0x41,
	0x42, 0x49, 0x4b, 0x4c, 0x4c, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x4e, 0x42, 0x41, 0x49, 0x49, 0x4e, 0x42, 0x42, 0x42, 0x50, 0x4b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x44, 0x47,
	0x4c, 0x4d, 0x48, 0x45, 0x43, 0x4b, 0x4b, 0x4a, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x4e, 0x42, 0x41, 0x49, 0x49, 0x4e, 0x42, 0x42, 0x42, 0x50, 0x4b, 0x12, 0x45, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x4a, 0x48, 0x4e, 0x42, 0x44, 0x4c, 0x4c,
	0x4c, 0x46, 0x4f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x4e, 0x4c, 0x46, 0x44, 0x4d, 0x4d, 0x46, 0x4e, 0x4d, 0x49, 0x4f, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x45, 0x4a, 0x48, 0x4e, 0x42, 0x44, 0x4c, 0x4c,
	0x4c, 0x46, 0x4f, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45,
	0x49, 0x4f, 0x50, 0x4f, 0x50, 0x41, 0x42, 0x42, 0x4e, 0x43, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x45, 0x49, 0x4f, 0x50, 0x4f, 0x50, 0x41,
	0x42, 0x42, 0x4e, 0x43, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescOnce sync.Once
	file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescData = file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDesc
)

func file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescGZIP() []byte {
	file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescOnce.Do(func() {
		file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescData)
	})
	return file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDescData
}

var file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_goTypes = []interface{}{
	(*Unk2700_HHGMCHANCBJ_ServerNotify)(nil), // 0: Unk2700_HHGMCHANCBJ_ServerNotify
	(Unk2700_KBBDJNLFAKD)(0),                 // 1: Unk2700_KBBDJNLFAKD
	(Unk2700_ADGLMHECKKJ)(0),                 // 2: Unk2700_ADGLMHECKKJ
	(*Unk2700_NLFDMMFNMIO)(nil),              // 3: Unk2700_NLFDMMFNMIO
}
var file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_depIdxs = []int32{
	1, // 0: Unk2700_HHGMCHANCBJ_ServerNotify.Unk2700_LGBODABIKLL:type_name -> Unk2700_KBBDJNLFAKD
	2, // 1: Unk2700_HHGMCHANCBJ_ServerNotify.Unk2700_NBAIINBBBPK:type_name -> Unk2700_ADGLMHECKKJ
	3, // 2: Unk2700_HHGMCHANCBJ_ServerNotify.Unk2700_EJHNBDLLLFO:type_name -> Unk2700_NLFDMMFNMIO
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_init() }
func file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_init() {
	if File_Unk2700_HHGMCHANCBJ_ServerNotify_proto != nil {
		return
	}
	file_Unk2700_ADGLMHECKKJ_proto_init()
	file_Unk2700_KBBDJNLFAKD_proto_init()
	file_Unk2700_NLFDMMFNMIO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_HHGMCHANCBJ_ServerNotify); i {
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
			RawDescriptor: file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_goTypes,
		DependencyIndexes: file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_depIdxs,
		MessageInfos:      file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_msgTypes,
	}.Build()
	File_Unk2700_HHGMCHANCBJ_ServerNotify_proto = out.File
	file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_rawDesc = nil
	file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_goTypes = nil
	file_Unk2700_HHGMCHANCBJ_ServerNotify_proto_depIdxs = nil
}
