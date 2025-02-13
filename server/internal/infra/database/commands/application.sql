------------------------------------COMMANDS--------------------------------------
-- name: AddApplication :exec
INSERT INTO
    "application" (
        id,
        organization_id,
        name,
        description,
        created_at,
        updated_at
    )
VALUES
    (
        sqlc.arg('id'),
        -- id
        sqlc.arg('organization_id'),
        -- organization_id
        sqlc.arg('name'),
        -- name
        sqlc.narg('description'),
        -- description
        sqlc.arg('created_at'),
        -- created_at
        sqlc.narg('updated_at') -- updated_at
    );

-- name: UpdateApplication :exec
UPDATE
    "application"
SET
    organization_id = sqlc.arg('organization_id'),
    name = sqlc.arg('name'),
    description = sqlc.narg('description'),
    updated_at = sqlc.arg('updated_at')
WHERE
    id = sqlc.arg('id');

-- name: DeleteApplication :exec
DELETE FROM
    "application"
WHERE
    id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------
-- name: CheckIfApplicationExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "application"
        WHERE
            id = sqlc.arg('id')
    );

-- name: GetApplicationByID :one
SELECT
    id,
    organization_id,
    name,
    description,
    badges,
    is_active,
    has_mfa_auth_app,
    has_mfa_email,
    password_hash_secret,
    created_at,
    updated_at
FROM
    "application"
WHERE
    id = sqlc.arg('id');

-- name: ListApplicationsFromOrganization :many
SELECT
    id,
    organization_id,
    name,
    description,
    badges,
    created_at,
    updated_at
FROM
    "application"
WHERE
    organization_id = sqlc.arg('organization_id');