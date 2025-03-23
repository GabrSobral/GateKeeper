package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type EmailMfaCodeRepository struct {
	Store *pgstore.Queries
}

func (r EmailMfaCodeRepository) Add(ctx context.Context, emailMfaCode *entities.EmailMfaCode) error {
	err := r.Store.AddEmailMfaCode(ctx, pgstore.AddEmailMfaCodeParams{
		ID:        emailMfaCode.ID,
		UserID:    emailMfaCode.UserID,
		Email:     emailMfaCode.Email,
		Token:     emailMfaCode.Token,
		CreatedAt: pgtype.Timestamp{Time: emailMfaCode.CreatedAt, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: emailMfaCode.ExpiresAt, Valid: true},
		IsUsed:    emailMfaCode.IsUsed,
	})

	return err
}

func (r EmailMfaCodeRepository) GetByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.EmailMfaCode, error) {

	emailConfirmation, err := r.Store.GetEmailMfaCodeByToken(ctx, pgstore.GetEmailMfaCodeByTokenParams{
		Token:  token,
		UserID: userID,
	})

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.EmailMfaCode{
		ID:        emailConfirmation.ID,
		UserID:    emailConfirmation.UserID,
		Email:     emailConfirmation.Email,
		Token:     emailConfirmation.Token,
		CreatedAt: emailConfirmation.CreatedAt.Time,
		ExpiresAt: emailConfirmation.ExpiresAt.Time,
		IsUsed:    emailConfirmation.IsUsed,
	}, nil
}

func (r EmailMfaCodeRepository) Update(ctx context.Context, emailMfaCode *entities.EmailMfaCode) error {
	err := r.Store.UpdateEmailMfaCode(ctx, pgstore.UpdateEmailMfaCodeParams{
		ID:        emailMfaCode.ID,
		UserID:    emailMfaCode.UserID,
		Email:     emailMfaCode.Email,
		Token:     emailMfaCode.Token,
		CreatedAt: pgtype.Timestamp{Time: emailMfaCode.CreatedAt, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: emailMfaCode.ExpiresAt, Valid: true},
		IsUsed:    emailMfaCode.IsUsed,
	})

	return err
}

func (r EmailMfaCodeRepository) DeleteByID(ctx context.Context, emailMfaCodeID uuid.UUID) error {
	err := r.Store.DeleteEmailMfaCode(ctx, emailMfaCodeID)

	return err
}
