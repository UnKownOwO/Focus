// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: Unk2800_JIPMJPAKIKE.proto

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

type Unk2800_JIPMJPAKIKE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFinished          bool   `protobuf:"varint,7,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	Unk2800_MMPELBBNFOD uint32 `protobuf:"varint,10,opt,name=Unk2800_MMPELBBNFOD,json=Unk2800MMPELBBNFOD,proto3" json:"Unk2800_MMPELBBNFOD,omitempty"`
	IsOpen              bool   `protobuf:"varint,5,opt,name=is_open,json=isOpen,proto3" json:"is_open,omitempty"`
	Unk2800_MGPEODNKEEC uint32 `protobuf:"varint,6,opt,name=Unk2800_MGPEODNKEEC,json=Unk2800MGPEODNKEEC,proto3" json:"Unk2800_MGPEODNKEEC,omitempty"`
}

func (x *Unk2800_JIPMJPAKIKE) Reset() {
	*x = Unk2800_JIPMJPAKIKE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2800_JIPMJPAKIKE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2800_JIPMJPAKIKE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2800_JIPMJPAKIKE) ProtoMessage() {}

func (x *Unk2800_JIPMJPAKIKE) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2800_JIPMJPAKIKE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2800_JIPMJPAKIKE.ProtoReflect.Descriptor instead.
func (*Unk2800_JIPMJPAKIKE) Descriptor() ([]byte, []int) {
	return file_Unk2800_JIPMJPAKIKE_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2800_JIPMJPAKIKE) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *Unk2800_JIPMJPAKIKE) GetUnk2800_MMPELBBNFOD() uint32 {
	if x != nil {
		return x.Unk2800_MMPELBBNFOD
	}
	return 0
}

func (x *Unk2800_JIPMJPAKIKE) GetIsOpen() bool {
	if x != nil {
		return x.IsOpen
	}
	return false
}

func (x *Unk2800_JIPMJPAKIKE) GetUnk2800_MGPEODNKEEC() uint32 {
	if x != nil {
		return x.Unk2800_MGPEODNKEEC
	}
	return 0
}

var File_Unk2800_JIPMJPAKIKE_proto protoreflect.FileDescriptor

var file_Unk2800_JIPMJPAKIKE_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4a, 0x49, 0x50, 0x4d, 0x4a, 0x50,
	0x41, 0x4b, 0x49, 0x4b, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4a, 0x49, 0x50, 0x4d, 0x4a, 0x50, 0x41, 0x4b,
	0x49, 0x4b, 0x45, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f,
	0x4d, 0x4d, 0x50, 0x45, 0x4c, 0x42, 0x42, 0x4e, 0x46, 0x4f, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x4d, 0x4d, 0x50, 0x45, 0x4c, 0x42,
	0x42, 0x4e, 0x46, 0x4f, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x2f,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4d, 0x47, 0x50, 0x45, 0x4f, 0x44,
	0x4e, 0x4b, 0x45, 0x45, 0x43, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x38, 0x30, 0x30, 0x4d, 0x47, 0x50, 0x45, 0x4f, 0x44, 0x4e, 0x4b, 0x45, 0x45, 0x43, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Unk2800_JIPMJPAKIKE_proto_rawDescOnce sync.Once
	file_Unk2800_JIPMJPAKIKE_proto_rawDescData = file_Unk2800_JIPMJPAKIKE_proto_rawDesc
)

func file_Unk2800_JIPMJPAKIKE_proto_rawDescGZIP() []byte {
	file_Unk2800_JIPMJPAKIKE_proto_rawDescOnce.Do(func() {
		file_Unk2800_JIPMJPAKIKE_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2800_JIPMJPAKIKE_proto_rawDescData)
	})
	return file_Unk2800_JIPMJPAKIKE_proto_rawDescData
}

var file_Unk2800_JIPMJPAKIKE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2800_JIPMJPAKIKE_proto_goTypes = []interface{}{
	(*Unk2800_JIPMJPAKIKE)(nil), // 0: Unk2800_JIPMJPAKIKE
}
var file_Unk2800_JIPMJPAKIKE_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2800_JIPMJPAKIKE_proto_init() }
func file_Unk2800_JIPMJPAKIKE_proto_init() {
	if File_Unk2800_JIPMJPAKIKE_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk2800_JIPMJPAKIKE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2800_JIPMJPAKIKE); i {
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
			RawDescriptor: file_Unk2800_JIPMJPAKIKE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2800_JIPMJPAKIKE_proto_goTypes,
		DependencyIndexes: file_Unk2800_JIPMJPAKIKE_proto_depIdxs,
		MessageInfos:      file_Unk2800_JIPMJPAKIKE_proto_msgTypes,
	}.Build()
	File_Unk2800_JIPMJPAKIKE_proto = out.File
	file_Unk2800_JIPMJPAKIKE_proto_rawDesc = nil
	file_Unk2800_JIPMJPAKIKE_proto_goTypes = nil
	file_Unk2800_JIPMJPAKIKE_proto_depIdxs = nil
}
