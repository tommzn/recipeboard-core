package core

import (
	"time"

	"gitlab.com/tommzn-go/aws/dynamodb"
	"gitlab.com/tommzn-go/utils/log"
)

// Type is used to group recipes e.g. for cooking or baking.
type RecipeType int

const (
	CookingRecipe RecipeType = iota
	BakingRecipe
)

// Action performed for a recipe.
type RecipeAction string

const (
	RecipeAdded   RecipeAction = "RecipeAddedd"
	RecipeUpdated              = "RecipeUpdated"
	RecipeDeleted              = "RecipeDeleted"
)

// DynamoDb ibject type.
type objectType string

const (
	objectType_Recipe objectType = "RECIPEBOARD_RECIPE"
	objectType_Index             = "RECIPEBOARD_INDEX"
)

// Central model for a recipe.
type Recipe struct {

	// Identifier of  a recipe.
	Id string

	// Type of a recipe.
	Type RecipeType

	// Title or headline for a recipe.
	Title string

	// List of ingredients.
	Ingredients string

	// Description, e.g. instructions for preparation.
	Description string

	// Timestamp a recipe has been created.
	CreatedAt time.Time
}

// DynamoDb recipe item.
type recipeItem struct {

	// Id for an item in DynamoDb.
	*dynamodb.ItemId

	// Type of a recipe.
	Type RecipeType

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

// A recipe message is published after an action for
// a recipe has been performed.
type RecipeMessage struct {

	// Message id.
	Id string

	// Action performed for a recipe.
	Action RecipeAction

	// Sequence id for messages, to avoid conflicts or ordering issues.
	Sequence int64

	// The recipe an action has been performed for.
	Recipe Recipe
}

// Domain service to manage recipe life circle
type RecipeManager struct {

	// Backend repository for recipes.
	repository Repository

	// Publisher to send notifications after actions for recipes has been performed.
	publisher MessagePublisher
}

// Adapter to persist recipes in AWS DynamoDb.
type DynamoDbRepository struct {

	// DynamoDb client.
	client dynamodb.Repository

	// Logger.
	logger log.Logger
}

// Mock for message publishing. Will be implemented later.
type PublisherMock struct {

	// Queue with all recipe messages
	queue []RecipeMessage
}
