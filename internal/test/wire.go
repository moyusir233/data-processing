//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package test

import (
	"gitee.com/moyusir/data-processing/internal/biz"
	"gitee.com/moyusir/data-processing/internal/conf"
	"gitee.com/moyusir/data-processing/internal/data"
	"gitee.com/moyusir/data-processing/internal/server"
	"gitee.com/moyusir/data-processing/internal/service"
	utilApi "gitee.com/moyusir/util/api/util/v1"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, []utilApi.DeviceStateRegisterInfo, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
