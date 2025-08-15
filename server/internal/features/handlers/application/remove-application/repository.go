package removeapplication

import (
	"context"

	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	RemoveApplication(ctx context.Context, applicationID uuid.UUID) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) RemoveApplication(ctx context.Context, applicationID uuid.UUID) error {
	err := r.Store.DeleteApplication(ctx, applicationID)

	if err != nil {
		return err
	}

	return nil
}
