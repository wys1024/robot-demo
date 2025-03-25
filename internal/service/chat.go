package service

import (
	"context"
	"fmt"
	pb "robot-demo/api/chat/v1"
	"robot-demo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// ChatService 聊天服务实现
type ChatService struct {
	pb.UnimplementedChatServer

	llm  *ollama.LLM
	chat *biz.Chat
	log  *log.Helper
}

// NewChatService 创建聊天服务实例
func NewChatService(chat *biz.Chat, logger log.Logger) *ChatService {
	return &ChatService{
		chat: chat,
		log:  log.NewHelper(logger),
	}
}

// InitLLM 初始化LLM客户端
func (s *ChatService) InitLLM(model string, apiHost string) error {
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
func (s *ChatService) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	s.log.WithContext(ctx).Infof("Received question: %s", req.Question)

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, req.Question),
	}

	completion, err := s.llm.GenerateContent(ctx, content)
	if err != nil {
		s.log.Errorf("Failed to generate content: %v", err)
		return nil, err
	}

	return &pb.SendMessageResponse{
		Answer: completion.Choices[0].Content,
	}, nil
}
