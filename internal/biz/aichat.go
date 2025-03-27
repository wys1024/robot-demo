package biz

import (
	"context"
	"robot-demo/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

type AichatRepo interface {
	// 定义与AichatRepo相关的方法

}
type AichatUsecase struct {
	repo AichatRepo
	llm  *ollama.LLM
	log  *log.Helper
}

// NewChat 创建聊天业务逻辑实例
func NewAichatUsecase(repo AichatRepo, conf *conf.Llm, logger log.Logger) *AichatUsecase {
	llm, err := ollama.New(
		ollama.WithModel(conf.Model),
		ollama.WithServerURL(conf.ApiHost),
	)
	if err != nil {
		return nil
	}

	return &AichatUsecase{
		repo: repo,
		llm:  llm,
		log:  log.NewHelper(logger),
	}
}

// SendMessage 发送聊天消息
func (c *AichatUsecase) SendMessage(ctx context.Context, content string, streamCallback func(ctx context.Context, chunk []byte) error) error {
	// 记录日志
	c.log.WithContext(ctx).Infof("收到问题: %s", content)

	// 按照ChatHandler的方式构造消息内容
	msgContent := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, content),
	}

	// 调用LLM生成内容，传入流式回调函数
	_, err := c.llm.GenerateContent(
		ctx,
		msgContent,
		llms.WithStreamingFunc(streamCallback), // 使用相同的流式参数
	)

	if err != nil {
		c.log.Errorf("生成内容失败: %v", err)
		return err
	}

	return nil
}
