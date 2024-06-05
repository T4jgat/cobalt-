package data

import (
	"baha.account/internal/validator"
	"testing"
)

func TestUserModel_ValidateEmail(t *testing.T) {
	v := validator.New()

	ValidateEmail(v, "user@example.com")
	assertValid(t, v)

	ValidateEmail(v, "user+test@example.com")
	assertValid(t, v)

	ValidateEmail(v, "")
	assertInvalid(t, v, "email", "must be provided")

	ValidateEmail(v, "user")
	assertInvalid(t, v, "email", "must be a valid email address")
}

func TestUserModel_ValidatePasswordPlainText(t *testing.T) {
	v := validator.New()

	ValidatePasswordPlainText(v, "validPassword123")
	assertValid(t, v)

	ValidatePasswordPlainText(v, "")
	assertInvalid(t, v, "password", "must be provided")

	ValidatePasswordPlainText(v, "short")
	assertInvalid(t, v, "password", "must be at least 8 bytes long")

	ValidatePasswordPlainText(v, "toolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpassword1")
	assertInvalid(t, v, "password", "must not be more than 71 bytes long")
}

func TestUserModel_ValidateUser(t *testing.T) {
	user := &User{
		Fname:     "John",
		Sname:     "Doe",
		Email:     "john.doe@example.com",
		Password:  password{plaintext: new(string)},
		Activated: false,
	}
	err := user.Password.Set("password123")
	if err != nil {
		t.Fatal(err)
	}
	v := validator.New()
	ValidateUser(v, user)
	assertValid(t, v)

	user.Fname = ""
	ValidateUser(v, user)
	assertInvalid(t, v, "fname", "must be provided")

	user.Fname = "John"
	user.Sname = ""
	ValidateUser(v, user)
	assertInvalid(t, v, "sname", "must be provided")

	user.Sname = "Doe"
	user.Email = "invalid-email"
	ValidateUser(v, user)
	assertInvalid(t, v, "email", "must be a valid email address")

	user.Email = "john.doe@example.com"
	user.Password.plaintext = nil
	user.Password.hash = nil
	ValidateUser(v, user)
	assertInvalid(t, v, "password", "must be provided")
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
