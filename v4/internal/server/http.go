package server

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http/server"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/service"
	v1 "github.com/yourusername/inventory-service/api/inventory/v1"
)

// NewHTTPServer creates a new HTTP server instance.
func NewHTTPServer(c *conf.Server, inventory *service.InventoryService, logger log.Logger) *server.Server {
	var opts = []server.ServerOption{
		server.WithAddress(c.Http.Addr),
		server.WithLogger(logger),
		server.WithMiddleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	srv := server.NewServer(opts...)

	v1.RegisterInventoryHTTPServer(srv, inventory)

	return srv
}

// StartHTTPServer starts the HTTP server.
func StartHTTPServer(ctx context.Context, srv *server.Server) error {
	return srv.Start(ctx)
}

// StopHTTPServer stops the HTTP server.
func StopHTTPServer(ctx context.Context, srv *server.Server) error {
	return srv.Stop(ctx)
}
