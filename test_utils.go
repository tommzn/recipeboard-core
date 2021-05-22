package core

import (
	"time"

	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/utils/common"
	"gitlab.com/tommzn-go/utils/config"
	"gitlab.com/tommzn-go/utils/log"
)

// Creates a new stdout logger for testing
func loggerForTest() log.Logger {
	return log.NewLogger(log.Debug, "dynamodb-test")
}

// Creates a config loader and returns retireved config
func loadConfigForTest() config.Config {
	configLoader := config.NewConfigLoader()
	return configLoader.Load()
}

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

// Returns a new recipe without an id and createdAt value
func recipeForServiceTest() model.Recipe {
	return model.Recipe{
		Type:        model.BakingRecipe,
		Title:       "Bake a Cake",
		Ingredients: "100g Mehl\n100g Zucker\n50ml Wasser",
		Description: "Einrühren.\nBacken.\nFertig!",
	}
}
