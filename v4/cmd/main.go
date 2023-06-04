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
	var (
		configFile = flag.String("conf", "config.yaml", "config file path")
	)
	flag.Parse()

	logger := log.NewStdLogger(os.Stdout)
	log := log.NewHelper(logger)

	c := config.New(
		config.WithSource(file.NewSource(*configFile)),
		config.WithDecoder(func(kv *config.KeyValue) (config.Value, error) {
			return config.NewValue(kv.Key, kv.Value), nil
		}),
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

	srv := server.NewMultiProtocolServer(app, httpSrv, grpcSrv, thriftSrv)

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	log.Infof("inventory-service is running on %s, %s, %s", bc.Server.HTTP.Addr, bc.Server.GRPC.Addr, bc.Server.Thrift.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	if err := srv.Stop(); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}
}
