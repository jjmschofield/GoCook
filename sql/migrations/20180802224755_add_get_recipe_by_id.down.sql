BEGIN;

DROP FUNCTION recipes.get_recipe_by_id(id uuid, user_id char(40));

COMMIT;