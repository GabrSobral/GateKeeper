package listorganizations

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
)

type IRepository interface {
	ListOrganizations(ctx context.Context) (*[]entities.Organization, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) ListOrganizations(ctx context.Context) (*[]entities.Organization, error) {
	organizations, err := r.Store.ListOrganizations(ctx)

	if err != nil && err != repositories.ErrNoRows {
		return nil, err
	}

	var organizationList []entities.Organization

	for _, organization := range organizations {
		organizationList = append(organizationList, entities.Organization{
			ID:          organization.ID,
			Name:        organization.Name,
			CreatedAt:   organization.CreatedAt.Time,
			UpdatedAt:   organization.UpdatedAt,
			Description: organization.Description,
		})
	}

	return &organizationList, nil
}
