package store

import (
	"errors"

	"github.com/rs/xid"
	"github.com/shopspring/decimal"
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
func (s Store) OfferProduct(productID, categoryID string, price decimal.Decimal, quantity int) (*Product, error) {
	if quantity == 0 {
		return nil, errors.New("zero quantity")
	}

	return NewProduct(productID, s.ID(), categoryID, price, quantity)
}

// NextID returns unique id for store.
func NextID() string {
	return xid.New().String()
}

// Repository is repository for store.
type Repository interface {
	Store(*Store) error
	StoreByID(id string) (*Store, error)
}
