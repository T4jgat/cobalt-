package app

import (
	"database/sql"
	"github.com/T4jgat/cobalt+/internal/controller/httpv1"
	"github.com/T4jgat/cobalt+/internal/usecase/repo"
	"github.com/T4jgat/cobalt+/pkg/logger"
	"github.com/T4jgat/cobalt+/pkg/rabbitmq/rmq_rpc/server"
)

type App struct {
	DB                *sql.DB
	RabbitMQ          *server.Server
	Logger            logger.Logger
	RentalsController *httpv1.RentalsController
}

func New(db *sql.DB, rabbitMQ *server.Server, log logger.Logger) *App {
	rentalRepo := repo.New(db)
	rentalsController := httpv1.New(rentalRepo)

	return &App{
		DB:                db,
		RabbitMQ:          rabbitMQ,
		Logger:            log,
		RentalsController: rentalsController,
	}
}
