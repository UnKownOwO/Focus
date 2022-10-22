// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ReunionPrivilegeInfo.proto

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

type ReunionPrivilegeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurDayCount uint32 `protobuf:"varint,7,opt,name=cur_day_count,json=curDayCount,proto3" json:"cur_day_count,omitempty"`
	TotalCount  uint32 `protobuf:"varint,10,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	PrivilegeId uint32 `protobuf:"varint,4,opt,name=privilege_id,json=privilegeId,proto3" json:"privilege_id,omitempty"`
}

func (x *ReunionPrivilegeInfo) Reset() {
	*x = ReunionPrivilegeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReunionPrivilegeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReunionPrivilegeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReunionPrivilegeInfo) ProtoMessage() {}

func (x *ReunionPrivilegeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ReunionPrivilegeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReunionPrivilegeInfo.ProtoReflect.Descriptor instead.
func (*ReunionPrivilegeInfo) Descriptor() ([]byte, []int) {
	return file_ReunionPrivilegeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ReunionPrivilegeInfo) GetCurDayCount() uint32 {
	if x != nil {
		return x.CurDayCount
	}
	return 0
}

func (x *ReunionPrivilegeInfo) GetTotalCount() uint32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ReunionPrivilegeInfo) GetPrivilegeId() uint32 {
	if x != nil {
		return x.PrivilegeId
	}
	return 0
}

var File_ReunionPrivilegeInfo_proto protoreflect.FileDescriptor

var file_ReunionPrivilegeInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x14,
	0x52, 0x65, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x5f, 0x64, 0x61, 0x79, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72,
	0x44, 0x61, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x49, 0x64, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReunionPrivilegeInfo_proto_rawDescOnce sync.Once
	file_ReunionPrivilegeInfo_proto_rawDescData = file_ReunionPrivilegeInfo_proto_rawDesc
)

func file_ReunionPrivilegeInfo_proto_rawDescGZIP() []byte {
	file_ReunionPrivilegeInfo_proto_rawDescOnce.Do(func() {
		file_ReunionPrivilegeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReunionPrivilegeInfo_proto_rawDescData)
	})
	return file_ReunionPrivilegeInfo_proto_rawDescData
}

var file_ReunionPrivilegeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReunionPrivilegeInfo_proto_goTypes = []interface{}{
	(*ReunionPrivilegeInfo)(nil), // 0: ReunionPrivilegeInfo
}
var file_ReunionPrivilegeInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReunionPrivilegeInfo_proto_init() }
func file_ReunionPrivilegeInfo_proto_init() {
	if File_ReunionPrivilegeInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ReunionPrivilegeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReunionPrivilegeInfo); i {
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
			RawDescriptor: file_ReunionPrivilegeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReunionPrivilegeInfo_proto_goTypes,
		DependencyIndexes: file_ReunionPrivilegeInfo_proto_depIdxs,
		MessageInfos:      file_ReunionPrivilegeInfo_proto_msgTypes,
	}.Build()
	File_ReunionPrivilegeInfo_proto = out.File
	file_ReunionPrivilegeInfo_proto_rawDesc = nil
	file_ReunionPrivilegeInfo_proto_goTypes = nil
	file_ReunionPrivilegeInfo_proto_depIdxs = nil
}