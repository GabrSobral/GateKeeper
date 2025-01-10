-- name: AddUser :exec

INSERT INTO "user" (
    id, 
    email, 
    password_hash, 
    created_at, 
    updated_at,
    is_active,
    is_email_confirmed,
    two_factor_enabled,
    two_factor_secret
) VALUES (
    sqlc.arg('id'), -- id
    sqlc.arg('email'), -- email
    sqlc.narg('password_hash'), -- password_hash
    sqlc.arg('created_at'), -- created_at
    sqlc.narg('updated_at'), -- updated_at
    sqlc.arg('is_active'), -- is_active
    sqlc.arg('is_email_confirmed'), -- is_email_confirmed
    sqlc.arg('two_factor_enabled'), -- two_factor_enabled
    sqlc.narg('two_factor_secret')  -- two_factor_secret
);

-- name: UpdateUser :exec

UPDATE "user" SET
    email = sqlc.arg('email'),
    password_hash = sqlc.narg('password_hash'),
    updated_at = sqlc.arg('updated_at'),
    is_active = sqlc.arg('is_active'),
    is_email_confirmed = sqlc.arg('is_email_confirmed'),
    two_factor_enabled = sqlc.arg('two_factor_enabled'),
    two_factor_secret = sqlc.narg('two_factor_secret')
WHERE id = sqlc.arg('id');