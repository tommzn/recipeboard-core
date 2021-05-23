package mock

import model "github.com/tommzn/recipeboard-core/model"

// PublisherMock is a mock for message publishing.
type PublisherMock struct {

	// Queue with all recipe messages
	Queue []model.RecipeMessage
}

// RepositoryMock is a mock with a local recipe storage.
type RepositoryMock struct {

	// Map with all avaiable recipes.
	Recipes map[string]model.Recipe
}
