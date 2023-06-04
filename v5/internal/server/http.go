package server

import (
	"context"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http/server"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/service"
)

func NewHTTPServer(c *conf.ServerConfig, inventorySvc *service.InventoryService, logger log.Logger) *server.Server {
	var opts = []server.ServerOption{
		server.WithAddress(c.HTTP.Addr),
		server.WithLogger(logger),
		server.WithMiddleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}

	srv := server.NewServer(opts...)

	httpHandler := http.NewServeMux()
	httpHandler.HandleFunc("/inventory", func(w http.ResponseWriter, r *http.Request) {
		// Implement your HTTP handler logic here
	})

	srv.HandlePrefix("/", httpHandler)

	return srv
}
