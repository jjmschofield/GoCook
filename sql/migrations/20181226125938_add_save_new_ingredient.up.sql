BEGIN;

CREATE FUNCTION ingredients.save_new_ingredient(new_id uuid, new_data json)
RETURNS TABLE(
        id uuid,
        name char(40),
        data jsonb,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
AS $$
BEGIN
	INSERT INTO ingredients.ingredients
	  VALUES (
	    new_id,
	    new_data
	    );
  RETURN QUERY
		SELECT
      ingredients.ingredients.id,
      ingredients.ingredients.name,
      ingredients.ingredients.data,
      ingredients.ingredients.created_at,
      ingredients.ingredients.updated_at
    FROM ingredients.ingredients
		WHERE ingredients.ingredients.id = new_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;