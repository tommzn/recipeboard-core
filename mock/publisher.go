// Package mock provides mocks for persistence layer and publisher for the recipe board project.
package mock

import (
	model "github.com/tommzn/recipeboard-core/model"
)

// NewPublisher creates a new message publisher mock which stores all passed messages
// in a local queue.
func NewPublisher() *PublisherMock {

	return &PublisherMock{
		Queue: []model.RecipeMessage{},
	}
}

// Send a messge, which will append the passed message to the internal queue.
func (mock *PublisherMock) Send(message model.RecipeMessage) error {
	mock.Queue = append(mock.Queue, message)
	return nil
}
