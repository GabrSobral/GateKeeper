package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IApplicationRepository interface {
	AddApplication(ctx context.Context, newApplication *entities.Application) error
	GetApplicationByID(ctx context.Context, applicationID uuid.UUID) (*entities.Application, error)
	RemoveApplication(ctx context.Context, applicationID uuid.UUID) error
	UpdateApplication(ctx context.Context, newApplication *entities.Application) error
	ListApplicationsFromOrganization(ctx context.Context, organizationID uuid.UUID) (*[]entities.Application, error)
	CheckIfApplicationExists(ctx context.Context, applicationID uuid.UUID) (bool, error)
}
