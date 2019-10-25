package inmem_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/product"

	"github.com/uudashr/marketplace/internal/repotest"

	"github.com/uudashr/marketplace/internal/inmem"
)

func TestProductSuite(t *testing.T) {
	repotest.ProductSuite(t, func(t *testing.T) repotest.ProductFixture {
		repo := inmem.NewProductRepository()
		return &productFixture{
			repo: repo,
		}
	})
}

type productFixture struct {
	repo *inmem.ProductRepository
}

func (fix *productFixture) Repository() product.Repository {
	return fix.repo
}

func (fix *productFixture) TearDown() {

}
