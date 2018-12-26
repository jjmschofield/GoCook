BEGIN;

CREATE FUNCTION ingredients.save_ingredient(ingredient_id uuid, new_data json)
RETURNS TABLE(
        id uuid,
        name char(40),
        data jsonb,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
AS $$
BEGIN
	UPDATE ingredients.ingredients
	  SET data = new_data
	  WHERE
	    ingredients.ingredients.id = ingredient_id;
  RETURN QUERY
    SELECT
      ingredients.ingredients.id,
      ingredients.ingredients.data,
      ingredients.ingredients.created_at,
      ingredients.ingredients.updated_at
    FROM ingredients.ingredients
    WHERE
      ingredients.ingredients.id = ingredient_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;
