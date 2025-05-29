package verifyappmfa

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
	DeleteAppMfaCodeByID(ctx context.Context, appMfaCodeID uuid.UUID) error
	GetAppMfaCodeByID(ctx context.Context, id uuid.UUID) (*entities.AppMfaCode, error)
	GetUserByEmail(ctx context.Context, email string, applicationID uuid.UUID) (*entities.ApplicationUser, error)
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) DeleteAppMfaCodeByID(ctx context.Context, appMfaCodeID uuid.UUID) error {
	err := r.Store.DeleteAppMfaCode(ctx, appMfaCodeID)

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

func (r Repository) GetAppMfaCodeByID(ctx context.Context, id uuid.UUID) (*entities.AppMfaCode, error) {

	emailConfirmation, err := r.Store.GetAppMfaCodeByID(ctx, id)

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.AppMfaCode{
		ID:        emailConfirmation.ID,
		UserID:    emailConfirmation.UserID,
		Email:     emailConfirmation.Email,
		CreatedAt: emailConfirmation.CreatedAt.Time,
		ExpiresAt: emailConfirmation.ExpiresAt.Time,
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
