package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/uudashr/marketplace/internal/repotest"

	"github.com/uudashr/marketplace/internal/inmem"

	"github.com/uudashr/marketplace/internal/category"
)

func TestCategorySuite(t *testing.T) {
	suite.Run(t, &repotest.CategoryTestSuite{
		SetupFixture: repotest.SetupCategoryFixtureFunc(func(t *testing.T) repotest.CategoryFixture {
			repo := inmem.NewCategoryRepository()
			return &categoryFixture{
				repo: repo,
			}
		}),
	})
}

type categoryFixture struct {
	repo *inmem.CategoryRepository
}

func (fix *categoryFixture) Repository() category.Repository {
	return fix.repo
}

func (fix *categoryFixture) TearDown() {

}
