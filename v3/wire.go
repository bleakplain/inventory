//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/data"
	"github.com/yourusername/inventory-service/internal/server"
	"github.com/yourusername/inventory-service/internal/service"
)

func initApp(*conf.Config) (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
