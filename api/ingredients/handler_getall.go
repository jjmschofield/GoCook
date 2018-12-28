package ingredients

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"go.uber.org/zap"
)

// swagger:route GET /ingredients Ingredients GetAllIngredients
//
// Get All Ingredients
//
// Lists all ingredients with no pagination or filtering(!).
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: []Ingredient
//       500: ErrorPayload
func getAllRequestHandler(context *gin.Context) {
	logger := context.MustGet("logger").(zap.Logger)
	ingredients, err := GetAllFromStore()


	if err != nil {
		logger.Error("Could not get ingredients from store", zap.Error(err))
		respond.InternalError(context, "Couldn't retrieve ingredients")
		return
	}

	respond.Ok(context, ingredients)
}
