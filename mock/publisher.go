package mock

import (
	core "github.com/tommzn/recipeboard-core"
)

// Creates a new message publisher mock which stores all passed messages
// in a local queue.
func NewPublisher() *PublisherMock {

	return &PublisherMock{
		Queue: []core.RecipeMessage{},
	}
}

// Adds the passed message to the internal queue.
func (mock *PublisherMock) Send(message core.RecipeMessage) error {
	mock.Queue = append(mock.Queue, message)
	return nil
}
