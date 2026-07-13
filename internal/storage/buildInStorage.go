package storage

import (
	"context"
	"errors"
	"url_shortener/internal/models"
)

type BuildInStorage struct {
	urlsMap map[string]models.Url
}

func NewBuiltInStorage() *BuildInStorage {
	return &BuildInStorage{
		urlsMap: make(map[string]models.Url),
	}
}

func (s *BuildInStorage) Save(ctx context.Context, url models.Url) error {
	s.urlsMap[url.Code] = url
	return nil
}

var ErrNotFound = errors.New("url not found")

func (s *BuildInStorage) GetByCode(ctx context.Context, code string) (models.Url, error) {
	shortURL, ok := s.urlsMap[code]
	if !ok {
		return models.Url{}, ErrNotFound
	}

	return shortURL, nil
}
