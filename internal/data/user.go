package data

import (
	"context"
	"fmt"
	"robot-demo/internal/biz"
	"robot-demo/model"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 创建新用户
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := model.User{
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
	// 写入db
	rv := r.data.db.Create(&user)

	// 写入redis
	set := r.data.redis.Set(ctx, "user:"+user.Username, user.ID, 0)
	if set.Err() != nil {
		return set.Err()
	}

	// 写入kafka
	err := r.data.kafkaProducer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(user.Username),
			Value: []byte(strconv.FormatUint(uint64(user.ID), 10)), // uint → string → []byte
		},
	)
	if err != nil {
		return err
	}

	// 读取kafka
	for {
		if m, err := r.data.kafkaConsumer.ReadMessage(ctx); err != nil {
			break
		} else {
			fmt.Printf("从Kafka读取到消息: Topic=%v, Partition=%v, Offset=%v, Key=%s, Value=%s\n",
				m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}

	// 写入elasticsearch
	_, err = r.data.esClient.Index().
		Index("users").
		BodyJson(user).
		Do(ctx)
	if err != nil {
		return err
	}
	return rv.Error
}

// 通过用户名查找用户

func (r *userRepo) GetUserByusername(ctx context.Context, username string) (*biz.User, error) {
	var user model.User
	rv := r.data.db.Where("username = ?", username).First(&user)
	if rv.Error != nil {
		return nil, rv.Error
	}
	return &biz.User{
		ID:           user.ID,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}, nil
}
