package mock

import (
	"errors"

	model "github.com/tommzn/recipeboard-core/model"
)

// Creates a new recipe repository mock which stored all recipes locally.
func NewRepository() *RepositoryMock {

	return &RepositoryMock{
		Recipes: make(map[string]model.Recipe),
	}
}

// Persist a recipe.
func (mock *RepositoryMock) Set(recipe model.Recipe) error {

	mock.Recipes[recipe.Id] = recipe
	return nil
}

// Try to retrieve a recipe for passed id.
func (mock *RepositoryMock) Get(id string) (*model.Recipe, error) {

	if recipe, ok := mock.Recipes[id]; ok {
		return &recipe, nil
	} else {
		return nil, errors.New("Not found.")
	}
}

// Lists all available recipes for passed type.
// It doesn't take care about ordering of recipes.
func (mock *RepositoryMock) List(recipeType model.RecipeType) ([]model.Recipe, error) {

	recipes := []model.Recipe{}
	for _, recipe := range mock.Recipes {
		if recipe.Type == recipeType {
			recipes = append(recipes, recipe)
		}
	}
	var err error
	if len(recipes) == 0 {
		err = errors.New("Not found")
	}
	return recipes, err
}

// Try to delete the passed recipe.
func (mock *RepositoryMock) Delete(recipe model.Recipe) error {

	if _, ok := mock.Recipes[recipe.Id]; ok {
		delete(mock.Recipes, recipe.Id)
		return nil
	} else {
		return errors.New("Not found.")
	}
}
