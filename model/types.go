package model

import (
	"time"
)

// RecipeType is used to group recipes e.g. for cooking or baking.
type RecipeType int

const (

	// CookingRecipe is used for cooking recipes.
	CookingRecipe RecipeType = iota

	// BakingRecipe is used for baking recipes.
	BakingRecipe
)

// RecipeAction tells a subscriber about the action performed for a recipe.
type RecipeAction string

const (

	// RecipeAdded is the a message type used after a new recipe has been created.
	RecipeAdded RecipeAction = "RecipeAddedd"

	// RecipeUpdated is is used if an existing recipe has been updates.
	RecipeUpdated = "RecipeUpdated"

	// RecipeDeleted will be send as actions after a recipe was deleted.
	RecipeDeleted = "RecipeDeleted"
)

// Recipe is the core model for a single baking or cooking recipe.
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

// RecipeMessage is published after an action for a recipe has been performed.
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
