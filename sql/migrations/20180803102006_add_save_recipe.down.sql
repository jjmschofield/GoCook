BEGIN;

DROP FUNCTION recipes.save_recipe(recipeId uuid, newData json);

COMMIT;