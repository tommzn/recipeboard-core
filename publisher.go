package core

import (
	"time"

	utils "github.com/tommzn/go-utils"
	model "github.com/tommzn/recipeboard-core/model"
)

// Creates a new recipe message for goven recipe.
// Will assign a new id and sequence number.
func newRecipeMessage(recipe model.Recipe, action model.RecipeAction) model.RecipeMessage {

	return model.RecipeMessage{
		Id:       utils.NewId(),
		Action:   action,
		Sequence: time.Now().UnixNano(),
		Recipe:   recipe,
	}
}
