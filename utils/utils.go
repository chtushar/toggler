package utils

import (
	"crypto/rand"
	"encoding/base64"
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

func GenerateAPIKey () (string, error) {
	const apiKeyLength = 28 // The desired length of the API key in bytes
	const prefix = "tog_"
	apiKeyBytes := make([]byte, apiKeyLength)

	_, err := rand.Read(apiKeyBytes)
	if err != nil {
		return "", err
	}

	apiKey := base64.URLEncoding.EncodeToString(apiKeyBytes)

	return prefix + apiKey, nil
}