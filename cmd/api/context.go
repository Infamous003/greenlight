package main

import (
	"context"
	"net/http"

	"github.com/Infamous003/greenlight/internal/data"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User) // need to assert the value user
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
