------------------------------------COMMANDS--------------------------------------
-- name: AddSecret :exec 
-- Add Secret to Application
INSERT INTO
    application_secret (
        id,
        application_id,
        name,
        value,
        created_at,
        updated_at,
        expires_at
    )
VALUES
    (
        sqlc.arg('id'),
        sqlc.arg('application_id'),
        sqlc.arg('name'),
        sqlc.arg('value'),
        sqlc.arg('created_at'),
        sqlc.arg('updated_at'),
        sqlc.arg('expires_at')
    );

-- name: RemoveSecret :exec
-- Remove Secret from Application
DELETE FROM
    application_secret
WHERE
    id = sqlc.arg('id');

------------------------------------QUERIES---------------------------------------
-- List Secrets from Application
-- name: ListSecretsFromApplication :many
SELECT
    id,
    application_id,
    name,
    value,
    created_at,
    updated_at,
    expires_at
FROM
    application_secret
WHERE
    application_id = sqlc.arg('application_id');