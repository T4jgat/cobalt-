package main

import (
	"fmt"
	"github.com/T4jgat/cobalt+/config"
	"github.com/T4jgat/cobalt+/internal/app"
	"github.com/T4jgat/cobalt+/pkg/httpserver"
	"github.com/T4jgat/cobalt+/pkg/logger"
	"github.com/T4jgat/cobalt+/pkg/postgres"
	"github.com/T4jgat/cobalt+/pkg/rabbitmq/rmq_rpc/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()
	log := logger.New(cfg.LogLevel.Level)

	db, err := postgres.New(cfg.Postgres)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	fmt.Println("postgres DSN ---- ", cfg.Postgres.DSN)

	rabbitMQ, err := server.New(cfg.RabbitMQ)
	if err != nil {
		log.Fatal("failed to connect to RabbitMQ", err)
	}

	application := app.New(db, rabbitMQ, *log)
	httpServer := httpserver.New(application.RentalsController, *log)

	log.Info("starting HTTP server on", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, httpServer); err != nil {
		log.Fatal("failed to start HTTP server", err)
	}
}
