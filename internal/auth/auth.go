package auth

import (
    "errors"
    "net/http"
)

// ErrNoAuthHeaderIncluded is returned when the Authorization header is missing
var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey is temporarily broken to trigger CI failure
func GetAPIKey(headers http.Header) (string, error) {
    // Broken version: always returns "BROKEN_KEY"
    return "BROKEN_KEY", nil
}
