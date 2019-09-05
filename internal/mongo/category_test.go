// +build integration

package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/uudashr/marketplace/internal/repotest"

	"github.com/uudashr/marketplace/internal/mongo"

	"github.com/uudashr/marketplace/internal/category"
)

func TestCategorySuite(t *testing.T) {
	suite.Run(t, &repotest.CategoryTestSuite{
		SetupFixture: repotest.SetupCategoryFixtureFunc(func(t *testing.T) repotest.CategoryFixture {
			dbFix := setupDBFixture(t)
			repo, err := mongo.NewCategoryRepository(dbFix.db)
			if err != nil {
				t.Fatal("err:", err)
			}

			return &categoryFixture{
				dbFix: dbFix,
				repo:  repo,
			}
		}),
	})
}

type categoryFixture struct {
	dbFix *dbFixture
	repo  *mongo.CategoryRepository
}

func (fix *categoryFixture) Repository() category.Repository {
	return fix.repo
}

func (fix *categoryFixture) TearDown() {
	fix.dbFix.tearDown()
}
