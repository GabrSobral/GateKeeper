-- Write your migrate up statements here

ALTER TABLE user_profile
ADD CONSTRAINT fk_user_profile_user
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

ALTER TABLE email_confirmation
ADD CONSTRAINT fk_user_email_confirmation
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

ALTER TABLE refresh_token
ADD CONSTRAINT fk_user_refresh_token
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

ALTER TABLE external_login
ADD CONSTRAINT fk_user_external_login
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

ALTER TABLE password_reset_token
ADD CONSTRAINT fk_user_password_reset_token
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

---- create above / drop below ----

ALTER TABLE user_profile
DROP CONSTRAINT fk_user_profile_user;

ALTER TABLE email_confirmation
DROP CONSTRAINT fk_user_email_confirmation;

ALTER TABLE refresh_token
DROP CONSTRAINT fk_user_refresh_token;

ALTER TABLE external_login
DROP CONSTRAINT fk_user_external_login;

ALTER TABLE password_reset_token
DROP CONSTRAINT fk_user_password_reset_token;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
