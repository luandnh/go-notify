package model

import "time"

type Client struct {
	ApplicationId string                       `json:"application_id" db:"application_id"`
	UserId        string                       `json:"user_id" db:"user_id"`
	ClientId      string                       `json:"client_id" db:"client_id"`
	Name          string                       `json:"name" db:"name"`
	ClientToken   string                       `json:"client_token" db:"client_token"`
	Extras        map[string]map[string]string `json:"extras" db:"extras"`
	IsActive      bool                         `json:"is_active" db:"is_active"`
	IsDeleted     bool                         `json:"is_deleted" db:"is_deleted"`
	CreatedAt     time.Time                    `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time                    `json:"updated_at" db:"updated_at"`
	ExpiredAt     time.Time                    `json:"expired_at" db:"expired_at"`
}
