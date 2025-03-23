-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS session_code (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    token VARCHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    is_used BOOLEAN NOT NULL
);

---- create above / drop below ----
DROP TABLE IF EXISTS session_code;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.