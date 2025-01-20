------------------------------------COMMANDS--------------------------------------

-- name: AddGroup :exec
INSERT INTO "group" (
    id,
    application_id,
    name,
    description,
    created_at,
    updated_at
) VALUES (
    sqlc.arg('id'), -- id
    sqlc.arg('application_id'), -- application_id
    sqlc.arg('name'), -- name
    sqlc.narg('description'), -- description
    sqlc.arg('created_at'), -- created_at
    sqlc.narg('updated_at') -- updated_at
);

-- name: RemoveGroup :exec
DELETE FROM "group" WHERE id = sqlc.arg('id');

-- name: UpdateGroup :exec
UPDATE "group" SET
    application_id = sqlc.arg('application_id'),
    name = sqlc.arg('name'),
    description = sqlc.narg('description'),
    updated_at = sqlc.arg('updated_at')
WHERE id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------

-- name: GetGroupById :one
SELECT
    id,
    application_id,
    name,
    description,
    created_at,
    updated_at
FROM "group"
WHERE
    id = sqlc.arg('id');

-- name: ListGroupsFromApplication :many
SELECT
    id,
    application_id,
    name,
    description,
    created_at,
    updated_at
FROM "group"
WHERE
    application_id = sqlc.arg('application_id');