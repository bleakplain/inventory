package domain

type inventoryService struct {
	repo InventoryRepository
}

func NewInventoryService(repo InventoryRepository) InventoryService {
	return &inventoryService{
		repo: repo,
	}
}

func (s *inventoryService) GetInventory(sku string, warehouseID int64, channel string) (*Inventory, error) {
	return s.repo.GetInventory(sku, warehouseID, channel)
}

func (s *inventoryService) UpdateInventory(sku string, warehouseID int64, channel string, quantity int32) (*Inventory, error) {
	return s.repo.UpdateInventory(sku, warehouseID, channel, quantity)
}

func (s *inventoryService) ListInventories() ([]*Inventory, error) {
	return s.repo.ListInventories()
}
