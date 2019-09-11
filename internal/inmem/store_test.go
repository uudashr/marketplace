package inmem_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/repotest"
	"github.com/uudashr/marketplace/internal/store"

	"github.com/uudashr/marketplace/internal/inmem"
)

func TestStoreSuite(t *testing.T) {
	repotest.StoreSuite(t, func(t *testing.T) repotest.StoreFixture {
		repo := inmem.NewStoreRepository()
		return &storeFixture{
			repo: repo,
		}
	})
}

type storeFixture struct {
	repo *inmem.StoreRepository
}

func (fix *storeFixture) Repository() store.Repository {
	return fix.repo
}

func (fix *storeFixture) TearDown() {

}
