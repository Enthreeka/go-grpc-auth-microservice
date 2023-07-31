package jwt

import "github.com/golang-jwt/jwt/v5"

// TODO создать ошибку
func ParseToken(accessToken string) (*customClaims, error) {
	token, _ := jwt.ParseWithClaims(accessToken, &customClaims{}, func(t *jwt.Token) (interface{}, error) {

		return nil, nil
	})

	if claims, ok := token.Claims.(*customClaims); ok {
		return claims, nil
	}

	return nil, nil
}
