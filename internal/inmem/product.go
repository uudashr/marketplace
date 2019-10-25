package inmem

import (
	"errors"

	"github.com/uudashr/marketplace/internal/product"
)

// ProductRepository is repository for product.
type ProductRepository struct {
	m map[string]product.Product
}

// NewProductRepository constructs new ProductRepository.
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		m: make(map[string]product.Product),
	}
}

// Store stores the product.
func (r *ProductRepository) Store(prd *product.Product) error {
	_, found := r.m[prd.ID()]
	if found {
		return errors.New("already exists")
	}

	r.m[prd.ID()] = *prd
	return nil
}

// ProductByID retrieves product by id.
func (r *ProductRepository) ProductByID(id string) (*product.Product, error) {
	prd, found := r.m[id]
	if !found {
		return nil, nil
	}

	return &prd, nil
}
