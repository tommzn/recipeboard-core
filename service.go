package core

import (
	"time"

	"github.com/tommzn/recipeboard-core/mock"
	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/utils/common"
	"gitlab.com/tommzn-go/utils/config"
	"gitlab.com/tommzn-go/utils/log"
)

func NewRecipeServiceFromConfig(conf config.Config, logger log.Logger) RecipeService {

	if logger == nil {
		logger = newLogger()
	}

	repository := newRepository(conf, logger)
	publisher := mock.NewPublisher()
	return NewRecipeService(repository, publisher, logger)
}

func NewRecipeService(repository Repository, publisher MessagePublisher, logger log.Logger) RecipeService {

	return &RecipeManager{
		repository: repository,
		publisher:  publisher,
	}
}

// Create a new recipe.
func (manager *RecipeManager) Create(recipe model.Recipe) (model.Recipe, error) {

	recipe.Id = common.NewId(nil)
	recipe.CreatedAt = time.Now()
	err := manager.repository.Set(recipe)
	if err != nil {
		return recipe, err
	}
	return recipe, manager.publisher.Send(newRecipeMessage(recipe, model.RecipeAdded))
}

// Update an existing recipe.
func (manager *RecipeManager) Update(recipe model.Recipe) error {

	_, err := manager.repository.Get(recipe.Id)
	if err != nil {
		return err
	}
	err = manager.repository.Set(recipe)
	if err != nil {
		return err
	}
	return manager.publisher.Send(newRecipeMessage(recipe, model.RecipeUpdated))
}

// Try to retrieve a recipe for passed id.
func (manager *RecipeManager) Get(id string) (*model.Recipe, error) {

	return manager.repository.Get(id)
}

// Lists all available recipes for passed type.
// It doesn't take care about ordering of recipes.
func (manager *RecipeManager) List(recipeType model.RecipeType) ([]model.Recipe, error) {

	return manager.repository.List(recipeType)
}

// Try to delete the passed recipe.
func (manager *RecipeManager) Delete(recipe model.Recipe) error {

	err := manager.repository.Delete(recipe)
	if err != nil {
		return err
	}
	return manager.publisher.Send(newRecipeMessage(recipe, model.RecipeDeleted))
}
