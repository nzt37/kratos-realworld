// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kartos-realworld/internal/biz"
	"kartos-realworld/internal/conf"
	"kartos-realworld/internal/data"
	"kartos-realworld/internal/server"
	"kartos-realworld/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
