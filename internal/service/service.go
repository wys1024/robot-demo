package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "robot-demo/api/aichat/v1"
	"robot-demo/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAichatService)

// ChatService 聊天服务
type AichatService struct {
	pb.UnimplementedAichatServer
	aiChat *biz.AichatUsecase
	user   *biz.UserUsecase
	log    *log.Helper
}

func NewAichatService(aiChat *biz.AichatUsecase, user *biz.UserUsecase, logger log.Logger) *AichatService {
	return &AichatService{
		aiChat: aiChat,
		user:   user,
		log:    log.NewHelper(log.With(logger, "module", "service/aichat")),
	}
}
