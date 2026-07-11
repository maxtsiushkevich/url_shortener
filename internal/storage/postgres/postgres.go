package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	conn *pgx.Conn
}

func NewPostgres(connString string) (*Postgres, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &Postgres{conn: conn}, nil
}

func (p *Postgres) Close() {
	p.conn.Close(context.Background())
}

func (p *Postgres) Save() error {
	return nil
}

func (p *Postgres) GetByCode(code string) (string, error) {
	return "Redirect", nil
}
