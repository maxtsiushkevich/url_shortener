package postgres

import "github.com/jackc/pgx/v5"

type Postgres struct {
	conn pgx.Conn
}
