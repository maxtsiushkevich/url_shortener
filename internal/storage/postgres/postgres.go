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
	query := `
		INSERT INTO urls (code, url, creation_time, clicks)
		VALUES ($1, $2, $3, $4)
	`

	_, err := p.conn.Exec(ctx, query, url.Code, url.URL, url.CreationTime, url.Clicks)
	if err != nil {
		return fmt.Errorf("failed to save url: %w", err)
	}
	return nil
}

func (p *Postgres) GetByCode(ctx context.Context, code string) (models.URL, error) {
	query := `
		SELECT code, url, creation_time, clicks
		FROM urls
		WHERE code = $1
	`

	var url models.URL
	err := p.conn.QueryRow(ctx, query, code).Scan(
		&url.Code,
		&url.URL,
		&url.CreationTime,
		&url.Clicks,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.URL{}, fmt.Errorf("url with code %s not found", code)
		}
		return models.URL{}, fmt.Errorf("failed to get url: %w", err)
	}

	return url, nil
}

func (p *Postgres) Update(ctx context.Context, url models.URL) error {
	query := `
		UPDATE urls
		SET url = $1, clicks = $2, creation_time = $3
		WHERE code = $4
	`

	result, err := p.conn.Exec(ctx, query, url.URL, url.Clicks, url.CreationTime, url.Code)
	if err != nil {
		return fmt.Errorf("failed to update url: %w", err)
	}

	// Проверить, был ли обновлён хотя бы один строк
	if result.RowsAffected() == 0 {
		return fmt.Errorf("url with code %s not found", url.Code)
	}

	return nil
}
