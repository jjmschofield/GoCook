package recipes

import (
	"github.com/satori/go.uuid"
)

var store = make(map [string]Recipe)

func GetAllFromStore() map [string]Recipe {
	return store
}

func GetFromStoreById(id string) (Recipe, bool) {
	recipe, found := store[id]
	return recipe, found
}

func SaveToStore(recipe Recipe) Recipe {
	if recipe.Id == "" {
		recipe.Id = uuid.Must(uuid.NewV4()).String()
	}

	store[recipe.Id] = recipe
	return recipe
}