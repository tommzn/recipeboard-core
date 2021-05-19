package core

// Persistence interface to manage recipe lifecircle.
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
