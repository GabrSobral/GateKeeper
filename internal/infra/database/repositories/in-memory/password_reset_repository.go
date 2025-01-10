package inmemory_repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/guard-service/internal/domain/entities"
)

type InMemoryPasswordResetRepository struct {
	PasswordTokens map[string]*entities.PasswordResetToken
}

func (r InMemoryPasswordResetRepository) GetByTokenID(ctx context.Context, tokenID uuid.UUID) (*entities.PasswordResetToken, error) {
	for _, login := range r.PasswordTokens {
		if login.ID == tokenID {
			return login, nil
		}
	}

	return nil, nil
}

func (r InMemoryPasswordResetRepository) CreatePasswordReset(ctx context.Context, passwordResetToken *entities.PasswordResetToken) error {
	r.PasswordTokens[passwordResetToken.ID.String()] = passwordResetToken

	return nil
}

func (r InMemoryPasswordResetRepository) DeletePasswordResetFromUser(ctx context.Context, userID uuid.UUID) error {
	for _, login := range r.PasswordTokens {
		if login.UserID == userID {
			delete(r.PasswordTokens, login.ID.String())
		}
	}

	return nil
}
