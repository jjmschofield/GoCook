package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
	"go.uber.org/zap"
)

// swagger:route POST /recipes{id} Recipes UpsertRecipe
//
// Save Recipe
//
// Carries out an upsert on a recipe.
//
// When recipe id is null, a new recipe is created. When it is populated the recipe is updated (if the caller has access)
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: Recipe
//		 400: MessagePayload
//       500: ErrorPayload
func saveRequestHandler(context *gin.Context) {
	logger := context.MustGet("logger").(zap.Logger)
	var requestRecipe Recipe

	bindError := context.Bind(&requestRecipe)

	if bindError != nil {
		respond.BadRequest(context, bindError.Error())
		return
	}

	validRequest, validationMessage := isValidSaveRequest(requestRecipe)

	if !validRequest {
		respond.BadRequest(context, validationMessage)
		return
	}

	savedRecipe, storeErr := SaveToStore(requestRecipe, context.MustGet("userId").(string))

	if storeErr != nil {
		logger.Error("Could not save recipe with id " + requestRecipe.Id, zap.Error(storeErr))
		respond.InternalError(context, "Failed writing to store")
		return
	}

	respond.Ok(context, savedRecipe)
}

func isValidSaveRequest(recipe Recipe) (validRequest bool, validationMessage string) {
	return validate.Struct(recipe)
}
