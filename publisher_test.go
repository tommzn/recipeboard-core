package core

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.com/tommzn-go/utils/common"
	"gitlab.com/tommzn-go/utils/config"
)

// Test suite for publsiher repository.
type PublisherTestSuite struct {
	suite.Suite
	conf      config.Config
	publisher MessagePublisher
}

func TestPublisherTestSuite(t *testing.T) {
	suite.Run(t, new(PublisherTestSuite))
}

// Setup test, load config and create a message publisher
func (suite *PublisherTestSuite) SetupTest() {
	suite.Nil(config.UseConfigFileIfNotExists("testconfig"))
	suite.conf = loadConfigForTest()
	suite.publisher = NewPublisher(suite.conf, nil)
}

// Test creating new recipe messages.
func (suite *PublisherTestSuite) TestCreateRecipeMessage() {

	recipe := recipeForTest()

	recipeMessage := NewRecipeMessage(recipe, RecipeAdded)
	suite.True(common.IsId(recipeMessage.Id))
	suite.True(recipeMessage.Sequence > 0)
	suite.Equal(recipe, recipeMessage.Recipe)
}

func (suite *PublisherTestSuite) TestSendMessage() {

	recipe := recipeForTest()

	recipeMessage := NewRecipeMessage(recipe, RecipeAdded)
	suite.publisher.Send(recipeMessage)
	suite.Len(suite.publisher.(*PublisherMock).queue, 1)
}
