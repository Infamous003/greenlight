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

	// overwriting the default repsponses with custom ones
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Post("/v1/movies", app.createMovieHandler)
	router.Get("/v1/movies/{id}", app.showMovieHandler)

	return router
}
