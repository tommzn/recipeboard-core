package mock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	model "github.com/tommzn/recipeboard-core/model"
)

// Test suite for repository mock.
type RepositoryMockTestSuite struct {
	suite.Suite
}

func TestRepositoryMockTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryMockTestSuite))
}

// Test adding a recipe. Assert recipe has been added to the repository
// and index is updated with recipe id.
func (suite *RepositoryMockTestSuite) TestAddRecipe() {

	repo := NewRepository()
	recipe := recipeForTest()

	suite.Nil(repo.Set(recipe))
	recipe2, err := repo.Get(recipe.Id)
	suite.Nil(err)
	suite.NotNil(recipe2)
	suite.assertRecipeIsEqual(recipe, *recipe2)
}

// Test updates for different recipe values.
func (suite *RepositoryMockTestSuite) TestUpdateRecipe() {

	repo := NewRepository()
	recipe := recipeForTest()

	suite.Nil(repo.Set(recipe))

	recipe.Title = "xxx"
	recipe.Ingredients = "yyy"
	recipe.Description = "zzz"
	suite.Nil(repo.Set(recipe))

	recipe2, err := repo.Get(recipe.Id)
	suite.Nil(err)
	suite.NotNil(recipe2)
	suite.assertRecipeIsEqual(recipe, *recipe2)
}

// Test list recipes by type.
func (suite *RepositoryMockTestSuite) TestListRecipes() {

	repo := NewRepository()
	recipe1 := recipeForTest()
	recipe2 := recipeForTest()
	recipes := []model.Recipe{recipe1, recipe2}
	for _, recipe := range recipes {
		suite.Nil(repo.Set(recipe))
	}

	recipes, err := repo.List(model.BakingRecipe)
	suite.Nil(err)
	suite.Len(recipes, 2)
	suite.Equal(recipe1.Id, recipes[0].Id)
	suite.Equal(recipe2.Id, recipes[1].Id)

	recipes2, err := repo.List(model.CookingRecipe)
	suite.NotNil(err)
	suite.Len(recipes2, 0)
}

// Test delete a recipe.
func (suite *RepositoryMockTestSuite) TestDeleteRecipe() {

	repo := NewRepository()
	recipe1 := recipeForTest()
	recipe2 := recipeForTest()
	recipes := []model.Recipe{recipe1, recipe2}
	for _, recipe := range recipes {
		suite.Nil(repo.Set(recipe))
	}

	recipes1, _ := repo.List(model.BakingRecipe)
	suite.Len(recipes1, 2)

	suite.Nil(repo.Delete(recipe1))
	recipes2, _ := repo.List(model.BakingRecipe)
	suite.Len(recipes2, 1)
	suite.Equal(recipe2.Id, recipes2[0].Id)
}

// Assert that expected recipe is equal to current recipe.
// Uses time format RFC3339 to companre CreatedAt value.
func (suite *RepositoryMockTestSuite) assertRecipeIsEqual(expectedRecipe, recipe model.Recipe) {
	suite.Equal(expectedRecipe.Id, recipe.Id)
	suite.Equal(expectedRecipe.Type, recipe.Type)
	suite.Equal(expectedRecipe.Title, recipe.Title)
	suite.Equal(expectedRecipe.Ingredients, recipe.Ingredients)
	suite.Equal(expectedRecipe.Description, recipe.Description)
	suite.Equal(expectedRecipe.CreatedAt.Format(time.RFC3339), recipe.CreatedAt.Format(time.RFC3339))
}
