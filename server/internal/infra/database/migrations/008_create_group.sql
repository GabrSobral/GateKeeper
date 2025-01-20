-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS "group" (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS "group";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
