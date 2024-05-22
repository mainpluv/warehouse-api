package api

import (
	"database/sql"
	"fmt"
)

// ReserveItems резервирует товары на складе.
func ReserveItems(db *sql.DB, warehouseID int, codes []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, code := range codes {
		result, err := tx.Exec(`
			UPDATE inventory
			SET quantity = quantity - 1
			WHERE item_id = (SELECT id FROM items WHERE code = $1)
			  AND warehouse_id = $2
			  AND quantity > 0
		`, code, warehouseID)
		if err != nil {
			return err
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil || rowsAffected == 0 {
			return fmt.Errorf("failed to reserve item with code %s", code)
		}
	}

	return tx.Commit()
}

// ReleaseItems освобождает резерв товаров на складе.
func ReleaseItems(db *sql.DB, warehouseID int, codes []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, code := range codes {
		_, err := tx.Exec(`
			UPDATE inventory
			SET quantity = quantity + 1
			WHERE item_id = (SELECT id FROM items WHERE code = $1)
			  AND warehouse_id = $2
		`, code, warehouseID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetRemainingItems получает количество оставшихся товаров на складе.
func GetRemainingItems(db *sql.DB, warehouseID int) ([]ItemQuantity, error) {
	rows, err := db.Query(`
		SELECT i.name, i.size, i.code, inv.quantity
		FROM items i
		JOIN inventory inv ON i.id = inv.item_id
		WHERE inv.warehouse_id = $1
	`, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []ItemQuantity
	for rows.Next() {
		var item ItemQuantity
		if err := rows.Scan(&item.Name, &item.Size, &item.Code, &item.Quantity); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

type ItemQuantity struct {
	Name     string
	Size     string
	Code     string
	Quantity int
}
