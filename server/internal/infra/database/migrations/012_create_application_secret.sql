-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS "application_secret" (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    expires_at TIMESTAMP NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS application_secret;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.