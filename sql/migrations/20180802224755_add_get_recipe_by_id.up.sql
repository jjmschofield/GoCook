BEGIN;

CREATE FUNCTION recipes.get_recipe_by_id(requested_id uuid, user_id char(40))
RETURNS TABLE(
        id uuid,
        data jsonb,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
AS $$
BEGIN
	RETURN QUERY
    SELECT
      recipes.recipes.id,
      recipes.recipes.data,
      recipes.recipes.created_at,
      recipes.recipes.updated_at
    FROM recipes.recipes
    WHERE
      recipes.recipes.id = requested_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;