package parser

import (
	"errors"
	v1 "gitee.com/moyusir/util/api/util/v1"
	"strconv"
)

// WarningFunc 预警函数，返回true时表示需要产生警告
type WarningFunc func(arg float64) bool

// GetWarningFunc 依据数据类型与比较规则构造相应参数的比较函数
func GetWarningFunc(fieldType v1.Type, cmp *v1.DeviceStateRegisterInfo_CmpRule) (WarningFunc, error) {
	// 由于redis ts存储的value为float64类型，
	// 因此这里统一将所有数据类型转换为float64进行比较
	float, err := strconv.ParseFloat(cmp.Arg, 64)
	if err != nil {
		return nil, err
	}
	switch cmp.Cmp {
	case v1.DeviceStateRegisterInfo_EQ:
		return func(arg float64) bool {
			return float == arg
		}, nil
	case v1.DeviceStateRegisterInfo_LT:
		return func(arg float64) bool {
			return arg < float
		}, nil
	case v1.DeviceStateRegisterInfo_GT:
		return func(arg float64) bool {
			return arg > float
		}, nil
	default:
		return nil, errors.New("unknown cmp rule")
	}
}
