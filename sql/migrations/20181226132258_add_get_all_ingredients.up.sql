BEGIN;

CREATE FUNCTION ingredients.get_all_ingredients()
RETURNS TABLE(
        id uuid,
        name char(40),
        data jsonb,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
AS $$
BEGIN
	RETURN QUERY
      SELECT
        ingredients.ingredients.id,
        ingredients.ingredients.name,
        ingredients.ingredients.data,
        ingredients.ingredients.created_at,
        ingredients.ingredients.updated_at
      FROM ingredients.ingredients;
END;
$$ LANGUAGE plpgsql;

COMMIT;