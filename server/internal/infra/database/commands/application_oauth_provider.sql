------------------------------------COMMANDS--------------------------------------
-- name: AddApplicationOauthProvider :exec 
-- Add a new application oauth provider
INSERT INTO
    application_oauth_provider (
        id,
        application_id,
        name,
        client_id,
        client_secret,
        redirect_uri,
        created_at,
        updated_at,
        "enabled"
    )
VALUES
    (
        sqlc.arg('id'),
        sqlc.arg('application_id'),
        sqlc.arg('name'),
        sqlc.arg('client_id'),
        sqlc.arg('client_secret'),
        sqlc.arg('redirect_uri'),
        sqlc.arg('created_at'),
        sqlc.arg('updated_at'),
        sqlc.arg('enabled')
    );

-- name: UpdateApplicationOauthProvider :exec 
-- Update the application oauth provider
UPDATE
    application_oauth_provider
SET
    name = sqlc.arg('name'),
    client_id = sqlc.arg('client_id'),
    client_secret = sqlc.arg('client_secret'),
    redirect_uri = sqlc.arg('redirect_uri'),
    updated_at = sqlc.arg('updated_at'),
    "enabled" = sqlc.arg('enabled')
WHERE
    id = sqlc.arg('id');

------------------------------------QUERIES--------------------------------------
-- name: GetApplicationOauthProviderByID :one
-- Get application oauth provider by ID
SELECT
    id,
    application_id,
    name,
    client_id,
    client_secret,
    redirect_uri,
    created_at,
    updated_at,
    "enabled"
FROM
    application_oauth_provider
WHERE
    id = sqlc.arg('id');

-- name: GetApplicationOauthProvidersByApplicationID :many
-- Get application oauth providers by application ID
SELECT
    id,
    application_id,
    name,
    client_id,
    client_secret,
    redirect_uri,
    created_at,
    updated_at,
    "enabled"
FROM
    application_oauth_provider
WHERE
    application_id = sqlc.arg('application_id');

-- name: GetApplicationOauthProviderByName :one
-- Get application oauth providers by application ID
SELECT
    id,
    application_id,
    name,
    client_id,
    client_secret,
    redirect_uri,
    created_at,
    updated_at,
    "enabled"
FROM
    application_oauth_provider
WHERE
    application_id = sqlc.arg('application_id')
    AND name = sqlc.arg('name');

-- name: CheckIfApplicationOauthProviderConfigurationExists :one
-- Check if application oauth provider configuration exists
SELECT
    EXISTS (
        SELECT
            1
        FROM
            application_oauth_provider
        WHERE
            application_id = sqlc.arg('application_id')
            AND name = sqlc.arg('name')
    ) AS EXISTS;