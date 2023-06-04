package infrastructure

import (
	"github.com/yourusername/inventory-service/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormInventoryRepository struct {
	db *gorm.DB
}

func NewGormInventoryRepository(dsn string) (domain.InventoryRepository, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Inventory{})
	if err != nil {
		return nil, err
	}

	return &GormInventoryRepository{db: db}, nil
}

func (r *GormInventoryRepository) GetInventory(sku string, warehouseID int64, channel string) (*domain.Inventory, error) {
	var inventory domain.Inventory
	err := r.db.Where("sku = ? AND warehouse_id = ? AND channel = ?", sku, warehouseID, channel).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (r *GormInventoryRepository) UpdateInventory(sku string, warehouseID int64, channel string, quantity int32) (*domain.Inventory, error) {
	inventory, err := r.GetInventory(sku, warehouseID, channel)
	if err != nil {
		inventory = &domain.Inventory{
			SKU:         sku,
			WarehouseID: warehouseID,
			Channel:     channel,
			Quantity:    quantity,
		}
		err = r.db.Create(inventory).Error
	} else {
		inventory.Quantity = quantity
		err = r.db.Save(inventory).Error
	}

	if err != nil {
		return nil, err
	}
	return inventory, nil
}

func (r *GormInventoryRepository) ListInventories() ([]*domain.Inventory, error) {
	var inventories []*domain.Inventory
	err := r.db.Find(&inventories).Error
	if err != nil {
		return nil, err
	}
	return inventories, nil
}
