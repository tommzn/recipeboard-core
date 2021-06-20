package core

import (
	"os"
	"time"

	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
	utils "github.com/tommzn/go-utils"
	model "github.com/tommzn/recipeboard-core/model"
)

// Creates a new stdout logger for testing
func loggerForTest() log.Logger {
	return log.NewLogger(log.Debug, nil, nil)
}

// Creates a config loader and returns retireved config
func loadConfigForTest() (config.Config, error) {
	configFile := "testconfig.yml"
	configSource := config.NewFileConfigSource(&configFile)
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

// runAtCI will return true if env var CI is set.
func runAtCI() bool {
	_, ok := os.LookupEnv("CI")
	return ok
}
