-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS external_login (
    user_id UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    provider VARCHAR(255) NOT NULL,
    provider_key VARCHAR(255) NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS external_login;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
