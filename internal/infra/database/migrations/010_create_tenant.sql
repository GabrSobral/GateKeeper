-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS "tenant" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS "tenant";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
