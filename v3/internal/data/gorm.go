package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yourusername/inventory-service/internal/conf"
)

// NewDB returns a new gorm.DB instance.
func NewDB(c *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", c.Database.DataSource)
	if err != nil {
		return nil, err
	}

	// Set up database settings
	db.DB().SetMaxIdleConns(c.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.Database.MaxOpenConns)
	db.DB().SetConnMaxLifetime(c.Database.ConnMaxLifetime)

	// Enable Logger
	if c.Database.LogMode {
		db.LogMode(true)
	}

	// AutoMigrate models
	db.AutoMigrate(&model.Inventory{})

	return db, nil
}
