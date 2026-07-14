package service

import (
	"context"
	"fmt"
	"time"
	"url_shortener/internal/models"
	"url_shortener/internal/storage"
	"url_shortener/internal/utils"
)

type UrlService struct {
	storage storage.Storage
}

func NewUrlService(storage storage.Storage) UrlService {
	return UrlService{
		storage: storage,
	}
}

func (srv *UrlService) GetFullUrl(code string) (models.URL, error) {
	shortUrlModel, err := srv.storage.GetByCode(context.Background(), code)
	if err != nil {
		fmt.Println(err)
		return models.URL{}, err
	}

	shortUrlModel.Clicks += 1
	err = srv.storage.Update(context.Background(), shortUrlModel)
	if err != nil {
		fmt.Println(err)
	}

	return shortUrlModel, nil
}

func (srv *UrlService) GetShortUrl(fullUrl string) (models.URL, error) {
	code := utils.GenerateCode(fullUrl)

	model := models.URL{
		Code:         code,
		URL:          fullUrl,
		CreationTime: time.Now(),
		Clicks:       0,
	}

	err := srv.storage.Save(context.Background(), model)
	if err != nil {
		return models.URL{}, err
	}
	return model, nil
}
