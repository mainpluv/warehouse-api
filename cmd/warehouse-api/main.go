package main

import (
	"log"
	"net/http"
	"warehouse-api/internal/api"
	"warehouse-api/internal/config"
	"warehouse-api/internal/db"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading cfg: %v", err)
	}
	db, err := db.Connect(cfg.DB)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	router := api.SetupRouter(db)
	log.Fatal(http.ListenAndServe(cfg.Server.Address, router))
}
