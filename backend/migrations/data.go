package migrations

import "github.com/jmoiron/sqlx"

type Migration struct {
	ID      string
	Migrate func(tx *sqlx.Tx) error
	Seed    func(tx *sqlx.Tx) error
	Upgrade func(tx *sqlx.Tx) error
}
