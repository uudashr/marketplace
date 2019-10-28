package app

import (
	"errors"

	"github.com/shopspring/decimal"

	"github.com/uudashr/marketplace/internal/store"

	"github.com/uudashr/marketplace/internal/product"
)

// Service is the application service.
type Service struct {
	categoryRepo product.CategoryRepository
	storeRepo    store.Repository
	productRepo  product.Repository
}

// NewService constructs new service.
func NewService(categoryRepo product.CategoryRepository, storeRepo store.Repository, productRepo product.Repository) (*Service, error) {
	if categoryRepo == nil {
		return nil, errors.New("nil categoryRepo")
	}

	if storeRepo == nil {
		return nil, errors.New("nil storeRepo")
	}

	if productRepo == nil {
		return nil, errors.New("nil productRepo")
	}

	return &Service{
		categoryRepo: categoryRepo,
		storeRepo:    storeRepo,
		productRepo:  productRepo,
	}, nil
}

// RegisterNewCategory registers new category.
func (svc *Service) RegisterNewCategory(cmd RegisterNewCategoryCommand) (*product.Category, error) {
	cat, err := product.NewCategory(product.NextCategoryID(), cmd.Name)
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
func (svc *Service) RetrieveCategories() ([]*product.Category, error) {
	return svc.categoryRepo.Categories()
}

// RetrieveCategoryByID retrieves category by id.
func (svc *Service) RetrieveCategoryByID(cmd RetrieveCategoryByIDCommand) (*product.Category, error) {
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

// RetrieveStores retrieves stores.
func (svc *Service) RetrieveStores() ([]*store.Store, error) {
	return svc.storeRepo.Stores()
}

// RetrieveStoreByID retrieves store by id.
func (svc *Service) RetrieveStoreByID(cmd RetrieveStoreByIDCommand) (*store.Store, error) {
	return svc.storeRepo.StoreByID(cmd.ID)
}

// OfferNewProduct offers new product.
func (svc *Service) OfferNewProduct(cmd OfferNewProductCommand) (*product.Product, error) {
	str, err := svc.storeRepo.StoreByID(cmd.StoreID)
	if err != nil {
		return nil, err
	}

	if str == nil {
		return nil, errors.New("store not found")
	}

	cat, err := svc.categoryRepo.CategoryByID(cmd.CategoryID)
	if err != nil {
		return nil, err
	}

	if cat == nil {
		return nil, errors.New("category not found")
	}

	prd, err := str.OfferProduct(product.NextID(), cat, cmd.Name, cmd.Price, cmd.Description, cmd.Quantity)
	if err != nil {
		return nil, err
	}

	err = svc.productRepo.Store(prd)
	if err != nil {
		return nil, err
	}

	return prd, nil
}

// RetrieveProducts retrieves products.
func (svc *Service) RetrieveProducts() ([]*product.Product, error) {
	return svc.productRepo.Products()
}

// RetrieveProductByID retrieves product by id.
func (svc *Service) RetrieveProductByID(cmd RetrieveProductByIDCommand) (*product.Product, error) {
	return svc.productRepo.ProductByID(cmd.ID)
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

// RetrieveStoreByIDCommand command for retrieving store by ID.
type RetrieveStoreByIDCommand struct {
	ID string
}

// OfferNewProductCommand command for offering new product.
type OfferNewProductCommand struct {
	StoreID     string
	CategoryID  string
	Name        string
	Price       decimal.Decimal
	Description string
	Quantity    int
}

// RetrieveProductByIDCommand command for retrieving product by ID.
type RetrieveProductByIDCommand struct {
	ID string
}

// RetrieveStoreProductsCommand command for retrieving store's products.
type RetrieveStoreProductsCommand struct {
	StoreID string
}
