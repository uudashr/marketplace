package product

import (
	"errors"

	"github.com/rs/xid"
	"github.com/shopspring/decimal"
)

// Product represents product offered by a store.
type Product struct {
	id          string
	storeID     string
	categoryID  string
	name        string
	price       decimal.Decimal
	quantity    int
	description string
}

// New constructs new store product instance.
func New(id, storeID, categoryID, name string, price decimal.Decimal, quantity int, description string) (*Product, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if categoryID == "" {
		return nil, errors.New("empty categoryID")
	}

	if name == "" {
		return nil, errors.New("empty name")
	}

	if price.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("zero or negative price")
	}

	if quantity < 0 {
		return nil, errors.New("negative quantity")
	}

	return &Product{
		id:          id,
		storeID:     storeID,
		categoryID:  categoryID,
		name:        name,
		price:       price,
		quantity:    quantity,
		description: description,
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

// Name of the product.
func (p Product) Name() string {
	return p.name
}

// Price of the store product.
func (p Product) Price() decimal.Decimal {
	return p.price
}

// Quantity of the product.
func (p Product) Quantity() int {
	return p.quantity
}

// Description of the product.
func (p Product) Description() string {
	return p.description
}

// NextID returns unique id for the product.
func NextID() string {
	return xid.New().String()
}

// ProductRepository is repository for the product.
type ProductRepository interface {
	Store(*Product) error
	ProductsByStoreID(storeID string) ([]*Product, error)
}
