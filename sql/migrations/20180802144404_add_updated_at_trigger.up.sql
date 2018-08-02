BEGIN;

CREATE FUNCTION recipes.set_updated_at_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_updated_at_timestamp
BEFORE UPDATE ON recipes.recipes
FOR EACH ROW
EXECUTE PROCEDURE recipes.set_updated_at_timestamp();

COMMIT;