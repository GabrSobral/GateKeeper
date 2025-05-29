-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS mfa_user_secret (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    secret VARCHAR(255) NOT NULL,
    is_validated BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

-- type MfaUserSecret struct {
-- 	ID          uuid.UUID
-- 	UserID      uuid.UUID
-- 	Secret      string
-- 	IsValidated bool
-- 	CreatedAt   time.Time
-- 	ExpiresAt   time.Time
-- }
---- create above / drop below ----
DROP TABLE IF EXISTS mfa_user_secret;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.