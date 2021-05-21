package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gitlab.com/tommzn-go/utils/common"
	"gitlab.com/tommzn-go/utils/config"
	testutils "gitlab.com/tommzn-go/utils/testing"
)

// Test suite for recipe service.
type RecipeServiceTestSuite struct {
	suite.Suite
	conf    config.Config
	service RecipeService
}

func TestRecipeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RecipeServiceTestSuite))
}

// Setup test, load config and create a recie service
func (suite *RecipeServiceTestSuite) SetupTest() {
	suite.Nil(config.UseConfigFileIfNotExists("testconfig"))
	suite.conf = loadConfigForTest()
	suite.service = NewRecipeServiceFromConfig(suite.conf, nil)
	tablename, region, endpoint := awsConfigForTest(suite.conf)
	suite.Nil(testutils.SetupTableForTest(tablename, region, endpoint))
}

// Tear down and delete DynamoDb table.
func (suite *RecipeServiceTestSuite) TearDownTest() {
	tablename, region, endpoint := awsConfigForTest(suite.conf)
	suite.Nil(testutils.TearDownTableForTest(tablename, region, endpoint))
}

// Test creating new recipe and assert a new message in queue.
func (suite *RecipeServiceTestSuite) TestCreateRecipe() {

	recipe := recipeForServiceTest()

	recipe2, err := suite.service.Create(recipe)
	suite.Nil(err)
	suite.True(common.IsId(recipe2.Id))
	suite.True(recipe2.CreatedAt.After(time.Now().Add(-5 * time.Second)))
	suite.assertQueueCount(1)
	suite.assertActionForMessage(0, RecipeAdded)
}

// Test update a recipe and assert a new message in queue.
func (suite *RecipeServiceTestSuite) TestUpdateRecipe() {

	recipe := recipeForServiceTest()

	recipe2, err := suite.service.Create(recipe)

	recipe2.Title = "New Title"
	err = suite.service.Update(recipe2)
	suite.Nil(err)

	suite.assertQueueCount(2)
	suite.assertActionForMessage(1, RecipeUpdated)

	recipe2.Id = "xxx"
	err = suite.service.Update(recipe2)
	suite.NotNil(err)

	suite.assertQueueCount(2)
}

// Test delete a recipe and assert a new message in queue.
func (suite *RecipeServiceTestSuite) TestDeleteRecipe() {

	recipe := recipeForServiceTest()

	recipe2, err := suite.service.Create(recipe)

	err = suite.service.Delete(recipe2)
	suite.Nil(err)

	suite.assertQueueCount(2)
	suite.assertActionForMessage(1, RecipeDeleted)
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
	suite.Len(suite.service.(*RecipeManager).publisher.(*PublisherMock).queue, expectedNumberOfMessages)
}

// Assert action for a recipe message.
func (suite *RecipeServiceTestSuite) assertActionForMessage(messageIndex int, expectedAction RecipeAction) {
	suite.Equal(expectedAction, suite.service.(*RecipeManager).publisher.(*PublisherMock).queue[messageIndex].Action)
}
