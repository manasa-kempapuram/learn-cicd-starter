package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -

	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}


func GetAPIKey(headers http.Header) (string, error) {
	// Broken version to test CI failure
	return "BROKEN_KEY", nil
}


