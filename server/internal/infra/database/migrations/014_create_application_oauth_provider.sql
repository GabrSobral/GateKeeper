-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS "application_oauth_provider" (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    client_id VARCHAR(255) NOT NULL,
    client_secret VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS application_oauth_provider;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.