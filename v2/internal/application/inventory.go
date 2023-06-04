package application

import (
	"context"

	"github.com/yourusername/inventory-service/api/inventory/v1"
	"github.com/yourusername/inventory-service/internal/domain"
)

type InventoryService struct {
	inventorySvc domain.InventoryService
}

func NewInventoryService(app *domain.Application) *InventoryService {
	return &InventoryService{
		inventorySvc: app.InventoryService,
	}
}

func (s *InventoryService) GetInventory(ctx context.Context, req *v1.GetInventoryRequest) (*v1.Inventory, error) {
	inventory, err := s.inventorySvc.GetInventory(ctx, req.Sku, req.WarehouseId, req.Channel)
	if err != nil {
		return nil, err
	}

	return &v1.Inventory{
		Sku:         inventory.Sku,
		WarehouseId: inventory.WarehouseID,
		Channel:     inventory.Channel,
		Quantity:    inventory.Quantity,
	}, nil
}

func (s *InventoryService) UpdateInventory(ctx context.Context, req *v1.UpdateInventoryRequest) (*v1.Inventory, error) {
	inventory, err := s.inventorySvc.UpdateInventory(ctx, req.Sku, req.WarehouseId, req.Channel, req.Quantity)
	if err != nil {
		return nil, err
	}

	return &v1.Inventory{
		Sku:         inventory.Sku,
		WarehouseId: inventory.WarehouseID,
		Channel:     inventory.Channel,
		Quantity:    inventory.Quantity,
	}, nil
}

func (s *InventoryService) ListInventories(ctx context.Context) (*v1.ListInventoriesResponse, error) {
	inventories, err := s.inventorySvc.ListInventories(ctx)
	if err != nil {
		return nil, err
	}

	apiInventories := make([]*v1.Inventory, len(inventories))
	for i, inventory := range inventories {
		apiInventories[i] = &v1.Inventory{
			Sku:         inventory.Sku,
			WarehouseId: inventory.WarehouseID,
			Channel:     inventory.Channel,
			Quantity:    inventory.Quantity,
		}
	}

	return &v1.ListInventoriesResponse{
		Inventories: apiInventories,
	}, nil
}
