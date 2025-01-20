------------------------------------COMMANDS--------------------------------------

-- name: AddExternalLogin :exec
INSERT INTO external_login (
    user_id,
    email,
    provider,
    provider_key
) VALUES (
    sqlc.arg('user_id'), -- user_id
    sqlc.arg('email'), -- email
    sqlc.arg('provider'), -- provider
    sqlc.arg('provider_key') -- provider_key
);

------------------------------------QUERIES--------------------------------------

-- name: GetExternalLoginByProviderKey :one
SELECT
    user_id,
    email,
    provider,
    provider_key
FROM external_login
WHERE
    provider = sqlc.arg('provider') AND
    provider_key = sqlc.arg('provider_key');

-- name: GetExternalLoginByUserID :one
SELECT
    user_id,
    email,
    provider,
    provider_key
FROM external_login
WHERE
    user_id = sqlc.arg('user_id');