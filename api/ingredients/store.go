package ingredients

import (
	"encoding/json"
	"errors"
	"github.com/jjmschofield/GoCook/common/db"
	"github.com/satori/go.uuid"
)

func GetAllFromStore() (ingredients []Ingredient, err error) {
	return multiRowQuery("SELECT data FROM ingredients.get_all_ingredients()")
}

func GetFromStoreById(id string) (Ingredient, error) {
	query := "SELECT data from ingredients.get_ingredient_by_id('" + id + "')"

	ingredient, err := singleRowQuery(query)

	return ingredient, err
}

func SaveToStore(ingredient Ingredient, userId string) (saved Ingredient, err error) {
	if ingredient.Id == "" {
		return saveNew(ingredient, userId)
	} else {
		return update(ingredient, userId)
	}
}

func saveNew(ingredient Ingredient, userId string) (saved Ingredient, err error) {
	ingredient.Id = uuid.Must(uuid.NewV4()).String()

	ingredientJson, err := json.Marshal(ingredient)

	query := "SELECT data from ingredients.save_new_ingredient('" + ingredient.Id + "','" + string(ingredientJson) + "')"

	return singleRowQuery(query)
}

func update(ingredient Ingredient, userId string) (saved Ingredient, err error) {
	ingredientJson, err := json.Marshal(ingredient)

	query := "SELECT data from ingredients.save_ingredient('" +
		ingredient.Id + "','" +
		string(ingredientJson) + "')"

	updatedIngredient, dbErr := singleRowQuery(query)

	if updatedIngredient.Id != ingredient.Id {
		return Ingredient{}, errors.New("ingredient could not be saved")
	}

	return updatedIngredient, dbErr
}

func singleRowQuery(query string) (ingredient Ingredient, err error) {
	dbResult, dbErr := db.SingleRowQuery(query)

	if dbErr != nil {
		return Ingredient{}, dbErr
	}

	err = json.Unmarshal(dbResult, &ingredient)

	if err != nil {
		return Ingredient{}, err
	}

	return ingredient, nil
}

func multiRowQuery(query string) (ingredients []Ingredient, err error) {
	dbResults, queryErr := db.MultiRowQuery(query)

	if queryErr != nil {
		return nil, queryErr
	}

	ingredients = make([]Ingredient, len(dbResults))

	for i := 0; i < len(dbResults); i++ {
		var ingredient Ingredient

		err = json.Unmarshal(dbResults[i], &ingredient)

		if err != nil {
			return nil, err
		}

		ingredients[i] = ingredient
	}

	return ingredients, nil
}
