package repository_interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
)

type IPasswordResetRepository interface {
	GetByTokenID(ctx context.Context, tokenID uuid.UUID) (*entities.PasswordResetToken, error)
	CreatePasswordReset(ctx context.Context, passwordResetToken *entities.PasswordResetToken) error
	DeletePasswordResetFromUser(ctx context.Context, userID uuid.UUID) error
}
