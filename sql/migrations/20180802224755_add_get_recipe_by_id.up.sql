BEGIN;

CREATE FUNCTION recipes.get_recipe_by_id(requested_id uuid)
RETURNS TABLE(
        id uuid,
        data jsonb,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
    )
AS $$
BEGIN
	RETURN QUERY
      SELECT * FROM recipes.recipes
      WHERE recipes.recipes.id = requested_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;