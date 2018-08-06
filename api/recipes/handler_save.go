package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
)

// @Summary Save Recipe
// @Description Saves a new recipe (when Id is null) or updates a recipe (when id is populated, valid and owned by the caller)
// @Security OAuth2Implicit
// @Tags Recipes
// @Accept json
// @Produce json
// @Param recipe body recipes.Recipe true "The recipe to save"
// @Success 200 {object} recipes.Recipe "The saved recipe - may include mutations, eg generated IDs"
// @Failure 400 {object} respond.ErrorPayload
// @Failure 500 {object} respond.ErrorPayload
// @Router /recipes [post]
func saveRequestHandler(context *gin.Context) {
	// TODO - the upsert style here is making this quite a long method
	// TODO - and more is needed to support sending a 403 when trying to update a record you don't have access too
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

	if storeErr != nil{
		respond.InternalError(context, "Failed writing to store")
		return
	}

	respond.Ok(context, savedRecipe)
}

func isValidSaveRequest(recipe Recipe) (validRequest bool, validationMessage string) {
	return validate.Struct(recipe)
}
