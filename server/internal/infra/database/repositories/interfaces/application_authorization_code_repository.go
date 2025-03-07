package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IApplicationAuthorizationCodeRepository interface {
	AddAuthorizationCode(ctx context.Context, newRole *entities.ApplicationAuthorizationCode) error
	RemoveAuthorizationCode(ctx context.Context, userID, applicationId uuid.UUID) error
	GetAuthorizationCodeById(ctx context.Context, code uuid.UUID) (*entities.ApplicationAuthorizationCode, error)
}
