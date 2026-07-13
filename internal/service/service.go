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

func (srv *UrlService) GetFullUrl(code string) (string, error) {
	shortUrlModel, err := srv.storage.GetByCode(context.Background(), code)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return shortUrlModel.Url, nil
}

func (srv *UrlService) GetShortUrlCode(fullUrl string) (string, error) {
	code := utils.GenerateCode(fullUrl)
	model := models.Url{
		Code:         code,
		Url:          fullUrl,
		CreationTime: time.Now(),
		Clicks:       0,
	}
	err := srv.storage.Save(context.Background(), model)
	if err != nil {
		return "", err
	}
	return model.Code, nil
}
