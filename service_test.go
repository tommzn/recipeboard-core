package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	mock "github.com/tommzn/recipeboard-core/mock"
	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/utils/common"
	"gitlab.com/tommzn-go/utils/config"
)

// Test suite for recipe service.
type RecipeServiceTestSuite struct {
	suite.Suite
	conf           config.Config
	repositoryMock *mock.RepositoryMock
	publisherMock  *mock.PublisherMock
	service        RecipeService
}

func TestRecipeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RecipeServiceTestSuite))
}

// Setup test, load config and create a recie service
func (suite *RecipeServiceTestSuite) SetupTest() {
	suite.repositoryMock = mock.NewRepository()
	suite.publisherMock = mock.NewPublisher()
	suite.service = NewRecipeService(suite.repositoryMock, suite.publisherMock, nil)
}

// Test creating new recipe and assert a new message in queue.
func (suite *RecipeServiceTestSuite) TestCreateRecipe() {

	recipe := recipeForServiceTest()

	recipe2, err := suite.service.Create(recipe)
	suite.Nil(err)
	suite.True(common.IsId(recipe2.Id))
	suite.True(recipe2.CreatedAt.After(time.Now().Add(-5 * time.Second)))
	suite.assertQueueCount(1)
	suite.assertActionForMessage(0, model.RecipeAdded)
}

// Test update a recipe and assert a new message in queue.
func (suite *RecipeServiceTestSuite) TestUpdateRecipe() {

	recipe := recipeForServiceTest()

	recipe2, err := suite.service.Create(recipe)
	suite.Nil(err)

	recipe2.Title = "New Title"
	err = suite.service.Update(recipe2)
	suite.Nil(err)

	suite.assertQueueCount(2)
	suite.assertActionForMessage(1, model.RecipeUpdated)

	recipe2.Id = "xxx"
	err = suite.service.Update(recipe2)
	suite.NotNil(err)

	suite.assertQueueCount(2)
}

// Test delete a recipe and assert a new message in queue.
func (suite *RecipeServiceTestSuite) TestDeleteRecipe() {

	recipe := recipeForServiceTest()

	recipe2, err := suite.service.Create(recipe)
	suite.Nil(err)

	err = suite.service.Delete(recipe2)
	suite.Nil(err)

	suite.assertQueueCount(2)
	suite.assertActionForMessage(1, model.RecipeDeleted)
}

// Test list existing recipes.
func (suite *RecipeServiceTestSuite) TestListRecipes() {

	recipe := recipeForServiceTest()

	recipe2, err := suite.service.Create(recipe)
	suite.Nil(err)

	recipes, err := suite.service.List(recipe.Type)
	suite.Nil(err)
	suite.Len(recipes, 1)
	suite.Equal(recipe2.Id, recipes[0].Id)
}

// Assert message queue in publisher has expected count.
func (suite *RecipeServiceTestSuite) assertQueueCount(expectedNumberOfMessages int) {
	suite.Len(suite.publisherMock.Queue, expectedNumberOfMessages)
}

// Assert action for a recipe message.
func (suite *RecipeServiceTestSuite) assertActionForMessage(messageIndex int, expectedAction model.RecipeAction) {
	suite.Equal(expectedAction, suite.publisherMock.Queue[messageIndex].Action)
}
