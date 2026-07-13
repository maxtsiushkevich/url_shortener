package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
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
	url, _ := io.ReadAll(r.Body)
	fmt.Printf("create handler: %s\n", url)
	code, err := h.urlService.GetShortUrlCode(string(url))

	if err != nil {
		http.Error(w, "failed to create short url", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(code))
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

	http.Redirect(w, r, url, http.StatusSeeOther)
	fmt.Println("Redirect handler")
}
