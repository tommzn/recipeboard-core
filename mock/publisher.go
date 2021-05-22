package mock

import (
	model "github.com/tommzn/recipeboard-core/model"
)

// Creates a new message publisher mock which stores all passed messages
// in a local queue.
func NewPublisher() *PublisherMock {

	return &PublisherMock{
		Queue: []model.RecipeMessage{},
	}
}

// Adds the passed message to the internal queue.
func (mock *PublisherMock) Send(message model.RecipeMessage) error {
	mock.Queue = append(mock.Queue, message)
	return nil
}
