package biz

import (
	"context"
	"errors"
	"robot-demo/internal/conf"
	auth "robot-demo/internal/pkg/middleware"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID           uint
	Username     string
	PasswordHash string
}

type UserLogin struct {
	Username string
	Token    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByusername(ctx context.Context, username string) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	jwtc *conf.Jwt
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, jwtc *conf.Jwt, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, jwtc: jwtc, log: log.NewHelper(logger)}
}

// 生成token
func (uc *UserUsecase) generateToken(userID uint) string {
	return auth.GenerateToken(uc.jwtc.Secret, userID)
}

// Register 注册用户
func (uc *UserUsecase) Register(ctx context.Context, username string, password string) (*UserLogin, error) {
	// 记录日志
	uc.log.WithContext(ctx).Infof("注册用户: %s", username)

	// 创建用户
	user := &User{
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	// 检查该用户是否已存在,如果存在返回该用户已存在
	if _, err := uc.repo.GetUserByusername(ctx, username); err == nil {
		return nil, errors.New("该用户已存在")
	}

	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &UserLogin{
		Username: username,
		Token:    uc.generateToken(user.ID),
	}, nil

}

// Login 登录用户
func (uc *AichatUsecase) Login(ctx context.Context, username string, password string) error {
	// 记录日志
	uc.log.WithContext(ctx).Infof("登录用户: %s", username)

	// 验证用户
	// ...

	return nil
}
