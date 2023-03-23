package data

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func (s *Store) Query(v string) ([]int64, error) {
	if s.db == nil {
		return nil, errors.New("database not initialized")
	}

	stmt, err := s.db.Prepare("SELECT val FROM counter WHERE id LIKE ?")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to prepare query statement")
	}

	rows, err := stmt.Query(fmt.Sprintf("%%%s%%", v))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrapf(err, "failed to execute select statement")
	}
	defer rows.Close()

	list := make([]int64, 0)
	for rows.Next() {
		var val int64
		if err := rows.Scan(&val); err != nil {
			return nil, errors.Wrapf(err, "failed to scan row")
		}
		list = append(list, val)
	}

	return list, nil
}
