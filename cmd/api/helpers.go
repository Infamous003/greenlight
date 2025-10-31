package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type envelope map[string]any

// doesnt have to be a method on application, just for consistency
func (app *application) readIDParam(r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil || id < 1 {
		return 0, errors.New("invalid id parmater")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for k, v := range headers {
		w.Header()[k] = v
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
