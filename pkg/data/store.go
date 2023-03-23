package data

import (
	"database/sql"
	"embed"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	// sqlite3 driver
	_ "modernc.org/sqlite"
)

const (
	dataDriver string = "sqlite"
)

var (
	//go:embed sql/*
	f embed.FS
)

// New initializes the database for a given name.
func New(path string) (*Store, error) {
	if path == "" {
		return nil, errors.New("path not specified")
	}

	wasCreated := false
	log.Debug().Msgf("data path: %s", path)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		log.Debug().Msg("data file does not exist, creating...")
		wasCreated = true
	}

	var err error
	db, err := sql.Open(dataDriver, path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open database: %s", path)
	}

	if wasCreated {
		log.Debug().Msg("creating schema...")

		b, err := f.ReadFile("sql/ddl.sql")
		if err != nil {
			return nil, errors.Wrap(err, "failed to read the schema creation file")
		}
		if _, err := db.Exec(string(b)); err != nil {
			return nil, errors.Wrapf(err, "failed to create database schema in: %s", path)
		}
	}

	s := &Store{
		db: db,
	}

	log.Debug().Msg("data initialized")
	return s, nil
}

// Close closes the database if one of previously created.
func (s *Store) Close() error {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			return errors.Wrap(err, "failed to close database")
		}
	}
	return nil
}

// Store is the data store.
type Store struct {
	db *sql.DB
}
