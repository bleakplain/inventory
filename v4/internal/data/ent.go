package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/gorm"
	"github.com/yourusername/inventory-service/internal/conf"
)

type Ent struct {
	data *Data
}

func NewEnt(data *Data) *Ent {
	return &Ent{data: data}
}

func (e *Ent) GetInventory(ctx context.Context, sku string, warehouseID int64, channel string) (*Inventory, error) {
	var inventory Inventory
	err := e.data.db.Where("sku = ? AND warehouse_id = ? AND channel = ?", sku, warehouseID, channel).First(&inventory).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &inventory, nil
}

func (e *Ent) UpdateInventory(ctx context.Context, inventory *Inventory) error {
	return e.data.db.Save(inventory).Error
}

func (e *Ent) CreateInventory(ctx context.Context, inventory *Inventory) error {
	return e.data.db.Create(inventory).Error
}

type Inventory struct {
	ID          int64  `gorm:"primary_key"`
	Sku         string `gorm:"type:varchar(255);not null"`
	WarehouseID int64  `gorm:"not null"`
	Channel     string `gorm:"type:varchar(255);not null"`
	Quantity    int64  `gorm:"not null"`
}

func (Inventory) TableName() string {
	return "inventory"
}

func initEnt(logger log.Logger, c *conf.Data) (db *gorm.DB, cleanup func(), err error) {
	db, err = gorm.Open("mysql", c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	db.LogMode(c.Database.Debug)

	if err = db.AutoMigrate(&Inventory{}).Error; err != nil {
		return nil, nil, err
	}

	cleanup = func() {
		db.Close()
	}
	return
}
