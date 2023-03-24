package data

import (
	"github.com/pkg/errors"
)

// Delete deletes the value of the given id.
func (s *Store) Delete(id string) error {
	if s.db == nil {
		return errDBNotInitialized
	}

	stmt, err := s.db.Prepare("DELETE FROM counter WHERE id = ?")
	if err != nil {
		return errors.Wrapf(err, "failed to prepare delete statement")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrapf(err, "failed to execute delete statement")
	}

	return nil
}
