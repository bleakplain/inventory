package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/thrift"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/server"
	"github.com/yourusername/inventory-service/internal/service"
	"github.com/yourusername/inventory-service/internal/wire"
)

func main() {
	flag.Parse()

	logger := log.NewStdLogger(os.Stdout)
	log := log.NewHelper(logger)

	c := config.New(
		config.WithSource(
			file.NewSource("config.yaml"),
		),
	)
	if err := c.Load(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		log.Fatalf("failed to scan config: %v", err)
	}

	app, cleanup, err := wire.InitApp(bc, logger)
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}
	defer cleanup()

	httpSrv := http.NewServer(http.Address(bc.Server.HTTP.Addr))
	grpcSrv := grpc.NewServer(grpc.Address(bc.Server.GRPC.Addr))
	thriftSrv := thrift.NewServer(thrift.Address(bc.Server.Thrift.Addr))

	srv := server.NewInventoryServer(app.InventoryService())
	httpSrv.HandlePrefix("/", srv)
	grpcSrv.Handle(srv)
	thriftSrv.Handle(srv)

	app.Run(
		kratos.Server(httpSrv),
		kratos.Server(grpcSrv),
		kratos.Server(thriftSrv),
	)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-ch

	log.Info("shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), bc.Server.ShutdownTimeout)
	defer cancel()

	if err := app.Stop(ctx); err != nil {
		log.Error(err)
	}
}
