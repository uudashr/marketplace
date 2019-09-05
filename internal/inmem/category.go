package inmem

import (
	"errors"

	"github.com/uudashr/marketplace/internal/category"
)

// CategoryRepository is repository for Category.
type CategoryRepository struct {
	m       map[string]category.Category
	nameIdx map[string]string
}

// NewCategoryRepository constructs new categorty repository.
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		m:       make(map[string]category.Category),
		nameIdx: make(map[string]string),
	}
}

// Store category to repository.
func (r *CategoryRepository) Store(cat *category.Category) error {
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

// CategoryByID on the repository.
func (r *CategoryRepository) CategoryByID(id string) (*category.Category, error) {
	cat, found := r.m[id]
	if !found {
		return nil, nil
	}

	return &cat, nil
}

// Categories on the repository.
func (r *CategoryRepository) Categories() ([]*category.Category, error) {
	i := 0
	cats := make([]*category.Category, len(r.m))
	for _, v := range r.m {
		cat := v
		cats[i] = &cat
		i++
	}
	return cats, nil
}
