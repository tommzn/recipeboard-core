package core

import (
	"time"

	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/aws/dynamodb"
	"gitlab.com/tommzn-go/utils/log"
)

// DynamoDb ibject type.
type objectType string

const (
	objectType_Recipe objectType = "RECIPEBOARD_RECIPE"
	objectType_Index             = "RECIPEBOARD_INDEX"
)

// DynamoDb recipe item.
type recipeItem struct {

	// Id for an item in DynamoDb.
	*dynamodb.ItemId

	// Type of a recipe.
	Type model.RecipeType

	// Title or headline for a recipe.
	Title string

	// List of ingredients.
	Ingredients string

	// Description, e.g. instructions for preparation.
	Description string

	// Timestamp a recipe has been created.
	CreatedAt time.Time
}

// A recipe index is maintainey by the repository
// to provide a faster access to items in DynamoDb.
type recipeIndex struct {

	// DynamoDb identifier.
	*dynamodb.ItemId

	// List of recipe ids.
	Ids map[string]bool
}

// Domain service to manage recipe life circle
type RecipeManager struct {

	// Backend repository for recipes.
	repository model.Repository

	// Publisher to send notifications after actions for recipes has been performed.
	publisher model.MessagePublisher
}

// Adapter to persist recipes in AWS DynamoDb.
type DynamoDbRepository struct {

	// DynamoDb client.
	client dynamodb.Repository

	// Logger.
	logger log.Logger
}
