package data

import (
	"baha.account/internal/validator"
	"testing"
)

func TestUserModel_ValidateEmail(t *testing.T) {
	v := validator.New()

	ValidateEmail(v, "user@example.com")
	if !v.Valid() {
		t.Error("Valid email marked as invalid")
	}

	v = validator.New() // Reset the validator for the next test
	ValidateEmail(v, "user+test@example.com")
	if !v.Valid() {
		t.Error("Valid email with '+' marked as invalid")
	}

	v = validator.New() // Reset the validator
	ValidateEmail(v, "")
	if v.Valid() {
		t.Error("Empty email should be invalid")
	} else if v.Errors["email"] != "must be provided" {
		t.Errorf("Incorrect error message for empty email: got %q, want %q", v.Errors["email"], "must be provided")
	}

	v = validator.New() // Reset the validator
	ValidateEmail(v, "user")
	if v.Valid() {
		t.Error("Invalid email format should be invalid")
	} else if v.Errors["email"] != "must be a valid email address" {
		t.Errorf("Incorrect error message for invalid email: got %q, want %q", v.Errors["email"], "must be a valid email address")
	}
}

func TestUserModel_ValidatePasswordPlainText(t *testing.T) {
	v := validator.New()

	ValidatePasswordPlainText(v, "validPassword123")
	if !v.Valid() {
		t.Error("Valid password marked as invalid")
	}

	v = validator.New() // Reset the validator
	ValidatePasswordPlainText(v, "")
	if v.Valid() {
		t.Error("Empty password should be invalid")
	} else if v.Errors["password"] != "must be provided" {
		t.Errorf("Incorrect error message for empty password: got %q, want %q", v.Errors["password"], "must be provided")
	}

	v = validator.New() // Reset the validator
	ValidatePasswordPlainText(v, "short")
	if v.Valid() {
		t.Error("Short password should be invalid")
	} else if v.Errors["password"] != "must be at least 8 bytes long" {
		t.Errorf("Incorrect error message for short password: got %q, want %q", v.Errors["password"], "must be at least 8 bytes long")
	}

	v = validator.New() // Reset the validator
	ValidatePasswordPlainText(v, "toolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpasswordtoolongpassword1")
	if v.Valid() {
		t.Error("Too long password should be invalid")
	} else if v.Errors["password"] != "must not be more than 71 bytes long" {
		t.Errorf("Incorrect error message for too long password: got %q, want %q", v.Errors["password"], "must not be more than 71 bytes long")
	}
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
	if !v.Valid() {
		t.Errorf("Valid user marked as invalid: %v", v.Errors)
	}

	user.Fname = ""
	v = validator.New()
	ValidateUser(v, user)
	if v.Valid() {
		t.Error("User with empty fname should be invalid")
	} else if v.Errors["fname"] != "must be provided" {
		t.Errorf("Incorrect error message for empty fname: got %q, want %q", v.Errors["fname"], "must be provided")
	}

	// ... [Add similar checks for other fields: Sname, Email, Password] ...
}
