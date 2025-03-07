package repository_handlers

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type EmailConfirmationRepository struct {
	Store *pgstore.Queries
}

func (r EmailConfirmationRepository) AddEmailConfirmation(ctx context.Context, emailConfirmation *entities.EmailConfirmation) error {
	err := r.Store.AddEmailConfirmation(ctx, pgstore.AddEmailConfirmationParams{
		ID:        emailConfirmation.ID,
		UserID:    emailConfirmation.UserID,
		Email:     emailConfirmation.Email,
		Token:     emailConfirmation.Token,
		CreatedAt: pgtype.Timestamp{Time: emailConfirmation.CreatedAt, Valid: true},
		CoolDown:  pgtype.Timestamp{Time: emailConfirmation.CoolDown, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: emailConfirmation.ExpiresAt, Valid: true},
		IsUsed:    emailConfirmation.IsUsed,
	})

	return err
}

func (r EmailConfirmationRepository) GetByEmail(ctx context.Context, email string, userID uuid.UUID) (*entities.EmailConfirmation, error) {
	emailConfirmation, err := r.Store.GetEmailConfirmationByEmail(ctx, pgstore.GetEmailConfirmationByEmailParams{
		Email:  email,
		UserID: userID,
	})

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.EmailConfirmation{
		ID:        emailConfirmation.ID,
		UserID:    emailConfirmation.UserID,
		Email:     emailConfirmation.Email,
		Token:     emailConfirmation.Token,
		CreatedAt: emailConfirmation.CreatedAt.Time,
		CoolDown:  emailConfirmation.CoolDown.Time,
		ExpiresAt: emailConfirmation.ExpiresAt.Time,
		IsUsed:    emailConfirmation.IsUsed,
	}, nil
}

func (r EmailConfirmationRepository) UpdateEmailConfirmation(ctx context.Context, emailConfirmation *entities.EmailConfirmation) error {
	err := r.Store.UpdateEmailConfirmation(ctx, pgstore.UpdateEmailConfirmationParams{
		ID:        emailConfirmation.ID,
		UserID:    emailConfirmation.UserID,
		Email:     emailConfirmation.Email,
		Token:     emailConfirmation.Token,
		CreatedAt: pgtype.Timestamp{Time: emailConfirmation.CreatedAt, Valid: true},
		CoolDown:  pgtype.Timestamp{Time: emailConfirmation.CoolDown, Valid: true},
		ExpiresAt: pgtype.Timestamp{Time: emailConfirmation.ExpiresAt, Valid: true},
		IsUsed:    emailConfirmation.IsUsed,
	})

	return err
}

func (r EmailConfirmationRepository) DeleteEmailConfirmation(ctx context.Context, emailConfirmationID uuid.UUID) error {
	err := r.Store.DeleteEmailConfirmation(ctx, emailConfirmationID)

	return err
}
