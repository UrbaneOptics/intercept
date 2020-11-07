package psql

import (
	"database/sql"

	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/models"
)

// MovingViolationModel type which wraps a sql.DB connection pool.
type MovingViolationModel struct {
	DB *sql.DB
}

// Get a specific moving violation
func (m *MovingViolationModel) Get(id int) (*models.MovingViolation, error) {
	stmt := `SELECT id, name
					 FROM moving_violations 
					 WHERE id = $1`

	mv := &models.MovingViolation{}

	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&mv.ID, &mv.Name)
	if err != nil {
		return nil, err
	}

	return mv, nil
}

// List returns a list of moving violations
func (m *MovingViolationModel) List() ([]*models.MovingViolation, error) {
	stmt := `SELECT id, name FROM moving_violations`

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
