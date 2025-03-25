package biz

import "time"

// Message 聊天消息模型
type Message struct {
	ID        uint64    `json:"id"`
	SessionID uint64    `json:"session_id"`
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// ChatSession 聊天会话模型
type ChatSession struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatRepo 聊天数据仓库接口
type ChatRepo interface {
	CreateMessage(message *Message) error
	CreateSession(session *ChatSession) error
	GetSessionMessages(sessionID uint64) ([]*Message, error)
	GetSessions() ([]*ChatSession, error)
}
