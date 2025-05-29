package listroles

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type IRepository interface {
	ListRolesFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.ApplicationRole, error)
	CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) ListRolesFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.ApplicationRole, error) {
	roles, err := r.Store.ListRolesFromApplication(ctx, applicationID)

	if err != nil && err != repositories.ErrNoRows {
		return nil, err
	}

	var applicationRoles []entities.ApplicationRole

	for _, role := range roles {
		applicationRoles = append(applicationRoles, entities.ApplicationRole{
			ID:            role.ID,
			ApplicationID: role.ApplicationID,
			Name:          role.Name,
			Description:   role.Description,
			CreatedAt:     role.CreatedAt.Time,
			UpdatedAt:     role.UpdatedAt,
		})
	}

	return &applicationRoles, nil
}

func (r Repository) CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error) {
	isApplicationExists, err := r.Store.CheckIfApplicationExists(ctx, applicationID)

	if err != nil {
		return false, err
	}

	return isApplicationExists, nil
}
