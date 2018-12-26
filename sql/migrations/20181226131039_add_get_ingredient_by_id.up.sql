BEGIN;

CREATE FUNCTION ingredients.get_ingredient_by_id(requested_id uuid)
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
      ingredients.name,
      ingredients.ingredients.data,
      ingredients.ingredients.created_at,
      ingredients.ingredients.updated_at
    FROM ingredients.ingredients
    WHERE
      ingredients.ingredients.id = requested_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;