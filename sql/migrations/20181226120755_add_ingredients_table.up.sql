BEGIN;

CREATE TABLE ingredients.ingredients
(
    id uuid NOT NULL,
    name char(40) NOT NULL,
    data jsonb NOT NULL,
	  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT ingredients_pkey PRIMARY KEY (id)
);

CREATE FUNCTION ingredients.set_updated_at_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_set_updated_at_timestamp
BEFORE UPDATE ON ingredients.ingredients
FOR EACH ROW
EXECUTE PROCEDURE ingredients.set_updated_at_timestamp();

CREATE OR REPLACE FUNCTION ingredients.extract_name() RETURNS TRIGGER AS $$
BEGIN
    NEW.name = (
        SELECT data->>'name'
        FROM json(NEW.data) data);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_ingredients_extract_name
BEFORE INSERT ON ingredients.ingredients
FOR EACH ROW EXECUTE PROCEDURE ingredients.extract_name();

COMMIT;