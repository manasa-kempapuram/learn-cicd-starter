package auth

import (
	"errors"
	"net/http"
)

// Define the specific error used in tests
var ErrNoAuthHeaderIncluded = errors.New("missing Authorization header")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}

	// Expect header format: "ApiKey <key>"
	const prefix = "ApiKey "
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		return "", errors.New("malformed Authorization header")
	}

	return authHeader[len(prefix):], nil
}

