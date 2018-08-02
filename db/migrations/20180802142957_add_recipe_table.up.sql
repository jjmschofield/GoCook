BEGIN;

CREATE TABLE recipes.recipes
(
    id uuid NOT NULL,
    data jsonb NOT NULL,
	  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT recipes_pkey PRIMARY KEY (id)
);

COMMIT;