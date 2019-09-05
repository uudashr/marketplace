package app

import (
	"errors"

	"github.com/uudashr/marketplace/internal/category"
)

// Service is the application service.
type Service struct {
	categoryRepo category.Repository
}

// NewService constructs new service.
func NewService(categoryRepo category.Repository) (*Service, error) {
	if categoryRepo == nil {
		return nil, errors.New("nil categoryRepo")
	}

	return &Service{
		categoryRepo: categoryRepo,
	}, nil
}

// RegisterNewCategory registers new category.
func (svc *Service) RegisterNewCategory(cmd RegisterNewCategoryCommand) (*category.Category, error) {
	cat, err := category.New(category.NextID(), cmd.Name)
	if err != nil {
		return nil, err
	}

	err = svc.categoryRepo.Store(cat)
	if err != nil {
		return nil, err
	}

	return cat, nil
}

// RetrieveCategories retrieves categories on the system.
func (svc *Service) RetrieveCategories() ([]*category.Category, error) {
	return svc.categoryRepo.Categories()
}

// RetrieveCategoryByID retrieves category on the system.
func (svc *Service) RetrieveCategoryByID(cmd RetrieveCategoryByIDCommand) (*category.Category, error) {
	return svc.categoryRepo.CategoryByID(cmd.ID)
}

// RegisterNewCategoryCommand is command for registering new category.
type RegisterNewCategoryCommand struct {
	Name string
}

// RetrieveCategoryByIDCommand is command for retrieving category by ID.
type RetrieveCategoryByIDCommand struct {
	ID string
}
