-- name: AddTenant :exec
INSERT INTO tenant (
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

-- name: RemoveTenant :exec
DELETE FROM tenant
WHERE id = sqlc.arg('tenant_id');

-- name: UpdateTenant :exec
UPDATE tenant
SET
    name = sqlc.arg('name'),
    description = sqlc.arg('description'),
    updated_at = sqlc.arg('updated_at')
WHERE id = sqlc.arg('id');

-- name: ListTenants :many
SELECT
    id,
    name,
    description,
    created_at,
    updated_at
FROM tenant
ORDER BY created_at DESC;