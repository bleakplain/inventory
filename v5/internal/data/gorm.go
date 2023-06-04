package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yourusername/inventory-service/internal/conf"
)

func NewGormDB(cfg *conf.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(cfg.Driver, cfg.Source)
	if err != nil {
		return nil, err
	}

	// Set up database settings
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}
