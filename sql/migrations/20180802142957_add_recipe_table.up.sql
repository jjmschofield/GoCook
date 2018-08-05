BEGIN;

CREATE TABLE recipes.recipes
(
    id uuid NOT NULL,
    data jsonb NOT NULL,
    owner char(40) NOT NULL,
	  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT recipes_pkey PRIMARY KEY (id)
);

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