package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IUserRoleRepository interface {
	AddUserRole(ctx context.Context, newUserRole *entities.UserRole) error
	RemoveUserRole(ctx context.Context, userRole *entities.UserRole) error
	GetRolesByUserID(ctx context.Context, userID uuid.UUID) ([]entities.ApplicationRole, error)
}
