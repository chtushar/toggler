package cmd

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func generateToken(Id int32, Uuid string, Email string, Name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    Id,
		"uuid":  Uuid,
		"email": Email,
		"name":  Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
