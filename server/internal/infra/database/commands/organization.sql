------------------------------------COMMANDS--------------------------------------

-- name: AddOrganization :exec
INSERT INTO organization (
    id,
    name,
    description,
    created_at
) VALUES (
    sqlc.arg('user_id'), -- user_id
    sqlc.arg('name'), -- name
    sqlc.arg('description'), -- description
    sqlc.arg('created_at') -- user_id
);

-- name: RemoveOrganization :exec
DELETE FROM organization
WHERE id = sqlc.arg('organization_id');

-- name: UpdateOrganization :exec
UPDATE organization
SET
    name = sqlc.arg('name'),
    description = sqlc.arg('description'),
    updated_at = sqlc.arg('updated_at')
WHERE id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------

-- name: GetOrganizationByID :one

SELECT
    id,
    name,
    description,
    created_at,
    updated_at
FROM "organization"
WHERE
    id = sqlc.arg('organization_id');

-- name: ListOrganizations :many
SELECT
    id,
    name,
    description,
    created_at,
    updated_at
FROM organization
ORDER BY created_at DESC;