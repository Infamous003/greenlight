package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Infamous003/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Creating a movie instance
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Pirates of the Caribbean",
		Year:      2003,
		Runtime:   118,
		Genres:    []string{"fantasy", "thriller", "pirates"},
		Version:   1,
	}

	// encoding the struct to responsewriter
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
