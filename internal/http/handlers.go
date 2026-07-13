// хендлеры
package http

import (
	"fmt"
	"io"
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
	h.urlService.GetShortUrl(string(url))
	fmt.Println("Create hanler")
}

func (h *UrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	h.urlService.GetFullUrl("123")
	fmt.Println("Redirect handler")
}
