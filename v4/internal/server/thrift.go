package server

import (
	"context"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/thrift"
	"github.com/go-kratos/kratos/v2/transport/thrift/proto/inventory"
	"github.com/go-kratos/kratos/v2/transport/thrift/server"
	"inventory/internal/conf"
	"inventory/internal/service"
)

type thriftServer struct {
	*server.Server
}

func NewThriftServer(c *conf.Config, inventoryService *service.InventoryService, logger log.Logger) *thriftServer {
	handler := &inventoryHandler{
		inventoryService: inventoryService,
	}
	processor := inventory.NewInventoryServiceProcessor(handler)

	serverOpts := []server.Option{
		server.WithAddress(c.Thrift.Addr),
		server.WithLogger(logger),
	}

	srv := server.NewServer(processor, serverOpts...)
	return &thriftServer{srv}
}

type inventoryHandler struct {
	inventoryService *service.InventoryService
}

func (h *inventoryHandler) GetInventory(ctx context.Context, sku string, warehouseID int64, channel string) (*inventory.Inventory, error) {
	inventory, err := h.inventoryService.GetInventory(ctx, sku, warehouseID, channel)
	if err != nil {
		return nil, err
	}

	return &inventory.Inventory{
		Sku:         inventory.Sku,
		WarehouseID: inventory.WarehouseID,
		Channel:     inventory.Channel,
		Quantity:    inventory.Quantity,
	}, nil
}

func (h *inventoryHandler) UpdateInventory(ctx context.Context, sku string, warehouseID int64, channel string, quantity int64) (bool, error) {
	err := h.inventoryService.UpdateInventory(ctx, sku, warehouseID, channel, quantity)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *thriftServer) Start(ctx context.Context) error {
	return s.Server.Start(ctx)
}

func (s *thriftServer) Stop(ctx context.Context) error {
	return s.Server.Stop(ctx)
}

func init() {
	thrift.RegisterTransport()
}
