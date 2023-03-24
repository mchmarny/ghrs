package data

import (
	"database/sql"
	"embed"
	"os"
	"path/filepath"

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

	dbNotInitializedErr = errors.New("database not initialized")
)

// New initializes the database for a given name.
func New(path string) (*Store, error) {
	if path == "" {
		return nil, errors.New("path not specified")
	}

	log.Debug().Msgf("data path: %s", path)
	wasCreated, err := ensureParentDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to ensure parent dir for: %s", path)
	}

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

// ensureParentDir ensures that the parent directory exists.
func ensureParentDir(s string) (bool, error) {
	if s == "" {
		return false, errors.New("path not specified")
	}

	if _, err := os.Stat(s); errors.Is(err, os.ErrNotExist) {
		d := filepath.Dir(s)
		log.Debug().Msgf("creating dir: %s", d)
		if err := os.MkdirAll(d, os.ModePerm); err != nil {
			return false, errors.Wrapf(err, "failed to create path: %s", d)
		}
		return true, nil
	}

	return false, nil
}
