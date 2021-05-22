package mock

import model "github.com/tommzn/recipeboard-core/model"

// Mock for message publishing.
type PublisherMock struct {

	// Queue with all recipe messages
	Queue []model.RecipeMessage
}

// Mock for recipe repository.
type RepositoryMock struct {

	// Map with all avaiable recipes.
	Recipes map[string]model.Recipe
}
