package psql

import (
	"database/sql"

	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/models"
)

// MovingViolation type which wraps a sql.DB connection pool.
type MovingViolationModel struct {
	DB *sql.DB
}

// Get a specific moving violation
func (m *MovingViolationModel) Get(id int) (*models.Tally, error) {
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

// List returns a list of moving violations
func (m *MovingViolationModel) List() ([]*models.MovingViolation, error) {
	stmt := `SELECT id, name FROM moving_violations`

	// m := &models.MovingViolation{}

	// row
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	movingViolations := []*models.MovingViolation{}

	for rows.Next() {
		m := &models.MovingViolation{}
		err = rows.Scan(&m.ID, &m.Name)
		if err != nil {
			return nil, err
		}
		movingViolations = append(movingViolations, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return movingViolations, nil
}
