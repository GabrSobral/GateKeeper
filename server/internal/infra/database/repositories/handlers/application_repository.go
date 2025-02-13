package repository_handlers

import (
	"context"
	"strings"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationRepository struct {
	Store *pgstore.Queries
}

func (r ApplicationRepository) CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error) {
	isApplicationExists, err := r.Store.CheckIfApplicationExists(ctx, applicationID)

	if err != nil {
		return false, err
	}

	return isApplicationExists, nil
}

func (r ApplicationRepository) AddApplication(ctx context.Context, newApplication *entities.Application) error {
	err := r.Store.AddApplication(ctx, pgstore.AddApplicationParams{
		ID:             newApplication.ID,
		Name:           newApplication.Name,
		Description:    newApplication.Description,
		OrganizationID: newApplication.OrganizationID,
		CreatedAt:      pgtype.Timestamp{Time: newApplication.CreatedAt, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (r ApplicationRepository) GetApplicationByID(ctx context.Context, applicationID uuid.UUID) (*entities.Application, error) {
	application, err := r.Store.GetApplicationByID(ctx, applicationID)

	if err != nil {
		return nil, err
	}

	return &entities.Application{
		ID:                 application.ID,
		Name:               application.Name,
		Description:        application.Description,
		OrganizationID:     application.OrganizationID,
		CreatedAt:          application.CreatedAt.Time,
		IsActive:           application.IsActive,
		HasMfaAuthApp:      application.HasMfaAuthApp,
		HasMfaEmail:        application.HasMfaEmail,
		PasswordHashSecret: application.PasswordHashSecret,
		UpdatedAt:          application.UpdatedAt,
		Badges:             strings.Split(*application.Badges, ","),
	}, nil
}

func (r ApplicationRepository) RemoveApplication(ctx context.Context, applicationID uuid.UUID) error {
	err := r.Store.DeleteApplication(ctx, applicationID)

	if err != nil {
		return err
	}

	return nil
}

func (r ApplicationRepository) UpdateApplication(ctx context.Context, newApplication *entities.Application) error {
	err := r.Store.UpdateApplication(ctx, pgstore.UpdateApplicationParams{
		ID:             newApplication.ID,
		Name:           newApplication.Name,
		OrganizationID: newApplication.OrganizationID,
		Description:    newApplication.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r ApplicationRepository) ListApplicationsFromOrganization(ctx context.Context, organizationID uuid.UUID) (*[]entities.Application, error) {
	applications, err := r.Store.ListApplicationsFromOrganization(ctx, organizationID)

	if err != nil {
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
		})
	}

	return &applicationList, nil
}
