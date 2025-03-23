package repository_interfaces

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/google/uuid"
)

type ISessionCodeRepository interface {
	Add(ctx context.Context, emailMfaCode *entities.SessionCode) error
	GetByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.SessionCode, error)
	Update(ctx context.Context, sessionCode *entities.SessionCode) error
	DeleteByID(ctx context.Context, id uuid.UUID) error
	RevokeByUserID(ctx context.Context, userID uuid.UUID) error
}
