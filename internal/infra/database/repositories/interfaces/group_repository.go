package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IGroupRepository interface {
	AddGroup(ctx context.Context, group *entities.Group) error
	RemoveGroup(ctx context.Context, groupID uuid.UUID) error
	UpdateGroup(ctx context.Context, group *entities.Group) error
	GetGroupByID(ctx context.Context, groupID uuid.UUID) (*entities.Group, error)
	ListGroupsFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.Group, error)
}
