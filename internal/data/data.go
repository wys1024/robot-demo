package data

import (
	"context"
	"github.com/olivere/elastic/v7"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"

	"robot-demo/internal/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAichatRepo, NewUserRepo)

// Data .
// TODO: 增加redis、kafka、elasticsearch等连接
type Data struct {
	db            *gorm.DB
	redis         *redis.Client
	kafkaProducer *kafka.Writer
	kafkaConsumer *kafka.Reader
	esClient      *elastic.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}

	// 初始化Redis连接
	redisClient := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Errorf("failed opening connection to redis: %v", err)
		return nil, nil, err
	}

	// 初始化Kafka连接
	kafkaProducer := &kafka.Writer{
		Addr:  kafka.TCP(c.Kafka.Brokers...),
		Topic: c.Kafka.Topic,
	}

	kafkaConsumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.Kafka.Brokers,
		Topic:   c.Kafka.Topic,
		GroupID: c.Kafka.GroupId,
	})

	// 初始化Elasticsearch客户端
	esClient, err := elastic.NewClient(
		elastic.SetURL(c.Elasticsearch.Addr...),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(true),
		elastic.SetHealthcheckTimeout(c.Elasticsearch.Timeout.AsDuration()),
		elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewExponentialBackoff(100, 5000))),
	)
	if err != nil {
		log.Errorf("failed opening connection to elasticsearch: %v", err)
		return nil, nil, err
	}
	
	// 执行健康检查
	ctx, cancel := context.WithTimeout(context.Background(), c.Elasticsearch.Timeout.AsDuration())
	defer cancel()
	
	_, _, err = esClient.Ping(c.Elasticsearch.Addr[0]).Do(ctx)
	if err != nil {
		log.Errorf("elasticsearch health check failed: %v", err)
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
		if kafkaProducer != nil {
			if err := kafkaProducer.Close(); err != nil {
				log.Error(err)
			}
		}
		if kafkaConsumer != nil {
			if err := kafkaConsumer.Close(); err != nil {
				log.Error(err)
			}
		}
		if esClient != nil {
			esClient.Stop()
		}
	}

	return &Data{
		db:            db,
		redis:         redisClient,
		kafkaProducer: kafkaProducer,
		kafkaConsumer: kafkaConsumer,
		esClient:      esClient,
	}, cleanup, nil
}
