BEGIN;

DROP FUNCTION recipes.save_recipe(recipe_id uuid, new_data json, user_id char(40));

COMMIT;