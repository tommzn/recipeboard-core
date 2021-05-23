package mock

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Test suite for publisher.
type PublisherMockTestSuite struct {
	suite.Suite
}

func TestPublisherMockTestSuite(t *testing.T) {
	suite.Run(t, new(PublisherMockTestSuite))
}

// Publish a recipe message and assert an existing message in internal queue.
func (suite *PublisherMockTestSuite) TestSendMessage() {

	publisher := NewPublisher()
	recipeMessage := recipeMessageForTest()

	suite.Nil(publisher.Send(recipeMessage))
	suite.Len(publisher.Queue, 1)
	suite.Equal(recipeMessage.Recipe.Id, publisher.Queue[0].Recipe.Id)
}
