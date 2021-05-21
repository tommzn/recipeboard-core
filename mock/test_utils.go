package mock

import (
	"time"

	core "github.com/tommzn/recipeboard-core"
	"gitlab.com/tommzn-go/utils/common"
)

// Creates a new recipe with dummy values for testing
func recipeForTest() core.Recipe {
	return core.Recipe{
		Id:          common.NewId(nil),
		Type:        core.BakingRecipe,
		Title:       "Bake a Cake",
		Ingredients: "100g Mehl\n100g Zucker\n50ml Wasser",
		Description: "Einrühren.\nBacken.\nFertig!",
		CreatedAt:   time.Now(),
	}
}

// Creates a new recipe message with dummy values for testing
func recipeMessageForTest() core.RecipeMessage {
	return core.RecipeMessage{
		Id:       common.NewId(nil),
		Action:   core.RecipeAdded,
		Sequence: time.Now().UnixNano(),
		Recipe:   recipeForTest(),
	}
}
