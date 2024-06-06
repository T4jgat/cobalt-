package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	t.Run("write JSON data", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := map[string]any{"message": "success"}
		headers := http.Header{"X-Custom-Header": []string{"value"}}

		err := WriteJSON(w, http.StatusOK, data, headers)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if w.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", w.Code)
		}
		if w.Header().Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type header to be application/json, got %s", w.Header().Get("Content-Type"))
		}
		if w.Header().Get("X-Custom-Header") != "value" {
			t.Errorf("expected X-Custom-Header header to be value, got %s", w.Header().Get("X-Custom-Header"))
		}

		var response map[string]any
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("unexpected error unmarshaling JSON: %v", err)
		}
		if response["message"] != "success" {
			t.Errorf("expected message to be success, got %v", response["message"])
		}
	})
}
