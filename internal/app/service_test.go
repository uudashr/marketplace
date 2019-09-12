package app_test

import (
	"fmt"
	"reflect"
	"testing"

	storemocks "github.com/uudashr/marketplace/internal/store/mocks"

	modelfixture "github.com/uudashr/marketplace/internal/fixture"

	"github.com/stretchr/testify/mock"

	"github.com/uudashr/marketplace/internal/app"
	prodmocks "github.com/uudashr/marketplace/internal/product/mocks"
)

func TestRegisterNewCategory(t *testing.T) {
	fix := setupFixture(t)
	defer fix.tearDown()

	cmd := app.RegisterNewCategoryCommand{
		Name: "Electronic",
	}

	fix.catRepo.On("Store", mock.Anything).Return(nil)
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

	fix.catRepo.On("Categories").Return(cats, nil)
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
	fix.catRepo.On("CategoryByID", cat.ID()).Return(cat, nil)

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

type testFixture struct {
	t         *testing.T
	catRepo   *prodmocks.CategoryRepository
	storeRepo *storemocks.Repository
	service   *app.Service
}

func setupFixture(t *testing.T) *testFixture {
	catRepo := new(prodmocks.CategoryRepository)
	storeRepo := new(storemocks.Repository)
	svc, err := app.NewService(catRepo, storeRepo)
	if err != nil {
		t.Fatal(fmt.Errorf("fail to create Service: %w", err))
	}

	return &testFixture{
		t:         t,
		catRepo:   catRepo,
		storeRepo: storeRepo,
		service:   svc,
	}
}

func (fix *testFixture) tearDown() {
	mock.AssertExpectationsForObjects(fix.t,
		fix.catRepo)
}
