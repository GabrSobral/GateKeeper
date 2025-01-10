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
    id = $1;

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
    email = $1;

-- name: IsUserExistsByEmail :one

SELECT
    EXISTS (
        SELECT 1 FROM "user"
        WHERE email = $1
    );

-- name: IsUserExistsById :one

SELECT
    EXISTS (
        SELECT 1 FROM "user"
        WHERE id = $1
    );