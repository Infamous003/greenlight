package data

import "time"

type Movie struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitzero"`
	// Runtime   int32     `json:"runtime,omitzero"`
	Runtime Runtime  `json:"runtime,omitzero"`
	Genres  []string `json:"genres,omitzero"`
	Version int32    `json:"-"`
}
