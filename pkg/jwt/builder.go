package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type builder struct {
	key []byte
}

func NewToken(key []byte) *builder {
	return &builder{
		key: key,
	}
}

func (b *builder) generateToken(userID string, subject string) (string, error) {
	customClaims := &customClaims{
		userID,
		jwt.RegisteredClaims{
			Issuer:    "Издатель токена",
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	//key := []byte("i52GpyuGaN.QMreM7V09f.l3sUPoUXNI")
	return token.SignedString(b.key)
}
