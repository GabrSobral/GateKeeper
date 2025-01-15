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

-- name: AddUserProfile :exec

INSERT INTO user_profile (
    user_id,
    first_name,
    last_name,
    phone_number,
    "address",
    photo_url
) VALUES (
    sqlc.arg('user_id'), -- user_id
    sqlc.arg('first_name'), -- first_name
    sqlc.arg('last_name'), -- last_name
    sqlc.narg('phone_number'), -- phone_number
    sqlc.narg('address'), -- address
    sqlc.narg('photo_url') -- photo_url
);

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