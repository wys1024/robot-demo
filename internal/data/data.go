package data

import (
	"robot-demo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Data ..
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData ..
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "data"))

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Errorf("failed to connect database: %v", err)
		return nil, nil, err
	}

	cleanup := func() {
		log.Info("closing the data resources")
		// 在这里添加清理资源的代码
	}

	return &Data{
		db:  db,
		log: log,
	}, cleanup, nil
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)
