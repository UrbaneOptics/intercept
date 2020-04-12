package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

// Precinct Model
type Precinct struct {
	ID   int
	Name string
}

// MovingViolation Model
type MovingViolation struct {
	ID   int
	Name string
}

// MovingViolationWritten model
type MovingViolationWritten struct {
	Count             int
	Month             int
	Year              int
	PrecinctID        int
	MovingViolationID int
}
