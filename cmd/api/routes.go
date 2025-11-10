package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Creates a router, registeres all API routes, and returns it
// refer Panic recovery chapter, pg 74
func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(app.recoverPanic)

	// rate limiter middleware
	router.Use(app.rateLimiter)

	// overwriting the default repsponses with custom ones
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	router.Get("/v1/movies/{id}", app.showMovieHandler)
	router.Patch("/v1/movies/{id}", app.updateMovieHandler)
	router.Delete("/v1/movies/{id}", app.deleteMovieHandler)
	router.Get("/v1/movies", app.listMoviesHandler)

	return router
}
