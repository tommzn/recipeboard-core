package mock

import core "github.com/tommzn/recipeboard-core"

// Mock for message publishing.
type PublisherMock struct {

	// Queue with all recipe messages
	Queue []core.RecipeMessage
}

// Mock for recipe repository.
type RepositoryMock struct {

	// Map with all avaiable recipes.
	Recipes map[string]core.Recipe
}
