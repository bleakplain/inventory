package domain

import (
	"errors"
	"sync"
)

type inMemoryInventoryRepository struct {
	inventories []*Inventory
	mu          sync.RWMutex
}

func NewInMemoryInventoryRepository() InventoryRepository {
	return &inMemoryInventoryRepository{
		inventories: make([]*Inventory, 0),
	}
}

func (r *inMemoryInventoryRepository) GetInventory(sku string, warehouseID int64, channel string) (*Inventory, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, inventory := range r.inventories {
		if inventory.SKU == sku && inventory.WarehouseID == warehouseID && inventory.Channel == channel {
			return inventory, nil
		}
	}

	return nil, errors.New("inventory not found")
}

func (r *inMemoryInventoryRepository) UpdateInventory(sku string, warehouseID int64, channel string, quantity int32) (*Inventory, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, inventory := range r.inventories {
		if inventory.SKU == sku && inventory.WarehouseID == warehouseID && inventory.Channel == channel {
			inventory.Quantity = quantity
			return inventory, nil
		}
	}

	newInventory := &Inventory{
		SKU:         sku,
		WarehouseID: warehouseID,
		Channel:     channel,
		Quantity:    quantity,
	}
	r.inventories = append(r.inventories, newInventory)
	return newInventory, nil
}

func (r *inMemoryInventoryRepository) ListInventories() ([]*Inventory, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.inventories, nil
}
