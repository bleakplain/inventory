package thrift

import (
	"context"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/yourusername/inventory-service/api/inventory/v1"
	"github.com/yourusername/inventory-service/internal/application"
	"github.com/yourusername/inventory-service/internal/domain"
)

type inventoryHandler struct {
	inventoryService *application.InventoryService
}

func NewInventoryHandler(app *domain.Application) *inventoryHandler {
	return &inventoryHandler{
		inventoryService: application.NewInventoryService(app),
	}
}

func (h *inventoryHandler) GetInventory(ctx context.Context, request *v1.GetInventoryRequest) (*v1.Inventory, error) {
	return h.inventoryService.GetInventory(ctx, request)
}

func (h *inventoryHandler) UpdateInventory(ctx context.Context, request *v1.UpdateInventoryRequest) (*v1.Inventory, error) {
	return h.inventoryService.UpdateInventory(ctx, request)
}

func (h *inventoryHandler) ListInventories(ctx context.Context) (*v1.ListInventoriesResponse, error) {
	return h.inventoryService.ListInventories(ctx)
}

func NewThriftServer(app *domain.Application) (*thrift.TSimpleServer, error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()

	handler := NewInventoryHandler(app)
	processor := v1.NewInventoryServiceProcessor(handler)

	transport, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		return nil, err
	}

	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return server, nil
}
