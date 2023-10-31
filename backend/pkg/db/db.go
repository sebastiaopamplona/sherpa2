package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/jmoiron/sqlx"
)

type config struct {
	DBURL      string `env:"DB_URL,required=true"`
	DBName     string `env:"DB_NAME,required=true"`
	DBUser     string `env:"DB_USER,required=true"`
	DBPassword string `env:"DB_PASSWORD,required=true"`
	DBSSLMode  bool   `env:"DB_SSL_MODE,default=false"`
}

type Client interface {
	Exec(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) (*sqlx.Row, error)
}

type client struct {
	conn *sqlx.DB
}

func New() (Client, error) {
	cfg := &config{}
	if _, err := env.UnmarshalFromEnviron(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal db config from env: %w", err)
	}

	conn, err := sqlx.Connect("postgres", fmt.Sprintf(
		"user=%s"+
			"user=%s"+
			"user=%s",
		cfg.DBUser,
		cfg.DBUser,
		cfg.DBUser,
	))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return &client{
		conn: conn,
	}
}

func (c *client) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {

}

func (c *client) Query(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {

}

func (c *client) QueryRow(ctx context.Context, query string, args ...any) (*sqlx.Row, error) {

}
