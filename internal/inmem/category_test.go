package inmem_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/repotest"

	"github.com/uudashr/marketplace/internal/inmem"

	"github.com/uudashr/marketplace/internal/product"
)

func TestCategorySuite(t *testing.T) {
	repotest.CategorySuite(t, func(t *testing.T) repotest.CategoryFixture {
		repo := inmem.NewCategoryRepository()
		return &categoryFixture{
			repo: repo,
		}
	})
}

type categoryFixture struct {
	repo *inmem.CategoryRepository
}

func (fix *categoryFixture) Repository() product.CategoryRepository {
	return fix.repo
}

func (fix *categoryFixture) TearDown() {

}
