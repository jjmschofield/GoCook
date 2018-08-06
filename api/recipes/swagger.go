package recipes

// swagger:response
type RecipeResponse struct{
	Body Recipe
}

// swagger:parameters GetRecipesById
type RecipeIdParam struct {
	// The ID of the recipe
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters UpsertRecipe
type RecipeBodyParam struct {
	// The pet to submit.
	//
	// in: body
	// required: true
	Recipe Recipe `json:"recipe"`
}

