BEGIN;

CREATE FUNCTION recipes.save_new_recipe(new_id uuid, new_data json)
RETURNS TABLE(
        id uuid,
        data jsonb,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
AS $$
BEGIN
	INSERT INTO recipes.recipes
	  VALUES (
	    new_id,
	    new_data,
	    new_data->>'owner'
	    );
  RETURN QUERY
		SELECT
      recipes.recipes.id,
      recipes.recipes.data,
      recipes.recipes.created_at,
      recipes.recipes.updated_at
    FROM recipes.recipes
		WHERE recipes.recipes.id = new_id;
END;
$$ LANGUAGE plpgsql;

COMMIT;