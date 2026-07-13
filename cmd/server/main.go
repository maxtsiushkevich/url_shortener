package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"url_shortener/internal/config"
	handlers "url_shortener/internal/http"
	"url_shortener/internal/service"
	"url_shortener/internal/storage"
	"url_shortener/internal/storage/postgres"
)

var configPath = flag.String("config", "config/config.yaml", "Path to configuration file")

func main() {
	//init config
	flag.Parse()
	config, err := config.Load(*configPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// init db
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Address,
		config.Database.DbName)

	db, err := postgres.NewPostgres(connString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// init service
	bis := storage.NewBuiltInStorage()

	// urlService := service.NewUrlService(db)
	urlService := service.NewUrlService(bis)

	handler := handlers.NewUrlHandler(&urlService)

	// init net/http
	// run server
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", handler.Create)
	mux.HandleFunc("GET /{code}", handler.Redirect)

	err = http.ListenAndServe(config.HTTPServer.Address, mux)

	if err != nil {
		log.Fatal(err)
	}
}
