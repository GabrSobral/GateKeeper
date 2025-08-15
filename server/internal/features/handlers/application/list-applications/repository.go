package listapplications

import (
	"context"
	"strings"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	GetOrganizationByID(ctx context.Context, organizationID uuid.UUID) (*entities.Organization, error)
	ListApplicationsFromOrganization(ctx context.Context, organizationID uuid.UUID) (*[]entities.Application, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) GetOrganizationByID(ctx context.Context, organizationID uuid.UUID) (*entities.Organization, error) {
	organization, err := r.Store.GetOrganizationByID(ctx, organizationID)

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.Organization{
		ID:          organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		CreatedAt:   organization.CreatedAt.Time,
		UpdatedAt:   organization.UpdatedAt,
	}, nil
}

func (r Repository) ListApplicationsFromOrganization(ctx context.Context, organizationID uuid.UUID) (*[]entities.Application, error) {
	applications, err := r.Store.ListApplicationsFromOrganization(ctx, organizationID)

	if err != nil && err != repositories.ErrNoRows {
		return nil, err
	}

	applicationList := make([]entities.Application, 0)

	for _, application := range applications {
		if application.Badges == nil {
			application.Badges = new(string)
		}

		applicationList = append(applicationList, entities.Application{
			ID:             application.ID,
			Name:           application.Name,
			Description:    application.Description,
			OrganizationID: application.OrganizationID,
			CreatedAt:      application.CreatedAt.Time,
			Badges:         strings.Split(*application.Badges, ","),
			UpdatedAt:      application.UpdatedAt,
		})
	}

	return &applicationList, nil
}
