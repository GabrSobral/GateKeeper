-- Write your migrate up statements here
/* user_profile -- user */
ALTER TABLE
    user_profile
ADD
    CONSTRAINT fk_user_profile_user FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* email_confirmation >- application_user */
ALTER TABLE
    email_confirmation
ADD
    CONSTRAINT fk_user_email_confirmation FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* email_mfa_code >- application_user */
ALTER TABLE
    email_mfa_code
ADD
    CONSTRAINT fk_user_email_mfa_code FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* session_code >- application_user */
ALTER TABLE
    session_code
ADD
    CONSTRAINT fk_user_session_code FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* change_password_code >- application_user */
ALTER TABLE
    change_password_code
ADD
    CONSTRAINT fk_user_change_password_code FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* refresh_token >- application_user */
ALTER TABLE
    refresh_token
ADD
    CONSTRAINT fk_user_refresh_token FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* external_login >- application_user */
ALTER TABLE
    external_login
ADD
    CONSTRAINT fk_user_external_login FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* authorization_code >- application_user */
ALTER TABLE
    application_authorization_code
ADD
    CONSTRAINT fk_application_authorization_code_user FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* application_authorization_code >- application */
ALTER TABLE
    application_authorization_code
ADD
    CONSTRAINT fk_application_authorization_code_application FOREIGN KEY (application_id) REFERENCES "application" (id) ON DELETE CASCADE;

/* password_reset_token >- application_user */
ALTER TABLE
    password_reset_token
ADD
    CONSTRAINT fk_user_password_reset_token FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* application >- organization */
ALTER TABLE
    "application"
ADD
    CONSTRAINT fk_organization_application FOREIGN KEY (organization_id) REFERENCES "organization" (id) ON DELETE CASCADE;

/* application_user >- application */
ALTER TABLE
    "application_user"
ADD
    CONSTRAINT fk_application_user_application FOREIGN KEY (application_id) REFERENCES "application" (id) ON DELETE CASCADE;

/* application_role >- application */
ALTER TABLE
    "application_role"
ADD
    CONSTRAINT fk_application_role_application FOREIGN KEY (application_id) REFERENCES "application" (id) ON DELETE CASCADE;

/* user_role >- application_role */
ALTER TABLE
    user_role
ADD
    CONSTRAINT fk_user_role_application_role FOREIGN KEY (role_id) REFERENCES "application_role" (id) ON DELETE CASCADE;

/* user_role >- application_user */
ALTER TABLE
    user_role
ADD
    CONSTRAINT fk_user_role_application_user FOREIGN KEY (user_id) REFERENCES "application_user" (id) ON DELETE CASCADE;

/* application_secret >- application */
ALTER TABLE
    application_secret
ADD
    CONSTRAINT fk_application_secret_application FOREIGN KEY (application_id) REFERENCES "application" (id) ON DELETE CASCADE;

/* application_mail_config -- application (one-to-one) */
ALTER TABLE
    application_mail_config
ADD
    CONSTRAINT fk_application_mail_config_application FOREIGN KEY (application_id) REFERENCES "application" (id) ON DELETE CASCADE;

/* application_oauth_provider >- application */
ALTER TABLE
    application_oauth_provider
ADD
    CONSTRAINT fk_application_oauth_provider_application FOREIGN KEY (application_id) REFERENCES "application" (id) ON DELETE CASCADE;

---- create above / drop below ----
ALTER TABLE
    user_profile DROP CONSTRAINT fk_user_profile_user;

ALTER TABLE
    email_confirmation DROP CONSTRAINT fk_user_email_confirmation;

ALTER TABLE
    refresh_token DROP CONSTRAINT fk_user_refresh_token;

ALTER TABLE
    external_login DROP CONSTRAINT fk_user_external_login;

ALTER TABLE
    password_reset_token DROP CONSTRAINT fk_user_password_reset_token;

ALTER TABLE
    "application" DROP CONSTRAINT fk_organization_application;

ALTER TABLE
    "application_role" DROP CONSTRAINT fk_application_role_application;

ALTER TABLE
    user_role DROP CONSTRAINT fk_user_role_application_role;

ALTER TABLE
    user_role DROP CONSTRAINT fk_user_role_application_user;

ALTER TABLE
    application_secret DROP CONSTRAINT fk_application_secret_application;

ALTER TABLE
    application_mail_config DROP CONSTRAINT fk_application_mail_config_application;

ALTER TABLE
    application_oauth_provider DROP CONSTRAINT fk_application_oauth_provider_application;

ALTER TABLE
    email_mfa_code DROP CONSTRAINT fk_user_email_mfa_code;

ALTER TABLE
    session_code DROP CONSTRAINT fk_user_session_code;

ALTER TABLE
    change_password_code DROP CONSTRAINT fk_user_change_password_code;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.