package psql

import (
	"database/sql"

	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/models"
)

// PrecinctModel type which wraps a sql.DB connection pool.
type PrecinctModel struct {
	DB *sql.DB
}

// Get a specific Precinct
func (m *PrecinctModel) Get(id int) (*models.Precinct, error) {
	stmt := `SELECT *
					 FROM precincts
					 WHERE id = $1`

	s := &models.Precinct{}

	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&s.ID, &s.Name)
	if err != nil {
		return nil, err
	}

	return s, nil
}
