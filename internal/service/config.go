package service

import (
	"context"
	"gitee.com/moyusir/data-processing/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "gitee.com/moyusir/data-processing/api/dataProcessing/v1"
	utilApi "gitee.com/moyusir/util/api/util/v1"
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

func (s *ConfigService) GetDeviceConfig(ctx context.Context, req *pb.GetDeviceConfigRequest) (*utilApi.TestedDeviceConfig, error) {
	// 构建查询选项
	info := &biz.DeviceGeneralInfo{
		DeviceID:      req.DeviceId,
		DeviceClassID: int(req.DeviceClassId),
	}

	config := new(utilApi.TestedDeviceConfig)
	err := s.configUsecase.GetDeviceConfig(info, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
