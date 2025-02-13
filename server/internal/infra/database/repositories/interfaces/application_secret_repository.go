package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IApplicationSecretRepository interface {
	AddSecret(ctx context.Context, newSecret *entities.ApplicationSecret) error
	RemoveSecret(ctx context.Context, roleID uuid.UUID) error
	ListSecretsFromApplication(ctx context.Context, applicationID uuid.UUID) (*[]entities.ApplicationSecret, error)
}
