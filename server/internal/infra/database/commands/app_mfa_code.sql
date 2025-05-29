------------------------------------COMMANDS--------------------------------------
-- name: AddAppMfaCode :exec
INSERT INTO
    app_mfa_code (
        id,
        user_id,
        email,
        created_at,
        expires_at
    )
VALUES
    (
        sqlc.arg('id'),
        sqlc.arg('user_id'),
        sqlc.arg('email'),
        sqlc.arg('created_at'),
        sqlc.arg('expires_at')
    );

-- name: UpdateAppMfaCode :exec
UPDATE
    app_mfa_code
SET
    user_id = sqlc.arg('user_id'),
    email = sqlc.arg('email'),
    created_at = sqlc.arg('created_at'),
    expires_at = sqlc.arg('expires_at')
WHERE
    id = sqlc.arg('id');

-- name: DeleteAppMfaCode :exec
DELETE FROM
    app_mfa_code
WHERE
    id = sqlc.arg('id');

-- name: RevokeAppMfaCodeByUserID :exec
DELETE FROM
    app_mfa_code
WHERE
    user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------
-- name: GetAppMfaCodeByID :one
SELECT
    id,
    user_id,
    email,
    created_at,
    expires_at
FROM
    app_mfa_code
WHERE
    id = sqlc.arg('id');