------------------------------------COMMANDS--------------------------------------
-- name: AddSessionCode :exec
INSERT INTO
    session_code (
        id,
        user_id,
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
        -- email
        sqlc.arg('token'),
        -- token
        sqlc.arg('created_at'),
        -- created_at
        sqlc.arg('expires_at'),
        -- expires_at
        sqlc.arg('is_used') -- is_used
    );

-- name: UpdateSessionCode :exec
UPDATE
    session_code
SET
    user_id = sqlc.arg('user_id'),
    token = sqlc.arg('token'),
    created_at = sqlc.arg('created_at'),
    expires_at = sqlc.arg('expires_at'),
    is_used = sqlc.arg('is_used')
WHERE
    id = sqlc.arg('id');

-- name: DeleteSessionCode :exec
DELETE FROM
    session_code
WHERE
    id = sqlc.arg('id');

-- name: RevokeSessionCodeByUserID :exec
DELETE FROM
    session_code
WHERE
    user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------
-- name: GetSessionCodeByToken :one
SELECT
    id,
    user_id,
    token,
    created_at,
    expires_at,
    is_used
FROM
    session_code
WHERE
    user_id = sqlc.arg('user_id')
    AND token = sqlc.arg('token');