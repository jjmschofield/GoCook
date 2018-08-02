package recipes

import (
	"github.com/satori/go.uuid"
	"github.com/jjmschofield/GoCook/common/db"
	"encoding/json"
	"errors"
)

var store = make(map[string]Recipe)

func GetAllFromStore() map[string]Recipe {
	return store
}

func GetFromStoreById(id string) (Recipe, error) {
	query := "SELECT data from recipes.get_recipe_by_id('" + id + "')"

	recipe, err := singleRecipeQuery(query)

	if err != nil {
		return recipe, err
	}

	if recipe.Id != id {
		return recipe, errors.New("recipe not found")
	}

	return recipe, nil
}

func SaveToStore(recipe Recipe) (saved Recipe, err error) {
	if recipe.Id == "" {
		return saveNewRecipe(recipe)
	}
	
	store[recipe.Id] = recipe
	return recipe, nil
}

func saveNewRecipe(recipe Recipe) (result Recipe, err error) {
	recipe.Id = uuid.Must(uuid.NewV4()).String()

	recipeJson, err := json.Marshal(recipe)

	query := "SELECT data from recipes.save_new_recipe('" + recipe.Id + "','" + string(recipeJson) + "')"

	return singleRecipeQuery(query)
}

// TODO - this should be generalized and abstracted down into the DB package
func singleRecipeQuery(query string) (recipe Recipe, err error) {
	var recipeJson []byte

	dbErr := db.GetConnection().QueryRow(query).Scan(&recipeJson)

	if dbErr != nil {
		return Recipe{}, dbErr
	}

	var storeRecipe Recipe
	jsonErr := json.Unmarshal(recipeJson, &storeRecipe)

	return storeRecipe, jsonErr
}
