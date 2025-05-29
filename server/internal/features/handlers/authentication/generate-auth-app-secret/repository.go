package generateauthappsecret

import (
	"context"
	"strings"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	GetApplicationByID(ctx context.Context, applicationID uuid.UUID) (*entities.Application, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.ApplicationUser, error)
	UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error)
	AddMfaUserSecret(ctx context.Context, mfaUserSecret *entities.MfaUserSecret) error
	RevokeMfaUserSecret(ctx context.Context, userID uuid.UUID) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) GetApplicationByID(ctx context.Context, applicationID uuid.UUID) (*entities.Application, error) {
	application, err := r.Store.GetApplicationByID(ctx, applicationID)

	if err == repositories.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entities.Application{
		ID:                 application.ID,
		Name:               application.Name,
		Description:        application.Description,
		OrganizationID:     application.OrganizationID,
		CreatedAt:          application.CreatedAt.Time,
		IsActive:           application.IsActive,
		HasMfaAuthApp:      application.HasMfaAuthApp,
		HasMfaEmail:        application.HasMfaEmail,
		PasswordHashSecret: application.PasswordHashSecret,
		UpdatedAt:          application.UpdatedAt,
		Badges:             strings.Split(*application.Badges, ","),
		CanSelfSignUp:      application.CanSelfSignUp,
		CanSelfForgotPass:  application.CanSelfForgotPass,
	}, nil
}

func (r Repository) UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error) {
	now := time.Now().UTC()

	err := r.Store.UpdateUser(ctx, pgstore.UpdateUserParams{
		ID:                  user.ID,
		Email:               user.Email,
		PasswordHash:        user.PasswordHash,
		UpdatedAt:           &now,
		IsActive:            user.IsActive,
		IsEmailConfirmed:    user.IsEmailConfirmed,
		IsMfaAuthAppEnabled: user.IsMfaAuthAppEnabled,
		IsMfaEmailEnabled:   user.IsMfaEmailEnabled,
		TwoFactorSecret:     user.TwoFactorSecret,
		ShouldChangePass:    user.ShouldChangePass,
	})

	return user, err
}

func (r Repository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.ApplicationUser, error) {
	user, err := r.Store.GetUserById(ctx, id)

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
		IsMfaEmailEnabled:   user.IsMfaEmailEnabled,
		ApplicationID:       user.ApplicationID,
		ShouldChangePass:    user.ShouldChangePass,
		TwoFactorSecret:     user.TwoFactorSecret,
	}, nil
}

func (r Repository) AddMfaUserSecret(ctx context.Context, mfaUserSecret *entities.MfaUserSecret) error {
	err := r.Store.AddMfaUserSecret(ctx, pgstore.AddMfaUserSecretParams{
		ID:          mfaUserSecret.ID,
		UserID:      mfaUserSecret.UserID,
		Secret:      mfaUserSecret.Secret,
		IsValidated: mfaUserSecret.IsValidated,
		CreatedAt:   pgtype.Timestamp{Time: mfaUserSecret.CreatedAt, Valid: true},
		ExpiresAt:   pgtype.Timestamp{Time: mfaUserSecret.ExpiresAt, Valid: true},
	})

	return err
}

func (r Repository) RevokeMfaUserSecret(ctx context.Context, userID uuid.UUID) error {
	err := r.Store.RevokeMfaUserSecretFromUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil
}
