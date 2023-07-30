package cmd

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func generateToken(Id int32, Uuid uuid.NullUUID, Email sql.NullString, Name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    Id,
		"uuid":  Uuid,
		"email": Email.String,
		"name":  Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWTSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
