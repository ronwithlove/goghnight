package models

import (
	"time"
)

// Message 消息模型
type Message struct {
	ID        int       `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// MessageResponse API响应结构
type MessageResponse struct {
	Message string   `json:"message"`
	Status  string   `json:"status"`
	Data    *Message `json:"data,omitempty"`
}
