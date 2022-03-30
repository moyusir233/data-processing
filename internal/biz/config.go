package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/proto"
)

type ConfigUsecase struct {
	repo   UnionRepo
	logger *log.Helper
}

type ConfigRepo interface {
	// GetDeviceConfig 查询保存在数据库中的设备配置的二进制信息
	GetDeviceConfig(key, field string) ([]byte, error)
}

func NewConfigUsecase(repo UnionRepo, logger log.Logger) *ConfigUsecase {
	return &ConfigUsecase{
		repo:   repo,
		logger: log.NewHelper(logger),
	}
}

// GetDeviceConfig 查询设备的配置信息，并解析到传入的proto message结构当中
func (u *ConfigUsecase) GetDeviceConfig(info *DeviceGeneralInfo, config proto.Message) error {
	// 以<用户id>:device_config:<device_class_id>:hash为键
	// ,以设备id为field,在redis hash中保存设备配置的protobuf二进制信息

	// 获得redis key
	key := GetDeviceConfigKey(info)

	// 调用repo层函数，查询信息
	deviceConfig, err := u.repo.GetDeviceConfig(key, info.DeviceID)
	if err != nil {
		return err
	}

	// 调用proto包的反序列化函数，将查询得到的二进制信息转换为proto message
	err = proto.Unmarshal(deviceConfig, config)
	if err != nil {
		return errors.Newf(
			500, "Biz_Config_Error",
			"反序列化设备配置的proto二进制信息时发生了错误%v", err)
	}

	return nil
}
