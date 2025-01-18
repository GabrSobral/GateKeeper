------------------------------------COMMANDS--------------------------------------

-- name: AddEmailConfirmation :exec
INSERT INTO email_confirmation (
    id,
    user_id,
    email,
    token,
    created_at,
    cool_down,
    expires_at,
    is_used
) VALUES (
    sqlc.arg('id'), -- id
    sqlc.arg('user_id'), -- user_id
    sqlc.arg('email'), -- email
    sqlc.arg('token'), -- token
    sqlc.arg('created_at'), -- created_at
    sqlc.arg('cool_down'), -- cool_down
    sqlc.arg('expires_at'), -- expires_at
    sqlc.arg('is_used') -- is_used
);

-- name: UpdateEmailConfirmation :exec
UPDATE email_confirmation SET
    user_id = sqlc.arg('user_id'),
    email = sqlc.arg('email'),
    token = sqlc.arg('token'),
    created_at = sqlc.arg('created_at'),
    cool_down = sqlc.arg('cool_down'),
    expires_at = sqlc.arg('expires_at'),
    is_used = sqlc.arg('is_used')
WHERE id = sqlc.arg('id');

-- name: DeleteEmailConfirmation :exec
DELETE FROM email_confirmation WHERE id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------

-- name: GetEmailConfirmationByEmail :one
SELECT
    id,
    user_id,
    email,
    token,
    created_at,
    cool_down,
    expires_at,
    is_used
FROM email_confirmation
WHERE 
    email = sqlc.arg('email') AND 
    user_id = sqlc.arg('user_id');