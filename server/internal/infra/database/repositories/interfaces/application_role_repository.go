package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IApplicationRoleRepository interface {
	AddRole(ctx context.Context, newRole *entities.ApplicationRole) error
	RemoveRole(ctx context.Context, roleID uuid.UUID) error
	ListRolesFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.ApplicationRole, error)
}
