package service

import (
	"context"
	"fmt"
	"url_shortener/internal/models"
	"url_shortener/internal/storage"
)

type UrlService struct {
	storage storage.Storage
}

func NewUrlService(storage storage.Storage) UrlService {
	return UrlService{
		storage: storage,
	}
}

func (srv *UrlService) GetFullUrl(shortUrl string) {
	srv.storage.GetByCode(context.Background(), shortUrl)
	fmt.Println("Произошо чтение из кэша, потом из БД и сохранение в кэш, вернулсяя короткий url")
}

func (srv *UrlService) GetShortUrl(fullUrl string) {
	srv.storage.Save(context.Background(), models.URL{})
	fmt.Println("Произошо сокращение ссылки, она записалась в БД")
}
