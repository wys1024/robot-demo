package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Chat 聊天业务逻辑
type Chat struct {
	llmModel string
	apiHost  string
	log      *log.Helper
}

// NewChat 创建聊天业务逻辑实例
func NewChat(llmModel, apiHost string, logger log.Logger) *Chat {
	return &Chat{
		llmModel: llmModel,
		apiHost:  apiHost,
		log:      log.NewHelper(logger),
	}
}

// SendMessage 发送聊天消息
func (c *Chat) SendMessage(ctx context.Context, question string) (string, error) {
	c.log.WithContext(ctx).Infof("Received question: %s", question)
	// TODO: 实现LLM调用逻辑
	return "", nil
}
