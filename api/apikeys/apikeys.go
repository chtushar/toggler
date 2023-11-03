package apikeys

import (
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"

	"github.com/chtushar/toggler/configs"
)

var (
	cfg = configs.Get()
)

func parseRSAPrivateKeyFromPEM(pemStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

type GenerateAPIKeyParams struct {
	allowed_domains []string
}

func generateAPIKey(config GenerateAPIKeyParams) (*string, error) {
	domainsStr := strings.Join(config.allowed_domains, ",")

	h := hmac.New(sha256.New, []byte(cfg.PrivateKey))

	h.Write([]byte(domainsStr))
	apiKey := base64.StdEncoding.EncodeToString([]byte(domainsStr)) + "." + base64.StdEncoding.EncodeToString(h.Sum(nil))
	return &apiKey, nil
}
