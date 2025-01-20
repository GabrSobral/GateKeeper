------------------------------------COMMANDS--------------------------------------

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

------------------------------------QUERIES--------------------------------------

-- name: GetUserProfileByUserId :one

SELECT
    user_id,
    first_name,
    last_name,
    phone_number,
    "address",
    photo_url
FROM
    user_profile
WHERE
    user_id = sqlc.arg('user_id');
