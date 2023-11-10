package entity

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	UserID       uuid.UUID `json:"user_id"`
	ExpiresAt    time.Time `json:"expires_at"`
	ID           string    `json:"id"`
	RefreshToken string    `json:"refresh_token"`
	Fingerprint  string    `json:"fingerprint"`
}
