package editorganization

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	UpdateOrganization(ctx context.Context, application *entities.Organization) error
	GetOrganizationById(ctx context.Context, id uuid.UUID) (*entities.Organization, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) GetOrganizationById(ctx context.Context, id uuid.UUID) (*entities.Organization, error) {
	org, err := r.Store.GetOrganizationByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return &entities.Organization{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		CreatedAt:   org.CreatedAt.Time,
		UpdatedAt:   org.UpdatedAt,
	}, nil
}

func (r Repository) UpdateOrganization(ctx context.Context, organization *entities.Organization) error {
	err := r.Store.UpdateOrganization(ctx, pgstore.UpdateOrganizationParams{
		ID:          organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		UpdatedAt:   organization.UpdatedAt,
	})

	if err != nil {
		return err
	}

	return nil
}
