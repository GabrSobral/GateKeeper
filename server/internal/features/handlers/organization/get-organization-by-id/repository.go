package getorganizationbyid

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entities.Organization, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*entities.Organization, error) {
	organization, err := r.Store.GetOrganizationByID(ctx, id)

	if err != nil && err != repositories.ErrNoRows {
		return nil, err
	}

	return &entities.Organization{
		ID:          organization.ID,
		Name:        organization.Name,
		CreatedAt:   organization.CreatedAt.Time,
		UpdatedAt:   organization.UpdatedAt,
		Description: organization.Description,
	}, nil
}
