package removeorganization

import (
	"context"

	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	RemoveOrganization(ctx context.Context, organizationID uuid.UUID) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) RemoveOrganization(ctx context.Context, organizationID uuid.UUID) error {
	err := r.Store.RemoveOrganization(ctx, organizationID)

	if err != nil {
		return err
	}

	return nil
}
