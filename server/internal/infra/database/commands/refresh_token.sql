------------------------------------COMMANDS--------------------------------------

-- name: AddRefreshToken :exec
INSERT INTO refresh_token (
    id,
    user_id,
    available_refreshes,
    expires_at,
    created_at
) VALUES (
    sqlc.arg('id'), -- id
    sqlc.arg('user_id'), -- user_id
    sqlc.arg('available_refreshes'), -- available_refreshes
    sqlc.arg('expires_at'), -- expires_at
    sqlc.arg('created_at') -- created_at
);

-- name: RevokeRefreshTokenFromUser :exec
DELETE FROM refresh_token WHERE user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------

-- name: GetRefreshTokensFromUser :many
SELECT
    id,
    user_id,
    available_refreshes,
    expires_at,
    created_at
FROM refresh_token
WHERE user_id = sqlc.arg('user_id');