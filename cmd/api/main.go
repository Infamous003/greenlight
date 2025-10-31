package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

const version = "1.0.0"

type config struct {
	port int    // port on which server is listening on
	env  string // prod, dev, testing, etc
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 9090, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
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

	err := s.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
