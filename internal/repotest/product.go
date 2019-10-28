package repotest

import (
	"testing"

	"github.com/uudashr/marketplace/internal/product"

	"github.com/uudashr/marketplace/internal/fixture"
)

// SetupProductFixtureFunc functions for setting up product fixture.
type SetupProductFixtureFunc func(t *testing.T) ProductFixture

// ProductFixture is test fixture for product.
type ProductFixture interface {
	Repository() product.Repository
	TearDown()
}

// ProductSuite runs the repository test for product.
func ProductSuite(t *testing.T, setupFixture SetupProductFixtureFunc) {
	t.Run("RetrieveStoredProduct", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		str := fixture.Store()
		prd := fixture.ProductOfStore(str)
		err := fix.Repository().Store(prd)
		if err != nil {
			t.Fatal("err:", err)
		}

		retPrd, err := fix.Repository().ProductByID(prd.ID())
		if err != nil {
			t.Fatal("err:", err)
		}

		if got, want := retPrd, prd; !got.Equal(want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("EnsureStoredProductOnTheList", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		str := fixture.Store()
		prd := fixture.ProductOfStore(str)
		err := fix.Repository().Store(prd)
		if err != nil {
			t.Fatal("err:", err)
		}

		retPrds, err := fix.Repository().Products()
		if err != nil {
			t.Fatal("err:", err)
		}

		if got := len(retPrds); got == 0 {
			t.Fatal("Expect not zero")
		}

		var found bool
		for _, v := range retPrds {
			if got, want := v, prd; got.Equal(want) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expect found %v on list %v", prd, retPrds)
		}
	})

	t.Run("ProducsOfStore", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		for i := 0; i < 2; i++ {
			str := fixture.Store()
			prds := fixture.ProductsOfStore(str, 3)

			for _, prd := range prds {
				err := fix.Repository().Store(prd)
				if err != nil {
					t.Fatal("err:", err)
				}
			}

			retPrds, err := fix.Repository().ProductsOfStore(str.ID())
			if err != nil {
				t.Fatal("err:", err)
			}

			if got, want := len(retPrds), len(prds); got != want {
				t.Errorf("Len got: %d, want: %d", got, want)
			}
		}

	})
}
