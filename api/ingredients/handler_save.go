package ingredients

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
	"go.uber.org/zap"
)

// swagger:route POST /ingredients{id} Ingredients UpsertIngredient
//
// Save Recipe
//
// Carries out an upsert on a ingredient.
//
// When recipe id is null, a new ingredient is created. When it is populated the recipe is updated (if the caller has access)
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
//       200: Ingredient
//		 400: MessagePayload
//       500: ErrorPayload
func saveRequestHandler(context *gin.Context) {
	logger := context.MustGet("logger").(zap.Logger)
	var requestIngredient Ingredient

	bindError := context.Bind(&requestIngredient)

	if bindError != nil {
		respond.BadRequest(context, bindError.Error())
		return
	}

	validRequest, validationMessage := isValidSaveRequest(requestIngredient)

	if !validRequest {
		respond.BadRequest(context, validationMessage)
		return
	}

	ingredient, storeErr := SaveToStore(requestIngredient, context.MustGet("userId").(string))

	if storeErr != nil {
		logger.Error("Failed to save ingredient with id " + requestIngredient.Id, zap.Error(storeErr))
		respond.InternalError(context, "Failed writing to store")
		return
	}

	respond.Ok(context, ingredient)
}

func isValidSaveRequest(ingredient Ingredient) (validRequest bool, validationMessage string) {
	return validate.Struct(ingredient)
}
