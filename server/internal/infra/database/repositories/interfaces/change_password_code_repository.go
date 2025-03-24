package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type IChangePasswordCodeRepository interface {
	Add(ctx context.Context, changePasswordCode *entities.ChangePasswordCode) error
	GetByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.ChangePasswordCode, error)
	Update(ctx context.Context, changePasswordCode *entities.ChangePasswordCode) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
	RevokeAllByID(ctx context.Context, userID uuid.UUID) error
}
