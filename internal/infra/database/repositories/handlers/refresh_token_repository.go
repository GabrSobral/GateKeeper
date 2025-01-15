package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type RefreshTokenRepository struct {
	Store *pgstore.Queries
}

func (r RefreshTokenRepository) AddRefreshToken(ctx context.Context, refreshToken *entities.RefreshToken) (*entities.RefreshToken, error) {
	err := r.Store.AddRefreshToken(ctx, pgstore.AddRefreshTokenParams{
		UserID:             refreshToken.UserID,
		ID:                 refreshToken.ID,
		AvailableRefreshes: int32(refreshToken.AvailableRefreshes),
		ExpiresAt:          pgtype.Timestamp{Time: refreshToken.ExpiresAt, Valid: true},
		CreatedAt:          pgtype.Timestamp{Time: refreshToken.CreatedAt, Valid: true},
	})

	if err != nil {
		return nil, err
	}

	return &entities.RefreshToken{
		ID:                 refreshToken.ID,
		UserID:             refreshToken.UserID,
		AvailableRefreshes: refreshToken.AvailableRefreshes,
		ExpiresAt:          refreshToken.ExpiresAt,
		CreatedAt:          refreshToken.CreatedAt,
	}, nil
}

func (r RefreshTokenRepository) GetRefreshTokensFromUser(ctx context.Context, userID uuid.UUID) (*[]entities.RefreshToken, error) {
	refreshTokens, err := r.Store.GetRefreshTokensFromUser(ctx, userID)

	if err != nil {
		return nil, err
	}

	var tokens []entities.RefreshToken

	for _, token := range refreshTokens {
		tokens = append(tokens, entities.RefreshToken{
			ID:                 token.ID,
			UserID:             token.UserID,
			AvailableRefreshes: uint8(token.AvailableRefreshes),
			ExpiresAt:          token.ExpiresAt.Time,
			CreatedAt:          token.CreatedAt.Time,
		})
	}

	return &tokens, nil
}

func (r RefreshTokenRepository) RevokeRefreshTokenFromUser(ctx context.Context, userID uuid.UUID) error {
	err := r.Store.RevokeRefreshTokenFromUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil
}
