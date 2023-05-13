package cmd

import (
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
