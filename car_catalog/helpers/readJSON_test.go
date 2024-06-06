package helpers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadJSON(t *testing.T) {
	t.Run("invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{invalid json}"))
		w := httptest.NewRecorder()
		var data map[string]any

		err := readJSON(w, req, &data)

		if err == nil {
			t.Error("expected error, got nil")
		}

		if !strings.Contains(err.Error(), "body contains badly-formed JSON") {
			t.Errorf("expected error message to contain 'body contains badly-formed JSON', got %s", err.Error())
		}
	})

	t.Run("empty body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(""))
		w := httptest.NewRecorder()
		var data map[string]any

		err := readJSON(w, req, &data)

		if err == nil {
			t.Error("expected error, got nil")
		}

		if !strings.Contains(err.Error(), "body must not be empty") {
			t.Errorf("expected error message to contain 'body must not be empty', got %s", err.Error())
		}
	})

	t.Run("multiple JSON values", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(`{"name": "John"} {"age": 30}`))
		w := httptest.NewRecorder()
		var data map[string]any

		err := readJSON(w, req, &data)

		if err == nil {
			t.Error("expected error, got nil")
		}

		if !strings.Contains(err.Error(), "body must only contain a single JSON value") {
			t.Errorf("expected error message to contain 'body must only contain a single JSON value', got %s", err.Error())
		}
	})
}
