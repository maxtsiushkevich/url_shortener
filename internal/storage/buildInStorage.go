package storage

import (
	"context"
	"errors"
	"url_shortener/internal/models"
)

type BuildInStorage struct {
	urlsMap map[string]models.URL
}

func NewBuiltInStorage() *BuildInStorage {
	return &BuildInStorage{
		urlsMap: make(map[string]models.URL),
	}
}

func (s *BuildInStorage) Save(ctx context.Context, url models.URL) error {
	_, ok := s.urlsMap[url.Code]
	if ok {
		return nil
	}

	s.urlsMap[url.Code] = url
	return nil
}

func (s *BuildInStorage) Update(ctx context.Context, url models.URL) error {
	s.urlsMap[url.Code] = url
	return nil
}

var ErrNotFound = errors.New("url not found")

func (s *BuildInStorage) GetByCode(ctx context.Context, code string) (models.URL, error) {
	shortURL, ok := s.urlsMap[code]
	if !ok {
		return models.URL{}, ErrNotFound
	}

	return shortURL, nil
}
