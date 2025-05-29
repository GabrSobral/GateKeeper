------------------------------------COMMANDS--------------------------------------
-- name: AddMfaUserSecret :exec
INSERT INTO
    mfa_user_secret (
        id,
        user_id,
        secret,
        is_validated,
        created_at,
        expires_at
    )
VALUES
    (
        sqlc.arg('id'),
        sqlc.arg('user_id'),
        sqlc.arg('secret'),
        sqlc.arg('is_validated'),
        sqlc.arg('created_at'),
        sqlc.arg('expires_at')
    );

-- name: RevokeMfaUserSecretFromUser :exec
DELETE FROM
    mfa_user_secret
WHERE
    user_id = sqlc.arg('user_id');

-- name: UpdateMfaUserSecret :exec
UPDATE
    mfa_user_secret
SET
    secret = sqlc.arg('secret'),
    is_validated = sqlc.arg('is_validated'),
    created_at = sqlc.arg('created_at'),
    expires_at = sqlc.arg('expires_at')
WHERE
    id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------
-- name: GetMfaUserSecretByUserId :one
SELECT
    id,
    user_id,
    secret,
    is_validated,
    created_at,
    expires_at
FROM
    mfa_user_secret
WHERE
    user_id = sqlc.arg('user_id');