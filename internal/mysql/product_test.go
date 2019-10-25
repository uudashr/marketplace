// +build integration

package mysql_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/product"

	"github.com/uudashr/marketplace/internal/repotest"

	"github.com/uudashr/marketplace/internal/mysql"
)

func TestProductSuite(t *testing.T) {
	repotest.ProductSuite(t, func(t *testing.T) repotest.ProductFixture {
		dbFix := setupDBFixture(t)
		repo, err := mysql.NewProductRepository(dbFix.db)
		if err != nil {
			t.Fatal("err:", err)
		}

		return &productRepository{
			dbFix: dbFix,
			repo:  repo,
		}
	})
}

type productRepository struct {
	dbFix *dbFixture
	repo  *mysql.ProductRepository
}

func (fix *productRepository) Repository() product.Repository {
	return fix.repo
}

func (fix *productRepository) TearDown() {
	fix.dbFix.tearDown()
}
