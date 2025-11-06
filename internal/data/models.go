package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// A single struct which holds all the models, like a container
type Models struct {
	Movies MovieModel
}

// we don't use pointers for `Models` or `MovieModel` cause they're
// lightweight containers that already hold a pointer to shared *sql.DB
func NewModels(db *sql.DB) Models {
	// initializing each model with a db
	return Models{
		Movies: MovieModel{DB: db},
	}
}
