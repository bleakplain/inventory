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
	v1 "github.com/yourusername/inventory-service/api/inventory/v1"
)

// NewHTTPServer creates a new HTTP server.
func NewHTTPServer(c *conf.ServerConfig, inventoryService *service.InventoryService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
		http.Filter(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				h.ServeHTTP(w, r)
			})
		}),
	}
	mux := http.NewServeMux()
	mux.Handle("/inventory/", v1.NewInventoryHTTPServer(inventoryService))
	return http.NewServer(c.HTTP.Addr, mux, opts...)
}

// StartHTTPServer starts the HTTP server.
func StartHTTPServer(ctx context.Context, srv *http.Server, logger log.Logger) error {
	log := log.NewHelper(logger)
	log.Infof("HTTP server listening on %s", srv.Addr())
	return srv.Start(ctx)
}

// StopHTTPServer stops the HTTP server.
func StopHTTPServer(ctx context.Context, srv *http.Server, logger log.Logger) error {
	log := log.NewHelper(logger)
	log.Infof("HTTP server stopping")
	return srv.Stop(ctx)
}
