// +build integration

package mongo_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/product"

	"github.com/uudashr/marketplace/internal/mongo"

	"github.com/uudashr/marketplace/internal/repotest"
)

func TestProduct(t *testing.T) {
	repotest.ProductSuite(t, func(t *testing.T) repotest.ProductFixture {
		dbFix := setupDBFixture(t)
		repo, err := mongo.NewProductRepository(dbFix.db)
		if err != nil {
			t.Fatal("err:", err)
		}

		return &productFixture{
			dbFix: dbFix,
			repo:  repo,
		}
	})
}

type productFixture struct {
	dbFix *dbFixture
	repo  *mongo.ProductRepository
}

func (fix *productFixture) Repository() product.Repository {
	return fix.repo
}

func (fix *productFixture) TearDown() {
	fix.dbFix.tearDown()
}
