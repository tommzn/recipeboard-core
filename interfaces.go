package core

import model "github.com/tommzn/recipeboard-core/model"

// Domain service to manage recipe live circls.
type RecipeService interface {

	// Create a new recipe.
	Create(model.Recipe) (model.Recipe, error)

	// Update an existing recipe.
	Update(model.Recipe) error

	// Try to retrieve a recipe for passed id.
	Get(string) (*model.Recipe, error)

	// Lists all available recipes for passed type.
	// It doesn't take care about ordering of recipes.
	List(model.RecipeType) ([]model.Recipe, error)

	// Try to delete the passed recipe.
	Delete(model.Recipe) error
}

// Persistence interface to manage recipe lifecircle.
type Repository interface {

	// Persist a recipe. Can be used to insert a new recipe
	// or update an existing one.
	Set(model.Recipe) error

	// Try to retrieve a recipe for passed id.
	Get(string) (*model.Recipe, error)

	// Lists all available recipes for passed type.
	// It doesn't take care about ordering of recipes.
	List(model.RecipeType) ([]model.Recipe, error)

	// Try to delete the passed recipe.
	Delete(model.Recipe) error
}

// Publishes messages for recipes to a queue or broker.
type MessagePublisher interface {

	// Sends given message to a message queue.
	Send(model.RecipeMessage) error
}
