package biz

import (
	"fmt"
	"gitee.com/moyusir/data-processing/internal/conf"
	"github.com/google/wire"
	"strings"
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

// GetDeviceStateKey 以<用户id>:device_state:<设备类别号>为键，在zset中保存
// 以timestamp为score，以设备状态二进制protobuf信息为value的键值对
func GetDeviceStateKey(deviceClassID int) string {
	return fmt.Sprintf("%s:device_state:%d", conf.Username, deviceClassID)
}

// GetDeviceStateFieldKey
// 每个预警字段保存到以<用户id>:device_state:<设备类别号>:<设备字段名>:<设备id>为key的ts中
func GetDeviceStateFieldKey(info *DeviceGeneralInfo, fieldName string) string {
	return fmt.Sprintf(
		"%s:device_state:%d:%s:%s",
		conf.Username,
		info.DeviceClassID,
		fieldName,
		info.DeviceID,
	)
}

// GetDeviceStateFieldLabel 每个预警字段保存到以field_id为标签名，以<用户id>:<设备类别号>:<字段名>为标签值的ts中
func GetDeviceStateFieldLabel(deviceClassID int, fieldName string) string {
	return fmt.Sprintf("%s:%d:%s",
		conf.Username,
		deviceClassID,
		fieldName,
	)
}

// GetDeviceIDFromDeviceStateFieldKey 从设备预警字段使用的key中提取设备的key(<用户名>:<设备类别号>:<设备id>)
func GetDeviceIDFromDeviceStateFieldKey(key string) string {
	args := strings.Split(key, ":")
	return fmt.Sprintf("%s:%s:%s", args[0], args[2], args[4])
}

// GetWarningsSaveKey
// 用户的警告消息存放在以<用户id>:warning为key的ZSet中，
// 并以timestamp为field，以警告消息的序列化二进制数据为value存储警告消息
func GetWarningsSaveKey() string {
	return fmt.Sprintf("%s:warning", conf.Username)
}
