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
				CREATE TABLE IF NOT EXISTS migrations (
					id TEXT PRIMARY KEY,
					created_at timestamptz NOT NULL default now()
				);

				CREATE TABLE users (
					id text NOT NULL,
					
					first_name text NOT NULL,
					last_name text NOT NULL,
					email text NOT NULL,
					avatar text NOT NULL,
					password text NOT NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					update_at timestamp(3) without time zone NOT NULL,
					deleted_at timestamp(3) without time zone NULL,
					
					CONSTRAINT users_pkey PRIMARY KEY (id),
					CONSTRAINT users_unique_email UNIQUE (email)
				);

				CREATE TABLE projects (
					id text NOT NULL,
					
					name text NOT NULL,
					description text NULL,
					github_url text NULL,
					jira_url text NULL,
					
					creator_id text NOT NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at timestamp(3) without time zone NOT NULL,
					deleted_at timestamp(3) without time zone NULL,
					
					CONSTRAINT projects_pkey PRIMARY KEY (id),
					CONSTRAINT projects_creator_fkey FOREIGN KEY (creator_id) REFERENCES users (id)
				);

				CREATE TABLE roles (
					id text NOT NULL,
					
					name text NOT NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at timestamp(3) without time zone NOT NULL,
					
					CONSTRAINT roles_pkey PRIMARY KEY (id)
				);

				CREATE TABLE permissions (
					id text NOT NULL,
					
					name text NOT NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at timestamp(3) without time zone NOT NULL,
					
					CONSTRAINT permissions_pkey PRIMARY KEY (id)
				);

				CREATE TABLE role_to_permissions (
					id text NOT NULL,
					
					role_id text NOT NULL,
					permission_id text NOT NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at timestamp(3) without time zone NOT NULL,
					
					CONSTRAINT rtp_pkey PRIMARY KEY (id),
					CONSTRAINT rtp_role_fkey FOREIGN KEY (role_id) REFERENCES roles (id),
					CONSTRAINT rtp_permission_fkey FOREIGN KEY (permission_id) REFERENCES permissions (id)
				);

				CREATE TABLE sprints (
					id text NOT NULL,
						
					title text NOT NULL,
					description text NULL,
					start_at timestamp(3) without time zone NOT NULL,
					end_at timestamp(3) without time zone NOT NULL,
					
					creator_id text NOT NULL,
					project_id text NOT NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at timestamp(3) without time zone NOT NULL,
					
					CONSTRAINT sprints_pkey PRIMARY KEY (id),
					CONSTRAINT sprints_creator_fkey FOREIGN KEY (creator_id) REFERENCES users (id),
					CONSTRAINT sprints_project_fkey FOREIGN KEY (project_id) REFERENCES projects (id)
				);

				CREATE TABLE stories (
					id text NOT NULL,
					
					title text NOT NULL,
					description text NOT NULL DEFAULT ''::text,
					estimate double precision NULL DEFAULT 0,
					type text NOT NULL,
					github_id text NULL,
					jira_id text NULL,
					state text NOT NULL,
					remaining_effort double precision NULL DEFAULT 0,
					
					project_id text NOT NULL,
					creator_id text NOT NULL,
					assignee_id text NULL,
					sprint_id text NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at timestamp(3) without time zone NOT NULL,
					deleted_at timestamp(3) without time zone NULL,
						
					CONSTRAINT stories_pkey PRIMARY KEY (id),
					CONSTRAINT stories_project_fkey FOREIGN KEY (project_id) REFERENCES projects (id),
					CONSTRAINT stories_creator_fkey FOREIGN KEY (creator_id) REFERENCES users (id),
					CONSTRAINT stories_assignee_fkey FOREIGN KEY (assignee_id) REFERENCES users (id),
					CONSTRAINT stories_sprint_fkey FOREIGN KEY (sprint_id) REFERENCES sprints (id)
				);

				CREATE TABLE story_worklogs (
					id text NOT NULL,
					
					description text NOT NULL,
					date timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					effort double precision NOT NULL,
					remaining_effort double precision NOT NULL,
					
					story_id text NOT NULL,
					creator_id text NOT NULL,
					
					created_at timestamp(3) without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
					updated_at timestamp(3) without time zone NOT NULL,
					
					CONSTRAINT worklogs_pkey PRIMARY KEY (id),
					CONSTRAINT sw_story_fkey FOREIGN KEY (story_id) REFERENCES stories (id),
					CONSTRAINT sw_creator_fkey FOREIGN KEY (creator_id) REFERENCES users (id)
				);
`

			_, err := tx.Exec(query)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
