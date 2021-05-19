package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gitlab.com/tommzn-go/utils/config"
	testutils "gitlab.com/tommzn-go/utils/testing"
)

// Test suite for DynamoDb repository.
// Runs tests for CRUD actions.
type RepositoryTestSuite struct {
	suite.Suite
	conf config.Config
	repo Repository
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

// Setup test. Loac config, create repository and init DynamoDb table.
func (suite *RepositoryTestSuite) SetupTest() {
	suite.Nil(config.UseConfigFileIfNotExists("testconfig"))
	suite.conf = loadConfigForTest()
	suite.repo = repositoryForTest(suite.conf)
	tablename, region, endpoint := suite.awsConfigForTest()
	suite.Nil(testutils.SetupTableForTest(tablename, region, endpoint))
}

// Tear down and delete DynamoDb table.
func (suite *RepositoryTestSuite) TearDownTest() {
	tablename, region, endpoint := suite.awsConfigForTest()
	suite.Nil(testutils.TearDownTableForTest(tablename, region, endpoint))
}

// Test adding a recipe. Assert recipe has been added to the repository
// and index is updated with recipe id.
func (suite *RepositoryTestSuite) TestAddRecipe() {

	recipe := recipeForTest()

	suite.Nil(suite.repo.Set(recipe))
	recipe2, err := suite.repo.Get(recipe.Id)
	suite.Nil(err)
	suite.NotNil(recipe2)
	suite.assertRecipeIsEqual(recipe, *recipe2)
	suite.assertIdExistsInIndex(recipe.Type, recipe.Id)
}

// Test updates for different recipe values.
func (suite *RepositoryTestSuite) TestUpdateRecipe() {

	recipe := recipeForTest()

	suite.Nil(suite.repo.Set(recipe))

	recipe.Title = "xxx"
	recipe.Ingredients = "yyy"
	recipe.Description = "zzz"
	suite.Nil(suite.repo.Set(recipe))

	recipe2, err := suite.repo.Get(recipe.Id)
	suite.Nil(err)
	suite.NotNil(recipe2)
	suite.assertRecipeIsEqual(recipe, *recipe2)
	suite.assertIdExistsInIndex(recipe.Type, recipe.Id)
}

// Test list recipes by type.
func (suite *RepositoryTestSuite) TestListRecipes() {

	recipe1_1 := recipeForTest()
	recipe1_2 := recipeForTest()
	recipe1_3 := recipeForTest()
	recipe2_1 := recipeForTest()
	recipe2_1.Type = CookingRecipe
	recipes := []Recipe{recipe1_1, recipe1_2, recipe1_3, recipe2_1}
	for _, recipe := range recipes {
		suite.Nil(suite.repo.Set(recipe))
	}

	suite.assertRecipeCountForType(BakingRecipe, 3)
	suite.assertRecipeCountForType(CookingRecipe, 1)
}

// Test delete a recipe.
func (suite *RepositoryTestSuite) TestDeleteRecipe() {

	recipe1 := recipeForTest()
	recipe2 := recipeForTest()
	recipes := []Recipe{recipe1, recipe2}
	for _, recipe := range recipes {
		suite.Nil(suite.repo.Set(recipe))
	}

	suite.assertRecipeCountForType(BakingRecipe, 2)
	suite.assertRecipeCountForType(CookingRecipe, 0)

	suite.repo.Delete(recipe2)
	suite.assertRecipeCountForType(BakingRecipe, 1)
	suite.assertRecipeCountForType(CookingRecipe, 0)
	suite.assertIdExistsInIndex(recipe1.Type, recipe1.Id)
	suite.assertIdNotExistsInIndex(recipe2.Type, recipe2.Id)
}

// Test update of recipe type. This should move the recipe id
// from one index to another.
func (suite *RepositoryTestSuite) TestUpdateRecipeType() {

	recipe := recipeForTest()

	suite.Nil(suite.repo.Set(recipe))

	recipe.Type = CookingRecipe
	suite.Nil(suite.repo.Set(recipe))

	recipe2, err := suite.repo.Get(recipe.Id)
	suite.Nil(err)
	suite.NotNil(recipe2)
	suite.assertRecipeIsEqual(recipe, *recipe2)
	suite.assertIdExistsInIndex(recipe.Type, recipe.Id)
	suite.assertIdNotExistsInIndex(BakingRecipe, recipe.Id)
}

// Assert that expected recipe is equal to current recipe.
// Uses time format RFC3339 to companre CreatedAt value.
func (suite *RepositoryTestSuite) assertRecipeIsEqual(expectedRecipe, recipe Recipe) {
	suite.Equal(expectedRecipe.Id, recipe.Id)
	suite.Equal(expectedRecipe.Type, recipe.Type)
	suite.Equal(expectedRecipe.Title, recipe.Title)
	suite.Equal(expectedRecipe.Ingredients, recipe.Ingredients)
	suite.Equal(expectedRecipe.Description, recipe.Description)
	suite.Equal(expectedRecipe.CreatedAt.Format(time.RFC3339), recipe.CreatedAt.Format(time.RFC3339))
}

// Fetch index for given type and assert passed id exists in.
func (suite *RepositoryTestSuite) assertIdExistsInIndex(recipeType RecipeType, id string) {

	recipeIndex := newRecipeIndex(recipeType)
	suite.Nil(suite.repo.(*DynamoDbRepository).client.Get(recipeIndex))
	_, ok := recipeIndex.Ids[id]
	suite.True(ok)
}

// Fetch index for given type and assert passed id not exists in.
func (suite *RepositoryTestSuite) assertIdNotExistsInIndex(recipeType RecipeType, id string) {

	recipeIndex := newRecipeIndex(recipeType)
	suite.Nil(suite.repo.(*DynamoDbRepository).client.Get(recipeIndex))
	_, ok := recipeIndex.Ids[id]
	suite.False(ok)
}

// Fetches all recipes by given type and assert expected number of them.
func (suite *RepositoryTestSuite) assertRecipeCountForType(recipeType RecipeType, expectedNumberOfRecipes int) {

	recipes, err := suite.repo.List(recipeType)
	if expectedNumberOfRecipes == 0 {
		suite.NotNil(err)
	} else {
		suite.Nil(err)
	}
	suite.Len(recipes, expectedNumberOfRecipes)
}

// Create a new repository for testing with passed config and
// a default stdout logger.
func repositoryForTest(conf config.Config) Repository {
	return NewRepository(conf, loggerForTest())
}

// Load DynamoDb settings from passed config
func (suite *RepositoryTestSuite) awsConfigForTest() (*string, *string, *string) {
	tablename := suite.conf.Get("aws.dynamodb.tablename", nil)
	region := suite.conf.Get("aws.dynamodb.awsregion", nil)
	endpoint := suite.conf.Get("aws.dynamodb.endpoint", nil)
	return tablename, region, endpoint
}
