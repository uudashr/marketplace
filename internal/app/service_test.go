package app_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"

	storemocks "github.com/uudashr/marketplace/internal/store/mocks"

	modelfixture "github.com/uudashr/marketplace/internal/fixture"

	"github.com/stretchr/testify/mock"

	"github.com/uudashr/marketplace/internal/app"
	prdmocks "github.com/uudashr/marketplace/internal/product/mocks"
)

func TestRegisterNewCategory(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	cmd := app.RegisterNewCategoryCommand{
		Name: "Electronic",
	}

	fix.categoryRepo.On("Store", mock.Anything).Return(nil)
	cat, err := fix.service.RegisterNewCategory(cmd)
	if err != nil {
		t.Fatal("err:", err)
	}

	if got, want := cat.Name(), cmd.Name; got != want {
		t.Errorf("Name got: %q, want: %q", got, want)
	}
}

func TestRetrieveCategories(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	cats := modelfixture.Categories(2)

	fix.categoryRepo.On("Categories").Return(cats, nil)
	retCats, err := fix.service.RetrieveCategories()
	if err != nil {
		t.Fatal("err:", err)
	}

	if got, want := retCats, cats; !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestRetrieveCategoryByID(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	cat := modelfixture.Category()
	fix.categoryRepo.On("CategoryByID", cat.ID()).Return(cat, nil)

	retCat, err := fix.service.RetrieveCategoryByID(app.RetrieveCategoryByIDCommand{
		ID: cat.ID(),
	})
	if err != nil {
		t.Fatal("err:", err)
	}

	if got, want := retCat, cat; got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestRegisterNewStore(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	cmd := app.RegisterNewStoreCommand{
		Name: "My Mart",
	}

	fix.storeRepo.On("Store", mock.Anything).Return(nil)
	str, err := fix.service.RegisterNewStore(cmd)
	if err != nil {
		t.Fatal("err:", err)
	}

	if got, want := str.Name(), cmd.Name; got != want {
		t.Errorf("Name got: %q, want: %q", got, want)
	}
}

func TestOfferNewProduct(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	str := modelfixture.Store()
	cat := modelfixture.Category()

	cmd := app.OfferNewProductCommand{
		StoreID:     str.ID(),
		CategoryID:  cat.ID(),
		Name:        "Sony Wf 1000MX3",
		Price:       decimal.NewFromFloat(3499000),
		Description: "Trully Wireless Earbud with Noise Cancelling",
		Quantity:    10,
	}

	fix.storeRepo.On("StoreByID", cmd.StoreID).Return(str, nil)
	fix.categoryRepo.On("CategoryByID", cmd.CategoryID).Return(cat, nil)
	fix.productRepo.On("Store", mock.Anything).Return(nil)

	prd, err := fix.service.OfferNewProduct(cmd)
	if err != nil {
		t.Fatal("err:", err)
	}

	if got := prd.ID(); got == "" {
		t.Errorf("ID got: %q, want not empty", got)
	}

	if got, want := prd.StoreID(), cmd.StoreID; got != want {
		t.Errorf("StoreID got: %q, want: %q", got, want)
	}

	if got, want := prd.CategoryID(), cmd.CategoryID; got != want {
		t.Errorf("CategoryID got: %q, want: %q", got, want)
	}

	if got, want := prd.Name(), cmd.Name; got != want {
		t.Errorf("Name got: %q, want: %q", got, want)
	}

	if got, want := prd.Price(), cmd.Price; !got.Equal(want) {
		t.Errorf("Price got: %q, want: %q", got, want)
	}

	if got, want := prd.Description(), cmd.Description; got != want {
		t.Errorf("Description got: %q, want: %q", got, want)
	}

	if got, want := prd.Quantity(), cmd.Quantity; got != want {
		t.Errorf("Name got: %d, want: %d", got, want)
	}
}

type testFixture struct {
	t            *testing.T
	categoryRepo *prdmocks.CategoryRepository
	storeRepo    *storemocks.Repository
	productRepo  *prdmocks.Repository
	service      *app.Service
}

func setupFixture(t *testing.T) *testFixture {
	categoryRepo := new(prdmocks.CategoryRepository)
	storeRepo := new(storemocks.Repository)
	productRepo := new(prdmocks.Repository)
	svc, err := app.NewService(categoryRepo, storeRepo, productRepo)
	if err != nil {
		t.Fatal(fmt.Errorf("fail to create Service: %w", err))
	}

	return &testFixture{
		t:            t,
		categoryRepo: categoryRepo,
		storeRepo:    storeRepo,
		productRepo:  productRepo,
		service:      svc,
	}
}

func (fix *testFixture) tearDown() {
	mock.AssertExpectationsForObjects(fix.t,
		fix.categoryRepo, fix.storeRepo, fix.productRepo)
}
