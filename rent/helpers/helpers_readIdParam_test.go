package helpers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestReadIDParam(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/rentals/123", nil)
		params := httprouter.Params{
			{Key: "id", Value: "123"},
		}
		ctx := context.WithValue(req.Context(), httprouter.ParamsKey, params)
		req = req.WithContext(ctx)

		id, err := ReadIDPAram(req)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		if id != 123 {
			t.Errorf("Expected ID to be 123, got: %v", id)
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/rentals/abc", nil)
		params := httprouter.Params{
			{Key: "id", Value: "abc"},
		}
		ctx := context.WithValue(req.Context(), httprouter.ParamsKey, params)
		req = req.WithContext(ctx)

		_, err := ReadIDPAram(req)
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}
