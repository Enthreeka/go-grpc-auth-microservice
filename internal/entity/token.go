package entity

import (
	"time"
)

type Token struct {
	ID        string    `json:"id"`
	ExpiresAt time.Time `json:"expires_at"`
	UserID    string    `json:"user_id"`
}
