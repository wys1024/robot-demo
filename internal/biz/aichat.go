package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tmc/langchaingo/llms/ollama"
)

type Aichat struct {
	llm *ollama.LLM
}

type AichatRepo interface {
	Save(context.Context, *Aichat) (*Aichat, error)
	Update(context.Context, *Aichat) (*Aichat, error)
	FindByID(context.Context, int64) (*Aichat, error)
	ListByHello(context.Context, string) ([]*Aichat, error)
	ListAll(context.Context) ([]*Aichat, error)
}

type AichatUsecase struct {
	repo AichatRepo
	log  *log.Helper
}

// NewChat 创建聊天业务逻辑实例
func NewAichatUsecase(repo AichatRepo, logger log.Logger) *AichatUsecase {
	return &Aichat{}
}

// SendMessage 发送聊天消息
func (c *AichatUsecase) SendMessage(ctx context.Context, content string) (string, error) {
	c.log.WithContext(ctx).Infof("Received question: %s", content)
	// 实现LLM调用逻辑
	completion, err := c.llm.GenerateContent(ctx, content)
	if err != nil {
		c.log.Errorf("Failed to generate content: %v", err)
		return nil, err
	}

	return completion, nil
}
