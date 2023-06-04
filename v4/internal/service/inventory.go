package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/yourusername/inventory-service/internal/conf"
	"github.com/yourusername/inventory-service/internal/data"
)

type InventoryService struct {
	log *log.Helper
	data *data.Data
}

type InventoryParams struct {
	Sku      string
	WarehouseID int64
	Channel  string
}

type Inventory struct {
	Sku         string
	WarehouseID int64
	Channel     string
	Quantity    int
}

func NewInventoryService(c *conf.Service, logger log.Logger, d *data.Data) *InventoryService {
	return &InventoryService{
		log:  log.NewHelper(logger),
		data: d,
	}
}

func (s *InventoryService) GetInventory(ctx context.Context, params *InventoryParams) (*Inventory, error) {
	// Implement the logic to get inventory details based on the given parameters
	// using the data layer (s.data)
	inventory, err := s.data.GetInventory(ctx, params.Sku, params.WarehouseID, params.Channel)
	if err != nil {
		return nil, err
	}

	return &Inventory{
		Sku:         inventory.Sku,
		WarehouseID: inventory.WarehouseID,
		Channel:     inventory.Channel,
		Quantity:    inventory.Quantity,
	}, nil
}

func (s *InventoryService) UpdateInventory(ctx context.Context, params *InventoryParams, newQuantity int) error {
	// Implement the logic to update inventory details based on the given parameters
	// and new quantity using the data layer (s.data)
	return s.data.UpdateInventory(ctx, params.Sku, params.WarehouseID, params.Channel, newQuantity)
}
