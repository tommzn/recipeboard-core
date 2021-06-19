package core

import (
	"time"

	dynamodb "github.com/tommzn/aws-dynamodb"
	sqs "github.com/tommzn/aws-sqs"
	log "github.com/tommzn/go-log"
	model "github.com/tommzn/recipeboard-core/model"
)

// DynamoDb ibject type.
type objectType string

const (
	objectTypeRecipe objectType = "RECIPEBOARD_RECIPE"
	objectTypeIndex             = "RECIPEBOARD_INDEX"
)

// DynamoDb recipe item.
type recipeItem struct {

	// Id for an item in DynamoDb.
	*dynamodb.ItemIdentifier

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
	*dynamodb.ItemIdentifier

	// List of recipe ids.
	Ids map[string]bool
}

// RecipeManager is the domain service which mananges recipes and all depending comoponents.
type RecipeManager struct {

	// Backend repository for recipes.
	repository model.Repository

	// Publisher to send notifications after actions for recipes has been performed.
	publisher model.MessagePublisher

	// Logger.
	logger log.Logger
}

// DynamoDbRepository is a persistence adapter to manage recipes in AWS DynamoDb.
type DynamoDbRepository struct {

	// DynamoDb client.
	client dynamodb.Repository

	// Logger.
	logger log.Logger
}

// sqsPublisher will publishs changes at recipes to a queue in AWS SQS.
type sqsPublisher struct {

	// client is used to publish messages on AWS SQS queues.
	client sqs.Publisher

	// queueName is the name of an AWS SQS queue messages for recipe actions will be send to.
	queueName *string

	// Logger.
	logger log.Logger
}
