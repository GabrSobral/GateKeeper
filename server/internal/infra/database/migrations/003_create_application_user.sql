-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS "application_user" (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_email_confirmed BOOLEAN NOT NULL DEFAULT FALSE,
    two_factor_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    two_factor_secret VARCHAR(255) NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS application_user;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.