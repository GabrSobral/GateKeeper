package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PasswordResetRepository struct {
	Store *pgstore.Queries
}

func (pr PasswordResetRepository) GetByTokenID(ctx context.Context, tokenID uuid.UUID) (*entities.PasswordResetToken, error) {
	passwordReset, err := pr.Store.GetPasswordResetByTokenID(ctx, tokenID)

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.PasswordResetToken{
		ID:        passwordReset.ID,
		UserID:    passwordReset.UserID,
		Token:     passwordReset.Token,
		CreatedAt: passwordReset.CreatedAt.Time,
		ExpiresAt: passwordReset.ExpiresAt.Time,
	}, nil
}

func (pr PasswordResetRepository) CreatePasswordReset(ctx context.Context, passwordResetToken *entities.PasswordResetToken) error {
	err := pr.Store.CreatePasswordReset(ctx, pgstore.CreatePasswordResetParams{
		ID:        passwordResetToken.ID,
		UserID:    passwordResetToken.UserID,
		Token:     passwordResetToken.Token,
		CreatedAt: pgtype.Timestamp{Time: passwordResetToken.CreatedAt, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: passwordResetToken.ExpiresAt, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (pr PasswordResetRepository) DeletePasswordResetFromUser(ctx context.Context, userID uuid.UUID) error {
	err := pr.Store.DeletePasswordResetFromUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil
}
