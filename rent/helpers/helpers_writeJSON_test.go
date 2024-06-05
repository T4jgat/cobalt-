package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	t.Run("Valid Data", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := map[string]interface{}{"message": "Test message"}

		err := WriteJSON(w, http.StatusOK, data, nil)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got: %v", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error reading response body: %v", err)
		}

		var actualData map[string]interface{}
		if err := json.Unmarshal(body, &actualData); err != nil {
			t.Errorf("Error unmarshalling response body: %v", err)
		}

		if actualData["message"] != "Test message" {
			t.Errorf("Expected message to be 'Test message', got: %v", actualData["message"])
		}
	})

	t.Run("Invalid Data", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := make(chan int)

		err := WriteJSON(w, http.StatusOK, data, nil)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
