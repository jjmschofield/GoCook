BEGIN;

DROP FUNCTION ingredients.save_ingredient(ingredient_id uuid, new_data json);

COMMIT;