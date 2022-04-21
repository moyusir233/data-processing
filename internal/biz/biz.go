package biz

import (
	"fmt"
	"gitee.com/moyusir/data-processing/internal/conf"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewConfigUsecase, NewWarningDetectUsecase)

// UnionRepo 方便wire注入而定义的repo合并接口
type UnionRepo interface {
	ConfigRepo
	WarningDetectRepo
}

// DeviceGeneralInfo 检索设备信息时使用的基本字段
type DeviceGeneralInfo struct {
	DeviceID      string
	DeviceClassID int
}

// GetDeviceConfigKey 以<用户id>:device_config:<device_class_id>:hash为键
// ,以设备id为field,在redis hash中保存设备配置的protobuf二进制信息
func GetDeviceConfigKey(info *DeviceGeneralInfo) string {
	return fmt.Sprintf("%s:device_config:%d:hash", conf.Username, info.DeviceClassID)
}
