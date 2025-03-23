-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS email_mfa_code (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    email VARCHAR(128) NOT NULL,
    token VARCHAR(12) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    is_used BOOLEAN NOT NULL
);

---- create above / drop below ----
DROP TABLE IF EXISTS email_mfa_code;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.