package data

import (
	"baha.account/internal/validator"
	"testing"
)

func TestTokenModel_ValidateTokenPlaintext(t *testing.T) {
	v := validator.New()

	ValidateTokenPlaintext(v, "abcdefghijklmnopqrstuvwxyz")
	assertValid(t, v)

	ValidateTokenPlaintext(v, "")
	assertInvalid(t, v, "token", "must be provided")

	ValidateTokenPlaintext(v, "short")
	assertInvalid(t, v, "token", "must be 26 bytes long")
}

func assertValid(t *testing.T, v *validator.Validator) {
	if !v.Valid() {
		t.Errorf("Validation failed: %v", v.Errors)
	}
}

func assertInvalid(t *testing.T, v *validator.Validator, key string, message string) {
	if v.Valid() {
		t.Error("Validation should have failed")
	}

	if v.Errors[key] != message {
		t.Errorf("Expected error message '%s' for key '%s', got '%s'", message, key, v.Errors[key])
	}
}
