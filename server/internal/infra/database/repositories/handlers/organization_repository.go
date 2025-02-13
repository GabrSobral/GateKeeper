package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type OrganizationRepository struct {
	Store *pgstore.Queries
}

func (r OrganizationRepository) GetOrganizationByID(ctx context.Context, organizationID uuid.UUID) (*entities.Organization, error) {
	organization, err := r.Store.GetOrganizationByID(ctx, organizationID)

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

func (r OrganizationRepository) AddOrganization(ctx context.Context, organization *entities.Organization) error {
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

func (r OrganizationRepository) RemoveOrganization(ctx context.Context, organizationID uuid.UUID) error {
	err := r.Store.RemoveOrganization(ctx, organizationID)
	if err != nil {
		return err
	}

	return nil
}

func (r OrganizationRepository) UpdateOrganization(ctx context.Context, organization *entities.Organization) error {
	err := r.Store.UpdateOrganization(ctx, pgstore.UpdateOrganizationParams{
		ID:          organization.ID,
		Name:        organization.Name,
		UpdatedAt:   organization.UpdatedAt,
		Description: organization.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r OrganizationRepository) ListOrganizations(ctx context.Context) (*[]entities.Organization, error) {
	organizations, err := r.Store.ListOrganizations(ctx)

	if err != nil {
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
