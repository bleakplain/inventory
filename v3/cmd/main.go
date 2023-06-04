package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/server"
	"github.com/yourusername/inventory-service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp loads the configuration, initializes the services, and returns the application.
func initApp(*conf.Config, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}

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
		log.Fatalf("failed to load configuration: %v", err)
	}

	var cfg conf.Config
	if err := c.Scan(&cfg); err != nil {
		log.Fatalf("failed to scan configuration: %v", err)
	}

	app, cleanup, err := initApp(&cfg, log)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}
	defer cleanup()

	// Wait for the interrupt signal.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	log.Infof("SIGNAL %d received, then shutting down...", <-ch)

	// Shutdown the application.
	if err := app.Stop(context.Background()); err != nil {
		log.Errorf("failed to stop application: %v", err)
	}
}
