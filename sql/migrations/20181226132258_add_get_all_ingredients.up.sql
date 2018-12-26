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
        id,
        name,
        data,
        created_at,
        updated_at
      FROM ingredients.ingredients;
END;
$$ LANGUAGE plpgsql;

COMMIT;