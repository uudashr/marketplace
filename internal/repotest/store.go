package repotest

import (
	"testing"

	"github.com/uudashr/marketplace/internal/store"

	"github.com/uudashr/marketplace/internal/fixture"
)

// SetupStoreFixtureFunc functions for setting up store fixture.
type SetupStoreFixtureFunc func(t *testing.T) StoreFixture

// StoreFixture is test fixture for store.
type StoreFixture interface {
	Repository() store.Repository
	TearDown()
}

// StoreSuite runs the repository test for store.
func StoreSuite(t *testing.T, setupFixture SetupStoreFixtureFunc) {
	t.Run("RetrieveStoredStore", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		str := fixture.Store()
		err := fix.Repository().Store(str)
		if err != nil {
			t.Fatal("err:", err)
		}

		retStr, err := fix.Repository().StoreByID(str.ID())
		if err != nil {
			t.Fatal("err:", err)
		}

		if got, want := retStr, str; !got.Equal(want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("EnsureUniqueCategoryName", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		str, err := store.New(store.NextID(), "My Mart")
		if err != nil {
			t.Fatal("err:", err)
		}

		err = fix.Repository().Store(str)
		if err != nil {
			t.Fatal("err:", err)
		}

		// Duplicate
		str, err = store.New(store.NextID(), "My Mart")
		if err != nil {
			t.Fatal("err:", err)
		}

		err = fix.Repository().Store(str)
		if err == nil {
			t.Error("Expect error on duplicate name")
		}
	})

	t.Run("EnsureStoredStoreOnTheList", func(t *testing.T) {
		fix := setupFixture(t)
		defer fix.TearDown()

		str := fixture.Store()
		err := fix.Repository().Store(str)
		if err != nil {
			t.Fatal("err:", err)
		}

		retStrs, err := fix.Repository().Stores()
		if err != nil {
			t.Fatal("err:", err)
		}

		var found bool
		for _, v := range retStrs {
			if got, want := v, str; got.Equal(want) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expect found %v on list %v", str, retStrs)
		}
	})
}
