package main

import (
	"github.com/T4jgat/cobalt+/config"
	"github.com/T4jgat/cobalt+/internal/app"
	"log"
)

func main() {

	// Load environment variables from .env file
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	//
	//log := logger.New()
	//cfg := config.LoadConfig()
	//
	//db, err := postgres.New(cfg.Postgres)
	//if err != nil {
	//	log.Fatal("failed to connect to database", err)
	//}
	//
	//rabbitMQ, err := server.New(cfg.RabbitMQ)
	//if err != nil {
	//	log.Fatal("failed to connect to RabbitMQ", err)
	//}
	//
	//application := app.New(db, rabbitMQ, log)
	//httpServer := httpv1.New(application.RentalsController, log)
	//
	//log.Info("starting HTTP server on", cfg.HTTPPort)
	//if err := http.ListenAndServe(":"+cfg.HTTPPort, httpServer); err != nil {
	//	log.Fatal("failed to start HTTP server", err)
	//}

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
