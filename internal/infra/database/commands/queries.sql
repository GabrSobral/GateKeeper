-- name: GetUserById :one

SELECT
    id,
    email,
    password_hash,
    created_at,
    updated_at,
    is_active,
    is_email_confirmed,
    two_factor_enabled,
    two_factor_secret
FROM
    "user"
WHERE
    id = sqlc.arg('id');

-- name: GetUserByEmail :one

SELECT
    id,
    email,
    password_hash,
    created_at,
    updated_at,
    is_active,
    is_email_confirmed,
    two_factor_enabled,
    two_factor_secret
FROM
    "user"
WHERE
    email = sqlc.arg('email');

-- name: IsUserExistsByEmail :one

SELECT
    EXISTS (
        SELECT 1 FROM "user"
        WHERE email = sqlc.arg('email')
    );

-- name: IsUserExistsById :one

SELECT
    EXISTS (
        SELECT 1 FROM "user"
        WHERE id = sqlc.arg('id')
    );

-- name: GetUserProfileByUserId :one

SELECT
    user_id,
    first_name,
    last_name,
    phone_number,
    "address",
    photo_url
FROM
    user_profile
WHERE
    user_id = sqlc.arg('user_id');


-- name: GetEmailConfirmationByEmail :one

SELECT
    id,
    user_id,
    email,
    token,
    created_at,
    cool_down,
    expires_at,
    is_used
FROM email_confirmation
WHERE 
    email = sqlc.arg('email') AND 
    user_id = sqlc.arg('user_id');

-- name: GetRefreshTokensFromUser :many

SELECT
    id,
    user_id,
    available_refreshes,
    expires_at,
    created_at
FROM refresh_token
WHERE user_id = sqlc.arg('user_id');

-- name: GetExternalLoginByProviderKey :one

SELECT
    user_id,
    email,
    provider,
    provider_key
FROM external_login
WHERE
    provider = sqlc.arg('provider') AND
    provider_key = sqlc.arg('provider_key');

-- name: GetExternalLoginByUserID :one

SELECT
    user_id,
    email,
    provider,
    provider_key
FROM external_login
WHERE
    user_id = sqlc.arg('user_id');

-- name: GetPasswordResetByTokenID :one

SELECT
    id,
    user_id,
    token,
    created_at,
    expires_at
FROM password_reset_token
WHERE
    id = sqlc.arg('id');