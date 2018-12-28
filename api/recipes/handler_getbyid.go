package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
	"go.uber.org/zap"
)

// swagger:route GET /recipes{id} Recipes GetRecipesById
//
// Get Recipe By Id
//
// Returns the requested recipe (if the caller has access).
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
//       200: Recipe
//		 400: MessagePayload
//		 404:
//       500: ErrorPayload
func getByIdRequestHandler(context *gin.Context) {
	id := context.Param("id")
	logger := context.MustGet("logger").(zap.Logger)

	validRequest, validationError := isValidGetByIdRequest(id)

	if !validRequest {
		respond.BadRequest(context, validationError)
		return
	}

	recipe, storeErr := GetFromStoreById(id, context.MustGet("userId").(string))

	if storeErr != nil {
		logger.Error("Could not get recipe with id " + id, zap.Error(storeErr))
		respond.NotFound(context)
		return
	}

	respond.Ok(context, recipe)
}

func isValidGetByIdRequest(id string) (isValid bool, validationMessage string) {
	return validate.UuidString(id)
}
