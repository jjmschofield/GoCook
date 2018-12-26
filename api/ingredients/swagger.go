package ingredients

// swagger:response
type IngredientResponse struct {
	Body Ingredient
}

// swagger:parameters GetIngredientsById
type IngredientIdParam struct {
	// The ID of the ingredient
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters UpsertIngredient
type IngredientBodyParam struct {
	// The pet to submit.
	//
	// in: body
	// required: true
	Ingredient Ingredient `json:"ingredient"`
}
