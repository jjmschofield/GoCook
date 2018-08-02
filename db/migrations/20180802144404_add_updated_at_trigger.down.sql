BEGIN;

DROP TRIGGER trigger_set_updated_at_timestamp on recipes.recipes;
DROP FUNCTION recipes.set_updated_at_timestamp();

COMMIT;