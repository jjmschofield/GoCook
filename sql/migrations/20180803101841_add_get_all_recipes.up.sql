BEGIN;

CREATE FUNCTION recipes.get_all_recipes()
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
      FROM recipes.recipes;
END;
$$ LANGUAGE plpgsql;

COMMIT;