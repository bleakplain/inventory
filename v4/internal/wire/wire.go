package wire

import (
	"github.com/google/wire"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/thrift"
	"inventory/internal/conf"
	"inventory/internal/data"
	"inventory/internal/server"
	"inventory/internal/service"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	conf.ProviderSet,
	data.ProviderSet,
	server.ProviderSet,
	service.ProviderSet,
	NewApp,
)

type App struct {
	httpSrv *http.Server
	grpcSrv *grpc.Server
	thriftSrv *thrift.Server
}

func NewApp(httpSrv *http.Server, grpcSrv *grpc.Server, thriftSrv *thrift.Server, logger log.Logger) *App {
	app := &App{
		httpSrv: httpSrv,
		grpcSrv: grpcSrv,
		thriftSrv: thriftSrv,
	}
	return app
}

func (a *App) Start() error {
	errChan := make(chan error)
	go func() {
		if err := a.httpSrv.Start(); err != nil {
			errChan <- err
		}
	}()
	go func() {
		if err := a.grpcSrv.Start(); err != nil {
			errChan <- err
		}
	}()
	go func() {
		if err := a.thriftSrv.Start(); err != nil {
			errChan <- err
		}
	}()
	return <-errChan
}

func (a *App) Stop() error {
	if err := a.httpSrv.Stop(); err != nil {
		return err
	}
	if err := a.grpcSrv.Stop(); err != nil {
		return err
	}
	if err := a.thriftSrv.Stop(); err != nil {
		return err
	}
	return nil
}
