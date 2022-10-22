// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: SceneAreaWeatherNotify.proto

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

// CmdId: 230
// EnetChannelId: 0
// EnetIsReliable: true
type SceneAreaWeatherNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WeatherAreaId   uint32            `protobuf:"varint,1,opt,name=weather_area_id,json=weatherAreaId,proto3" json:"weather_area_id,omitempty"`
	WeatherGadgetId uint32            `protobuf:"varint,9,opt,name=weather_gadget_id,json=weatherGadgetId,proto3" json:"weather_gadget_id,omitempty"`
	ClimateType     uint32            `protobuf:"varint,14,opt,name=climate_type,json=climateType,proto3" json:"climate_type,omitempty"`
	TransDuration   float32           `protobuf:"fixed32,15,opt,name=trans_duration,json=transDuration,proto3" json:"trans_duration,omitempty"`
	WeatherValueMap map[uint32]string `protobuf:"bytes,10,rep,name=weather_value_map,json=weatherValueMap,proto3" json:"weather_value_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SceneAreaWeatherNotify) Reset() {
	*x = SceneAreaWeatherNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneAreaWeatherNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneAreaWeatherNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneAreaWeatherNotify) ProtoMessage() {}

func (x *SceneAreaWeatherNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SceneAreaWeatherNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneAreaWeatherNotify.ProtoReflect.Descriptor instead.
func (*SceneAreaWeatherNotify) Descriptor() ([]byte, []int) {
	return file_SceneAreaWeatherNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SceneAreaWeatherNotify) GetWeatherAreaId() uint32 {
	if x != nil {
		return x.WeatherAreaId
	}
	return 0
}

func (x *SceneAreaWeatherNotify) GetWeatherGadgetId() uint32 {
	if x != nil {
		return x.WeatherGadgetId
	}
	return 0
}

func (x *SceneAreaWeatherNotify) GetClimateType() uint32 {
	if x != nil {
		return x.ClimateType
	}
	return 0
}

func (x *SceneAreaWeatherNotify) GetTransDuration() float32 {
	if x != nil {
		return x.TransDuration
	}
	return 0
}

func (x *SceneAreaWeatherNotify) GetWeatherValueMap() map[uint32]string {
	if x != nil {
		return x.WeatherValueMap
	}
	return nil
}

var File_SceneAreaWeatherNotify_proto protoreflect.FileDescriptor

var file_SceneAreaWeatherNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x57, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4,
	0x02, 0x0a, 0x16, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x64,
	0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x11, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61,
	0x70, 0x1a, 0x42, 0x0a, 0x14, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneAreaWeatherNotify_proto_rawDescOnce sync.Once
	file_SceneAreaWeatherNotify_proto_rawDescData = file_SceneAreaWeatherNotify_proto_rawDesc
)

func file_SceneAreaWeatherNotify_proto_rawDescGZIP() []byte {
	file_SceneAreaWeatherNotify_proto_rawDescOnce.Do(func() {
		file_SceneAreaWeatherNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneAreaWeatherNotify_proto_rawDescData)
	})
	return file_SceneAreaWeatherNotify_proto_rawDescData
}

var file_SceneAreaWeatherNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SceneAreaWeatherNotify_proto_goTypes = []interface{}{
	(*SceneAreaWeatherNotify)(nil), // 0: SceneAreaWeatherNotify
	nil,                            // 1: SceneAreaWeatherNotify.WeatherValueMapEntry
}
var file_SceneAreaWeatherNotify_proto_depIdxs = []int32{
	1, // 0: SceneAreaWeatherNotify.weather_value_map:type_name -> SceneAreaWeatherNotify.WeatherValueMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneAreaWeatherNotify_proto_init() }
func file_SceneAreaWeatherNotify_proto_init() {
	if File_SceneAreaWeatherNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneAreaWeatherNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneAreaWeatherNotify); i {
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
			RawDescriptor: file_SceneAreaWeatherNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneAreaWeatherNotify_proto_goTypes,
		DependencyIndexes: file_SceneAreaWeatherNotify_proto_depIdxs,
		MessageInfos:      file_SceneAreaWeatherNotify_proto_msgTypes,
	}.Build()
	File_SceneAreaWeatherNotify_proto = out.File
	file_SceneAreaWeatherNotify_proto_rawDesc = nil
	file_SceneAreaWeatherNotify_proto_goTypes = nil
	file_SceneAreaWeatherNotify_proto_depIdxs = nil
}
