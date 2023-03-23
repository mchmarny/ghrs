package data

import (
	"github.com/pkg/errors"
)

// Delete deletes the value of the given id.
func (s *Store) Delete(id string) (deleted bool, err error) {
	if s.db == nil {
		return false, errors.New("database not initialized")
	}

	stmt, err := s.db.Prepare("DELETE FROM counter WHERE id = ?")
	if err != nil {
		return false, errors.Wrapf(err, "failed to prepare delete statement")
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false, errors.Wrapf(err, "failed to execute delete statement")
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return false, errors.Wrapf(err, "failed to get affected rows")
	}

	return affect > 0, nil
}
