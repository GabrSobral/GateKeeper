------------------------------------COMMANDS--------------------------------------
-- name: AddUserProfile :exec
INSERT INTO
    user_profile (
        user_id,
        display_name,
        first_name,
        last_name,
        phone_number,
        "address",
        photo_url
    )
VALUES
    (
        sqlc.arg('user_id'),
        -- user_id
        sqlc.arg('display_name'),
        -- display_name
        sqlc.arg('first_name'),
        -- first_name
        sqlc.arg('last_name'),
        -- last_name
        sqlc.narg('phone_number'),
        -- phone_number
        sqlc.narg('address'),
        -- address
        sqlc.narg('photo_url') -- photo_url
    );

-- name: UpdateUserProfile :exec
UPDATE
    user_profile
SET
    display_name = sqlc.arg('display_name'),
    first_name = sqlc.arg('first_name'),
    last_name = sqlc.arg('last_name'),
    phone_number = sqlc.narg('phone_number'),
    "address" = sqlc.narg('address'),
    photo_url = sqlc.narg('photo_url')
WHERE
    user_id = sqlc.arg('user_id');

------------------------------------QUERIES--------------------------------------
-- name: GetUserProfileByUserId :one
SELECT
    user_id,
    display_name,
    first_name,
    last_name,
    phone_number,
    "address",
    photo_url
FROM
    user_profile
WHERE
    user_id = sqlc.arg('user_id');