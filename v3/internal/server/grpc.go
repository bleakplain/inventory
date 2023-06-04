package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/yourusername/inventory-service/api/inventory/v1"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/service"
)

// NewGRPCServer creates a new gRPC server.
func NewGRPCServer(c *conf.ServerConfig, inventoryService *service.InventoryService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
		grpc.Address(c.GRPC.Addr),
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterInventoryServer(srv, inventoryService)
	return srv
}

// StartGRPCServer starts the gRPC server.
func StartGRPCServer(srv *grpc.Server) error {
	return srv.Start(context.Background())
}

// StopGRPCServer stops the gRPC server.
func StopGRPCServer(srv *grpc.Server) error {
	return srv.Stop(context.Background())
}
