package ingredients

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmschofield/GoCook/common/respond"
	"github.com/jjmschofield/GoCook/common/validate"
)

// swagger:route GET /ingredients{id} Ingredients GetIngredientsById
//
// Get Recipe By Id
//
// Returns the requested recipe (if the caller has access).
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
//		 404:
//       500: ErrorPayload
func getByIdRequestHandler(context *gin.Context) {
	id := context.Param("id")

	validRequest, validationError := isValidGetByIdRequest(id)

	if !validRequest {
		respond.BadRequest(context, validationError)
		return
	}

	recipe, storeErr := GetFromStoreById(id)

	if storeErr != nil {
		respond.NotFound(context)
		return
	}

	respond.Ok(context, recipe)

}

func isValidGetByIdRequest(id string) (isValid bool, validationMessage string) {
	return validate.UuidString(id)
}
