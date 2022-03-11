package biz

type ConfigUsecase struct {
}
type ConfigRepo interface {
	// GetDeviceConfig 查询保存在数据库中的设备配置的二进制信息
	GetDeviceConfig(key, field string) ([]byte, error)
}

func NewConfigUsecase() *ConfigUsecase {
	return nil
}
