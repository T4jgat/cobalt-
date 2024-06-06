package integration_test

import (
	"testing"

	"github.com/T4jgat/cobalt/config"
	"github.com/T4jgat/cobalt/pkg/postgres"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseMigrations(t *testing.T) {
	// Connect to the database
	cfg := config.LoadConfig()
	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Check if the 'catalog' table exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_name = 'catalog'").Scan(&count)
	if err != nil {
		t.Fatalf("failed to query database: %v", err)
	}
	assert.Equal(t, count, 1, "expected 'catalog' table to exist")
}
