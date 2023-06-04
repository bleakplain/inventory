package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yourusername/inventory-service/internal/data"
	"github.com/yourusername/inventory-service/internal/data/model"
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

func (s *InventoryService) GetInventory(ctx context.Context, sku string, warehouseID uint64, channel string) (*model.Inventory, error) {
	var inventory model.Inventory
	err := s.data.DB().Where("sku = ? AND warehouse_id = ? AND channel = ?", sku, warehouseID, channel).First(&inventory).Error
	if err != nil {
		s.log.Errorf("Failed to get inventory: %v", err)
		return nil, err
	}
	return &inventory, nil
}

func (s *InventoryService) UpdateInventory(ctx context.Context, sku string, warehouseID uint64, channel string, quantity int64) error {
	err := s.data.DB().Model(&model.Inventory{}).Where("sku = ? AND warehouse_id = ? AND channel = ?", sku, warehouseID, channel).Update("quantity", quantity).Error
	if err != nil {
		s.log.Errorf("Failed to update inventory: %v", err)
		return err
	}
	return nil
}
