// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: application_user.sql

package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addUser = `-- name: AddUser :exec
INSERT INTO
    "application_user" (
        id,
        email,
        password_hash,
        application_id,
        created_at,
        updated_at,
        is_active,
        is_email_confirmed,
        is_mfa_auth_app_enabled,
        is_mfa_email_enabled,
        two_factor_secret,
        should_change_pass
    )
VALUES
    (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11,
        $12
    )
`

type AddUserParams struct {
	ID                  uuid.UUID        `db:"id"`
	Email               string           `db:"email"`
	PasswordHash        *string          `db:"password_hash"`
	ApplicationID       uuid.UUID        `db:"application_id"`
	CreatedAt           pgtype.Timestamp `db:"created_at"`
	UpdatedAt           *time.Time       `db:"updated_at"`
	IsActive            bool             `db:"is_active"`
	IsEmailConfirmed    bool             `db:"is_email_confirmed"`
	IsMfaAuthAppEnabled bool             `db:"is_mfa_auth_app_enabled"`
	IsMfaEmailEnabled   bool             `db:"is_mfa_email_enabled"`
	TwoFactorSecret     *string          `db:"two_factor_secret"`
	ShouldChangePass    bool             `db:"should_change_pass"`
}

// ----------------------------------COMMANDS--------------------------------------
// Add user
func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) error {
	_, err := q.db.Exec(ctx, addUser,
		arg.ID,
		arg.Email,
		arg.PasswordHash,
		arg.ApplicationID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.IsActive,
		arg.IsEmailConfirmed,
		arg.IsMfaAuthAppEnabled,
		arg.IsMfaEmailEnabled,
		arg.TwoFactorSecret,
		arg.ShouldChangePass,
	)
	return err
}

const deleteApplicationUser = `-- name: DeleteApplicationUser :exec
DELETE FROM
    "application_user"
WHERE
    id = $1
    AND application_id = $2
`

type DeleteApplicationUserParams struct {
	ID            uuid.UUID `db:"id"`
	ApplicationID uuid.UUID `db:"application_id"`
}

func (q *Queries) DeleteApplicationUser(ctx context.Context, arg DeleteApplicationUserParams) error {
	_, err := q.db.Exec(ctx, deleteApplicationUser, arg.ID, arg.ApplicationID)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT
    id,
    email,
    application_id,
    password_hash,
    created_at,
    updated_at,
    is_active,
    is_email_confirmed,
    is_mfa_auth_app_enabled,
    is_mfa_email_enabled,
    should_change_pass,
    two_factor_secret
FROM
    "application_user"
WHERE
    email = $1
    AND application_id = $2
`

type GetUserByEmailParams struct {
	Email         string    `db:"email"`
	ApplicationID uuid.UUID `db:"application_id"`
}

type GetUserByEmailRow struct {
	ID                  uuid.UUID        `db:"id"`
	Email               string           `db:"email"`
	ApplicationID       uuid.UUID        `db:"application_id"`
	PasswordHash        *string          `db:"password_hash"`
	CreatedAt           pgtype.Timestamp `db:"created_at"`
	UpdatedAt           *time.Time       `db:"updated_at"`
	IsActive            bool             `db:"is_active"`
	IsEmailConfirmed    bool             `db:"is_email_confirmed"`
	IsMfaAuthAppEnabled bool             `db:"is_mfa_auth_app_enabled"`
	IsMfaEmailEnabled   bool             `db:"is_mfa_email_enabled"`
	ShouldChangePass    bool             `db:"should_change_pass"`
	TwoFactorSecret     *string          `db:"two_factor_secret"`
}

// Get user by email
func (q *Queries) GetUserByEmail(ctx context.Context, arg GetUserByEmailParams) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, arg.Email, arg.ApplicationID)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.ApplicationID,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsActive,
		&i.IsEmailConfirmed,
		&i.IsMfaAuthAppEnabled,
		&i.IsMfaEmailEnabled,
		&i.ShouldChangePass,
		&i.TwoFactorSecret,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT
    id,
    email,
    application_id,
    password_hash,
    created_at,
    updated_at,
    is_active,
    is_email_confirmed,
    is_mfa_auth_app_enabled,
    is_mfa_email_enabled,
    should_change_pass,
    two_factor_secret
FROM
    "application_user"
WHERE
    id = $1
`

type GetUserByIdRow struct {
	ID                  uuid.UUID        `db:"id"`
	Email               string           `db:"email"`
	ApplicationID       uuid.UUID        `db:"application_id"`
	PasswordHash        *string          `db:"password_hash"`
	CreatedAt           pgtype.Timestamp `db:"created_at"`
	UpdatedAt           *time.Time       `db:"updated_at"`
	IsActive            bool             `db:"is_active"`
	IsEmailConfirmed    bool             `db:"is_email_confirmed"`
	IsMfaAuthAppEnabled bool             `db:"is_mfa_auth_app_enabled"`
	IsMfaEmailEnabled   bool             `db:"is_mfa_email_enabled"`
	ShouldChangePass    bool             `db:"should_change_pass"`
	TwoFactorSecret     *string          `db:"two_factor_secret"`
}

// ----------------------------------QUERIES--------------------------------------
// Get user by id
func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (GetUserByIdRow, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.ApplicationID,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsActive,
		&i.IsEmailConfirmed,
		&i.IsMfaAuthAppEnabled,
		&i.IsMfaEmailEnabled,
		&i.ShouldChangePass,
		&i.TwoFactorSecret,
	)
	return i, err
}

const getUsersByApplicationID = `-- name: GetUsersByApplicationID :many
SELECT
    au.id,
    au.email,
    au.application_id,
    up.display_name,
    au.created_at,
    au.updated_at,
    au.is_active,
    au.is_email_confirmed,
    au.is_mfa_auth_app_enabled,
    au.is_mfa_email_enabled,
    COALESCE(r.roles, '[]' :: jsonb) AS roles,
    COUNT(*) OVER () AS total_users
FROM
    "application_user" au
    LEFT JOIN "user_profile" up ON up.user_id = au.id
    LEFT JOIN LATERAL (
        SELECT
            jsonb_agg(
                jsonb_build_object(
                    'id',
                    ar.id,
                    'name',
                    ar.name,
                    'description',
                    ar.description
                )
            ) AS roles
        FROM
            "user_role" ur
            JOIN "application_role" ar ON ar.id = ur.role_id
        WHERE
            ur.user_id = au.id
    ) r ON TRUE
WHERE
    au.application_id = $1
ORDER BY
    au.created_at
LIMIT
    $3 OFFSET $2
`

type GetUsersByApplicationIDParams struct {
	ApplicationID uuid.UUID `db:"application_id"`
	Offset        int32     `db:"offset"`
	Limit         int32     `db:"limit"`
}

type GetUsersByApplicationIDRow struct {
	ID                  uuid.UUID        `db:"id"`
	Email               string           `db:"email"`
	ApplicationID       uuid.UUID        `db:"application_id"`
	DisplayName         *string          `db:"display_name"`
	CreatedAt           pgtype.Timestamp `db:"created_at"`
	UpdatedAt           *time.Time       `db:"updated_at"`
	IsActive            bool             `db:"is_active"`
	IsEmailConfirmed    bool             `db:"is_email_confirmed"`
	IsMfaAuthAppEnabled bool             `db:"is_mfa_auth_app_enabled"`
	IsMfaEmailEnabled   bool             `db:"is_mfa_email_enabled"`
	Roles               []byte           `db:"roles"`
	TotalUsers          int64            `db:"total_users"`
}

// Get users by application id paged, and ordered by created_at, that includes the application roles
func (q *Queries) GetUsersByApplicationID(ctx context.Context, arg GetUsersByApplicationIDParams) ([]GetUsersByApplicationIDRow, error) {
	rows, err := q.db.Query(ctx, getUsersByApplicationID, arg.ApplicationID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersByApplicationIDRow
	for rows.Next() {
		var i GetUsersByApplicationIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.ApplicationID,
			&i.DisplayName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsActive,
			&i.IsEmailConfirmed,
			&i.IsMfaAuthAppEnabled,
			&i.IsMfaEmailEnabled,
			&i.Roles,
			&i.TotalUsers,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const isUserExistsByEmail = `-- name: IsUserExistsByEmail :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "application_user"
        WHERE
            email = $1
            AND application_id = $2
    )
`

type IsUserExistsByEmailParams struct {
	Email         string    `db:"email"`
	ApplicationID uuid.UUID `db:"application_id"`
}

// Check if user exists by email
func (q *Queries) IsUserExistsByEmail(ctx context.Context, arg IsUserExistsByEmailParams) (bool, error) {
	row := q.db.QueryRow(ctx, isUserExistsByEmail, arg.Email, arg.ApplicationID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isUserExistsById = `-- name: IsUserExistsById :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            "application_user"
        WHERE
            id = $1
    )
`

// Check if user exists by id
func (q *Queries) IsUserExistsById(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, isUserExistsById, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE
    "application_user"
SET
    email = $1,
    password_hash = $2,
    updated_at = $3,
    is_active = $4,
    is_email_confirmed = $5,
    is_mfa_auth_app_enabled = $6,
    is_mfa_email_enabled = $7,
    two_factor_secret = $8,
    should_change_pass = $9
WHERE
    id = $10
`

type UpdateUserParams struct {
	Email               string     `db:"email"`
	PasswordHash        *string    `db:"password_hash"`
	UpdatedAt           *time.Time `db:"updated_at"`
	IsActive            bool       `db:"is_active"`
	IsEmailConfirmed    bool       `db:"is_email_confirmed"`
	IsMfaAuthAppEnabled bool       `db:"is_mfa_auth_app_enabled"`
	IsMfaEmailEnabled   bool       `db:"is_mfa_email_enabled"`
	TwoFactorSecret     *string    `db:"two_factor_secret"`
	ShouldChangePass    bool       `db:"should_change_pass"`
	ID                  uuid.UUID  `db:"id"`
}

// Update user
func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Email,
		arg.PasswordHash,
		arg.UpdatedAt,
		arg.IsActive,
		arg.IsEmailConfirmed,
		arg.IsMfaAuthAppEnabled,
		arg.IsMfaEmailEnabled,
		arg.TwoFactorSecret,
		arg.ShouldChangePass,
		arg.ID,
	)
	return err
}
