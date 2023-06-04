package model

import (
	"time"
)

type Inventory struct {
	ID        uint64 `gorm:"primary_key"`
	Sku       string `gorm:"type:varchar(255);not null"`
	WarehouseID uint64 `gorm:"not null"`
	Channel   string `gorm:"type:varchar(255);not null"`
	Quantity  int64  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Inventory) TableName() string {
	return "inventory"
}
