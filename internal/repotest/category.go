package repotest

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/uudashr/marketplace/internal/category"
)

// SetupCategoryFixtureFunc functions for setting up category fixture.
type SetupCategoryFixtureFunc func(t *testing.T) CategoryFixture

// CategoryFixture is test fixture for category.
type CategoryFixture interface {
	Repository() category.Repository
	TearDown()
}

// CategoryTestSuite is the test suite for category repository.
type CategoryTestSuite struct {
	suite.Suite
	SetupFixture SetupCategoryFixtureFunc
	fix          CategoryFixture
}

// SetupTest setup function, runs before each test.
func (suite *CategoryTestSuite) SetupTest() {
	suite.fix = suite.SetupFixture(suite.T())
}

// TearDownTest tear down function , runs before each test.
func (suite *CategoryTestSuite) TearDownTest() {
	suite.fix.TearDown()
}

// TestRetrieveStoredCategory test whether we can retrieve stored category.
func (suite *CategoryTestSuite) TestRetrieveStoredCategory() {
	cat, err := category.New(category.NextID(), "Utilities")
	suite.Require().NoError(err)

	err = suite.fix.Repository().Store(cat)
	suite.Require().NoError(err)

	retCat, err := suite.fix.Repository().CategoryByID(cat.ID())
	suite.Require().NoError(err)
	suite.Assert().Equal(cat, retCat)
}

// TestUniqueCategoryName test whether the category name is unique.
func (suite *CategoryTestSuite) TestUniqueCategoryName() {
	cat, err := category.New(category.NextID(), "Utilities")
	suite.Require().NoError(err)

	err = suite.fix.Repository().Store(cat)
	suite.Require().NoError(err)

	// Duplicate
	cat, err = category.New(category.NextID(), "Utilities")
	suite.Require().NoError(err)

	err = suite.fix.Repository().Store(cat)
	suite.Require().Error(err)
}

// TestStoredCategoryOnTheList test whether stored category shows on the list.
func (suite *CategoryTestSuite) TestStoredCategoryOnTheList() {
	cat, err := category.New(category.NextID(), "Utilities")
	suite.Require().NoError(err)

	err = suite.fix.Repository().Store(cat)
	suite.Require().NoError(err)

	retCats, err := suite.fix.Repository().Categories()
	suite.Require().NoError(err)
	suite.Require().Contains(retCats, cat)
}
