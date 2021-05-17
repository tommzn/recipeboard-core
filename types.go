package core

import (
	"time"

	"gitlab.com/tommzn-go/aws/dynamodb"
)

type RecipeType int
type RecipeAction string
type ObjectType string

const (
	CookingRecipe RecipeType = iota
	BakingRecipe
)

const (
	ObjectType_Recipe ObjectType = "RECIPEBOARD_RECIPE"
	ObjectType_Index             = "RECIPEBOARD_INDEX"
)

const (
	RecipeAdded   RecipeAction = "RecipeAddedd"
	RecipeUpdated              = "RecipeUpdated"
	RecipeDeleted              = "RecipeDeleted"
)

type Recipe struct {
	Id          string
	Type        RecipeType
	Title       string
	Ingredients string
	Description string
	CreatedAt   time.Time
}

type RecipeItem struct {
	dynamodb.ItemId
	Type        RecipeType
	Title       string
	Ingredients string
	Description string
	CreatedAt   time.Time
}

type RecipeIndex struct {
	dynamodb.ItemId
	Ids []string
}

type RecipeMessage struct {
	Id       string
	Action   RecipeAction
	Sequence int
	Recipe   Recipe
}
