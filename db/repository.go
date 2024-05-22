package db

import (
	"database/sql"
	"fmt"
	"lamoda_test/config"
)

func Connect(cfg config.DBConfig) (*sql.DB, error) {
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open(cfg.Driver, dns)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
