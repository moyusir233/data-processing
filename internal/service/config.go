package service

import (
	"context"
	"gitee.com/moyusir/data-processing/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
)

type ConfigService struct {
	pb.UnimplementedConfigServer
	configUsecase *biz.ConfigUsecase
	logger        *log.Helper
}

func NewConfigService(cu *biz.ConfigUsecase, logger log.Logger) *ConfigService {
	return &ConfigService{
		configUsecase: cu,
		logger:        log.NewHelper(logger),
	}
}

func (s *ConfigService) GetDeviceConfig0(ctx context.Context, req *pb.GetDeviceConfigRequest) (*pb.DeviceConfig0, error) {
	// 构建查询选项
	info := &biz.DeviceGeneralInfo{
		DeviceID: req.DeviceId,
		// 代码注入deviceClassID
		DeviceClassID: 0,
	}

	config := new(pb.DeviceConfig0)
	err := s.configUsecase.GetDeviceConfig(info, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (s *ConfigService) GetDeviceConfig1(ctx context.Context, req *pb.GetDeviceConfigRequest) (*pb.DeviceConfig1, error) {
	// 构建查询选项
	info := &biz.DeviceGeneralInfo{
		DeviceID: req.DeviceId,
		// 代码注入deviceClassID
		DeviceClassID: 1,
	}

	config := new(pb.DeviceConfig1)
	err := s.configUsecase.GetDeviceConfig(info, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
