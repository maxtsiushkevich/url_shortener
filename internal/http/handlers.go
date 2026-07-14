package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"url_shortener/internal/models"
	"url_shortener/internal/service"
)

type UrlHandler struct {
	urlService *service.UrlService
}

func NewUrlHandler(service *service.UrlService) UrlHandler {
	return UrlHandler{
		urlService: service,
	}
}

func (h *UrlHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"URL"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println(request.URL)

	urlModel, err := h.urlService.GetShortUrl(request.URL)

	if err != nil {
		http.Error(w, "failed to create short url", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	shortUrl := r.Host + "/" + urlModel.Code
	shortUrlResponse := models.ShortURLResponse{
		URL: shortUrl,
	}
	urlJson, _ := json.Marshal(shortUrlResponse)

	_, err = w.Write([]byte(urlJson))
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func (h *UrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	url, err := h.urlService.GetFullUrl(string(code))
	if err != nil {
		http.Error(w, "failed to fetch full url", http.StatusNotFound)
		return
	}
	fmt.Println(url)

	http.Redirect(w, r, url.URL, http.StatusSeeOther)
}
