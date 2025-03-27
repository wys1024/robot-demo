package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"robot-demo/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAichatRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}

	cleanup := func() {
		log.Info("closing the data resources")
		sqlDB, err := db.DB()
		if err != nil {
			log.Error(err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Error(err)
		}
	}
	return &Data{
		db: db,
	}, cleanup, nil
}
