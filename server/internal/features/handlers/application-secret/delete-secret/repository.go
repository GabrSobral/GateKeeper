package deletesecret

import (
	"context"

	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error)
	RemoveSecret(ctx context.Context, secretID uuid.UUID) error
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

func (r Repository) RemoveSecret(ctx context.Context, secretID uuid.UUID) error {
	err := r.Store.RemoveSecret(ctx, secretID)

	return err
}
