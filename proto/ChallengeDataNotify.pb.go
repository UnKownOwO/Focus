// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ChallengeDataNotify.proto

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

// CmdId: 953
// EnetChannelId: 0
// EnetIsReliable: true
type ChallengeDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value          uint32 `protobuf:"varint,8,opt,name=value,proto3" json:"value,omitempty"`
	ChallengeIndex uint32 `protobuf:"varint,2,opt,name=challenge_index,json=challengeIndex,proto3" json:"challenge_index,omitempty"`
	ParamIndex     uint32 `protobuf:"varint,9,opt,name=param_index,json=paramIndex,proto3" json:"param_index,omitempty"`
}

func (x *ChallengeDataNotify) Reset() {
	*x = ChallengeDataNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeDataNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeDataNotify) ProtoMessage() {}

func (x *ChallengeDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeDataNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeDataNotify.ProtoReflect.Descriptor instead.
func (*ChallengeDataNotify) Descriptor() ([]byte, []int) {
	return file_ChallengeDataNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeDataNotify) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ChallengeDataNotify) GetChallengeIndex() uint32 {
	if x != nil {
		return x.ChallengeIndex
	}
	return 0
}

func (x *ChallengeDataNotify) GetParamIndex() uint32 {
	if x != nil {
		return x.ParamIndex
	}
	return 0
}

var File_ChallengeDataNotify_proto protoreflect.FileDescriptor

var file_ChallengeDataNotify_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x13, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeDataNotify_proto_rawDescOnce sync.Once
	file_ChallengeDataNotify_proto_rawDescData = file_ChallengeDataNotify_proto_rawDesc
)

func file_ChallengeDataNotify_proto_rawDescGZIP() []byte {
	file_ChallengeDataNotify_proto_rawDescOnce.Do(func() {
		file_ChallengeDataNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeDataNotify_proto_rawDescData)
	})
	return file_ChallengeDataNotify_proto_rawDescData
}

var file_ChallengeDataNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeDataNotify_proto_goTypes = []interface{}{
	(*ChallengeDataNotify)(nil), // 0: ChallengeDataNotify
}
var file_ChallengeDataNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChallengeDataNotify_proto_init() }
func file_ChallengeDataNotify_proto_init() {
	if File_ChallengeDataNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChallengeDataNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeDataNotify); i {
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
			RawDescriptor: file_ChallengeDataNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeDataNotify_proto_goTypes,
		DependencyIndexes: file_ChallengeDataNotify_proto_depIdxs,
		MessageInfos:      file_ChallengeDataNotify_proto_msgTypes,
	}.Build()
	File_ChallengeDataNotify_proto = out.File
	file_ChallengeDataNotify_proto_rawDesc = nil
	file_ChallengeDataNotify_proto_goTypes = nil
	file_ChallengeDataNotify_proto_depIdxs = nil
}
