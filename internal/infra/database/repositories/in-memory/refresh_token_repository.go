package inmemory_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
)

type InMemoryRefreshTokenRepository struct {
	RefreshTokens map[string]*entities.RefreshToken
}

func (r InMemoryRefreshTokenRepository) AddRefreshToken(ctx context.Context, newRefreshToken *entities.RefreshToken) (*entities.RefreshToken, error) {
	r.RefreshTokens[newRefreshToken.ID.String()] = newRefreshToken

	return newRefreshToken, nil
}

func (r InMemoryRefreshTokenRepository) RevokeRefreshTokenFromUser(ctx context.Context, userID uuid.UUID) error {
	for _, refreshToken := range r.RefreshTokens {
		if refreshToken.UserID == userID {
			delete(r.RefreshTokens, refreshToken.ID.String())
		}
	}

	return nil
}

func (r InMemoryRefreshTokenRepository) GetRefreshTokensFromUser(ctx context.Context, userID uuid.UUID) (*[]entities.RefreshToken, error) {
	var refreshTokens []entities.RefreshToken

	for _, refreshToken := range r.RefreshTokens {
		if refreshToken.UserID == userID {
			refreshTokens = append(refreshTokens, *refreshToken)
		}
	}

	return &refreshTokens, nil
}
