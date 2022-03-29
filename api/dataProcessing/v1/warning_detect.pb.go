// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/dataProcessing/v1/warning_detect.proto

package v1

import (
	v1 "gitee.com/moyusir/util/api/util/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceClassId int32                  `protobuf:"varint,1,opt,name=device_class_id,json=deviceClassId,proto3" json:"device_class_id,omitempty"`
	DeviceId      string                 `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Time          *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Fields        map[string]float64     `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Tags          map[string]string      `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DeviceState) Reset() {
	*x = DeviceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceState) ProtoMessage() {}

func (x *DeviceState) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceState.ProtoReflect.Descriptor instead.
func (*DeviceState) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceState) GetDeviceClassId() int32 {
	if x != nil {
		return x.DeviceClassId
	}
	return 0
}

func (x *DeviceState) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DeviceState) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DeviceState) GetFields() map[string]float64 {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *DeviceState) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// 设备状态信息批量查询请求
type BatchGetDeviceStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设备类别号
	DeviceClassId int32 `protobuf:"varint,1,opt,name=device_class_id,json=deviceClassId,proto3" json:"device_class_id,omitempty"`
	// 时间范围 开始时间
	Start *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	// 结束时间
	End *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
	// 相对时间范围查询，用于查询过去若干时间内的记录
	Past *durationpb.Duration `protobuf:"bytes,4,opt,name=past,proto3" json:"past,omitempty"`
	// 查询过滤的条件，可以针对_measurement、_field或tag进行过滤
	Filter map[string]string `protobuf:"bytes,5,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BatchGetDeviceStateRequest) Reset() {
	*x = BatchGetDeviceStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetDeviceStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetDeviceStateRequest) ProtoMessage() {}

func (x *BatchGetDeviceStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetDeviceStateRequest.ProtoReflect.Descriptor instead.
func (*BatchGetDeviceStateRequest) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{1}
}

func (x *BatchGetDeviceStateRequest) GetDeviceClassId() int32 {
	if x != nil {
		return x.DeviceClassId
	}
	return 0
}

func (x *BatchGetDeviceStateRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *BatchGetDeviceStateRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *BatchGetDeviceStateRequest) GetPast() *durationpb.Duration {
	if x != nil {
		return x.Past
	}
	return nil
}

func (x *BatchGetDeviceStateRequest) GetFilter() map[string]string {
	if x != nil {
		return x.Filter
	}
	return nil
}

// 设备状态信息批量查询响应
type BatchGetDeviceStateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	States []*DeviceState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty"`
}

func (x *BatchGetDeviceStateReply) Reset() {
	*x = BatchGetDeviceStateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetDeviceStateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetDeviceStateReply) ProtoMessage() {}

func (x *BatchGetDeviceStateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetDeviceStateReply.ProtoReflect.Descriptor instead.
func (*BatchGetDeviceStateReply) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{2}
}

func (x *BatchGetDeviceStateReply) GetStates() []*DeviceState {
	if x != nil {
		return x.States
	}
	return nil
}

// 警告消息批量查询请求
type BatchGetWarningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 时间范围 开始时间
	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	// 结束时间
	End *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	// 相对时间范围查询，用于查询过去若干时间内的记录
	Past *durationpb.Duration `protobuf:"bytes,3,opt,name=past,proto3" json:"past,omitempty"`
	// 查询过滤的条件，可以针对_measurement、_field或tag进行过滤
	Filter map[string]string `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BatchGetWarningRequest) Reset() {
	*x = BatchGetWarningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetWarningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetWarningRequest) ProtoMessage() {}

func (x *BatchGetWarningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetWarningRequest.ProtoReflect.Descriptor instead.
func (*BatchGetWarningRequest) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{3}
}

func (x *BatchGetWarningRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *BatchGetWarningRequest) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

func (x *BatchGetWarningRequest) GetPast() *durationpb.Duration {
	if x != nil {
		return x.Past
	}
	return nil
}

func (x *BatchGetWarningRequest) GetFilter() map[string]string {
	if x != nil {
		return x.Filter
	}
	return nil
}

// 警告消息批量查询响应
type BatchGetWarningReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Warnings []*v1.Warning `protobuf:"bytes,1,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *BatchGetWarningReply) Reset() {
	*x = BatchGetWarningReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetWarningReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetWarningReply) ProtoMessage() {}

func (x *BatchGetWarningReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetWarningReply.ProtoReflect.Descriptor instead.
func (*BatchGetWarningReply) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{4}
}

func (x *BatchGetWarningReply) GetWarnings() []*v1.Warning {
	if x != nil {
		return x.Warnings
	}
	return nil
}

// 设备状态注册信息查询请求
type GetDeviceStateRegisterInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设备类别号
	DeviceClassId int64 `protobuf:"varint,1,opt,name=device_class_id,json=deviceClassId,proto3" json:"device_class_id,omitempty"`
}

func (x *GetDeviceStateRegisterInfoRequest) Reset() {
	*x = GetDeviceStateRegisterInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStateRegisterInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStateRegisterInfoRequest) ProtoMessage() {}

func (x *GetDeviceStateRegisterInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceStateRegisterInfoRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceStateRegisterInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeviceStateRegisterInfoRequest) GetDeviceClassId() int64 {
	if x != nil {
		return x.DeviceClassId
	}
	return 0
}

type DeviceState0 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Voltage     float64                `protobuf:"fixed64,3,opt,name=voltage,proto3" json:"voltage,omitempty"`
	Current     float64                `protobuf:"fixed64,4,opt,name=current,proto3" json:"current,omitempty"`
	Temperature float64                `protobuf:"fixed64,5,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *DeviceState0) Reset() {
	*x = DeviceState0{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceState0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceState0) ProtoMessage() {}

func (x *DeviceState0) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceState0.ProtoReflect.Descriptor instead.
func (*DeviceState0) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{6}
}

func (x *DeviceState0) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceState0) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DeviceState0) GetVoltage() float64 {
	if x != nil {
		return x.Voltage
	}
	return 0
}

func (x *DeviceState0) GetCurrent() float64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *DeviceState0) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

type DeviceState1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Voltage     float64                `protobuf:"fixed64,3,opt,name=voltage,proto3" json:"voltage,omitempty"`
	Current     float64                `protobuf:"fixed64,4,opt,name=current,proto3" json:"current,omitempty"`
	Temperature float64                `protobuf:"fixed64,5,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *DeviceState1) Reset() {
	*x = DeviceState1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceState1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceState1) ProtoMessage() {}

func (x *DeviceState1) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warning_detect_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceState1.ProtoReflect.Descriptor instead.
func (*DeviceState1) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP(), []int{7}
}

func (x *DeviceState1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeviceState1) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DeviceState1) GetVoltage() float64 {
	if x != nil {
		return x.Voltage
	}
	return 0
}

func (x *DeviceState1) GetCurrent() float64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *DeviceState1) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

var File_api_dataProcessing_v1_warning_detect_proto protoreflect.FileDescriptor

var file_api_dataProcessing_v1_warning_detect_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x74, 0x69,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x40, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09,
	0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe5, 0x02, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x2d, 0x0a, 0x04,
	0x70, 0x61, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x61, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x56, 0x0a,
	0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x2d, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x61, 0x73, 0x74, 0x12,
	0x51, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x48, 0x0a,
	0x14, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x77,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x4b, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0c,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x31, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x76,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x32, 0xe5, 0x03, 0x0a, 0x0d, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x12, 0xa0, 0x01, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x0f, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x6e,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x57, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12,
	0x09, 0x2f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0xad, 0x01, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x29, 0x12, 0x27, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x66,
	0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x55, 0x0a, 0x15, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x6f, 0x79, 0x75, 0x73, 0x69, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_dataProcessing_v1_warning_detect_proto_rawDescOnce sync.Once
	file_api_dataProcessing_v1_warning_detect_proto_rawDescData = file_api_dataProcessing_v1_warning_detect_proto_rawDesc
)

func file_api_dataProcessing_v1_warning_detect_proto_rawDescGZIP() []byte {
	file_api_dataProcessing_v1_warning_detect_proto_rawDescOnce.Do(func() {
		file_api_dataProcessing_v1_warning_detect_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dataProcessing_v1_warning_detect_proto_rawDescData)
	})
	return file_api_dataProcessing_v1_warning_detect_proto_rawDescData
}

var file_api_dataProcessing_v1_warning_detect_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_dataProcessing_v1_warning_detect_proto_goTypes = []interface{}{
	(*DeviceState)(nil),                       // 0: api.dataProcessing.v1.DeviceState
	(*BatchGetDeviceStateRequest)(nil),        // 1: api.dataProcessing.v1.BatchGetDeviceStateRequest
	(*BatchGetDeviceStateReply)(nil),          // 2: api.dataProcessing.v1.BatchGetDeviceStateReply
	(*BatchGetWarningRequest)(nil),            // 3: api.dataProcessing.v1.BatchGetWarningRequest
	(*BatchGetWarningReply)(nil),              // 4: api.dataProcessing.v1.BatchGetWarningReply
	(*GetDeviceStateRegisterInfoRequest)(nil), // 5: api.dataProcessing.v1.GetDeviceStateRegisterInfoRequest
	(*DeviceState0)(nil),                      // 6: api.dataProcessing.v1.DeviceState0
	(*DeviceState1)(nil),                      // 7: api.dataProcessing.v1.DeviceState1
	nil,                                       // 8: api.dataProcessing.v1.DeviceState.FieldsEntry
	nil,                                       // 9: api.dataProcessing.v1.DeviceState.TagsEntry
	nil,                                       // 10: api.dataProcessing.v1.BatchGetDeviceStateRequest.FilterEntry
	nil,                                       // 11: api.dataProcessing.v1.BatchGetWarningRequest.FilterEntry
	(*timestamppb.Timestamp)(nil),             // 12: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),               // 13: google.protobuf.Duration
	(*v1.Warning)(nil),                        // 14: api.util.v1.Warning
	(*v1.DeviceStateRegisterInfo)(nil),        // 15: api.util.v1.DeviceStateRegisterInfo
}
var file_api_dataProcessing_v1_warning_detect_proto_depIdxs = []int32{
	12, // 0: api.dataProcessing.v1.DeviceState.time:type_name -> google.protobuf.Timestamp
	8,  // 1: api.dataProcessing.v1.DeviceState.fields:type_name -> api.dataProcessing.v1.DeviceState.FieldsEntry
	9,  // 2: api.dataProcessing.v1.DeviceState.tags:type_name -> api.dataProcessing.v1.DeviceState.TagsEntry
	12, // 3: api.dataProcessing.v1.BatchGetDeviceStateRequest.start:type_name -> google.protobuf.Timestamp
	12, // 4: api.dataProcessing.v1.BatchGetDeviceStateRequest.end:type_name -> google.protobuf.Timestamp
	13, // 5: api.dataProcessing.v1.BatchGetDeviceStateRequest.past:type_name -> google.protobuf.Duration
	10, // 6: api.dataProcessing.v1.BatchGetDeviceStateRequest.filter:type_name -> api.dataProcessing.v1.BatchGetDeviceStateRequest.FilterEntry
	0,  // 7: api.dataProcessing.v1.BatchGetDeviceStateReply.states:type_name -> api.dataProcessing.v1.DeviceState
	12, // 8: api.dataProcessing.v1.BatchGetWarningRequest.start:type_name -> google.protobuf.Timestamp
	12, // 9: api.dataProcessing.v1.BatchGetWarningRequest.end:type_name -> google.protobuf.Timestamp
	13, // 10: api.dataProcessing.v1.BatchGetWarningRequest.past:type_name -> google.protobuf.Duration
	11, // 11: api.dataProcessing.v1.BatchGetWarningRequest.filter:type_name -> api.dataProcessing.v1.BatchGetWarningRequest.FilterEntry
	14, // 12: api.dataProcessing.v1.BatchGetWarningReply.warnings:type_name -> api.util.v1.Warning
	12, // 13: api.dataProcessing.v1.DeviceState0.time:type_name -> google.protobuf.Timestamp
	12, // 14: api.dataProcessing.v1.DeviceState1.time:type_name -> google.protobuf.Timestamp
	1,  // 15: api.dataProcessing.v1.WarningDetect.BatchGetDeviceStateInfo:input_type -> api.dataProcessing.v1.BatchGetDeviceStateRequest
	3,  // 16: api.dataProcessing.v1.WarningDetect.BatchGetWarning:input_type -> api.dataProcessing.v1.BatchGetWarningRequest
	5,  // 17: api.dataProcessing.v1.WarningDetect.GetDeviceStateRegisterInfo:input_type -> api.dataProcessing.v1.GetDeviceStateRegisterInfoRequest
	2,  // 18: api.dataProcessing.v1.WarningDetect.BatchGetDeviceStateInfo:output_type -> api.dataProcessing.v1.BatchGetDeviceStateReply
	4,  // 19: api.dataProcessing.v1.WarningDetect.BatchGetWarning:output_type -> api.dataProcessing.v1.BatchGetWarningReply
	15, // 20: api.dataProcessing.v1.WarningDetect.GetDeviceStateRegisterInfo:output_type -> api.util.v1.DeviceStateRegisterInfo
	18, // [18:21] is the sub-list for method output_type
	15, // [15:18] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_api_dataProcessing_v1_warning_detect_proto_init() }
func file_api_dataProcessing_v1_warning_detect_proto_init() {
	if File_api_dataProcessing_v1_warning_detect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceState); i {
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
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetDeviceStateRequest); i {
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
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetDeviceStateReply); i {
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
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetWarningRequest); i {
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
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetWarningReply); i {
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
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceStateRegisterInfoRequest); i {
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
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceState0); i {
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
		file_api_dataProcessing_v1_warning_detect_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceState1); i {
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
			RawDescriptor: file_api_dataProcessing_v1_warning_detect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dataProcessing_v1_warning_detect_proto_goTypes,
		DependencyIndexes: file_api_dataProcessing_v1_warning_detect_proto_depIdxs,
		MessageInfos:      file_api_dataProcessing_v1_warning_detect_proto_msgTypes,
	}.Build()
	File_api_dataProcessing_v1_warning_detect_proto = out.File
	file_api_dataProcessing_v1_warning_detect_proto_rawDesc = nil
	file_api_dataProcessing_v1_warning_detect_proto_goTypes = nil
	file_api_dataProcessing_v1_warning_detect_proto_depIdxs = nil
}
