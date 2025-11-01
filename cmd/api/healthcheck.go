package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	headers := http.Header{}
	headers.Set("Content-Language", "en-UK")

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
