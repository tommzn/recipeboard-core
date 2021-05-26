package core

import (
	"time"

	config "github.com/tommzn/go-config"
	utils "github.com/tommzn/go-utils"
	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/utils/log"
)

// Creates a new stdout logger for testing
func loggerForTest() log.Logger {
	return log.NewLogger(log.Debug, "dynamodb-test")
}

// Creates a config loader and returns retireved config
func loadConfigForTest() (config.Config, error) {
	configSource := config.NewConfigSource()
	return configSource.Load()
}

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

// Returns a new recipe without an id and createdAt value
func recipeForServiceTest() model.Recipe {
	return model.Recipe{
		Type:        model.BakingRecipe,
		Title:       "Bake a Cake",
		Ingredients: "100g Mehl\n100g Zucker\n50ml Wasser",
		Description: "Einrühren.\nBacken.\nFertig!",
	}
}
