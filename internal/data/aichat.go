package data

import (
	"context"

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

func (r *aichatRepo) Save(ctx context.Context, g *biz.Aichat) (*biz.Aichat, error) {
	return g, nil
}

func (r *aichatRepo) Update(ctx context.Context, g *biz.Aichat) (*biz.Aichat, error) {
	return g, nil
}

func (r *aichatRepo) FindByID(context.Context, int64) (*biz.Aichat, error) {
	return nil, nil
}

func (r *aichatRepo) ListByHello(context.Context, string) ([]*biz.Aichat, error) {
	return nil, nil
}

func (r *aichatRepo) ListAll(context.Context) ([]*biz.Aichat, error) {
	return nil, nil
}
