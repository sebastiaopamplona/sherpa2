package migrations

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var migrations = []*Migration{}

// Run runs database migrations
func Run(sqlDB *sql.DB, dialect string) error {
	db := sqlx.NewDb(sqlDB, dialect)

	_, _ = fmt.Println("Running migrations...")
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id TEXT PRIMARY KEY,
			created_at timestamptz NOT NULL default now()
		);`)
	if err != nil {
		return fmt.Errorf("creating migrations table: %w", err)
	}

	for _, m := range migrations {
		var found string
		err := db.Get(&found, `SELECT id FROM migrations WHERE id=$1`, m.ID)
		switch err {
		case sql.ErrNoRows:
			_, _ = fmt.Printf("Running migration: %v\n", m.ID)
		case nil:
			_, _ = fmt.Printf("Skipping migration: %v\n", m.ID)
			continue
		default:
			return fmt.Errorf("looking up migration by id: %w", err)
		}
		errorf := func(err error) error { return fmt.Errorf("running migration: %w", err) }

		migrationSteps := []func(tx *sqlx.Tx) error{m.Migrate, m.Seed, m.Upgrade}
		for _, step := range migrationSteps {
			if step == nil {
				continue
			}

			tx, err := db.Beginx()
			if err != nil {
				return errorf(err)
			}

			if err := step(tx); err != nil {
				if err := tx.Rollback(); err != nil {
					return errorf(err)
				}
				return errorf(err)
			}

			if err := tx.Commit(); err != nil {
				return errorf(err)
			}
		}

		tx, err := db.Beginx()
		if err != nil {
			return errorf(err)
		}

		_, err = tx.Exec(`INSERT INTO migrations (id, created_at) VALUES ($1,now())`, m.ID)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				return errorf(err)
			}
			return errorf(err)
		}

		if err := tx.Commit(); err != nil {
			return errorf(err)
		}

		return nil
	}
	return nil
}
