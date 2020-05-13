package psql

import (
	"database/sql"

	"github.com/lib/pq"
	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/models"
	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/queries"
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
func (m *TallyModel) List(req *queries.TalliesRequest) ([]*models.Tally, error) {
	stmt := `SELECT id, count, month, year, precinct_id, moving_violation_id
					 FROM tallies
					 WHERE precinct_id=ANY($1)
					 OFFSET $2 ROWS
					 LIMIT $3
					`

	rows, err := m.DB.Query(stmt, pq.Array(req.PrecinctIDs), (req.Page-1)*req.PerPage, req.PerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tallies := []*models.Tally{}
	for rows.Next() {
		t := &models.Tally{}
		err = rows.Scan(&t.ID, &t.Count, &t.Month, &t.Year, &t.PrecinctID, &t.MovingViolationID)
		if err != nil {
			return nil, err
		}
		tallies = append(tallies, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tallies, nil
}
