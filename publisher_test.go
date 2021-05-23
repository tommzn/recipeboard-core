package core

import (
	"testing"

	"github.com/stretchr/testify/suite"
	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/utils/common"
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
	suite.True(common.IsId(recipeMessage.Id))
	suite.True(recipeMessage.Sequence > 0)
	suite.Equal(recipe, recipeMessage.Recipe)
}
