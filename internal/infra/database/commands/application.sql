------------------------------------COMMANDS--------------------------------------

-- name: AddApplication :exec
INSERT INTO "application" (
    id,
    tenant_id,
    name,
    description,
    created_at,
    updated_at
) VALUES (
    sqlc.arg('id'), -- id
    sqlc.arg('tenant_id'), -- tenant_id
    sqlc.arg('name'), -- name
    sqlc.narg('description'), -- description
    sqlc.arg('created_at'), -- created_at
    sqlc.narg('updated_at') -- updated_at
);

-- name: UpdateApplication :exec
UPDATE "application" SET
    tenant_id = sqlc.arg('tenant_id'),
    name = sqlc.arg('name'),
    description = sqlc.narg('description'),
    updated_at = sqlc.arg('updated_at')
WHERE id = sqlc.arg('id');

-- name: DeleteApplication :exec
DELETE FROM "application" WHERE id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------

-- name: GetApplicationByID :one
SELECT
    id,
    tenant_id,
    name,
    description,
    created_at,
    updated_at
FROM "application"
WHERE
    id = sqlc.arg('id');

-- name: ListApplicationsFromTenant :many
SELECT
    id,
    tenant_id,
    name,
    description,
    created_at,
    updated_at
FROM "application"
WHERE
    tenant_id = sqlc.arg('tenant_id');