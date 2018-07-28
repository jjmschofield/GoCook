package recipies

import (
	"github.com/satori/go.uuid"
	"github.com/jjmschofield/GoCook/recipes/models"
)

var store = make(map [string]recipes.Recipe);

func GetAllFromStore() map [string]recipes.Recipe {
	return store;
}

func GetFromStoreById(id string) recipes.Recipe {
	recipe := store[id]
	return recipe
}

func SaveToStore(recipe recipes.Recipe) recipes.Recipe {
	if(recipe.Id == ""){
		recipe.Id = uuid.Must(uuid.NewV4()).String()
	}

	store[recipe.Id] = recipe
	return recipe
}