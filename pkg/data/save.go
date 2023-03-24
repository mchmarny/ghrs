package data

import (
	"github.com/pkg/errors"
)

// Save saves the value of the given id.
func (s *Store) Save(id string, val int64) error {
	if s.db == nil {
		return errDBNotInitialized
	}

	stmt, err := s.db.Prepare("INSERT INTO counter (id, val) VALUES (?, ?)")
	if err != nil {
		return errors.Wrapf(err, "failed to prepare insert statement")
	}

	_, err = stmt.Exec(id, val)
	if err != nil {
		return errors.Wrapf(err, "failed to execute insert statement")
	}

	return nil
}

// SaveAll saves all the ids in the database.
func (s *Store) SaveAll(ids map[string]int64) error {
	if s.db == nil {
		return errDBNotInitialized
	}

	stmt, err := s.db.Prepare("INSERT INTO counter (id, val) VALUES (?, ?)")
	if err != nil {
		return errors.Wrapf(err, "failed to prepare batch statement")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrapf(err, "failed to begin transaction")
	}

	for id, val := range ids {
		_, err = tx.Stmt(stmt).Exec(id, val)
		if err != nil {
			if err = tx.Rollback(); err != nil {
				return errors.Wrapf(err, "failed to rollback transaction")
			}
			return errors.Wrapf(err, "failed to execute batch statement")
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "failed to commit transaction")
	}

	return nil
}
