package core

import (
	"time"

	"gitlab.com/tommzn-go/utils/common"
	"gitlab.com/tommzn-go/utils/config"
	"gitlab.com/tommzn-go/utils/log"
)

// Creates a new message publisher.
// At the moment only a mock with an internal slice of all published messages.
func newPublisher(conf config.Config, logger log.Logger) MessagePublisher {

	return &PublisherMock{
		queue: []RecipeMessage{},
	}
}

// Adds the passed message to the internal queue.
func (mock *PublisherMock) Send(message RecipeMessage) error {
	mock.queue = append(mock.queue, message)
	return nil
}

// Creates a new recipe message for goven recipe.
// Will assign a new id and sequence number.
func newRecipeMessage(recipe Recipe, action RecipeAction) RecipeMessage {

	return RecipeMessage{
		Id:       common.NewId(nil),
		Action:   action,
		Sequence: time.Now().UnixNano(),
		Recipe:   recipe,
	}
}
