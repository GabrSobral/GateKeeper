-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS "application_role" (
    id UUID PRIMARY KEY,
    application_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    /* application_role >- application = fk_application_role_application */
    CONSTRAINT fk_application_role_application FOREIGN KEY (application_id) REFERENCES "application" (id) ON DELETE CASCADE
);

---- create above / drop below ----
DROP TABLE IF EXISTS application_role;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.