package core

import (
	"time"

	model "github.com/tommzn/recipeboard-core/model"
	"gitlab.com/tommzn-go/utils/common"
)

// Creates a new recipe message for goven recipe.
// Will assign a new id and sequence number.
func newRecipeMessage(recipe model.Recipe, action model.RecipeAction) model.RecipeMessage {

	return model.RecipeMessage{
		Id:       common.NewId(nil),
		Action:   action,
		Sequence: time.Now().UnixNano(),
		Recipe:   recipe,
	}
}
