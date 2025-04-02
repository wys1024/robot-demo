package service

import (
	"context"
	pb "robot-demo/api/aichat/v1"
)

// Register 实现注册接口
func (s *AichatService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.log.WithContext(ctx).Infof("Received username: %s", req.Username)
	register, err := s.user.Register(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		Token: register.Token,
	}, nil
}

// Login 实现登录接口
func (s *AichatService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.log.WithContext(ctx).Infof("Received username: %s", req.Username)
	//  TODO: Implement the Login function
	return &pb.LoginResponse{
		Token: "1234567890",
	}, nil
}
