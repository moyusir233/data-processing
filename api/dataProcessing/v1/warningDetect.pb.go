// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/dataProcessing/v1/warningDetect.proto

package v1

import (
	v1 "gitee.com/moyusir/util/api/util/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// 设备状态信息分页查询请求
type BatchGetDeviceStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 设备类别号
	DeviceClassId int64 `protobuf:"varint,2,opt,name=device_class_id,json=deviceClassId,proto3" json:"device_class_id,omitempty"`
	// 时间范围 开始时间
	Start *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	// 结束时间
	End *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	// 分页查询
	Page  int64 `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Count int64 `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BatchGetDeviceStateRequest) Reset() {
	*x = BatchGetDeviceStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetDeviceStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetDeviceStateRequest) ProtoMessage() {}

func (x *BatchGetDeviceStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[0]
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
	return file_api_dataProcessing_v1_warningDetect_proto_rawDescGZIP(), []int{0}
}

func (x *BatchGetDeviceStateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BatchGetDeviceStateRequest) GetDeviceClassId() int64 {
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

func (x *BatchGetDeviceStateRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *BatchGetDeviceStateRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 设备状态信息分页查询响应
type BatchGetDeviceStateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceStateInfos []*BatchGetDeviceStateReply_DeviceStateInfo `protobuf:"bytes,1,rep,name=device_state_infos,json=deviceStateInfos,proto3" json:"device_state_infos,omitempty"`
}

func (x *BatchGetDeviceStateReply) Reset() {
	*x = BatchGetDeviceStateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetDeviceStateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetDeviceStateReply) ProtoMessage() {}

func (x *BatchGetDeviceStateReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[1]
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
	return file_api_dataProcessing_v1_warningDetect_proto_rawDescGZIP(), []int{1}
}

func (x *BatchGetDeviceStateReply) GetDeviceStateInfos() []*BatchGetDeviceStateReply_DeviceStateInfo {
	if x != nil {
		return x.DeviceStateInfos
	}
	return nil
}

// 设备状态注册信息查询请求
type GetDeviceStateRegisterInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属用户id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 设备类别号
	DeviceClassId int64 `protobuf:"varint,2,opt,name=device_class_id,json=deviceClassId,proto3" json:"device_class_id,omitempty"`
	// 设备id
	DeviceId int64 `protobuf:"varint,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetDeviceStateRegisterInfoRequest) Reset() {
	*x = GetDeviceStateRegisterInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceStateRegisterInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceStateRegisterInfoRequest) ProtoMessage() {}

func (x *GetDeviceStateRegisterInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[2]
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
	return file_api_dataProcessing_v1_warningDetect_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeviceStateRegisterInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetDeviceStateRegisterInfoRequest) GetDeviceClassId() int64 {
	if x != nil {
		return x.DeviceClassId
	}
	return 0
}

func (x *GetDeviceStateRegisterInfoRequest) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

type BatchGetDeviceStateReply_DeviceStateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BatchGetDeviceStateReply_DeviceStateInfo) Reset() {
	*x = BatchGetDeviceStateReply_DeviceStateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetDeviceStateReply_DeviceStateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetDeviceStateReply_DeviceStateInfo) ProtoMessage() {}

func (x *BatchGetDeviceStateReply_DeviceStateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_dataProcessing_v1_warningDetect_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetDeviceStateReply_DeviceStateInfo.ProtoReflect.Descriptor instead.
func (*BatchGetDeviceStateReply_DeviceStateInfo) Descriptor() ([]byte, []int) {
	return file_api_dataProcessing_v1_warningDetect_proto_rawDescGZIP(), []int{1, 0}
}

var File_api_dataProcessing_v1_warningDetect_proto protoreflect.FileDescriptor

var file_api_dataProcessing_v1_warningDetect_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x74,
	0x69, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x6d, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x1a, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x78, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x32, 0x88, 0x02,
	0x0a, 0x0d, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12,
	0x79, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x7c, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x55, 0x0a, 0x15, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x6f, 0x79, 0x75, 0x73, 0x69, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_dataProcessing_v1_warningDetect_proto_rawDescOnce sync.Once
	file_api_dataProcessing_v1_warningDetect_proto_rawDescData = file_api_dataProcessing_v1_warningDetect_proto_rawDesc
)

func file_api_dataProcessing_v1_warningDetect_proto_rawDescGZIP() []byte {
	file_api_dataProcessing_v1_warningDetect_proto_rawDescOnce.Do(func() {
		file_api_dataProcessing_v1_warningDetect_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dataProcessing_v1_warningDetect_proto_rawDescData)
	})
	return file_api_dataProcessing_v1_warningDetect_proto_rawDescData
}

var file_api_dataProcessing_v1_warningDetect_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_dataProcessing_v1_warningDetect_proto_goTypes = []interface{}{
	(*BatchGetDeviceStateRequest)(nil),               // 0: api.dataProcessing.v1.BatchGetDeviceStateRequest
	(*BatchGetDeviceStateReply)(nil),                 // 1: api.dataProcessing.v1.BatchGetDeviceStateReply
	(*GetDeviceStateRegisterInfoRequest)(nil),        // 2: api.dataProcessing.v1.GetDeviceStateRegisterInfoRequest
	(*BatchGetDeviceStateReply_DeviceStateInfo)(nil), // 3: api.dataProcessing.v1.BatchGetDeviceStateReply.DeviceStateInfo
	(*timestamppb.Timestamp)(nil),                    // 4: google.protobuf.Timestamp
	(*v1.DeviceStateRegisterInfo)(nil),               // 5: api.util.v1.DeviceStateRegisterInfo
}
var file_api_dataProcessing_v1_warningDetect_proto_depIdxs = []int32{
	4, // 0: api.dataProcessing.v1.BatchGetDeviceStateRequest.start:type_name -> google.protobuf.Timestamp
	4, // 1: api.dataProcessing.v1.BatchGetDeviceStateRequest.end:type_name -> google.protobuf.Timestamp
	3, // 2: api.dataProcessing.v1.BatchGetDeviceStateReply.device_state_infos:type_name -> api.dataProcessing.v1.BatchGetDeviceStateReply.DeviceStateInfo
	0, // 3: api.dataProcessing.v1.WarningDetect.BatchGetDeviceState:input_type -> api.dataProcessing.v1.BatchGetDeviceStateRequest
	2, // 4: api.dataProcessing.v1.WarningDetect.GetDeviceStateRegisterInfo:input_type -> api.dataProcessing.v1.GetDeviceStateRegisterInfoRequest
	1, // 5: api.dataProcessing.v1.WarningDetect.BatchGetDeviceState:output_type -> api.dataProcessing.v1.BatchGetDeviceStateReply
	5, // 6: api.dataProcessing.v1.WarningDetect.GetDeviceStateRegisterInfo:output_type -> api.util.v1.DeviceStateRegisterInfo
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_dataProcessing_v1_warningDetect_proto_init() }
func file_api_dataProcessing_v1_warningDetect_proto_init() {
	if File_api_dataProcessing_v1_warningDetect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_dataProcessing_v1_warningDetect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_dataProcessing_v1_warningDetect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_dataProcessing_v1_warningDetect_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_dataProcessing_v1_warningDetect_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetDeviceStateReply_DeviceStateInfo); i {
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
			RawDescriptor: file_api_dataProcessing_v1_warningDetect_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dataProcessing_v1_warningDetect_proto_goTypes,
		DependencyIndexes: file_api_dataProcessing_v1_warningDetect_proto_depIdxs,
		MessageInfos:      file_api_dataProcessing_v1_warningDetect_proto_msgTypes,
	}.Build()
	File_api_dataProcessing_v1_warningDetect_proto = out.File
	file_api_dataProcessing_v1_warningDetect_proto_rawDesc = nil
	file_api_dataProcessing_v1_warningDetect_proto_goTypes = nil
	file_api_dataProcessing_v1_warningDetect_proto_depIdxs = nil
}
