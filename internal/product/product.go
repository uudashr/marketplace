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
	description string
	quantity    int
}

// New constructs new store product instance.
func New(id, storeID, categoryID, name string, price decimal.Decimal, description string, quantity int) (*Product, error) {
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
		description: description,
		quantity:    quantity,
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

// Description of the product.
func (p Product) Description() string {
	return p.description
}

// Quantity of the product.
func (p Product) Quantity() int {
	return p.quantity
}

// Equal checks whether equal to p2.
func (p Product) Equal(p2 *Product) bool {
	return p.id == p2.id &&
		p.storeID == p2.storeID &&
		p.categoryID == p2.categoryID &&
		p.name == p2.name &&
		p.price.Equal(p2.price) &&
		p.description == p2.description &&
		p.quantity == p2.quantity
}

// NextID returns unique id for the product.
func NextID() string {
	return xid.New().String()
}

// Repository is repository for the product.
//go:generate mockery -name=Repository
type Repository interface {
	Store(*Product) error
	ProductByID(id string) (*Product, error)
	Products() ([]*Product, error)
	ProductsOfStore(storeID string) ([]*Product, error)
}
