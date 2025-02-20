------------------------------------COMMANDS--------------------------------------
/*
 CREATE TABLE IF NOT EXISTS "application" (
 id UUID PRIMARY KEY,
 organization_id UUID NOT NULL,
 name VARCHAR(255) NOT NULL,
 description TEXT NULL,
 is_active BOOLEAN NOT NULL DEFAULT TRUE,
 has_mfa_auth_app BOOLEAN NOT NULL DEFAULT FALSE,
 has_mfa_email BOOLEAN NOT NULL DEFAULT FALSE,
 password_hash_secret VARCHAR(255) NOT NULL,
 badges TEXT NULL,
 created_at TIMESTAMP NOT NULL,
 updated_at TIMESTAMP NULL
 );
 */
-- name: AddApplication :exec
INSERT INTO
    "application" (
        id,
        organization_id,
        name,
        is_active,
        has_mfa_auth_app,
        has_mfa_email,
        password_hash_secret,
        badges,
        description,
        created_at,
        updated_at
    )
VALUES
    (
        sqlc.arg('id'),
        sqlc.arg('organization_id'),
        sqlc.arg('name'),
        sqlc.arg('is_active'),
        sqlc.arg('has_mfa_auth_app'),
        sqlc.arg('has_mfa_email'),
        sqlc.arg('password_hash_secret'),
        sqlc.narg('badges'),
        sqlc.narg('description'),
        sqlc.arg('created_at'),
        sqlc.arg('updated_at')
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