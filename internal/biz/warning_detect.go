package biz

import "time"

const WarningDetectFieldLabelName = "field_id"

type WarningDetectUsecase struct {
}
type WarningDetectRepo interface {
	// BatchGetDeviceStateInfo 批量查询设备状态信息，返回保存在数据库中设备状态的二进制信息
	BatchGetDeviceStateInfo(key string, option *QueryOption) ([][]byte, error)
	// BatchGetDeviceWarningDetectField 批量查询具有指定label的TS中保存的预警字段值信息
	BatchGetDeviceWarningDetectField(label string, option *TSQueryOption) ([]interface{}, error)
	// GetWarningMessage 查询当前存储在数据库中的警告信息
	GetWarningMessage(key string, option *QueryOption) ([][]byte, error)
	// SaveWarningMessage 保存警告信息
	SaveWarningMessage(key string, timestamp int64, value []byte) error
}

// QueryOption 查询ZSet时可以附加的查询参数
type QueryOption struct {
	// 查询的时间范围，以时间戳表示
	Begin, End int64
	// 偏移量和记录数限制
	Offset, Count int64
}

// TSQueryOption 批量查询TS时可以附加的查询参数
type TSQueryOption struct {
	// 查询的时间范围，以时间戳表示
	Begin, End int64
	// 查询的聚合类型
	AggregationType string
	// 聚合查询时使用的时间间隔，以毫秒为单位
	TimeBucket time.Duration
}

func NewWarningDetectUsecase() *WarningDetectUsecase {
	return nil
}
