package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IOrganizationRepository interface {
	AddOrganization(ctx context.Context, organization *entities.Organization) error
	RemoveOrganization(ctx context.Context, organizationID uuid.UUID) error
	UpdateOrganization(ctx context.Context, organization *entities.Organization) error
	GetOrganizationByID(ctx context.Context, organizationID uuid.UUID) (*entities.Organization, error)
	ListOrganizations(ctx context.Context) (*[]entities.Organization, error)
}
