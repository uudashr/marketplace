package product

import (
	"errors"

	"github.com/rs/xid"
)

// Category represents product category.
type Category struct {
	id   string
	name string
}

// NewCategory constructs new category instance.
func NewCategory(id, name string) (*Category, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	if name == "" {
		return nil, errors.New("empty name")
	}

	return &Category{
		id:   id,
		name: name,
	}, nil
}

// ID of the category.
func (c Category) ID() string {
	return c.id
}

// Name of the category.
func (c Category) Name() string {
	return c.name
}

// NextID returns unique id for category.
func NextCategoryID() string {
	return xid.New().String()
}

// CategoryRepository is repository for category.
//go:generate mockery -name=CategoryRepository
type CategoryRepository interface {
	Store(*Category) error
	CategoryByID(id string) (*Category, error)
	Categories() ([]*Category, error)
}
