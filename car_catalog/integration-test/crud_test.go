package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/T4jgat/cobalt/internal/entity"
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

func TestCatalogCRUD(t *testing.T) {
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

	// Create a new catalog entry
	client := ts.Client()
	req, err := http.NewRequest(http.MethodPost, "/catalogs",
		bytes.NewBuffer([]byte(`{"model": "Model X", "brand": "Tesla", "color": "Red", "price": 100000}`)))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status code 201, got %d", resp.StatusCode)
	}

	// Get the newly created catalog entry
	req, err = http.NewRequest(http.MethodGet, "/catalogs", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	var cars []entity.Catalog
	if err := json.NewDecoder(resp.Body).Decode(&cars); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	assert.Greater(t, len(cars), 0, "expected at least one catalog entry")

	// Update the catalog entry
	catalogID := cars[0].ID
	req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("/catalogs/%d", catalogID),
		bytes.NewBuffer([]byte(`{"model": "Model S", "brand": "Tesla", "color": "Blue", "price": 90000}`)))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	// Delete the catalog entry
	req, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/catalogs/%d", catalogID), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
}
