// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SummerTimeSprintBoatRecord.proto

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

type SummerTimeSprintBoatRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BestScore     uint32   `protobuf:"varint,3,opt,name=best_score,json=bestScore,proto3" json:"best_score,omitempty"`
	StartTime     uint32   `protobuf:"varint,13,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	IsTouched     bool     `protobuf:"varint,7,opt,name=is_touched,json=isTouched,proto3" json:"is_touched,omitempty"`
	WatcherIdList []uint32 `protobuf:"varint,10,rep,packed,name=watcher_id_list,json=watcherIdList,proto3" json:"watcher_id_list,omitempty"`
	GroupId       uint32   `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *SummerTimeSprintBoatRecord) Reset() {
	*x = SummerTimeSprintBoatRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SummerTimeSprintBoatRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SummerTimeSprintBoatRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SummerTimeSprintBoatRecord) ProtoMessage() {}

func (x *SummerTimeSprintBoatRecord) ProtoReflect() protoreflect.Message {
	mi := &file_SummerTimeSprintBoatRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SummerTimeSprintBoatRecord.ProtoReflect.Descriptor instead.
func (*SummerTimeSprintBoatRecord) Descriptor() ([]byte, []int) {
	return file_SummerTimeSprintBoatRecord_proto_rawDescGZIP(), []int{0}
}

func (x *SummerTimeSprintBoatRecord) GetBestScore() uint32 {
	if x != nil {
		return x.BestScore
	}
	return 0
}

func (x *SummerTimeSprintBoatRecord) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *SummerTimeSprintBoatRecord) GetIsTouched() bool {
	if x != nil {
		return x.IsTouched
	}
	return false
}

func (x *SummerTimeSprintBoatRecord) GetWatcherIdList() []uint32 {
	if x != nil {
		return x.WatcherIdList
	}
	return nil
}

func (x *SummerTimeSprintBoatRecord) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_SummerTimeSprintBoatRecord_proto protoreflect.FileDescriptor

var file_SummerTimeSprintBoatRecord_proto_rawDesc = []byte{
	0x0a, 0x20, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x72, 0x69,
	0x6e, 0x74, 0x42, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x1a, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x42, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x65, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x74, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SummerTimeSprintBoatRecord_proto_rawDescOnce sync.Once
	file_SummerTimeSprintBoatRecord_proto_rawDescData = file_SummerTimeSprintBoatRecord_proto_rawDesc
)

func file_SummerTimeSprintBoatRecord_proto_rawDescGZIP() []byte {
	file_SummerTimeSprintBoatRecord_proto_rawDescOnce.Do(func() {
		file_SummerTimeSprintBoatRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_SummerTimeSprintBoatRecord_proto_rawDescData)
	})
	return file_SummerTimeSprintBoatRecord_proto_rawDescData
}

var file_SummerTimeSprintBoatRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SummerTimeSprintBoatRecord_proto_goTypes = []interface{}{
	(*SummerTimeSprintBoatRecord)(nil), // 0: SummerTimeSprintBoatRecord
}
var file_SummerTimeSprintBoatRecord_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SummerTimeSprintBoatRecord_proto_init() }
func file_SummerTimeSprintBoatRecord_proto_init() {
	if File_SummerTimeSprintBoatRecord_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SummerTimeSprintBoatRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SummerTimeSprintBoatRecord); i {
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
			RawDescriptor: file_SummerTimeSprintBoatRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SummerTimeSprintBoatRecord_proto_goTypes,
		DependencyIndexes: file_SummerTimeSprintBoatRecord_proto_depIdxs,
		MessageInfos:      file_SummerTimeSprintBoatRecord_proto_msgTypes,
	}.Build()
	File_SummerTimeSprintBoatRecord_proto = out.File
	file_SummerTimeSprintBoatRecord_proto_rawDesc = nil
	file_SummerTimeSprintBoatRecord_proto_goTypes = nil
	file_SummerTimeSprintBoatRecord_proto_depIdxs = nil
}