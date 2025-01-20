package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationRepository struct {
	Store *pgstore.Queries
}

func (r ApplicationRepository) AddApplication(ctx context.Context, newApplication *entities.Application) error {
	err := r.Store.AddApplication(ctx, pgstore.AddApplicationParams{
		ID:          newApplication.ID,
		Name:        newApplication.Name,
		Description: newApplication.Description,
		TenantID:    newApplication.TenantID,
		CreatedAt:   pgtype.Timestamp{Time: newApplication.CreatedAt, Valid: true},
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
		ID:          application.ID,
		Name:        application.Name,
		Description: application.Description,
		TenantID:    application.TenantID,
		CreatedAt:   application.CreatedAt.Time,
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
		ID:          newApplication.ID,
		Name:        newApplication.Name,
		TenantID:    newApplication.TenantID,
		Description: newApplication.Description,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r ApplicationRepository) ListApplicationsFromTenant(ctx context.Context, tenantID uuid.UUID) (*[]entities.Application, error) {
	applications, err := r.Store.ListApplicationsFromTenant(ctx, tenantID)

	if err != nil {
		return nil, err
	}

	applicationList := make([]entities.Application, 0)

	for _, application := range applications {
		applicationList = append(applicationList, entities.Application{
			ID:          application.ID,
			Name:        application.Name,
			Description: application.Description,
			TenantID:    application.TenantID,
			CreatedAt:   application.CreatedAt.Time,
		})
	}

	return &applicationList, nil
}
