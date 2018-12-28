package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"go.uber.org/zap"
)

// swagger:route GET /recipes Recipes GetAllRecipes
//
// Get All Recipes
//
// Lists all recipes with no pagination or filtering(!).
//
// This will show only those recipes which the caller has been granted access to.
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
//       200: []Recipe
//       500: ErrorPayload
func getAllRequestHandler(context *gin.Context) {
	logger := context.MustGet("logger").(zap.Logger)
	userId := context.MustGet("userId").(string)

	recipeMap, err := GetAllFromStore(userId)

	if err != nil {
		logger.Error("Failed to get recipes from store", zap.Error(err))
		respond.InternalError(context, "Couldn't retrieve recipeMap")
		return
	}

	recipes := []Recipe{}
	for _, recipe := range recipeMap {
		recipes = append(recipes, recipe)
	}

	respond.Ok(context, recipes)
}
