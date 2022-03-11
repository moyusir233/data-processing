package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewConfigUsecase, NewWarningDetectUsecase)

// UnionRepo 方便wire注入而定义的repo合并接口
type UnionRepo interface {
	ConfigRepo
	WarningDetectRepo
}
