package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
)

// @Summary Get All Recipes
// @Description Gets all recipes which the caller has access to, note pagination is not implemented yet.
// @Security OAuth2Implicit
// @Tags Recipes
// @Accept json
// @Produce json
// @Success 200 {array} recipes.Recipe
// @Failure 400 {object} respond.ErrorPayload
// @Failure 404 {object} respond.ErrorPayload
// @Router /recipes [get]
func getAllRequestHandler(context *gin.Context) {

	recipes, err := GetAllFromStore(context.MustGet("userId").(string))

	if err != nil{
		respond.InternalError(context,"Couldn't retrieve recipes")
		return
	}

	respond.Ok(context, recipes)
}
