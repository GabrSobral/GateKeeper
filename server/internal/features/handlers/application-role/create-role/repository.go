package createrole

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	AddRole(ctx context.Context, role *entities.ApplicationRole) error
	CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) AddRole(ctx context.Context, newRole *entities.ApplicationRole) error {
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

func (r Repository) CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error) {
	isApplicationExists, err := r.Store.CheckIfApplicationExists(ctx, applicationID)

	if err != nil {
		return false, err
	}

	return isApplicationExists, nil
}
