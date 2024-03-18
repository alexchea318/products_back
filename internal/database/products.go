package database

import (
	"database/sql"
	"server/internal/models"
	"time"
)

type ProductsTableActions interface {
	Insert(product models.NewProduct) error
	Update(product models.NewProduct, id uint64) error
	Delete(id uint64) error
}

type ProductsTable struct {
	Db *sql.DB
}

func (p ProductsTable) Insert(product models.NewProduct) error {
	db := p.Db

	// Insert the product into the database.
	stmt, err := db.Prepare("INSERT INTO products (name, description, price, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	created := time.Now().Format(time.RFC3339)

	_, err = stmt.Exec(product.Name, product.Description, product.Price, created)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductsTable) Update(product models.NewProduct, id uint64) error {
	db := p.Db

	// Insert the product into the database.
	stmt, err := db.Prepare("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Description, product.Price, id)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductsTable) GetOne(id uint64) (models.Product, error) {
	db := p.Db

	// Get the product from the database.
	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	var product models.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
		return product, err
	}

	return product, nil
}

func (p ProductsTable) Delete(id uint64) error {
	db := p.Db

	// Delete the product from the database.
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (p ProductsTable) Get(search string, limit uint64, offset uint64) ([]models.Product, error) {
	db := p.Db
	var products []models.Product

	// Search for products in the database.
	rows, err := db.Query("SELECT * FROM products WHERE name LIKE ? ORDER BY created_at DESC LIMIT ? OFFSET ?", "%"+search+"%", limit, offset)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
			continue
		}
		products = append(products, product)
	}

	return products, nil
}
