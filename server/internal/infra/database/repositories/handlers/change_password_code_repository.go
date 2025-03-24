package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ChangePasswordCodeRepository struct {
	Store *pgstore.Queries
}

func (r ChangePasswordCodeRepository) Add(ctx context.Context, changePasswordCode *entities.ChangePasswordCode) error {
	err := r.Store.AddChangePasswordCode(ctx, pgstore.AddChangePasswordCodeParams{
		ID:        changePasswordCode.ID,
		UserID:    changePasswordCode.UserID,
		Email:     changePasswordCode.Email,
		Token:     changePasswordCode.Token,
		CreatedAt: pgtype.Timestamp{Time: changePasswordCode.CreatedAt, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: changePasswordCode.ExpiresAt, Valid: true},
	})

	return err
}

func (r ChangePasswordCodeRepository) RevokeAllByID(ctx context.Context, userID uuid.UUID) error {
	err := r.Store.RevokeChangePasswordCodeByUserID(ctx, userID)

	return err
}

func (r ChangePasswordCodeRepository) GetByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.ChangePasswordCode, error) {

	emailConfirmation, err := r.Store.GetChangePasswordCodeByToken(ctx, pgstore.GetChangePasswordCodeByTokenParams{
		Token:  token,
		UserID: userID,
	})

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.ChangePasswordCode{
		ID:        emailConfirmation.ID,
		UserID:    emailConfirmation.UserID,
		Email:     emailConfirmation.Email,
		Token:     emailConfirmation.Token,
		CreatedAt: emailConfirmation.CreatedAt.Time,
		ExpiresAt: emailConfirmation.ExpiresAt.Time,
	}, nil
}

func (r ChangePasswordCodeRepository) Update(ctx context.Context, d *entities.ChangePasswordCode) error {
	err := r.Store.UpdateChangePasswordCode(ctx, pgstore.UpdateChangePasswordCodeParams{
		ID:        d.ID,
		UserID:    d.UserID,
		Email:     d.Email,
		Token:     d.Token,
		CreatedAt: pgtype.Timestamp{Time: d.CreatedAt, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: d.ExpiresAt, Valid: true},
	})

	return err
}

func (r ChangePasswordCodeRepository) DeleteByID(ctx context.Context, changePasswordID uuid.UUID) error {
	err := r.Store.DeleteChangePasswordCode(ctx, changePasswordID)

	return err
}
