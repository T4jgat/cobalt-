package main

import (
	"baha.account/internal/data"
	"baha.account/internal/mailer"
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	limiter struct {
		enabled bool
		rps     float64
		burst   int
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

// Dependency injection
type application struct {
	config config
	logger *log.Logger
	models data.Models
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {

	// config struct instance
	var cfg config

	// defining command line flags (if flags was not inserted, server uses default values)
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("BNURLANULY_DB_DSN"), "PostgreSQL DSN")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "smtp-mail.outlook.com", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "220068@astanait.edu.kz", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "CuCher5xXbVVO", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "220068@astanait.edu.kz", "SMTP sender")

	flag.Parse()

	// logger instance
	// writes logs with prefixes - date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Printf("database connection pool established")

	migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Fatal(err, nil)
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://C:/Users/baqti/OneDrive/Desktop/adv-prog-assign-1/migrations",
		"postgres",
		migrationDriver,
		//C:/Users/baqti/OneDrive/Desktop/adv-prog-assign-1/migrations
	)
	if err != nil {
		logger.Fatal(err, nil)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal(err, nil)
	}

	logger.Printf("database migrations applied")

	// application instance
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// server instance
	// server staring log message
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	// error logs
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
