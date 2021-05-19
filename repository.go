package core

import (
	"gitlab.com/tommzn-go/aws/dynamodb"
	"gitlab.com/tommzn-go/utils/config"
	"gitlab.com/tommzn-go/utils/log"
)

// Creates a new DynamoDb client. Passed config have to contain at least the
// DynamoDb table name. Region and endpoint for DynamoDb are optional.
// If no logger is passed a new stdout logger will be ued.
//
// Example config, YAML.
//
// aws:
//   dynamodb:
//     tablename: DynamoDbTable
//     region: eu-west-1
//     endpoint: http://localhost:8000
func NewRepository(conf config.Config, logger log.Logger) Repository {

	if logger == nil {
		logger = newLogger()
	}
	return &DynamoDbRepository{
		client: dynamodb.NewRepository(conf, logger),
		logger: logger,
	}
}

// Persist a recipe.
func (repo *DynamoDbRepository) Set(recipe Recipe) error {

	recipeItem := toDynamoDbItem(recipe)
	currentRecipe, _ := repo.Get(recipe.Id)

	if currentRecipe == nil {

		err := repo.appendToIndex(recipe)
		if err != nil {
			return err
		}

	} else if currentRecipe.Type != recipe.Type {

		err := repo.removeFromIndex(*currentRecipe)
		if err != nil {
			return err
		}
		err = repo.appendToIndex(recipe)
		if err != nil {
			return err
		}
	}
	return repo.client.Add(recipeItem)
}

// Try to retrieve a recipe for passed id.
func (repo *DynamoDbRepository) Get(id string) (*Recipe, error) {

	recipeItem := &recipeItem{ItemId: newRecipeIdForDynamoDb(id)}
	err := repo.client.Get(recipeItem)
	if err != nil {
		return nil, err
	}
	recipe := fromDynamoDbItem(recipeItem)
	return &recipe, nil
}

// Lists all available recipes for passed type.
// It doesn't take care about ordering of recipes.
func (repo *DynamoDbRepository) List(recipeType RecipeType) ([]Recipe, error) {

	recipes := []Recipe{}
	recipeIndex := newRecipeIndex(recipeType)
	err := repo.client.Get(recipeIndex)
	if err != nil {
		return recipes, err
	}

	for id, _ := range recipeIndex.Ids {
		recipeItem := &recipeItem{ItemId: newRecipeIdForDynamoDb(id)}
		if err := repo.client.Get(recipeItem); err == nil {
			recipes = append(recipes, fromDynamoDbItem(recipeItem))
		} else {
			return recipes, err
		}
	}
	return recipes, nil
}

// Try to delete the passed recipe.
func (repo *DynamoDbRepository) Delete(recipe Recipe) error {

	err := repo.removeFromIndex(recipe)
	if err != nil {
		return err
	}
	recipeItem := toDynamoDbItem(recipe)
	return repo.client.Delete(recipeItem)
}

// Adds the passed recipe id to an index depending on it's type.
func (repo *DynamoDbRepository) appendToIndex(recipe Recipe) error {

	recipeIndex := newRecipeIndex(recipe.Type)
	repo.client.Get(recipeIndex)
	recipeIndex.Ids[recipe.Id] = true
	return repo.client.Add(recipeIndex)
}

// Removes the id of a recipe id from an index depending on it's type.
func (repo *DynamoDbRepository) removeFromIndex(recipe Recipe) error {

	recipeIndex := newRecipeIndex(recipe.Type)
	repo.client.Get(recipeIndex)
	delete(recipeIndex.Ids, recipe.Id)
	return repo.client.Add(recipeIndex)
}

// Converts the passed recipe into a DynamoDb item.
func toDynamoDbItem(recipe Recipe) *recipeItem {
	return &recipeItem{
		ItemId:      newRecipeIdForDynamoDb(recipe.Id),
		Type:        recipe.Type,
		Title:       recipe.Title,
		Ingredients: recipe.Ingredients,
		Description: recipe.Description,
		CreatedAt:   recipe.CreatedAt,
	}
}

// Convert the passed DynamoDb item into a recipe.
func fromDynamoDbItem(recipeItem *recipeItem) Recipe {
	return Recipe{
		Id:          recipeItem.GetId(),
		Type:        recipeItem.Type,
		Title:       recipeItem.Title,
		Ingredients: recipeItem.Ingredients,
		Description: recipeItem.Description,
		CreatedAt:   recipeItem.CreatedAt,
	}
}

// Returns a new DynamoDb item id for passed ceipe id.
func newRecipeIdForDynamoDb(id string) *dynamodb.ItemId {
	return dynamodb.NewItemId(string(objectType_Recipe), &id)
}

// Returns a new DynamoDb item id for passed ceipe type.
func newIndexIdForDynamoDb(recipeType RecipeType) *dynamodb.ItemId {
	id := string(recipeType)
	return dynamodb.NewItemId(string(objectType_Index), &id)
}

// Creates a new recipe index DynamoDb item
func newRecipeIndex(recipeType RecipeType) *recipeIndex {
	return &recipeIndex{
		ItemId: newIndexIdForDynamoDb(recipeType),
		Ids:    make(map[string]bool),
	}
}
