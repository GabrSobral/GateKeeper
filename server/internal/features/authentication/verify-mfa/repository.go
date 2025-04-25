package verifymfa

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	AddSessionCode(ctx context.Context, sessionCode *entities.SessionCode) error
	DeleteEmailMfaCodeByID(ctx context.Context, emailMfaCodeID uuid.UUID) error
	GetEmailMfaCodeByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.EmailMfaCode, error)
	GetUserByEmail(ctx context.Context, email string, applicationID uuid.UUID) (*entities.ApplicationUser, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) DeleteEmailMfaCodeByID(ctx context.Context, emailMfaCodeID uuid.UUID) error {
	err := r.Store.DeleteEmailMfaCode(ctx, emailMfaCodeID)

	return err
}

func (r Repository) AddSessionCode(ctx context.Context, sessionCode *entities.SessionCode) error {
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

func (r Repository) GetEmailMfaCodeByToken(ctx context.Context, userID uuid.UUID, token string) (*entities.EmailMfaCode, error) {

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

func (r Repository) GetUserByEmail(ctx context.Context, email string, applicationID uuid.UUID) (*entities.ApplicationUser, error) {
	user, err := r.Store.GetUserByEmail(ctx, pgstore.GetUserByEmailParams{
		Email:         email,
		ApplicationID: applicationID,
	})

	if err != nil {
		return nil, err
	}

	return &entities.ApplicationUser{
		ID:                  user.ID,
		Email:               user.Email,
		PasswordHash:        user.PasswordHash,
		CreatedAt:           user.CreatedAt.Time,
		UpdatedAt:           user.UpdatedAt,
		IsActive:            user.IsActive,
		IsEmailConfirmed:    user.IsEmailConfirmed,
		IsMfaAuthAppEnabled: user.IsMfaAuthAppEnabled,
		ApplicationID:       user.ApplicationID,
		ShouldChangePass:    user.ShouldChangePass,
		IsMfaEmailEnabled:   user.IsMfaEmailEnabled,
		TwoFactorSecret:     user.TwoFactorSecret,
	}, nil
}
