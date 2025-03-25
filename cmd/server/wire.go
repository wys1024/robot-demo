//go:build wireinject
// +build wireinject

package main

import (
	"robot-demo/internal/biz"
	"robot-demo/internal/conf"
	"robot-demo/internal/data"
	"robot-demo/internal/server"
	"robot-demo/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}