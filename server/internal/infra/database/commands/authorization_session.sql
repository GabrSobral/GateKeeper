------------------------------------COMMANDS--------------------------------------
-- name: AddAuthorizationSession :exec
INSERT INTO
    authorization_session (
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

-- name: UpdateAuthorizationSession :exec
UPDATE
    authorization_session
SET
    user_id = sqlc.arg('user_id'),
    token = sqlc.arg('token'),
    created_at = sqlc.arg('created_at'),
    expires_at = sqlc.arg('expires_at'),
    is_used = sqlc.arg('is_used')
WHERE
    id = sqlc.arg('id');

-- name: DeleteAuthorizationSession :exec
DELETE FROM
    authorization_session
WHERE
    id = sqlc.arg('id');

-- name: RevokeAuthorizationSession :exec
DELETE FROM
    authorization_session
WHERE
    user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------
-- name: GetAuthorizationSession :one
SELECT
    id,
    user_id,
    token,
    created_at,
    expires_at,
    is_used
FROM
    authorization_session
WHERE
    user_id = sqlc.arg('user_id')
    AND token = sqlc.arg('token');