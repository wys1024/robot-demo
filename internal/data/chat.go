package data

import (
	"time"

	"robot-demo/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// chatRepo 聊天数据仓库实现
type chatRepo struct {
	data *Data
	log  *log.Helper
}

// NewChatRepo 创建聊天数据仓库实例
func NewChatRepo(data *Data, logger log.Logger) biz.ChatRepo {
	return &chatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateMessage 创建聊天消息
func (r *chatRepo) CreateMessage(message *biz.Message) error {
	r.log.Info("Creating message")
	return r.data.db.Create(message).Error
}

// CreateSession 创建聊天会话
func (r *chatRepo) CreateSession(session *biz.ChatSession) error {
	r.log.Info("Creating chat session")
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()
	return r.data.db.Create(session).Error
}

// GetSessionMessages 获取会话消息列表
func (r *chatRepo) GetSessionMessages(sessionID uint64) ([]*biz.Message, error) {
	r.log.Info("Getting session messages")
	var messages []*biz.Message
	err := r.data.db.Where("session_id = ?", sessionID).Find(&messages).Error
	return messages, err
}

// GetSessions 获取会话列表
func (r *chatRepo) GetSessions() ([]*biz.ChatSession, error) {
	r.log.Info("Getting chat sessions")
	var sessions []*biz.ChatSession
	err := r.data.db.Find(&sessions).Error
	return sessions, err
}
