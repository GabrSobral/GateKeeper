package createsecret

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error)
	AddSecret(ctx context.Context, secret *entities.ApplicationSecret) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error) {
	isApplicationExists, err := r.Store.CheckIfApplicationExists(ctx, applicationID)

	if err != nil {
		return false, err
	}

	return isApplicationExists, nil
}

func (r Repository) AddSecret(ctx context.Context, newSecret *entities.ApplicationSecret) error {
	err := r.Store.AddSecret(ctx, pgstore.AddSecretParams{
		ID:            newSecret.ID,
		ApplicationID: newSecret.ApplicationID,
		Name:          newSecret.Name,
		Value:         newSecret.Value,
		CreatedAt:     pgtype.Timestamp{Time: newSecret.CreatedAt, Valid: true},
		UpdatedAt:     newSecret.UpdatedAt,
		ExpiresAt:     newSecret.ExpiresAt,
	})

	return err
}
