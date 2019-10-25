package repotest

import (
	"reflect"
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

		if got, want := retPrd, prd; !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
