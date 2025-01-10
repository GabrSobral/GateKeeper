package repository_interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
)

type IEmailConfirmationRepository interface {
	AddEmailConfirmation(ctx context.Context, emailConfirmation *entities.EmailConfirmation) error
	GetByEmail(ctx context.Context, email string, userID uuid.UUID) (*entities.EmailConfirmation, error)
	UpdateEmailConfirmation(ctx context.Context, emailConfirmation *entities.EmailConfirmation) error
	DeleteEmailConfirmation(ctx context.Context, emailConfirmationID uuid.UUID) error
}
