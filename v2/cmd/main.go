package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/yourusername/inventory-service/internal/application"
	"github.com/yourusername/inventory-service/internal/infrastructure"
	"github.com/yourusername/inventory-service/internal/transport"
)

func main() {
	// Load configuration
	c := config.New(
		config.WithSource(
			file.NewSource("config.yaml"),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	// Initialize logger
	var logger log.Logger
	if err := c.Scan(&logger); err != nil {
		panic(err)
	}

	// Initialize application services
	app, cleanup, err := infrastructure.InitApp(c)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// Initialize HTTP server
	httpSrv := http.NewServer(
		http.Address(app.HTTPAddr),
		http.WithLogger(logger),
	)
	transport.RegisterHTTPServer(httpSrv, application.NewInventoryService(app))

	// Initialize gRPC server
	grpcSrv := grpc.NewServer(
		grpc.Address(app.GRPCAddr),
		grpc.WithLogger(logger),
	)
	transport.RegisterGRPCServer(grpcSrv, application.NewInventoryService(app))

	// Create Kratos application
	kratosApp := kratos.New(
		kratos.Name("inventory-service"),
		kratos.Version("v1.0.0"),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
		kratos.WithLogger(logger),
	)

	// Run the application
	if err := kratosApp.Run(); err != nil {
		panic(err)
	}

	// Graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), app.ShutdownTimeout)
	defer cancel()

	if err := kratosApp.Stop(ctx); err != nil {
		panic(err)
	}
}
