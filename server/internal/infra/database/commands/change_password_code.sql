------------------------------------COMMANDS--------------------------------------
-- name: AddChangePasswordCode :exec
INSERT INTO
    change_password_code (
        id,
        user_id,
        email,
        token,
        created_at,
        expires_at
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
        sqlc.arg('expires_at') -- expires_at
    );

-- name: UpdateChangePasswordCode :exec
UPDATE
    change_password_code
SET
    user_id = sqlc.arg('user_id'),
    email = sqlc.arg('email'),
    token = sqlc.arg('token'),
    created_at = sqlc.arg('created_at'),
    expires_at = sqlc.arg('expires_at')
WHERE
    id = sqlc.arg('id');

-- name: DeleteChangePasswordCode :exec
DELETE FROM
    change_password_code
WHERE
    id = sqlc.arg('id');

-- name: RevokeChangePasswordCodeByUserID :exec
DELETE FROM
    change_password_code
WHERE
    user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------
-- name: GetChangePasswordCodeByToken :one
SELECT
    id,
    user_id,
    email,
    token,
    created_at,
    expires_at
FROM
    change_password_code
WHERE
    user_id = sqlc.arg('user_id')
    AND token = sqlc.arg('token');