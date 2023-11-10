package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id",`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Password string    `json:"password"`
	Token    *Token    `json:"token"`
}
