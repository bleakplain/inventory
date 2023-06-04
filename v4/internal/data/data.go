package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/gorm"
	"github.com/yourusername/inventory-service/internal/conf"
)

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	db, err := initDB(c.Database)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.DefaultLogger.Error("failed to close database connection", "error", err)
		}
	}

	return &Data{db: db}, cleanup, nil
}

func (d *Data) WithContext(ctx context.Context) *Data {
	return &Data{
		db: d.db.WithContext(ctx),
	}
}

func initDB(c *conf.Database) (*gorm.DB, error) {
	db, err := gorm.Open(c.Driver, c.Source)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	db.DB().SetConnMaxLifetime(c.ConnMaxLifetime)

	return db, nil
}
