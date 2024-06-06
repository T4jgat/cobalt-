package helpers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestReadIDParam(t *testing.T) {
	t.Run("valid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/catalogs/123", nil)
		ctx := context.WithValue(req.Context(), httprouter.ParamsKey, httprouter.Params{
			{Key: "id", Value: "123"},
		})
		req = req.WithContext(ctx)

		id, err := readIDPAram(req)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if id != 123 {
			t.Errorf("expected ID to be 123, got %d", id)
		}
	})

	t.Run("invalid ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/catalogs/abc", nil)
		ctx := context.WithValue(req.Context(), httprouter.ParamsKey, httprouter.Params{
			{Key: "id", Value: "abc"},
		})
		req = req.WithContext(ctx)

		_, err := readIDPAram(req)

		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
