package psql

import (
	"database/sql"

	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/models"
)

// TallyModel type which wraps a sql.DB connection pool.
type TallyModel struct {
	DB *sql.DB
}

// Get a specific tally
func (m *TallyModel) Get(id int) (*models.Tally, error) {
	stmt := `SELECT id, count, month, year, precinct_id, moving_violation_id 
					 FROM tallies 
					 WHERE id = $1`

	t := &models.Tally{}

	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&t.ID, &t.Count, &t.Month, &t.Year, &t.PrecinctID, &t.MovingViolationID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// List returns a list of tallies filtered by params
// TODO: Implement
func (m *TallyModel) List() ([]*models.Tally, error) {
	return nil, nil
}
