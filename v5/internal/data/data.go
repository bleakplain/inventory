package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	"github.com/bleakplain/inventory/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

type Data struct {
	db *gorm.DB
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	conn, err := gorm.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}

	sqlDB := conn.DB()
	sqlDB.SetMaxIdleConns(int(c.Database.IdleConn))
	sqlDB.SetMaxOpenConns(int(c.Database.OpenConn))
	d := &Data{
		db: conn,
	}
	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			log.NewHelper(logger).Error("failed to close database", err)

		}
	}

	return d, cleanup, nil
}

