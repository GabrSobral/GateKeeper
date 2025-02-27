------------------------------------COMMANDS--------------------------------------
-- name: AddUserRole :exec
INSERT INTO
    user_role (user_id, role_id, created_at)
VALUES
    (
        sqlc.arg('user_id'),
        -- user_id
        sqlc.arg('role_id'),
        -- role_id
        sqlc.arg('created_at') -- created_at
    );

-- name: RemoveUserRole :exec
DELETE FROM
    user_role
WHERE
    user_id = sqlc.arg('user_id')
    AND role_id = sqlc.arg('role_id');

------------------------------------QUERIES--------------------------------------