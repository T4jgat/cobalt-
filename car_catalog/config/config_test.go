package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("load config from env", func(t *testing.T) {
		os.Setenv("HTTP_PORT", "8080")
		os.Setenv("POSTGRES_DSN", "postgres://user:password@host:port/database")
		os.Setenv("RABBITMQ_URL", "amqp://user:password@host:port")
		os.Setenv("LOG_LEVEL", "info")

		cfg := LoadConfig()

		if cfg.HTTPPort != "8080" {
			t.Errorf("expected HTTP_PORT to be 8080, got %s", cfg.HTTPPort)
		}
		if cfg.Postgres.DSN != "postgres://user:password@host:port/database" {
			t.Errorf("expected POSTGRES_DSN to be postgres://user:password@host:port/database, got %s", cfg.Postgres.DSN)
		}
		if cfg.RabbitMQ.URL != "amqp://user:password@host:port" {
			t.Errorf("expected RABBITMQ_URL to be amqp://user:password@host:port, got %s", cfg.RabbitMQ.URL)
		}
		if cfg.LogLevel.Level != "info" {
			t.Errorf("expected LOG_LEVEL to be info, got %s", cfg.LogLevel.Level)
		}
	})
}
