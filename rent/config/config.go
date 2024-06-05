package config

import "os"

type Config struct {
	HTTPPort string
	Postgres PostgresConfig
	RabbitMQ RabbitMQConfig
	LogLevel LogLevelConfig
}

type PostgresConfig struct {
	DSN string
}

type RabbitMQConfig struct {
	URL string
}

type LogLevelConfig struct {
	Level string
}

func LoadConfig() Config {
	return Config{
		HTTPPort: os.Getenv("HTTP_PORT"),
		Postgres: PostgresConfig{
			DSN: os.Getenv("POSTGRES_DSN"),
		},
		RabbitMQ: RabbitMQConfig{
			URL: os.Getenv("RABBITMQ_URL"),
		},
		LogLevel: LogLevelConfig{
			Level: os.Getenv("LOG_LEVEL"),
		},
	}
}
