package data

import (
	"robot-demo/internal/biz"

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
