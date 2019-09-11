package inmem

import (
	"errors"

	"github.com/uudashr/marketplace/internal/store"
)

// StoreRepository is repository for Store.
type StoreRepository struct {
	m       map[string]store.Store
	nameIdx map[string]string
}

// NewStoreRepository constructs new store repository.
func NewStoreRepository() *StoreRepository {
	return &StoreRepository{
		m:       make(map[string]store.Store),
		nameIdx: make(map[string]string),
	}
}

// Store stores/puts store.
func (r *StoreRepository) Store(str *store.Store) error {
	if _, exists := r.m[str.ID()]; exists {
		return errors.New("already exists")
	}

	if _, exists := r.nameIdx[str.Name()]; exists {
		return errors.New("already exists")
	}

	r.m[str.ID()] = *str
	r.nameIdx[str.Name()] = str.ID()
	return nil
}

// StoreByID retrieves store by ID.
func (r *StoreRepository) StoreByID(id string) (*store.Store, error) {
	str, found := r.m[id]
	if !found {
		return nil, nil
	}

	return &str, nil
}
