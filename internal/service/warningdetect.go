package service

import (
	"context"

	pb "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	util "gitee.com/moyusir/util/api/util/v1"
)

type WarningDetectService struct {
	pb.UnimplementedWarningDetectServer
}

func NewWarningDetectService() *WarningDetectService {
	return &WarningDetectService{}
}

func (s *WarningDetectService) BatchGetDeviceState(ctx context.Context, req *pb.BatchGetDeviceStateRequest) (*pb.BatchGetDeviceStateReply, error) {
	return &pb.BatchGetDeviceStateReply{}, nil
}
func (s *WarningDetectService) GetDeviceStateRegisterInfo(ctx context.Context, req *pb.GetDeviceStateRegisterInfoRequest) (*pb.api.util.v1.DeviceStateRegisterInfo, error) {
	return &pb.api.util.v1.DeviceStateRegisterInfo{}, nil
}
