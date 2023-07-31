package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID               string `json:"user_id"`
	Role                 string `json:"role"`
	jwt.RegisteredClaims `json:"registered_claims"`
}

func (c *CustomClaims) Validate() error {
	if (c.Role != "user") && (c.Role != "admin") {
		return errors.New("must be user or admin") // TODO create new error
	}

	return nil
}
