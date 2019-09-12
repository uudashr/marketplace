package inmem

import (
	"errors"

	"github.com/uudashr/marketplace/internal/product"
)

// CategoryRepository is repository for product category.
type CategoryRepository struct {
	m       map[string]product.Category
	nameIdx map[string]string
}

// NewCategoryRepository constructs new product category repository.
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		m:       make(map[string]product.Category),
		nameIdx: make(map[string]string),
	}
}

// Store stores the product category.
func (r *CategoryRepository) Store(cat *product.Category) error {
	if _, exists := r.m[cat.ID()]; exists {
		return errors.New("already exists")
	}

	if _, exists := r.nameIdx[cat.Name()]; exists {
		return errors.New("already exists")
	}

	r.m[cat.ID()] = *cat
	r.nameIdx[cat.Name()] = cat.ID()
	return nil
}

// CategoryByID retrieves product category by ID.
func (r *CategoryRepository) CategoryByID(id string) (*product.Category, error) {
	cat, found := r.m[id]
	if !found {
		return nil, nil
	}

	return &cat, nil
}

// Categories retrieves product categories.
func (r *CategoryRepository) Categories() ([]*product.Category, error) {
	i := 0
	cats := make([]*product.Category, len(r.m))
	for _, v := range r.m {
		cat := v
		cats[i] = &cat
		i++
	}
	return cats, nil
}
