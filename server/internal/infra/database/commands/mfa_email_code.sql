------------------------------------COMMANDS--------------------------------------
-- name: AddMfaEmailCode :exec
INSERT INTO
    mfa_email_code (
        id,
        mfa_method_id,
        token,
        created_at,
        expires_at,
        verified
    )
VALUES
    (
        sqlc.arg('id'),
        sqlc.arg('mfa_method_id'),
        sqlc.arg('token'),
        sqlc.arg('created_at'),
        sqlc.arg('expires_at'),
        sqlc.arg('verified')
    );

-- name: UpdateMfaEmailCode :exec
UPDATE
    mfa_email_code
SET
    mfa_method_id = sqlc.arg('mfa_method_id'),
    token = sqlc.arg('token'),
    created_at = sqlc.arg('created_at'),
    expires_at = sqlc.arg('expires_at'),
    verified = sqlc.arg('verified')
WHERE
    id = sqlc.arg('id');

-- name: DeleteMfaEmailCode :exec
DELETE FROM
    mfa_email_code
WHERE
    id = sqlc.arg('id');

-- name: RevokeMfaEmailCodeByMFaMethodId :exec
DELETE FROM
    mfa_email_code
WHERE
    mfa_method_id = sqlc.arg('mfa_method_id');

------------------------------------QUERIES--------------------------------------
-- name: GetMfaEmailCodeByToken :one
SELECT
    id,
    mfa_method_id,
    token,
    created_at,
    expires_at,
    verified
FROM
    mfa_email_code
WHERE
    mfa_method_id = sqlc.arg('mfa_method_id')
    AND token = sqlc.arg('token');