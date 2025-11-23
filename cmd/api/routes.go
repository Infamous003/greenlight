package main

import (
	"expvar"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Creates a router, registeres all API routes, and returns it
// refer Panic recovery chapter, pg 74
func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(app.recoverPanic)

	router.Use(app.enableCORS)

	router.Use(app.authenticate)

	// rate limiter middleware
	router.Use(app.rateLimiter)

	// overwriting the default repsponses with custom ones
	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Get("/v1/movies", app.requirePermission("movies:read", app.listMoviesHandler))
	router.Post("/v1/movies", app.requirePermission("movies:write", app.createMovieHandler))
	router.Get("/v1/movies/{id}", app.requirePermission("movies:read", app.showMovieHandler))
	router.Patch("/v1/movies/{id}", app.requirePermission("movies:write", app.updateMovieHandler))
	router.Delete("/v1/movies/{id}", app.requirePermission("movies:write", app.deleteMovieHandler))

	router.Post("/v1/users", app.registerUserHandler)
	router.Put("/v1/users/activated", app.activateUserHandler)

	router.Post("/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	router.Get("/v1/metrics", expvar.Handler().ServeHTTP)

	return router
}
