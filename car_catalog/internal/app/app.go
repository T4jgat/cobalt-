package app

import (
	"database/sql"
	v1 "github.com/T4jgat/cobalt/internal/controller/http/v1"
	"github.com/T4jgat/cobalt/internal/usecase/repo"
	"github.com/T4jgat/cobalt/pkg/logger"
	"github.com/T4jgat/cobalt/pkg/rabbitmq/rmq_rpc/server"
)

type App struct {
	DB                 *sql.DB
	Logger             logger.Logger
	CatalogsController *v1.CatalogsController
}

func New(db *sql.DB, rabbitMQ *server.Server, log logger.Logger) *App {
	catalogRepo := repo.New(db)
	catalogsController := v1.New(catalogRepo)

	return &App{
		DB:                 db,
		Logger:             log,
		CatalogsController: catalogsController,
	}
}
