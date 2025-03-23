package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SessionCodeRepository struct {
	Store *pgstore.Queries
}

func (r SessionCodeRepository) Add(ctx context.Context, sessionCode *entities.SessionCode) error {
	err := r.Store.AddSessionCode(ctx, pgstore.AddSessionCodeParams{
		ID:        sessionCode.ID,
		UserID:    sessionCode.UserID,
		Token:     sessionCode.Token,
		CreatedAt: pgtype.Timestamp{Time: sessionCode.CreatedAt, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: sessionCode.ExpiresAt, Valid: true},
		IsUsed:    sessionCode.IsUsed,
	})

	return err
}

func (r SessionCodeRepository) GetByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.SessionCode, error) {

	emailConfirmation, err := r.Store.GetSessionCodeByToken(ctx, pgstore.GetSessionCodeByTokenParams{
		Token:  token,
		UserID: userID,
	})

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.SessionCode{
		ID:        emailConfirmation.ID,
		UserID:    emailConfirmation.UserID,
		Token:     emailConfirmation.Token,
		CreatedAt: emailConfirmation.CreatedAt.Time,
		ExpiresAt: emailConfirmation.ExpiresAt.Time,
		IsUsed:    emailConfirmation.IsUsed,
	}, nil
}

func (r SessionCodeRepository) Update(ctx context.Context, sessionCode *entities.SessionCode) error {
	err := r.Store.UpdateSessionCode(ctx, pgstore.UpdateSessionCodeParams{
		ID:        sessionCode.ID,
		UserID:    sessionCode.UserID,
		Token:     sessionCode.Token,
		CreatedAt: pgtype.Timestamp{Time: sessionCode.CreatedAt, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: sessionCode.ExpiresAt, Valid: true},
		IsUsed:    sessionCode.IsUsed,
	})

	return err
}

func (r SessionCodeRepository) DeleteByID(ctx context.Context, emailMfaCodeID uuid.UUID) error {
	err := r.Store.DeleteSessionCode(ctx, emailMfaCodeID)

	return err
}

func (r SessionCodeRepository) RevokeByUserID(ctx context.Context, userID uuid.UUID) error {
	err := r.Store.RevokeSessionCodeByUserID(ctx, userID)

	return err
}
