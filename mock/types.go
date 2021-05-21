package mock

// Mock for message publishing. Will be implemented later.
type PublisherMock struct {

	// Queue with all recipe messages
	queue []RecipeMessage
}
