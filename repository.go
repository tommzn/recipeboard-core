package core

import (
	dynamodb "github.com/tommzn/aws-dynamodb"
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	model "github.com/tommzn/recipeboard-core/model"
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
func newRepository(conf config.Config, logger log.Logger) model.Repository {

	return &DynamoDbRepository{
		client: dynamodb.NewRepository(conf, logger),
		logger: logger,
	}
}

// Set can be used to create a new recipe or update an existing one.
func (repo *DynamoDbRepository) Set(recipe model.Recipe) error {

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

// Get will try to retrieve a recipe for passed id.
func (repo *DynamoDbRepository) Get(id string) (*model.Recipe, error) {

	recipeItem := &recipeItem{ItemIdentifier: newRecipeIdForDynamoDb(id)}
	err := repo.client.Get(recipeItem)
	if err != nil {
		return nil, err
	}
	recipe := fromDynamoDbItem(recipeItem)
	return &recipe, nil
}

// List returns all available recipes for passed type.
// It doesn't take care about ordering of recipes.
func (repo *DynamoDbRepository) List(recipeType model.RecipeType) ([]model.Recipe, error) {

	recipes := []model.Recipe{}
	recipeIndex := newRecipeIndex(recipeType)
	err := repo.client.Get(recipeIndex)
	if err != nil {
		return recipes, err
	}

	for id := range recipeIndex.Ids {
		recipeItem := &recipeItem{ItemIdentifier: newRecipeIdForDynamoDb(id)}
		if err := repo.client.Get(recipeItem); err == nil {
			recipes = append(recipes, fromDynamoDbItem(recipeItem))
		} else {
			return recipes, err
		}
	}
	return recipes, nil
}

// Delete will try to remove the passed recipe.
func (repo *DynamoDbRepository) Delete(recipe model.Recipe) error {

	err := repo.removeFromIndex(recipe)
	if err != nil {
		return err
	}
	recipeItem := toDynamoDbItem(recipe)
	return repo.client.Delete(recipeItem)
}

// Adds the passed recipe id to an index depending on it's type.
func (repo *DynamoDbRepository) appendToIndex(recipe model.Recipe) error {

	recipeIndex := newRecipeIndex(recipe.Type)
	repo.client.Get(recipeIndex)
	recipeIndex.Ids[recipe.Id] = true
	return repo.client.Add(recipeIndex)
}

// Removes the id of a recipe id from an index depending on it's type.
func (repo *DynamoDbRepository) removeFromIndex(recipe model.Recipe) error {

	recipeIndex := newRecipeIndex(recipe.Type)
	repo.client.Get(recipeIndex)
	delete(recipeIndex.Ids, recipe.Id)
	return repo.client.Add(recipeIndex)
}

// Converts the passed recipe into a DynamoDb item.
func toDynamoDbItem(recipe model.Recipe) *recipeItem {
	return &recipeItem{
		ItemIdentifier: newRecipeIdForDynamoDb(recipe.Id),
		Type:           recipe.Type,
		Title:          recipe.Title,
		Ingredients:    recipe.Ingredients,
		Description:    recipe.Description,
		CreatedAt:      recipe.CreatedAt,
	}
}

// Convert the passed DynamoDb item into a recipe.
func fromDynamoDbItem(recipeItem *recipeItem) model.Recipe {
	return model.Recipe{
		Id:          recipeItem.GetId(),
		Type:        recipeItem.Type,
		Title:       recipeItem.Title,
		Ingredients: recipeItem.Ingredients,
		Description: recipeItem.Description,
		CreatedAt:   recipeItem.CreatedAt,
	}
}

// Returns a new DynamoDb item id for passed ceipe id.
func newRecipeIdForDynamoDb(id string) *dynamodb.ItemIdentifier {
	return dynamodb.NewItemIdentifier(id, string(objectTypeRecipe))
}

// Returns a new DynamoDb item id for passed ceipe type.
func newIndexIdForDynamoDb(recipeType model.RecipeType) *dynamodb.ItemIdentifier {
	id := string(recipeType)
	return dynamodb.NewItemIdentifier(id, string(objectTypeIndex))
}

// Creates a new recipe index DynamoDb item
func newRecipeIndex(recipeType model.RecipeType) *recipeIndex {
	return &recipeIndex{
		ItemIdentifier: newIndexIdForDynamoDb(recipeType),
		Ids:            make(map[string]bool),
	}
}
