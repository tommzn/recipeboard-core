package mock

import (
	"time"

	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/utils/common"
)

// Creates a new recipe with dummy values for testing
func recipeForTest() model.Recipe {
	return model.Recipe{
		Id:          common.NewId(nil),
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
		Id:       common.NewId(nil),
		Action:   model.RecipeAdded,
		Sequence: time.Now().UnixNano(),
		Recipe:   recipeForTest(),
	}
}
