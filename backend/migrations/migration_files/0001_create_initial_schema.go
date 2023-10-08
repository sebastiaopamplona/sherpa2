package migrator

import (
	"github.com/jmoiron/sqlx"
	"github.com/sebastiaopamplona/sherpa2/migrations"
)

func M0001CreateInitialSchema() migrations.Migration {
	return migrations.Migration{
		ID: "0001_create_initial_schema",
		Migrate: func(tx *sqlx.Tx) error {
			query := `
				CREATE TABLE public."User" (
					id text NOT NULL,
					first_name text NOT NULL,
					last_name text NOT NULL,
					email text NOT NULL,
					avatar text NOT NULL,
					password text NOT NULL,
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at timestamp(3) without time zone NOT NULL,
					deleted_at timestamp(3) without time zone NULL,
					
					CONSTRAINT "user_pkey" PRIMARY KEY (id),
					CONSTRAINT "user_unique_email" UNIQUE (email)
				);

				ALTER TABLE public."User" ADD CONSTRAINT "User_pkey" PRIMARY KEY (id);


				);`

			_, err := tx.Exec(query)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
