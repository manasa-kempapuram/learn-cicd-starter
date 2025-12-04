package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey apikey_12345")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "apikey_12345" {
		t.Fatalf("expected 'apikey_12345', got '%s'", key)
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer not-an-api-key")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error for malformed header, got nil")
	}
}
