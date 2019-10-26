package repotest

import (
	"testing"

	"github.com/uudashr/marketplace/internal/fixture"
	"github.com/uudashr/marketplace/internal/product"
)

// SetupCategoryFixtureFunc functions for setting up category fixture.
type SetupCategoryFixtureFunc func(t *testing.T) CategoryFixture

// CategoryFixture is test fixture for category.
type CategoryFixture interface {
	Repository() product.CategoryRepository
	TearDown()
}

// CategorySuite runs the repository test for category.
func CategorySuite(t *testing.T, setupFixture SetupCategoryFixtureFunc) {
	t.Run("RetrieveStoredCategory", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		cat := fixture.Category()
		err := fix.Repository().Store(cat)
		if err != nil {
			t.Fatal("err:", err)
		}

		retCat, err := fix.Repository().CategoryByID(cat.ID())
		if err != nil {
			t.Fatal("err:", err)
		}

		if got, want := retCat, cat; !got.Equal(want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("EnsureUniqueCategoryName", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		cat, err := product.NewCategory(product.NextCategoryID(), "Utilities")
		if err != nil {
			t.Fatal("err:", err)
		}

		err = fix.Repository().Store(cat)
		if err != nil {
			t.Fatal("err:", err)
		}

		// Duplicate
		cat, err = product.NewCategory(product.NextCategoryID(), "Utilities")
		if err != nil {
			t.Fatal("err:", err)
		}

		err = fix.Repository().Store(cat)
		if err == nil {
			t.Error("Expect error on duplicate name")
		}
	})

	t.Run("EnsureStoredCategoryOnTheList", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		cat := fixture.Category()
		err := fix.Repository().Store(cat)
		if err != nil {
			t.Fatal("err:", err)
		}

		retCats, err := fix.Repository().Categories()
		if err != nil {
			t.Fatal("err:", err)
		}

		var found bool
		for _, v := range retCats {
			if got, want := v, cat; got.Equal(want) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expect found %v on list %v", cat, retCats)
		}
	})
}
