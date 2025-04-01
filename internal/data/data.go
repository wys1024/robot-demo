package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/go-redis/redis/v8"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"robot-demo/internal/conf"

)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAichatRepo, NewUserRepo)

// Data .
type Data struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Errorf("failed opening connection to redis: %v", err)
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
		if err := redisClient.Close(); err != nil {
			log.Error(err)
		}
	}
	return &Data{
		db:    db,
		redis: redisClient,
	}, cleanup, nil
}
