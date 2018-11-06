package webapputil

import (
	"crypto/rand"
	"encoding/base64"
)

func secureRandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

//GenerateSecureRandom creates a base64 url enconding
func GenerateSecureRandom() (string, error) {
	byteLength := 32
	b, err := secureRandomBytes(byteLength)

	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
