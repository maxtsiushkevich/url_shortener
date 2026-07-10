package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"url_shortener/internal/config"

	"github.com/jackc/pgx/v5"
)

var configPath = flag.String("config", "config/config.yaml", "Path to configuration file")

func main() {
	flag.Parse()
	config, err := config.Load(*configPath)
	fmt.Println(config)
	fmt.Println(err)

	// init logger
	// init db
	// init net/http
	// run server
}
