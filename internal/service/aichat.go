package service

import (
	"context"
	pb "robot-demo/api/aichat/v1"
)

// SendMessage 实现聊天消息发送接口
func (s *AichatService) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	s.log.WithContext(ctx).Infof("Received question: %s", req.Question)

	// 创建一个channel用于接收流式响应
	responseChan := make(chan string, 100)
	errChan := make(chan error, 1)

	// 异步调用流式生成
	go func() {
		err := s.aiChat.SendMessage(ctx, req.Question, func(ctx context.Context, chunk []byte) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case responseChan <- string(chunk):
				return nil
			}
		})
		if err != nil {
			errChan <- err
		}
		close(responseChan)
	}()

	// 收集所有响应
	var fullResponse string
	for {
		select {
		case chunk, ok := <-responseChan:
			if !ok {
				// 通道已关闭，发送完整响应
				return &pb.SendMessageResponse{
					Answer: fullResponse,
				}, nil
			}
			fullResponse += chunk
		case err := <-errChan:
			s.log.Errorf("Failed to generate response: %v", err)
			return nil, err
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
