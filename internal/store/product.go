package store

import (
	"errors"

	"github.com/rs/xid"
	"github.com/shopspring/decimal"
)

// Product represents product offered by a store.
type Product struct {
	id         string
	storeID    string
	categoryID string
	price      decimal.Decimal
	quantity   int
}

// NewProduct constructs new store product instance.
func NewProduct(id, storeID, categoryID string, price decimal.Decimal, quantity int) (*Product, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if categoryID == "" {
		return nil, errors.New("empty categoryID")
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("zero or negative price")
	}

	if quantity < 0 {
		return nil, errors.New("negative quantity")
	}

	return &Product{
		id:         id,
		storeID:    storeID,
		categoryID: categoryID,
		price:      price,
		quantity:   quantity,
	}, nil
}

// ID of the product.
func (p Product) ID() string {
	return p.id
}

// StoreID of the product.
func (p Product) StoreID() string {
	return p.storeID
}

// CategoryID of the product.
func (p Product) CategoryID() string {
	return p.categoryID
}

// Price of the store product.
func (p Product) Price() decimal.Decimal {
	return p.price
}

// Quantity of the product.
func (p Product) Quantity() int {
	return p.quantity
}

// NextProductID returns unique id for the product.
func NextProductID() string {
	return xid.New().String()
}

// ProductRepository is repository for the product.
type ProductRepository interface {
	Store(*Product) error
	ProductsByStoreID(storeID string) ([]*Product, error)
}
