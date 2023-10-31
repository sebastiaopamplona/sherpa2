package migrations

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	mf "github.com/sebastiaopamplona/sherpa2/migrations/migration_files"
)

var migrations = []Migration{
	mf.M0001CreateInitialSchema(),
}

// Migrate runs database migrations
func Migrate(sqlDB *sql.DB, dialect string) error {
	_, _ = fmt.Println("Running migrations...")
	_, err := sqlDB.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id TEXT PRIMARY KEY,
			created_at timestamptz NOT NULL default now()
		);`)
	if err != nil {
		return fmt.Errorf("creating migrations table: %w", err)
	}

	for _, m := range migrations {
		var found string
		err := sqlDB.QueryRow(`SELECT id FROM migrations WHERE id=$1`, m.ID).Scan(&found)
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

			tx, err := sqlDB.Begin()
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

		tx, err := sqlDB.Begin()
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

// Truncate runs database migrations
func Truncate(sqlDB *sql.DB) error {
	tx, err := sqlDB.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query(`
		SELECT
			table_name
		FROM
			information_schema.tables
		WHERE
			table_type = 'BASE TABLE' AND table_schema = ANY (current_schemas(false));
	`)
	if err != nil {
		return err
	}

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return err
		}
		tables = append(tables, table)
	}

	for _, table := range tables {
		_, err = tx.Exec(`DROP TABLE ?;`, table)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
