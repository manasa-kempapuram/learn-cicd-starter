package auth

import (
	"errors"
	"net/http"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing Authorization header")
	}

	// Assuming header format: "ApiKey <key>"
	const prefix = "ApiKey "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		return "", errors.New("malformed Authorization header")
	}

	return authHeader[len(prefix):], nil
}

