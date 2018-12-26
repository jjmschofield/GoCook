BEGIN;

DROP TRIGGER trigger_ingredients_extract_name ON ingredients.ingredients;
DROP TRIGGER trigger_set_updated_at_timestamp on ingredients.ingredients;
DROP FUNCTION ingredients.extract_name;
DROP FUNCTION ingredients.set_updated_at_timestamp();
DROP TABLE ingredients.ingredients;

COMMIT;