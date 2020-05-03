package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

// Precinct Model
type Precinct struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	ShortName   string `json:"short_name" db:"short_name"`
	IsAggregate string `json:"is_aggregate" db:"is_aggregate"`
}

// MovingViolation Model
type MovingViolation struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Tally model
type Tally struct {
	ID                int `json:"id" db:"id"`
	Count             int `json:"count" db:"count"`
	Month             int `json:"month" db:"month"`
	Year              int `json:"year" db:"year"`
	PrecinctID        int `json:"precinct_id" db:"precinct_id"`
	MovingViolationID int `json:"moving_violation_id" db:"moving_violation_id"`
}
