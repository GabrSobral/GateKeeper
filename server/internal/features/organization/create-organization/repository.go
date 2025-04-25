package createorganization

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	AddOrganization(ctx context.Context, application *entities.Organization) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) AddOrganization(ctx context.Context, organization *entities.Organization) error {
	err := r.Store.AddOrganization(ctx, pgstore.AddOrganizationParams{
		UserID:      organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		CreatedAt:   pgtype.Timestamp{Time: organization.CreatedAt, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}
