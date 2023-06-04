package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yourusername/inventory-service/internal/data"
)

type InventoryService struct {
	log *log.Helper
	data *data.Data
}

func NewInventoryService(logger log.Logger, d *data.Data) *InventoryService {
	return &InventoryService{
		log:  log.NewHelper(logger),
		data: d,
	}
}

func (s *InventoryService) CreateInventory(ctx context.Context, sku string, warehouseID uint64, channel string, quantity int) (*data.Inventory, error) {
	inventory := &data.Inventory{
		Sku:        sku,
		WarehouseID: warehouseID,
		Channel:    channel,
		Quantity:   quantity,
	}
	err := s.data.CreateInventory(inventory)
	return inventory, err
}

func (s *InventoryService) UpdateInventory(ctx context.Context, id uint64, sku string, warehouseID uint64, channel string, quantity int) (*data.Inventory, error) {
	inventory := &data.Inventory{
		ID:         id,
		Sku:        sku,
		WarehouseID: warehouseID,
		Channel:    channel,
		Quantity:   quantity,
	}
	err := s.data.UpdateInventory(inventory)
	return inventory, err
}

func (s *InventoryService) DeleteInventory(ctx context.Context, id uint64) error {
	return s.data.DeleteInventory(id)
}

func (s *InventoryService) GetInventory(ctx context.Context, id uint64) (*data.Inventory, error) {
	return s.data.GetInventory(id)
}

func (s *InventoryService) ListInventories(ctx context.Context, sku string, warehouseID uint64, channel string) ([]*data.Inventory, error) {
	return s.data.ListInventories(sku, warehouseID, channel)
}
