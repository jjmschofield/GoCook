package recipes

import (
	"github.com/satori/go.uuid"
	"github.com/jjmschofield/GoCook/common/db"
	"encoding/json"
	"errors"
	"database/sql"
)

func GetAllFromStore(userId string) (recipes map[string]Recipe, err error) {
	return multiRecipeQuery("SELECT data FROM recipes.get_all_recipes('" + userId + "')")
}

func GetFromStoreById(id string, userId string) (Recipe, error) {
	query := "SELECT data from recipes.get_recipe_by_id('" + id + "','"+ userId+ "')"

	recipe, err := singleRecipeQuery(query)

	if err != nil {
		return recipe, err
	}

	if recipe.Id != id {
		return recipe, errors.New("recipe not found")
	}

	return recipe, nil
}

func SaveToStore(recipe Recipe, userId string) (saved Recipe, err error) {
	if recipe.Id == "" {
		return saveNewRecipe(recipe, userId)
	} else {
		return updateRecipe(recipe, userId)
	}
}

func saveNewRecipe(recipe Recipe, owner string) (result Recipe, err error) {
	recipe.Owner = owner
	recipe.Id = uuid.Must(uuid.NewV4()).String()

	recipeJson, err := json.Marshal(recipe)

	query := "SELECT data from recipes.save_new_recipe('" + recipe.Id + "','" + string(recipeJson) + "')"

	return singleRecipeQuery(query)
}

func updateRecipe(recipe Recipe, userId string) (result Recipe, err error) {
	// TODO - this should make it easier to determine an error as a query error or a permissions error

	recipeJson, err := json.Marshal(recipe)

	query := "SELECT data from recipes.save_recipe('" +
		recipe.Id + "','" +
		string(recipeJson) + "','" +
		userId + "')"

	updatedRecipe, dbErr := singleRecipeQuery(query)

	if updatedRecipe.Id != recipe.Id {
		return Recipe{}, errors.New("recipe could not be saved")
	}

	return updatedRecipe, dbErr
}

// TODO - after this line should be generalized and abstracted down into the DB package
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

func multiRecipeQuery(query string) (recipes map[string]Recipe, queryErr error) {
	rows, queryErr := db.GetConnection().Query(query)

	if queryErr != nil {
		return nil, queryErr
	}

	defer rows.Close()

	return getAllRecipesFromRows(rows)
}

func getAllRecipesFromRows(rows *sql.Rows) (recipes map[string]Recipe, err error){
	storeRecipes := make(map[string]Recipe)

	for rows.Next() {

		var recipeJson []byte

		scanErr := rows.Scan(&recipeJson)
		if scanErr != nil {
			return nil, scanErr
		}

		var storeRecipe Recipe
		jsonErr := json.Unmarshal(recipeJson, &storeRecipe)

		if jsonErr != nil{
			return nil, jsonErr
		}

		storeRecipes[storeRecipe.Id] = storeRecipe
	}

	rowsErr := rows.Err()

	return storeRecipes, rowsErr
}
