package model

import "time"

type Application struct {
	ApplicationId    string    `json:"application_id" db:"application_id"`
	ApplicationName  string    `json:"application_name" db:"application_name"`
	ApplicationToken string    `json:"application_token" db:"application_token"`
	Description      string    `json:"description" db:"description"`
	IsActive         bool      `json:"is_active" db:"is_active"`
	IsDeleted        bool      `json:"is_deleted" db:"is_deleted"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	ExpiredAt        time.Time `json:"expired_at" db:"expired_at"`
}
