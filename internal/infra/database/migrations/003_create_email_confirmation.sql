-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS email_confirmations (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    email VARCHAR(255) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    provider_key VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

---- create above / drop below ----

DROP TABLE IF EXISTS email_confirmations;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
