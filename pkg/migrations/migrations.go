package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
)

// RunMigrations runs migrations
func RunMigrations(cfg Config) (version uint, dirty bool, err error) {
	if !cfg.Enable {
		return 0, false, nil
	}

	// Create the migrate instance with the sourceURL and databaseURL
	m, err := migrate.New(cfg.SourceURL, cfg.DbURL)
	if err != nil {
		return 0, false, err
	}
	// Close the database connection after the migrations have finished
	defer func() {
		sourceErr, dbErr := m.Close()
		if sourceErr != nil {
			err = sourceErr
		}
		if dbErr != nil {
			err = dbErr
		}
	}()

	// Recreate the database if the recreate flag is set
	if cfg.Recreate {
		if err := m.Down(); err != nil {
			return 0, false, err
		}
	}

	// Run the migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return 0, false, err
	}

	// Return the current version and whether or not the database is dirty
	return m.Version()
}
