package app

import (
	"errors"

	"github.com/uudashr/marketplace/internal/store"

	"github.com/uudashr/marketplace/internal/category"
)

// Service is the application service.
type Service struct {
	categoryRepo category.Repository
	storeRepo    store.Repository
}

// NewService constructs new service.
func NewService(categoryRepo category.Repository, storeRepo store.Repository) (*Service, error) {
	if categoryRepo == nil {
		return nil, errors.New("nil categoryRepo")
	}

	if storeRepo == nil {
		return nil, errors.New("nil storeRepo")
	}

	return &Service{
		categoryRepo: categoryRepo,
		storeRepo:    storeRepo,
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

// RetrieveCategories retrieves categories.
func (svc *Service) RetrieveCategories() ([]*category.Category, error) {
	return svc.categoryRepo.Categories()
}

// RetrieveCategoryByID retrieves category by id.
func (svc *Service) RetrieveCategoryByID(cmd RetrieveCategoryByIDCommand) (*category.Category, error) {
	return svc.categoryRepo.CategoryByID(cmd.ID)
}

// RegisterNewStore registers new store.
func (svc *Service) RegisterNewStore(cmd RegisterNewStoreCommand) (*store.Store, error) {
	str, err := store.New(store.NextID(), cmd.Name)
	if err != nil {
		return nil, err
	}

	if err = svc.storeRepo.Store(str); err != nil {
		return nil, err
	}

	return str, nil
}

// RegisterNewCategoryCommand command for registering new category.
type RegisterNewCategoryCommand struct {
	Name string
}

// RetrieveCategoryByIDCommand command for retrieving category by ID.
type RetrieveCategoryByIDCommand struct {
	ID string
}

// RegisterNewStoreCommand command for registering new store.
type RegisterNewStoreCommand struct {
	Name string
}
