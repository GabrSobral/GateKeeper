package signin

import (
	"context"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
)

func assignRefreshToken(ctx context.Context, service *SignInService, user entities.ApplicationUser) (*entities.RefreshToken, error) {
	currentDate := time.Now().UTC()
	futureDate := currentDate.Add(time.Hour * 24 * 7).UTC() // 7 days from now

	service.RefreshTokenRepository.RevokeRefreshTokenFromUser(ctx, user.ID)
	refreshToken, err := entities.CreateRefreshToken(user.ID, 5, futureDate)

	if err != nil {
		return nil, err
	}

	if _, err := service.RefreshTokenRepository.AddRefreshToken(ctx, refreshToken); err != nil {
		return nil, err
	}

	return refreshToken, nil
}
