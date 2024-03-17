package database

import (
	"database/sql"
	"server/internal/models"
	"time"
)

type ProductsTableActions interface {
}

type ProductsTable struct {
	Db *sql.DB
}

func (p ProductsTable) Insert(product models.Product) error {
	db := p.Db

	// Insert the product into the database.
	stmt, err := db.Prepare("INSERT INTO products (name, description, price, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Description, product.Price, time.Now())
	if err != nil {
		return err
	}

	return nil
}
