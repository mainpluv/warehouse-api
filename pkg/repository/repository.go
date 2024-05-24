package repository

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"warehouse-api/config"
	"warehouse-api/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

// репозиторий для взаимодействия с бд
type Repository struct {
	conn *pgxpool.Pool
}

// новый экземпляр репозитория с указанным соединением
func NewRepository(conn *pgxpool.Pool) *Repository {
	return &Repository{conn: conn}
}

// создает новое соединение с бд
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
	// открытие транзакции
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	// итерируемся по кодам
	for _, code := range productCodes {
		// проверяем доступность склада
		var isAvailable bool
		err := tx.QueryRow(ctx, "SELECT is_available FROM warehouses WHERE id = (SELECT warehouse_id FROM products WHERE code = $1)", code).Scan(&isAvailable)
		if err != nil {
			return err
		}
		// проверка на доступность склада
		if !isAvailable {
			return fmt.Errorf("warehouse is not available for product with code %s now", code)
		}
		// проверяем есть ли уже запись о резервации для данного товара
		var reservedQty int
		err = tx.QueryRow(ctx, "SELECT quantity FROM reserved WHERE code = $1", code).Scan(&reservedQty)
		if err != nil {
			// если нет создаем новую запись о резервации
			_, err := tx.Exec(ctx, "INSERT INTO reserved (name, size, code, quantity, warehouse_id) SELECT name, size, code, 1, warehouse_id FROM products WHERE code = $1", code)
			if err != nil {
				return err
			}
		} else {
			// если есть обновляем количество резерва
			_, err := tx.Exec(ctx, "UPDATE reserved SET quantity = $1 WHERE code = $2", reservedQty+1, code)
			if err != nil {
				return err
			}
		}

		// уменьшаем количество товара на складе на 1
		_, err = tx.Exec(ctx, "UPDATE products SET quantity = quantity - 1 WHERE code = $1", code)
		if err != nil {
			return err
		}
	}
	// фиксируем транзакцию
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
		var isAvailable bool
		err := tx.QueryRow(ctx, "SELECT is_available FROM warehouses WHERE id = (SELECT warehouse_id FROM products WHERE code = $1)", code).Scan(&isAvailable)
		if err != nil {
			return err
		}
		if !isAvailable {
			return fmt.Errorf("warehouse is not available for product with code %s now", code)
		}
		var reservedQty int
		err = tx.QueryRow(ctx, "SELECT quantity FROM reserved WHERE code = $1", code).Scan(&reservedQty)
		if err != nil {
			return fmt.Errorf("no reservation found for product with code %s", code)
		}

		_, err = tx.Exec(ctx, "UPDATE reserved SET quantity = $1 WHERE code = $2", reservedQty-1, code)
		if err != nil {
			return err
		}

		if reservedQty-1 == 0 {
			_, err = tx.Exec(ctx, "DELETE FROM reserved WHERE code = $1", code)
			if err != nil {
				return err
			}
		}

		_, err = tx.Exec(ctx, "UPDATE products SET quantity = quantity + 1 WHERE code = $1", code)
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
	idInt, err := strconv.Atoi(warehouseID)
	if err != nil {
		// возвращаем статус bad request и сообщение об ошибке
		return nil, fmt.Errorf("invalid warehouse ID: %v", err)
	}

	rows, err := r.conn.Query(ctx, "SELECT * FROM products WHERE warehouse_id = $1", idInt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stock []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Size, &product.Code, &product.Quantity, &product.WarehouseID); err != nil {
			return nil, err
		}
		stock = append(stock, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stock, nil
}

func (r *Repository) GetReservedStock(ctx context.Context, warehouseID int) ([]model.Product, error) {
	rows, err := r.conn.Query(ctx, "SELECT * FROM reserved WHERE warehouse_id = $1", warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stock []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Size, &product.Code, &product.Quantity, &product.WarehouseID); err != nil {
			return nil, err
		}
		stock = append(stock, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stock, nil
}
