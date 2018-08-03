BEGIN;

CREATE FUNCTION recipes.save_recipe(recipe_id uuid, new_data json)
RETURNS TABLE(
        id uuid,
        data jsonb,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
AS $$
BEGIN
	UPDATE recipes.recipes
	  SET data = new_data
	  WHERE recipes.recipes.id = recipe_id;
  RETURN QUERY
    SELECT
      recipes.recipes.id,
      recipes.recipes.data,
      recipes.recipes.created_at,
      recipes.recipes.updated_at
    FROM recipes.recipes
    WHERE recipes.recipes.id = recipe_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;
