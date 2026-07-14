package storage

import (
	"context"
	"url_shortener/internal/models"
)

type Storage interface {
	Save(ctx context.Context, url models.URL) error
	Update(ctx context.Context, url models.URL) error
	GetByCode(ctx context.Context, code string) (models.URL, error)
}
