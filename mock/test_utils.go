package mock

import (
	"time"

	utils "github.com/tommzn/go-utils"
	model "github.com/tommzn/recipeboard-core/model"
)

// Creates a new recipe with dummy values for testing
func recipeForTest() model.Recipe {
	return model.Recipe{
		Id:          utils.NewId(),
		Type:        model.BakingRecipe,
		Title:       "Bake a Cake",
		Ingredients: "100g Mehl\n100g Zucker\n50ml Wasser",
		Description: "Einrühren.\nBacken.\nFertig!",
		CreatedAt:   time.Now(),
	}
}

// Creates a new recipe message with dummy values for testing
func recipeMessageForTest() model.RecipeMessage {
	return model.RecipeMessage{
		Id:       utils.NewId(),
		Action:   model.RecipeAdded,
		Sequence: time.Now().UnixNano(),
		Recipe:   recipeForTest(),
	}
}
