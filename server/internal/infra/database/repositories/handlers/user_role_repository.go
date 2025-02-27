package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRoleRepository struct {
	Store *pgstore.Queries
}

func (r UserRoleRepository) AddUserRole(ctx context.Context, newUserRole *entities.UserRole) error {
	err := r.Store.AddUserRole(ctx, pgstore.AddUserRoleParams{
		UserID:    newUserRole.UserID,
		RoleID:    newUserRole.RoleID,
		CreatedAt: pgtype.Timestamp{Time: newUserRole.CreatedAt, Valid: true},
	})

	return err
}

func (r UserRoleRepository) RemoveUserRole(ctx context.Context, userRole *entities.UserRole) error {
	err := r.Store.RemoveUserRole(ctx, pgstore.RemoveUserRoleParams{
		UserID: userRole.UserID,
		RoleID: userRole.RoleID,
	})

	return err
}
