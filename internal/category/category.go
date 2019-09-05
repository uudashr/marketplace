package category

import (
	"errors"

	"github.com/rs/xid"
)

// Category represents product category.
type Category struct {
	id   string
	name string
}

// New constructs new category instance.
func New(id, name string) (*Category, error) {
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
func NextID() string {
	return xid.New().String()
}

// Repository is repository for category.
type Repository interface {
	Store(*Category) error
	CategoryByID(id string) (*Category, error)
	Categories() ([]*Category, error)
}
