package postgres

import (
	"testing"

	"github.com/T4jgat/cobalt/config"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.PostgresConfig{
		DSN: "postgres://user:password@host:port/database",
	}

	db, err := New(cfg)

	assert.NotNil(t, db)
	assert.NoError(t, err)
}
