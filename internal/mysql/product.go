package mysql

import (
	"database/sql"
	"errors"

	"github.com/shopspring/decimal"

	"github.com/uudashr/marketplace/internal/product"
)

// ProductRepository is repository for product.
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository constructs new product repository.
func NewProductRepository(db *sql.DB) (*ProductRepository, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	return &ProductRepository{
		db: db,
	}, nil
}

// Store stores product to repository.
func (r *ProductRepository) Store(prd *product.Product) error {
	res, err := r.db.Exec("INSERT INTO products (id, store_id, category_id, name, price, description, quantity) VALUES (?, ?, ?, ?, ?, ?, ?)", prd.ID(), prd.StoreID(), prd.CategoryID(), prd.Name(), prd.Price().String(), prd.Description(), prd.Quantity())
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

// ProductByID retrieves product by id.
func (r *ProductRepository) ProductByID(id string) (*product.Product, error) {
	var (
		storeID     string
		categoryID  string
		name        string
		price       decimal.Decimal
		description string
		quantity    int
	)

	err := r.db.QueryRow("SELECT store_id, category_id, name, price, description, quantity FROM products WHERE id = ?", id).Scan(
		&storeID,
		&categoryID,
		&name,
		&price,
		&description,
		&quantity,
	)
	if err != nil {
		return nil, err
	}

	return product.New(id, storeID, categoryID, name, price, description, quantity)
}

// Products retrieves products.
func (r *ProductRepository) Products() ([]*product.Product, error) {
	rows, err := r.db.Query("SELECT id, store_id, category_id, name, price, description, quantity FROM products")
	if err != nil {
		return nil, err
	}

	var out []*product.Product
	for rows.Next() {
		var (
			id          string
			storeID     string
			categoryID  string
			name        string
			price       decimal.Decimal
			description string
			quantity    int
		)
		if err := rows.Scan(
			&id,
			&storeID,
			&categoryID,
			&name,
			&price,
			&description,
			&quantity,
		); err != nil {
			return nil, err
		}

		prd, err := product.New(id, storeID, categoryID, name, price, description, quantity)
		if err != nil {
			return nil, err
		}

		out = append(out, prd)
	}

	return out, nil
}
