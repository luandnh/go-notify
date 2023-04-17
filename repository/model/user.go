package model

import "time"

type User struct {
	ApplicationId string         `json:"application_id" db:"application_id"`
	UserId        string         `json:"user_id" db:"user_id"`
	RefUserId     string         `json:"ref_user_id" db:"ref_user_id"`
	Username      string         `json:"username" db:"username"`
	Password      string         `json:"password" db:"password"`
	UserToken     string         `json:"user_token" db:"user_token"`
	Level         string         `json:"level" db:"level"`
	Extras        map[string]any `json:"extras" db:"extras"`
	IsActive      bool           `json:"is_active" db:"is_active"`
	IsDeleted     bool           `json:"is_deleted" db:"is_deleted"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
	ExpiredAt     time.Time      `json:"expired_at" db:"expired_at"`
}
