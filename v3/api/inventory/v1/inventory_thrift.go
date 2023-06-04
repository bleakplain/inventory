package v1

import (
	"context"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/yourusername/inventory-service/internal/service"
)

type inventoryThriftServer struct {
	invSvc *service.InventoryService
}

func (s *inventoryThriftServer) GetInventory(ctx context.Context, sku string, warehouseID int64, channel string) (*Inventory, error) {
	inventory, err := s.invSvc.GetInventory(ctx, sku, uint64(warehouseID), channel)
	if err != nil {
		return nil, err
	}
	return &Inventory{
		Sku:         inventory.Sku,
		WarehouseID: int64(inventory.WarehouseID),
		Channel:     inventory.Channel,
		Quantity:    inventory.Quantity,
	}, nil
}

func (s *inventoryThriftServer) UpdateInventory(ctx context.Context, sku string, warehouseID int64, channel string, quantity int64) (bool, error) {
	err := s.invSvc.UpdateInventory(ctx, sku, uint64(warehouseID), channel, quantity)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewInventoryThriftServer(invSvc *service.InventoryService) *inventoryThriftServer {
	return &inventoryThriftServer{
		invSvc: invSvc,
	}
}

func RegisterInventoryThriftServer(processor *thrift.TMultiplexedProcessor, invSvc *service.InventoryService) {
	server := NewInventoryThriftServer(invSvc)
	processor.RegisterProcessor("InventoryService", NewInventoryServiceProcessor(server))
}
