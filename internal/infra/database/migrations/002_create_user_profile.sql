-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS user_profile (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255),
    address VARCHAR(255),
    photo_url VARCHAR(255)
);

---- create above / drop below ----

DROP TABLE IF EXISTS user_profile;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.