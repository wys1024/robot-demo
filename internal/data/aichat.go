package data

import (
	"context"
	"robot-demo/internal/biz"
	"robot-demo/model"

	"github.com/go-kratos/kratos/v2/log"
)

type aichatRepo struct {
	data *Data
	log  *log.Helper
}

func NewAichatRepo(data *Data, logger log.Logger) biz.AichatRepo {
	return &aichatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 添加聊天记录
func (r *aichatRepo) Save(ctx context.Context, aichat *biz.Aichat) (*biz.Aichat, error) {
	r.log.WithContext(ctx).Infof("添加聊天记录: %v", aichat)

	result := r.data.db.Create(&model.Aichats{
		Id:        aichat.Id,
		Model:     aichat.Model,
		Questions: aichat.Questions,
		Answers:   aichat.Answers,
		Time:      aichat.Time,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return aichat, nil
}
