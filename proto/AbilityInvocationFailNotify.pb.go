// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: AbilityInvocationFailNotify.proto

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

// CmdId: 1107
// EnetChannelId: 0
// EnetIsReliable: true
type AbilityInvocationFailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason   string              `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	EntityId uint32              `protobuf:"varint,13,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Invoke   *AbilityInvokeEntry `protobuf:"bytes,3,opt,name=invoke,proto3" json:"invoke,omitempty"`
}

func (x *AbilityInvocationFailNotify) Reset() {
	*x = AbilityInvocationFailNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AbilityInvocationFailNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbilityInvocationFailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbilityInvocationFailNotify) ProtoMessage() {}

func (x *AbilityInvocationFailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_AbilityInvocationFailNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbilityInvocationFailNotify.ProtoReflect.Descriptor instead.
func (*AbilityInvocationFailNotify) Descriptor() ([]byte, []int) {
	return file_AbilityInvocationFailNotify_proto_rawDescGZIP(), []int{0}
}

func (x *AbilityInvocationFailNotify) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *AbilityInvocationFailNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *AbilityInvocationFailNotify) GetInvoke() *AbilityInvokeEntry {
	if x != nil {
		return x.Invoke
	}
	return nil
}

var File_AbilityInvocationFailNotify_proto protoreflect.FileDescriptor

var file_AbilityInvocationFailNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a,
	0x1b, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x2b, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AbilityInvocationFailNotify_proto_rawDescOnce sync.Once
	file_AbilityInvocationFailNotify_proto_rawDescData = file_AbilityInvocationFailNotify_proto_rawDesc
)

func file_AbilityInvocationFailNotify_proto_rawDescGZIP() []byte {
	file_AbilityInvocationFailNotify_proto_rawDescOnce.Do(func() {
		file_AbilityInvocationFailNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityInvocationFailNotify_proto_rawDescData)
	})
	return file_AbilityInvocationFailNotify_proto_rawDescData
}

var file_AbilityInvocationFailNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AbilityInvocationFailNotify_proto_goTypes = []interface{}{
	(*AbilityInvocationFailNotify)(nil), // 0: AbilityInvocationFailNotify
	(*AbilityInvokeEntry)(nil),          // 1: AbilityInvokeEntry
}
var file_AbilityInvocationFailNotify_proto_depIdxs = []int32{
	1, // 0: AbilityInvocationFailNotify.invoke:type_name -> AbilityInvokeEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AbilityInvocationFailNotify_proto_init() }
func file_AbilityInvocationFailNotify_proto_init() {
	if File_AbilityInvocationFailNotify_proto != nil {
		return
	}
	file_AbilityInvokeEntry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AbilityInvocationFailNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbilityInvocationFailNotify); i {
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
			RawDescriptor: file_AbilityInvocationFailNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityInvocationFailNotify_proto_goTypes,
		DependencyIndexes: file_AbilityInvocationFailNotify_proto_depIdxs,
		MessageInfos:      file_AbilityInvocationFailNotify_proto_msgTypes,
	}.Build()
	File_AbilityInvocationFailNotify_proto = out.File
	file_AbilityInvocationFailNotify_proto_rawDesc = nil
	file_AbilityInvocationFailNotify_proto_goTypes = nil
	file_AbilityInvocationFailNotify_proto_depIdxs = nil
}
