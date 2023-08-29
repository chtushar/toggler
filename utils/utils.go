package utils

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

func StringPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}

func YesNoPrompt(label string, def bool) bool {
    choices := "Y/n"
    if !def {
        choices = "y/N"
    }

    r := bufio.NewReader(os.Stdin)
    var s string

    for {
        fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)
        s, _ = r.ReadString('\n')
        s = strings.TrimSpace(s)
        if s == "" {
            return def
        }
        s = strings.ToLower(s)
        if s == "y" || s == "yes" {
            return true
        }
        if s == "n" || s == "no" {
            return false
        }
    }
}