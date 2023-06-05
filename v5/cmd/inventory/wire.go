//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/bleakplain/inventory/internal/conf"
	"github.com/bleakplain/inventory/internal/data"
	"github.com/bleakplain/inventory/internal/server"
	"github.com/bleakplain/inventory/internal/service"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
