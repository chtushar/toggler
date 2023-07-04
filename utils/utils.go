package utils

import (
	"encoding/json"
	"fmt"

	"github.com/jackc/pgtype"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateEmptyJSONB() pgtype.JSONB {
	jsonb := &pgtype.JSONB{}

	err := json.Unmarshal([]byte(`{}`), jsonb)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return *jsonb
}