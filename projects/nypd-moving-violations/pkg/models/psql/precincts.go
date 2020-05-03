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
	stmt := `SELECT id, name, short_name, is_aggregate
					 FROM precincts
					 WHERE id = $1`

	s := &models.Precinct{}

	row := m.DB.QueryRow(stmt, id)

	err := row.Scan(&s.ID, &s.Name, &s.ShortName, &s.IsAggregate)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// List returns a list of precincts
func (m *PrecinctModel) List() ([]*models.Precinct, error) {
	stmt := `SELECT id, name, short_name, is_aggregate
					 FROM precincts`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	precincts := []*models.Precinct{}
	for rows.Next() {
		p := &models.Precinct{}
		err = rows.Scan(&p.ID, &p.Name, &p.ShortName, &p.IsAggregate)
		if err != nil {
			return nil, err
		}
		precincts = append(precincts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return precincts, nil
}
