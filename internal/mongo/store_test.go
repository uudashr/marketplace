// +build integration

package mongo_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/store"

	"github.com/uudashr/marketplace/internal/mongo"

	"github.com/uudashr/marketplace/internal/repotest"
)

func TestStore(t *testing.T) {
	repotest.StoreSuite(t, func(t *testing.T) repotest.StoreFixture {
		dbFix := setupDBFixture(t)
		repo, err := mongo.NewStoreRepository(dbFix.db)
		if err != nil {
			t.Fatal("err:", err)
		}

		return &storeFixture{
			dbFix: dbFix,
			repo:  repo,
		}
	})
}

type storeFixture struct {
	dbFix *dbFixture
	repo  *mongo.StoreRepository
}

func (fix *storeFixture) Repository() store.Repository {
	return fix.repo
}

func (fix *storeFixture) TearDown() {
	fix.dbFix.tearDown()
}
