package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"robot-demo/internal/biz"
)

type User struct {
	//gorm.Model
	ID           uint   `gorm:"primarykey"`
	Username     string `gorm:"type:varchar(255);uniqueIndex"`
	PasswordHash string `gorm:"type:varchar(255)"`
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 创建新用户
func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	user := User{
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
	rv := r.data.db.Create(&user)
	return rv.Error
}

// 通过用户名查找用户

func (r *userRepo) GetUserByusername(ctx context.Context, username string) (*biz.User, error) {
	var user User
	rv := r.data.db.Where("username = ?", username).First(&user)
	if rv.Error != nil {
		return nil, rv.Error
	}
	return &biz.User{
		ID:           user.ID,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}, nil
}
