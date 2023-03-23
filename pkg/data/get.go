package data

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	NotSetValue = int64(0)
)

// Get gets the value of the given id.
func (s *Store) Get(id string) (val int64, err error) {
	if s.db == nil {
		return NotSetValue, errors.New("database not initialized")
	}

	stmt, err := s.db.Prepare("SELECT val FROM counter WHERE id = ?")
	if err != nil {
		return NotSetValue, errors.Wrapf(err, "failed to prepare select statement")
	}

	row := stmt.QueryRow(id)

	var v int64
	err = row.Scan(&v)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug().Err(err).Msgf("failed to find record for id: %s", id)
			return NotSetValue, nil
		}
		return NotSetValue, errors.Wrapf(err, "failed to scan row")
	}

	return v, nil
}
