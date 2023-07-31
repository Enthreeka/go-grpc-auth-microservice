package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	UserID               string `json:"user_id"`
	Admin                bool   `json:"admin"`
	jwt.RegisteredClaims `json:"registered_claims"`
}

//func (c *customClaims) Validate() error {
//	if (c.Role != "user") && (c.Role != "admin") {
//		return errors.New("must be user or admin") // TODO create new error
//	}
//
//	return nil
//}
