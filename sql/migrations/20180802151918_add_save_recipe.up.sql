BEGIN;

CREATE FUNCTION recipes.save_new_recipe(newId uuid, newData json)
RETURNS TABLE(
        id uuid,
        data jsonb,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    )
AS $$
BEGIN
	INSERT INTO recipes.recipes VALUES (newId, newData);
    RETURN QUERY
		SELECT * FROM recipes.recipes
		WHERE recipes.recipes.id = newId;
END;
$$ LANGUAGE plpgsql;

COMMIT;