package main

import (
	"context"
	"log"
	"warehouse-api/config"
	"warehouse-api/pkg/api"
	"warehouse-api/pkg/repository"
	"warehouse-api/pkg/service"

	"github.com/spf13/viper"
)

// точка входа
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.LoadConfig()
	db, err := repository.NewPostgresDB(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	handler := api.NewHandler(srv)

	router := api.NewRouter(handler)

	log.Println("starting server on port 8080")
	if err := router.Start(viper.GetString("server.address")); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
