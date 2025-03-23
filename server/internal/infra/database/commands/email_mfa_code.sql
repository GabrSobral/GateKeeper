------------------------------------COMMANDS--------------------------------------
-- name: AddEmailMfaCode :exec
INSERT INTO
    email_mfa_code (
        id,
        user_id,
        email,
        token,
        created_at,
        expires_at,
        is_used
    )
VALUES
    (
        sqlc.arg('id'),
        -- id
        sqlc.arg('user_id'),
        -- user_id
        sqlc.arg('email'),
        -- email
        sqlc.arg('token'),
        -- token
        sqlc.arg('created_at'),
        -- created_at
        sqlc.arg('expires_at'),
        -- expires_at
        sqlc.arg('is_used') -- is_used
    );

-- name: UpdateEmailMfaCode :exec
UPDATE
    email_mfa_code
SET
    user_id = sqlc.arg('user_id'),
    email = sqlc.arg('email'),
    token = sqlc.arg('token'),
    created_at = sqlc.arg('created_at'),
    expires_at = sqlc.arg('expires_at'),
    is_used = sqlc.arg('is_used')
WHERE
    id = sqlc.arg('id');

-- name: DeleteEmailMfaCode :exec
DELETE FROM
    email_mfa_code
WHERE
    id = sqlc.arg('id');

-- name: RevokeEmailMfaCodeByUserID :exec
DELETE FROM
    email_mfa_code
WHERE
    user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------
-- name: GetEmailMfaCodeByToken :one
SELECT
    id,
    user_id,
    email,
    token,
    created_at,
    expires_at,
    is_used
FROM
    email_mfa_code
WHERE
    user_id = sqlc.arg('user_id')
    AND token = sqlc.arg('token');