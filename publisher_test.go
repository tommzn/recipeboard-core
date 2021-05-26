package core

import (
	"testing"

	"github.com/stretchr/testify/suite"
	utils "github.com/tommzn/go-utils"
	model "github.com/tommzn/recipeboard-core/model"
)

// Test suite for publiher repository.
type PublisherTestSuite struct {
	suite.Suite
}

func TestPublisherTestSuite(t *testing.T) {
	suite.Run(t, new(PublisherTestSuite))
}

// Test creating new recipe messages.
func (suite *PublisherTestSuite) TestCreateRecipeMessage() {

	recipe := recipeForTest()

	recipeMessage := newRecipeMessage(recipe, model.RecipeAdded)
	suite.True(utils.IsId(recipeMessage.Id))
	suite.True(recipeMessage.Sequence > 0)
	suite.Equal(recipe, recipeMessage.Recipe)
}
