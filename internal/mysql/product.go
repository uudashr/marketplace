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
	var pb productBuilder
	row := r.db.QueryRow("SELECT id, store_id, category_id, name, price, description, quantity FROM products WHERE id = ?", id)
	return pb.build(row)
}

// Products retrieves products.
func (r *ProductRepository) Products() ([]*product.Product, error) {
	rows, err := r.db.Query("SELECT id, store_id, category_id, name, price, description, quantity FROM products")
	if err != nil {
		return nil, err
	}

	var out []*product.Product
	for rows.Next() {
		var b productBuilder
		prd, err := b.build(rows)
		if err != nil {
			return nil, err
		}

		out = append(out, prd)
	}

	return out, nil
}

// ProductsOfStore retrieves products.
func (r *ProductRepository) ProductsOfStore(storeID string) ([]*product.Product, error) {
	rows, err := r.db.Query("SELECT id, store_id, category_id, name, price, description, quantity FROM products WHERE store_id = ?", storeID)
	if err != nil {
		return nil, err
	}

	var out []*product.Product
	for rows.Next() {
		var b productBuilder
		prd, err := b.build(rows)
		if err != nil {
			return nil, err
		}

		out = append(out, prd)
	}

	return out, nil
}

type productBuilder struct {
	id          string
	storeID     string
	categoryID  string
	name        string
	price       decimal.Decimal
	description string
	quantity    int
}

func (b productBuilder) build(s scanner) (*product.Product, error) {
	if err := s.Scan(
		&b.id,
		&b.storeID,
		&b.categoryID,
		&b.name,
		&b.price,
		&b.description,
		&b.quantity,
	); err != nil {
		return nil, err
	}

	return product.New(b.id, b.storeID, b.categoryID, b.name, b.price, b.description, b.quantity)
}
