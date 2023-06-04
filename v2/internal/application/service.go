package application

import (
	"github.com/yourusername/yourproject/internal/domain"
)

type InventoryService interface {
	GetInventory(sku string, warehouseID int64, channel string) (*domain.Inventory, error)
	UpdateInventory(sku string, warehouseID int64, channel string, quantity int32) (*domain.Inventory, error)
	ListInventories() ([]*domain.Inventory, error)
}

type inventoryService struct {
	domainService domain.InventoryService
}

func NewInventoryService(domainService domain.InventoryService) InventoryService {
	return &inventoryService{
		domainService: domainService,
	}
}

func (s *inventoryService) GetInventory(sku string, warehouseID int64, channel string) (*domain.Inventory, error) {
	return s.domainService.GetInventory(sku, warehouseID, channel)
}

func (s *inventoryService) UpdateInventory(sku string, warehouseID int64, channel string, quantity int32) (*domain.Inventory, error) {
	return s.domainService.UpdateInventory(sku, warehouseID, channel, quantity)
}

func (s *inventoryService) ListInventories() ([]*domain.Inventory, error) {
	return s.domainService.ListInventories()
}
