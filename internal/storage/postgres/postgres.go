package postgres

import (
	"context"
	"fmt"
	"url_shortener/internal/models"

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

func (p *Postgres) Save(ctx context.Context, url models.URL) error {
	fmt.Println("Save to db")
	return nil
}

func (p *Postgres) GetByCode(ctx context.Context, code string) (models.URL, error) {
	fmt.Println("Read from db")
	return models.URL{}, nil
}
