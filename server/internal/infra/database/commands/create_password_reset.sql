------------------------------------COMMANDS--------------------------------------

-- name: CreatePasswordReset :exec
INSERT INTO password_reset_token (
    id,
    user_id,
    token,
    created_at,
    expires_at
) VALUES (
    sqlc.arg('id'), -- id
    sqlc.arg('user_id'), -- user_id
    sqlc.arg('token'), -- token
    sqlc.arg('created_at'), -- created_at
    sqlc.arg('expires_at') -- expires_at
);

-- name: DeletePasswordResetFromUser :exec
DELETE FROM password_reset_token WHERE user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------

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