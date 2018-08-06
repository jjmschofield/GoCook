package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
)

// @Summary Get Recipe by Id
// @Description Gets a recipe by ID which the caller has access too
// @Security OAuth2Implicit
// @Tags Recipes
// @Accept json
// @Produce json
// @Param id path string true "The uuid of the recipe"
// @Success 200 {object} recipes.Recipe
// @Failure 400 {object} respond.ErrorPayload
// @Failure 404 {object} respond.ErrorPayload
// @Router /recipes/{id} [get]
func getByIdRequestHandler(context *gin.Context) {
	id := context.Param("id")

	validRequest, validationError := isValidGetByIdRequest(id)

	if !validRequest {
		respond.BadRequest(context, validationError)
		return
	}

	recipe, storeErr := GetFromStoreById(id, context.MustGet("userId").(string))

	if storeErr != nil {
		respond.NotFound(context)
		return
	}

	respond.Ok(context, recipe)

}

func isValidGetByIdRequest(id string) (isValid bool, validationMessage string) {
	return validate.UuidString(id)
}
