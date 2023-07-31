package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userID string, role string, subject string) (string, error) {

	customClaims := CustomClaims{
		userID,
		role,
		jwt.RegisteredClaims{
			Issuer:    "Издатель токена",
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	//TODO Create string
	return token.SignedString("")
}
