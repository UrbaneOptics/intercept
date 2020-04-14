package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

// Precinct Model
type Precinct struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// MovingViolation Model
type MovingViolation struct {
	ID   int
	Name string
}

// Tally model
type Tally struct {
	Count             int
	Month             int
	Year              int
	PrecinctID        int
	MovingViolationID int
}
