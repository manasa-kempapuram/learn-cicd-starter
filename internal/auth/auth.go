package auth

import (
    "errors"
    "net/http"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func GetAPIKey(headers http.Header) (string, error) {
    // Broken version: always returns "BROKEN_KEY"
    return "BROKEN_KEY", nil
}

