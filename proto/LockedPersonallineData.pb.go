// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: LockedPersonallineData.proto

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

type LockedPersonallineData_LockReason int32

const (
	LockedPersonallineData_LOCK_REASON_LEVEL LockedPersonallineData_LockReason = 0
	LockedPersonallineData_LOCK_REASON_QUEST LockedPersonallineData_LockReason = 1
)

// Enum value maps for LockedPersonallineData_LockReason.
var (
	LockedPersonallineData_LockReason_name = map[int32]string{
		0: "LOCK_REASON_LEVEL",
		1: "LOCK_REASON_QUEST",
	}
	LockedPersonallineData_LockReason_value = map[string]int32{
		"LOCK_REASON_LEVEL": 0,
		"LOCK_REASON_QUEST": 1,
	}
)

func (x LockedPersonallineData_LockReason) Enum() *LockedPersonallineData_LockReason {
	p := new(LockedPersonallineData_LockReason)
	*p = x
	return p
}

func (x LockedPersonallineData_LockReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockedPersonallineData_LockReason) Descriptor() protoreflect.EnumDescriptor {
	return file_LockedPersonallineData_proto_enumTypes[0].Descriptor()
}

func (LockedPersonallineData_LockReason) Type() protoreflect.EnumType {
	return &file_LockedPersonallineData_proto_enumTypes[0]
}

func (x LockedPersonallineData_LockReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LockedPersonallineData_LockReason.Descriptor instead.
func (LockedPersonallineData_LockReason) EnumDescriptor() ([]byte, []int) {
	return file_LockedPersonallineData_proto_rawDescGZIP(), []int{0, 0}
}

type LockedPersonallineData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockReason     LockedPersonallineData_LockReason `protobuf:"varint,2,opt,name=lock_reason,json=lockReason,proto3,enum=LockedPersonallineData_LockReason" json:"lock_reason,omitempty"`
	PersonalLineId uint32                            `protobuf:"varint,13,opt,name=personal_line_id,json=personalLineId,proto3" json:"personal_line_id,omitempty"`
	// Types that are assignable to Param:
	//
	//	*LockedPersonallineData_ChapterId
	//	*LockedPersonallineData_Level
	Param isLockedPersonallineData_Param `protobuf_oneof:"param"`
}

func (x *LockedPersonallineData) Reset() {
	*x = LockedPersonallineData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockedPersonallineData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockedPersonallineData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockedPersonallineData) ProtoMessage() {}

func (x *LockedPersonallineData) ProtoReflect() protoreflect.Message {
	mi := &file_LockedPersonallineData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockedPersonallineData.ProtoReflect.Descriptor instead.
func (*LockedPersonallineData) Descriptor() ([]byte, []int) {
	return file_LockedPersonallineData_proto_rawDescGZIP(), []int{0}
}

func (x *LockedPersonallineData) GetLockReason() LockedPersonallineData_LockReason {
	if x != nil {
		return x.LockReason
	}
	return LockedPersonallineData_LOCK_REASON_LEVEL
}

func (x *LockedPersonallineData) GetPersonalLineId() uint32 {
	if x != nil {
		return x.PersonalLineId
	}
	return 0
}

func (m *LockedPersonallineData) GetParam() isLockedPersonallineData_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (x *LockedPersonallineData) GetChapterId() uint32 {
	if x, ok := x.GetParam().(*LockedPersonallineData_ChapterId); ok {
		return x.ChapterId
	}
	return 0
}

func (x *LockedPersonallineData) GetLevel() uint32 {
	if x, ok := x.GetParam().(*LockedPersonallineData_Level); ok {
		return x.Level
	}
	return 0
}

type isLockedPersonallineData_Param interface {
	isLockedPersonallineData_Param()
}

type LockedPersonallineData_ChapterId struct {
	ChapterId uint32 `protobuf:"varint,3,opt,name=chapter_id,json=chapterId,proto3,oneof"`
}

type LockedPersonallineData_Level struct {
	Level uint32 `protobuf:"varint,1,opt,name=level,proto3,oneof"`
}

func (*LockedPersonallineData_ChapterId) isLockedPersonallineData_Param() {}

func (*LockedPersonallineData_Level) isLockedPersonallineData_Param() {}

var File_LockedPersonallineData_proto protoreflect.FileDescriptor

var file_LockedPersonallineData_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85,
	0x02, 0x0a, 0x16, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x0b, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x6c,
	0x69, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x10, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x3a, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LockedPersonallineData_proto_rawDescOnce sync.Once
	file_LockedPersonallineData_proto_rawDescData = file_LockedPersonallineData_proto_rawDesc
)

func file_LockedPersonallineData_proto_rawDescGZIP() []byte {
	file_LockedPersonallineData_proto_rawDescOnce.Do(func() {
		file_LockedPersonallineData_proto_rawDescData = protoimpl.X.CompressGZIP(file_LockedPersonallineData_proto_rawDescData)
	})
	return file_LockedPersonallineData_proto_rawDescData
}

var file_LockedPersonallineData_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LockedPersonallineData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LockedPersonallineData_proto_goTypes = []interface{}{
	(LockedPersonallineData_LockReason)(0), // 0: LockedPersonallineData.LockReason
	(*LockedPersonallineData)(nil),         // 1: LockedPersonallineData
}
var file_LockedPersonallineData_proto_depIdxs = []int32{
	0, // 0: LockedPersonallineData.lock_reason:type_name -> LockedPersonallineData.LockReason
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LockedPersonallineData_proto_init() }
func file_LockedPersonallineData_proto_init() {
	if File_LockedPersonallineData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_LockedPersonallineData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockedPersonallineData); i {
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
	file_LockedPersonallineData_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*LockedPersonallineData_ChapterId)(nil),
		(*LockedPersonallineData_Level)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LockedPersonallineData_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LockedPersonallineData_proto_goTypes,
		DependencyIndexes: file_LockedPersonallineData_proto_depIdxs,
		EnumInfos:         file_LockedPersonallineData_proto_enumTypes,
		MessageInfos:      file_LockedPersonallineData_proto_msgTypes,
	}.Build()
	File_LockedPersonallineData_proto = out.File
	file_LockedPersonallineData_proto_rawDesc = nil
	file_LockedPersonallineData_proto_goTypes = nil
	file_LockedPersonallineData_proto_depIdxs = nil
}