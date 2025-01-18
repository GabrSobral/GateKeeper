-- Write your migrate up statements here

/* user_profile -- user */
ALTER TABLE user_profile
ADD CONSTRAINT fk_user_profile_user
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

/* email_confirmation >- user */
ALTER TABLE email_confirmation
ADD CONSTRAINT fk_user_email_confirmation
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

/* refresh_token >- user */
ALTER TABLE refresh_token
ADD CONSTRAINT fk_user_refresh_token
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

/* external_login >- user */
ALTER TABLE external_login
ADD CONSTRAINT fk_user_external_login
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

/* password_reset_token >- user */
ALTER TABLE password_reset_token
ADD CONSTRAINT fk_user_password_reset_token
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

/* group >- application */
ALTER TABLE "group"
ADD CONSTRAINT fk_application_group
FOREIGN KEY (application_id)
REFERENCES "application" (id)
ON DELETE CASCADE;

/* group_participation >- group */
ALTER TABLE "group_participation"
ADD CONSTRAINT fk_group_group_application
FOREIGN KEY (group_id)
REFERENCES "group" (id)
ON DELETE CASCADE;

/* group_participation >- user */
ALTER TABLE "group_participation"
ADD CONSTRAINT fk_user_group_application
FOREIGN KEY (user_id)
REFERENCES "user" (id)
ON DELETE CASCADE;

/* application >- tenant */
ALTER TABLE "application"
ADD CONSTRAINT fk_tenant_application
FOREIGN KEY (tenant_id)
REFERENCES "tenant" (id)
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

ALTER TABLE group
DROP CONSTRAINT fk_application_group;

ALTER TABLE group_participation
DROP CONSTRAINT fk_group_group_application;

ALTER TABLE group_participation
DROP CONSTRAINT fk_user_group_application;

ALTER TABLE "application"
DROP CONSTRAINT fk_tenant_application;

ALTER TABLE tenant
DROP CONSTRAINT fk_application_tenant;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
