package integration_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T4jgat/cobalt/config"
	"github.com/T4jgat/cobalt/internal/app"
	v1 "github.com/T4jgat/cobalt/internal/controller/http/v1"
	"github.com/T4jgat/cobalt/pkg/logger"
	"github.com/T4jgat/cobalt/pkg/postgres"
	"github.com/T4jgat/cobalt/pkg/rabbitmq/rmq_rpc/server"
	"github.com/stretchr/testify/assert"
)

func TestErrorHandling(t *testing.T) {
	// Connect to the database
	cfg := config.LoadConfig()
	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Connect to RabbitMQ
	rabbitMQ, err := server.New(cfg.RabbitMQ)
	if err != nil {
		t.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Create a new app instance
	log := logger.New("info")
	application := app.New(db, rabbitMQ, *log)
	httpServer := v1.NewRouter(application.CatalogsController, *log)

	// Create a test server
	ts := httptest.NewServer(httpServer)
	defer ts.Close()

	// Test invalid JSON body
	client := ts.Client()
	req, err := http.NewRequest(http.MethodPost, "/catalogs",
		bytes.NewBuffer([]byte(`{"model": "Model X", "brand": "Tesla", "color": "Red"`)))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusBadRequest, "expected status code 400")

	// Test non-existent catalog entry
	req, err = http.NewRequest(http.MethodGet, "/catalogs/12345", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, http.StatusNotFound, "expected status code 404")

	// Test internal server error (TODO: Implement an error scenario for testing)
	// ...
}
