package cmd

import (
	"fmt"
	"time"

	"github.com/chtushar/toggler/db/queries"
	"github.com/golang-jwt/jwt"
)

func generateToken(Id int32, Email string, Name string, Role queries.UserRole) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    Id,
		"email": Email,
		"name":  Name,
		"role":  Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return cfg.JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
