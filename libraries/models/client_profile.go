package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type ClientProfile struct {
	ID           uuid.UUID `json:"profile_id" gorm:"primary_key;type:uuid;"`
	DisplayName  string    `json:"display_name"`
	RoleName     string    `json:"role_name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at"`
	TokenType    string    `json:"token_type"`
	Domain       string    `json:"domain"`
	LoginedBy    string    `json:"logined_by"`
}
