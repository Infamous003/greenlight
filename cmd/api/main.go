package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Infamous003/greenlight/internal/data"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int    // port on which server is listening on
	env  string // prod, dev, testing, etc
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	config config
	logger *slog.Logger
	models data.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 9090, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// reading the dsn value
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// creating a connection pool
	db, err := openDB(&cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db), // injecting db to our models
	}

	r := chi.NewRouter()

	appRouter := app.routes()
	r.Mount("/api", appRouter) // appending `/api/` to all the appROuter endpoints

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      r,
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", s.Addr, "env", cfg.env)

	err = s.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg *config) (*sql.DB, error) {
	// Open creates an empty connection pool, using the driver name, and dsn string
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// max num of idle + in-use conns, val 0 or less is unlimited
	db.SetMaxOpenConns(cfg.db.maxOpenConns)

	// max idle conns, is always <= Max open conns
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	// idle timeout for conns in pool
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// if the connecton couldn't be established successfully within 5s, then an error is raised
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
