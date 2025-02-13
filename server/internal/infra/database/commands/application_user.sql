------------------------------------COMMANDS--------------------------------------
-- name: AddUser :exec
-- Add user
INSERT INTO
    "application_user" (
        id,
        email,
        password_hash,
        application_id,
        created_at,
        updated_at,
        is_active,
        is_email_confirmed,
        two_factor_enabled,
        two_factor_secret
    )
VALUES
    (
        sqlc.arg('id'),
        -- id
        sqlc.arg('email'),
        -- email
        sqlc.narg('password_hash'),
        -- password_hash
        sqlc.arg('application_id'),
        -- application_id
        sqlc.arg('created_at'),
        -- created_at
        sqlc.narg('updated_at'),
        -- updated_at
        sqlc.arg('is_active'),
        -- is_active
        sqlc.arg('is_email_confirmed'),
        -- is_email_confirmed
        sqlc.arg('two_factor_enabled'),
        -- two_factor_enabled
        sqlc.narg('two_factor_secret') -- two_factor_secret
    );

-- name: UpdateUser :exec
-- Update user
UPDATE
    "application_user"
SET
    email = sqlc.arg('email'),
    password_hash = sqlc.narg('password_hash'),
    updated_at = sqlc.arg('updated_at'),
    is_active = sqlc.arg('is_active'),
    is_email_confirmed = sqlc.arg('is_email_confirmed'),
    two_factor_enabled = sqlc.arg('two_factor_enabled'),
    two_factor_secret = sqlc.narg('two_factor_secret')
WHERE
    id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------
-- name: GetUserById :one
-- Get user by id
SELECT
    id,
    email,
    application_id,
    password_hash,
    created_at,
    updated_at,
    is_active,
    is_email_confirmed,
    two_factor_enabled,
    two_factor_secret
FROM
    "application_user"
WHERE
    id = sqlc.arg('id');

-- name: GetUserByEmail :one
-- Get user by email
SELECT
    id,
    email,
    application_id,
    password_hash,
    created_at,
    updated_at,
    is_active,
    is_email_confirmed,
    two_factor_enabled,
    two_factor_secret
FROM
    "application_user"
WHERE
    email = sqlc.arg('email')
    AND application_id = sqlc.arg('application_id');

-- name: IsUserExistsByEmail :one
-- Check if user exists by email
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "application_user"
        WHERE
            email = sqlc.arg('email')
            AND application_id = sqlc.arg('application_id')
    );

-- name: IsUserExistsById :one
-- Check if user exists by id
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "application_user"
        WHERE
            id = sqlc.arg('id')
    );