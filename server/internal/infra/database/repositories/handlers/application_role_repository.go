package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ApplicationRoleRepository struct {
	Store *pgstore.Queries
}

func (r ApplicationRoleRepository) AddRole(ctx context.Context, newRole *entities.ApplicationRole) error {
	err := r.Store.AddRole(ctx, pgstore.AddRoleParams{
		ID:            newRole.ID,
		ApplicationID: newRole.ApplicationID,
		Name:          newRole.Name,
		Description:   newRole.Description,
		CreatedAt:     pgtype.Timestamp{Time: newRole.CreatedAt, Valid: true},
		UpdatedAt:     newRole.UpdatedAt,
	})

	return err
}

func (r ApplicationRoleRepository) RemoveRole(ctx context.Context, roleID uuid.UUID) error {
	err := r.Store.RemoveRole(ctx, roleID)

	return err
}

func (r ApplicationRoleRepository) ListRolesFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.ApplicationRole, error) {
	roles, err := r.Store.ListRolesFromApplication(ctx, applicationID)

	if err != nil {
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
