// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: application_authorization_code.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addAuthorizationCode = `-- name: AddAuthorizationCode :exec
INSERT INTO
    application_authorization_code (
        id,
        application_id,
        user_id,
        expired_at,
        code,
        redirect_uri,
        code_challenge,
        code_challenge_method
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
        $8
    )
`

type AddAuthorizationCodeParams struct {
	ID                  uuid.UUID        `db:"id"`
	ApplicationID       uuid.UUID        `db:"application_id"`
	UserID              uuid.UUID        `db:"user_id"`
	ExpiredAt           pgtype.Timestamp `db:"expired_at"`
	Code                string           `db:"code"`
	RedirectUri         string           `db:"redirect_uri"`
	CodeChallenge       string           `db:"code_challenge"`
	CodeChallengeMethod string           `db:"code_challenge_method"`
}

// ----------------------------------COMMANDS--------------------------------------
// Add Authorization Code to Application
func (q *Queries) AddAuthorizationCode(ctx context.Context, arg AddAuthorizationCodeParams) error {
	_, err := q.db.Exec(ctx, addAuthorizationCode,
		arg.ID,
		arg.ApplicationID,
		arg.UserID,
		arg.ExpiredAt,
		arg.Code,
		arg.RedirectUri,
		arg.CodeChallenge,
		arg.CodeChallengeMethod,
	)
	return err
}

const getAuthorizationCodeById = `-- name: GetAuthorizationCodeById :one
SELECT
    id,
    application_id,
    user_id,
    expired_at,
    code,
    redirect_uri,
    code_challenge,
    code_challenge_method
FROM
    application_authorization_code
WHERE
    id = $1
`

// ----------------------------------QUERIES---------------------------------------
// List Authorization Codes by Application Id
func (q *Queries) GetAuthorizationCodeById(ctx context.Context, id uuid.UUID) (ApplicationAuthorizationCode, error) {
	row := q.db.QueryRow(ctx, getAuthorizationCodeById, id)
	var i ApplicationAuthorizationCode
	err := row.Scan(
		&i.ID,
		&i.ApplicationID,
		&i.UserID,
		&i.ExpiredAt,
		&i.Code,
		&i.RedirectUri,
		&i.CodeChallenge,
		&i.CodeChallengeMethod,
	)
	return i, err
}

const removeAuthorizationCode = `-- name: RemoveAuthorizationCode :exec
DELETE FROM
    application_authorization_code
WHERE
    application_id = $1
    AND user_id = $2
`

type RemoveAuthorizationCodeParams struct {
	ApplicationID uuid.UUID `db:"application_id"`
	UserID        uuid.UUID `db:"user_id"`
}

// Remove Authorization Code from Application
func (q *Queries) RemoveAuthorizationCode(ctx context.Context, arg RemoveAuthorizationCodeParams) error {
	_, err := q.db.Exec(ctx, removeAuthorizationCode, arg.ApplicationID, arg.UserID)
	return err
}
