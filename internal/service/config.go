package service

import (
	"context"

	pb "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
)

type ConfigService struct {
	pb.UnimplementedConfigServer
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (s *ConfigService) GetDeviceConfig(ctx context.Context, req *pb.GetDeviceConfigRequest) (*pb.DeviceConfig, error) {
	return &pb.DeviceConfig{}, nil
}
