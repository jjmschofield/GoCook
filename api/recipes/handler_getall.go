package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
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
	recipes, err := GetAllFromStore(context.MustGet("userId").(string))

	if err != nil{
		respond.InternalError(context,"Couldn't retrieve recipes")
		return
	}

	respond.Ok(context, recipes)
}
