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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.LoadConfig()
	db, err := repository.NewPostgresDB(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repository.NewRepository(db)
	srv := service.NewService(repo)
	handler := api.NewHandler(srv)

	router := api.NewRouter(handler)

	log.Println("Starting server on port 8080")
	if err := router.Start(viper.GetString("server.address")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
