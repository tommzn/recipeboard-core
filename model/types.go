package model

import (
	"time"
)

// Type is used to group recipes e.g. for cooking or baking.
type RecipeType int

const (
	CookingRecipe RecipeType = iota // Cooking recipes
	BakingRecipe                    // Baking recipes
)

// Action performed for a recipe.
type RecipeAction string

const (
	RecipeAdded   RecipeAction = "RecipeAddedd"  // RecipeAdded is the a message type used after a new recipe has been created.
	RecipeUpdated              = "RecipeUpdated" // RecipeUpdated is is used if an existing recipe has been updates.
	RecipeDeleted              = "RecipeDeleted" // RecipeDeleted will be send as actions after a recipe was deleted.
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
