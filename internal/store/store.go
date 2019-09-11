package store

import (
	"errors"

	"github.com/rs/xid"
	"github.com/shopspring/decimal"
	"github.com/uudashr/marketplace/internal/product"
)

// Store represents store.
type Store struct {
	id   string
	name string
}

// New constructs new store instance.
func New(id, name string) (*Store, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if name == "" {
		return nil, errors.New("empty name")
	}

	return &Store{
		id:   id,
		name: name,
	}, nil
}

// ID of the store.
func (s Store) ID() string {
	return s.id
}

// Name of the store.
func (s Store) Name() string {
	return s.name
}

// OfferProduct offers a product.
func (s Store) OfferProduct(productID, categoryID, name string, price decimal.Decimal, quantity int, description string) (*product.Product, error) {
	if quantity == 0 {
		return nil, errors.New("zero quantity")
	}

	return product.New(productID, s.ID(), categoryID, name, price, quantity, description)
}

// NextID returns unique id for store.
func NextID() string {
	return xid.New().String()
}

// Repository is repository for store.
//go:generate mockery -name=Repository
type Repository interface {
	Store(*Store) error
	StoreByID(id string) (*Store, error)
}
