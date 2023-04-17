package model

import "time"

type Message struct {
	ApplicationId string            `json:"application_id" db:"application_id"`
	UserId        string            `json:"user_id" db:"user_id"`
	MessageId     string            `json:"message_id" db:"message_id"`
	Title         string            `json:"title" db:"title"`
	Content       string            `json:"content" db:"content"`
	Priority      int               `json:"priority" db:"priority"`
	Extras        map[string]string `json:"extras" db:"extras"`
	IsActive      bool              `json:"is_active" db:"is_active"`
	IsDeleted     bool              `json:"is_deleted" db:"is_deleted"`
	CreatedAt     time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at" db:"updated_at"`
}
