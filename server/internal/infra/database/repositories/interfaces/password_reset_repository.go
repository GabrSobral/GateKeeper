package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IPasswordResetRepository interface {
	GetByTokenID(ctx context.Context, tokenID uuid.UUID) (*entities.PasswordResetToken, error)
	CreatePasswordReset(ctx context.Context, passwordResetToken *entities.PasswordResetToken) error
	DeletePasswordResetFromUser(ctx context.Context, userID uuid.UUID) error
}
