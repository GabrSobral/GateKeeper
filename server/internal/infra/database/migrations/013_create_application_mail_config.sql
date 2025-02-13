-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS "application_mail_config" (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    host VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    port INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS application_mail_config;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.