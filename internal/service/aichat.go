package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	pb "robot-demo/api/aichat/v1"
	"robot-demo/internal/biz"
	"robot-demo/internal/conf"
)

// ChatService 聊天服务
type AichatService struct {
	pb.UnimplementedAichatServer
	llm  *ollama.LLM
	chat *biz.AichatUsecase
	log  *log.Helper
}

func NewAichatService(chat *biz.AichatUsecase) *AichatService {
	return &AichatService{
		chat: chat,
		log:  log.NewHelper(log.With(logger, "module", "service/aichat")),
	}
}

// InitLLM 初始化LLM客户端
func (s *AichatService) InitLLM(model string, apiHost string) error {
	var err error
	options := []ollama.Option{
		ollama.WithModel(model),
	}

	if apiHost != "" {
		options = append(options, ollama.WithServerURL(apiHost))
	}

	s.llm, err = ollama.New(options...)
	if err != nil {
		return fmt.Errorf("failed to initialize LLM: %v", err)
	}
	return nil
}

// SendMessage 实现聊天消息发送接口
func (s *AichatService) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	s.log.WithContext(ctx).Infof("Received question: %s", req.Question)

	// 初始化llm客户端
	if err := s.InitLLM(conf., conf.); err != nil {
		s.log.Errorf("Failed to initialize LLM: %v", err)
		return nil, err
	}

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, req.Question),
	}

	completion, err := s.chat.SendMessage(ctx, content)

	return &pb.SendMessageResponse{
		Answer: completion.Choices[0].Content,
	}, nil
}
