package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/bleakplain/inventory/api/inventory/v1"
	"github.com/bleakplain/inventory/internal/data"
)

type InventoryService struct {
	v1.UnimplementedInventoryServiceServer

	log  *log.Helper
	data *data.Data
}

func NewInventoryService(d *data.Data, logger log.Logger) *InventoryService {
	return &InventoryService{
		log:  log.NewHelper(logger),
		data: d,
	}
}

func (s *InventoryService) UpdateInventory(ctx context.Context,
	req *v1.UpdateInventoryRequest) (*v1.Inventory, error) {
	err := s.data.UpdateInventory(&data.Inventory{
		ID:          req.GetId(),
		Sku:         req.GetSku(),
		WarehouseID: req.GetWarehouseId(),
		Channel:     req.GetChannel(),
		Quantity:    int(req.GetQuantity()),
	})
	if err != nil {
		return nil, err
	}
	return &v1.Inventory{
		Id:          req.GetId(),
		Sku:         req.GetSku(),
		WarehouseId: req.GetWarehouseId(),
		Channel:     req.GetChannel(),
		Quantity:    req.GetQuantity(),
	}, nil
}

func (s *InventoryService) GetInventory(ctx context.Context,
	req *v1.GetInventoryRequest) (*v1.Inventory, error) {
	inventory, err := s.data.GetInventory(req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.Inventory{
		Id:          inventory.ID,
		Sku:         inventory.Sku,
		WarehouseId: inventory.WarehouseID,
		Channel:     inventory.Channel,
		Quantity:    int32(inventory.Quantity),
	}, nil
}

func (s *InventoryService) ListInventories(ctx context.Context,
	req *v1.ListInventoriesRequest) (*v1.ListInventoriesResponse, error) {
	inventories, err := s.data.ListInventories(req.GetSku(), req.GetWarehouseId(), req.GetChannel())
	if err != nil {
		return nil, err
	}
	var list []*v1.Inventory
	for _, inventory := range inventories {
		list = append(list, &v1.Inventory{
			Id:          inventory.ID,
			Sku:         inventory.Sku,
			WarehouseId: inventory.WarehouseID,
			Channel:     inventory.Channel,
			Quantity:    int32(inventory.Quantity),
		})
	}
	return &v1.ListInventoriesResponse{Inventories: list}, nil
}
