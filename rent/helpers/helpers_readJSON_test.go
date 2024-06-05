package helpers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadJSON(t *testing.T) {
	t.Run("Valid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"message": "Test message"}`)))
		var data map[string]interface{}

		err := ReadJSON(httptest.NewRecorder(), req, &data)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		if data["message"] != "Test message" {
			t.Errorf("Expected message to be 'Test message', got: %v", data["message"])
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"message": "Test message"`)))
		var data map[string]interface{}

		err := ReadJSON(httptest.NewRecorder(), req, &data)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("Empty Body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		var data map[string]interface{}

		err := ReadJSON(httptest.NewRecorder(), req, &data)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("Multiple JSON Values", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`
		{"message": "Test message 1"}
		{"message": "Test message 2"}
	`)))
		var data map[string]interface{}

		err := ReadJSON(httptest.NewRecorder(), req, &data)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	t.Run("Large Body", func(t *testing.T) {
		largeBody := make([]byte, 1_048_577)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(largeBody))
		var data map[string]interface{}

		err := ReadJSON(httptest.NewRecorder(), req, &data)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
