// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: application_secret.sql

package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addSecret = `-- name: AddSecret :exec
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
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7
    )
`

type AddSecretParams struct {
	ID            uuid.UUID        `db:"id"`
	ApplicationID uuid.UUID        `db:"application_id"`
	Name          string           `db:"name"`
	Value         string           `db:"value"`
	CreatedAt     pgtype.Timestamp `db:"created_at"`
	UpdatedAt     *time.Time       `db:"updated_at"`
	ExpiresAt     *time.Time       `db:"expires_at"`
}

// ----------------------------------COMMANDS--------------------------------------
// Add Secret to Application
func (q *Queries) AddSecret(ctx context.Context, arg AddSecretParams) error {
	_, err := q.db.Exec(ctx, addSecret,
		arg.ID,
		arg.ApplicationID,
		arg.Name,
		arg.Value,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ExpiresAt,
	)
	return err
}

const listSecretsFromApplication = `-- name: ListSecretsFromApplication :many
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
    application_id = $1
`

type ListSecretsFromApplicationRow struct {
	ID            uuid.UUID        `db:"id"`
	ApplicationID uuid.UUID        `db:"application_id"`
	Name          string           `db:"name"`
	Value         string           `db:"value"`
	CreatedAt     pgtype.Timestamp `db:"created_at"`
	UpdatedAt     *time.Time       `db:"updated_at"`
	ExpiresAt     *time.Time       `db:"expires_at"`
}

// ----------------------------------QUERIES---------------------------------------
// List Secrets from Application
func (q *Queries) ListSecretsFromApplication(ctx context.Context, applicationID uuid.UUID) ([]ListSecretsFromApplicationRow, error) {
	rows, err := q.db.Query(ctx, listSecretsFromApplication, applicationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSecretsFromApplicationRow
	for rows.Next() {
		var i ListSecretsFromApplicationRow
		if err := rows.Scan(
			&i.ID,
			&i.ApplicationID,
			&i.Name,
			&i.Value,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ExpiresAt,
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

const removeSecret = `-- name: RemoveSecret :exec
DELETE FROM
    application_secret
WHERE
    id = $1
`

// Remove Secret from Application
func (q *Queries) RemoveSecret(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, removeSecret, id)
	return err
}
