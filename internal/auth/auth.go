package auth

import (
	"errors"
	"net/http"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey temporarily broken to test CI
func GetAPIKey(headers http.Header) (string, error) {
	return "BROKEN_KEY", nil
}
