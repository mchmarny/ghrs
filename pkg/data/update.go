package data

import (
	"github.com/pkg/errors"
)

// Update updates the value of the given id.
func (s *Store) Update(id string, val int64) error {
	if s.db == nil {
		return errDBNotInitialized
	}

	stmt, err := s.db.Prepare("UPDATE counter SET val = ? WHERE id = ?")
	if err != nil {
		return errors.Wrapf(err, "failed to prepare update statement")
	}

	_, err = stmt.Exec(val, id)
	if err != nil {
		return errors.Wrapf(err, "failed to execute update statement")
	}

	return nil
}

// Upsert updates the value of the given id.
func (s *Store) Upsert(id string, val int64) error {
	if s.db == nil {
		return errors.New("database not initialized")
	}

	stmt, err := s.db.Prepare(`INSERT INTO counter (id, val) VALUES(?,?)
							   ON CONFLICT(id) DO UPDATE SET val=excluded.val`)
	if err != nil {
		return errors.Wrapf(err, "failed to prepare upsert statement")
	}

	_, err = stmt.Exec(id, val)
	if err != nil {
		return errors.Wrapf(err, "failed to execute upsert statement")
	}

	return nil
}
