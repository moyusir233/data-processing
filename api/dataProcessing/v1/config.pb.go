// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/dataProcessing/v1/config.proto

package v1

import (
	v1 "gitee.com/moyusir/util/api/util/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 配置查询请求
type GetDeviceConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设备类别号
	DeviceClassId int64 `protobuf:"varint,1,opt,name=device_class_id,json=deviceClassId,proto3" json:"device_class_id,omitempty"`
	// 设备id
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetDeviceConfigRequest) Reset() {
	*x = GetDeviceConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceConfigRequest) ProtoMessage() {}

func (x *GetDeviceConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceConfigRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *GetDeviceConfigRequest) GetDeviceClassId() int64 {
	if x != nil {
		return x.DeviceClassId
	}
	return 0
}

func (x *GetDeviceConfigRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

var File_api_dataProcessing_v1_config_proto protoreflect.FileDescriptor

var file_api_dataProcessing_v1_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x74, 0x69, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x32, 0x7d, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x73, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x42, 0x55, 0x0a, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f,
	0x79, 0x75, 0x73, 0x69, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_dataProcessing_v1_config_proto_rawDescOnce sync.Once
	file_api_dataProcessing_v1_config_proto_rawDescData = file_api_dataProcessing_v1_config_proto_rawDesc
)

func file_api_dataProcessing_v1_config_proto_rawDescGZIP() []byte {
	file_api_dataProcessing_v1_config_proto_rawDescOnce.Do(func() {
		file_api_dataProcessing_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dataProcessing_v1_config_proto_rawDescData)
	})
	return file_api_dataProcessing_v1_config_proto_rawDescData
}

var file_api_dataProcessing_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_dataProcessing_v1_config_proto_goTypes = []interface{}{
	(*GetDeviceConfigRequest)(nil), // 0: api.dataProcessing.v1.GetDeviceConfigRequest
	(*v1.TestedDeviceConfig)(nil),  // 1: api.util.v1.TestedDeviceConfig
}
var file_api_dataProcessing_v1_config_proto_depIdxs = []int32{
	0, // 0: api.dataProcessing.v1.Config.GetDeviceConfig:input_type -> api.dataProcessing.v1.GetDeviceConfigRequest
	1, // 1: api.dataProcessing.v1.Config.GetDeviceConfig:output_type -> api.util.v1.TestedDeviceConfig
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_dataProcessing_v1_config_proto_init() }
func file_api_dataProcessing_v1_config_proto_init() {
	if File_api_dataProcessing_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_dataProcessing_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceConfigRequest); i {
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
			RawDescriptor: file_api_dataProcessing_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dataProcessing_v1_config_proto_goTypes,
		DependencyIndexes: file_api_dataProcessing_v1_config_proto_depIdxs,
		MessageInfos:      file_api_dataProcessing_v1_config_proto_msgTypes,
	}.Build()
	File_api_dataProcessing_v1_config_proto = out.File
	file_api_dataProcessing_v1_config_proto_rawDesc = nil
	file_api_dataProcessing_v1_config_proto_goTypes = nil
	file_api_dataProcessing_v1_config_proto_depIdxs = nil
}
