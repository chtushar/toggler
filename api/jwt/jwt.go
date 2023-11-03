package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)



func GenerateToken(claims jwt.MapClaims, JWTSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWTSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
