// Contains the core model and interfaces for the recipe board project.
package core

// Persistence interface to manage recipe life circle.
type Repository interface {

	// Persist a recipe. Can be used to insert a new recipe
	// or update an existing one.
	Set(Recipe) error

	// Try to retrieve a recipe for passed id.
	Get(string) (*Recipe, error)

	// Lists all available recipes for passed type.
	// It doesn't take care about ordering of recipes.
	List(RecipeType) ([]Recipe, error)

	// Try to delete the passed recipe.
	Delete(Recipe) error
}

// Publishes messages for recipes to a queue or broker.
type MessagePublisher interface {

	// Sends given message to a message queue.
	Send(RecipeMessage) error
}
