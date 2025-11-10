package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         ":9090",
		Handler:      app.routes(),
		IdleTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1) // a buffered(size 1) channel that stores os.Signal type

		// Notify will keep track of signal interrupt and terminate
		// and sends it to the channel, also this is non blocking
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit // reading the signal from the channel, this is a blocking operation

		app.logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// shutdown may cause an error, so we send this to a channel
		// if the ctx expires before the srv shutsdown, we need to grab that err
		shutdownError <- srv.Shutdown(ctx)
	}()

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	// it may cause http.ErrServerClosed, which is expected, but if it is something else,
	// then that's a problem
	// Shutdown stops ListenAndServe and it sends an ErrServerCLosed
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.logger.Info("stopped server", "addr", srv.Addr)

	return nil
}
