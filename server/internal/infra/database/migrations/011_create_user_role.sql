-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS user_role (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS user_role;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.