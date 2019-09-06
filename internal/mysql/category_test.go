// +build integration

package mysql_test

import (
	"testing"

	"github.com/uudashr/marketplace/internal/repotest"

	"github.com/uudashr/marketplace/internal/mysql"

	"github.com/uudashr/marketplace/internal/category"
)

func TestCategorySuite(t *testing.T) {
	repotest.CategorySuite(t, func(t *testing.T) repotest.CategoryFixture {
		dbFix := setupDBFixture(t)
		repo, err := mysql.NewCategoryRepository(dbFix.db)
		if err != nil {
			t.Fatal("err:", err)
		}

		return &categoryFixture{
			dbFix: dbFix,
			repo:  repo,
		}
	})
}

type categoryFixture struct {
	dbFix *dbFixture
	repo  *mysql.CategoryRepository
}

func (fix *categoryFixture) Repository() category.Repository {
	return fix.repo
}

func (fix *categoryFixture) TearDown() {
	fix.dbFix.tearDown()
}
