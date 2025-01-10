package repository_interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
)

type IRefreshTokenRepository interface {
	AddRefreshToken(ctx context.Context, refreshToken *entities.RefreshToken) (*entities.RefreshToken, error)
	GetRefreshTokensFromUser(ctx context.Context, userID uuid.UUID) (*[]entities.RefreshToken, error)
	RevokeRefreshTokenFromUser(ctx context.Context, userID uuid.UUID) error
}
