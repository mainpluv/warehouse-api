package repository

import (
	"context"
	"fmt"
	"os"
	"warehouse-api/config"
	"warehouse-api/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) *Repository {
	return &Repository{conn: conn}
}

func NewPostgresDB(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	strgPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	conn, err := pgxpool.New(ctx, strgPath)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (r *Repository) ReserveProducts(ctx context.Context, productCodes []string) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, code := range productCodes {
		_, err := tx.Exec(ctx, "INSERT INTO reservation (product_code) VALUES ($1)", code)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ReleaseProducts(ctx context.Context, productCodes []string) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, code := range productCodes {
		_, err := tx.Exec(ctx, "DELETE FROM reservation WHERE product_code = $1", code)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetStock(ctx context.Context, warehouseID string) ([]model.Product, error) {
	rows, err := r.conn.Query(ctx, "SELECT * FROM products WHERE warehouse_id = $1", warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stock []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Size, &product.Code, &product.Quantity); err != nil {
			return nil, err
		}
		stock = append(stock, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stock, nil
}
