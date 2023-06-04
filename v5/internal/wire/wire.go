package wire

import (
	"github.com/google/wire"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/data"
	"github.com/yourusername/inventory-service/internal/server"
	"github.com/yourusername/inventory-service/internal/service"
)

func InitApp(*conf.Config) (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}

type App struct {
	Server  *server.Server
	Service *service.Service
	Data    *data.Data
}

func newApp(server *server.Server, service *service.Service, data *data.Data) *App {
	return &App{
		Server:  server,
		Service: service,
		Data:    data,
	}
}
