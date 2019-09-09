package app_test

import (
	"fmt"
	"reflect"
	"testing"

	modelfixture "github.com/uudashr/marketplace/internal/fixture"

	"github.com/stretchr/testify/mock"

	"github.com/uudashr/marketplace/internal/app"
	catmocks "github.com/uudashr/marketplace/internal/category/mocks"
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

type testFixture struct {
	t       *testing.T
	catRepo *catmocks.Repository
	service *app.Service
}

func setupFixture(t *testing.T) *testFixture {
	catRepo := new(catmocks.Repository)
	svc, err := app.NewService(catRepo)
	if err != nil {
		t.Fatal(fmt.Errorf("fail to create Service: %w", err))
	}

	return &testFixture{
		t:       t,
		catRepo: catRepo,
		service: svc,
	}
}

func (fix *testFixture) tearDown() {
	mock.AssertExpectationsForObjects(fix.t,
		fix.catRepo)
}
