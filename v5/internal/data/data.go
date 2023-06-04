package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/gorm"
	"github.com/yourusername/inventory-service/internal/conf"
)

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.DataConfig, logger log.Logger) (*Data, error) {
	db, err := gorm.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, err
	}

	return &Data{
		db: db,
	}, nil
}

func (d *Data) Close() error {
	return d.db.Close()
}

func (d *Data) DB() *gorm.DB {
	return d.db
}
