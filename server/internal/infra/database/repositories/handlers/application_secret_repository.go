package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationSecretRepository struct {
	Store *pgstore.Queries
}

func (r ApplicationSecretRepository) AddSecret(ctx context.Context, newSecret *entities.ApplicationSecret) error {
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

func (r ApplicationSecretRepository) RemoveSecret(ctx context.Context, secretID uuid.UUID) error {
	err := r.Store.RemoveSecret(ctx, secretID)

	return err
}

func (r ApplicationSecretRepository) ListSecretsFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.ApplicationSecret, error) {
	secrets, err := r.Store.ListSecretsFromApplication(ctx, applicationID)

	if err != nil && err != repositories.ErrNoRows {
		return nil, err
	}

	var applicationSecrets []entities.ApplicationSecret

	for _, secret := range secrets {
		applicationSecrets = append(applicationSecrets, entities.ApplicationSecret{
			ID:            secret.ID,
			ApplicationID: secret.ApplicationID,
			Name:          secret.Name,
			Value:         secret.Value,
			CreatedAt:     secret.CreatedAt.Time,
			UpdatedAt:     secret.UpdatedAt,
			ExpiresAt:     secret.ExpiresAt,
		})
	}

	return &applicationSecrets, nil
}
