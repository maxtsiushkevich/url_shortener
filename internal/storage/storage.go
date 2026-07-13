package storage

import (
	"context"
	"url_shortener/internal/models"
)

type Storage interface {
	Save(ctx context.Context, url models.Url) error
	GetByCode(ctx context.Context, code string) (models.Url, error)
}
