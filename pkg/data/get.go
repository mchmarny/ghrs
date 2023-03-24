package data

import (
	"database/sql"

	"github.com/pkg/errors"
)

const (
	NotSetValue int64 = 0
)

// Get gets the value of the given id.
func (s *Store) Get(id string) (val int64, err error) {
	if s.db == nil {
		return NotSetValue, dbNotInitializedErr
	}

	stmt, err := s.db.Prepare("SELECT val FROM counter WHERE id = ?")
	if err != nil {
		return NotSetValue, errors.Wrapf(err, "error preparing select statement")
	}

	row := stmt.QueryRow(id)

	var v int64
	err = row.Scan(&v)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NotSetValue, nil
		}
		return NotSetValue, errors.Wrapf(err, "error scanning row")
	}

	return v, nil
}
