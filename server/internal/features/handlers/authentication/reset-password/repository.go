package resetpassword

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
	RevokeRefreshTokenFromUser(ctx context.Context, userID uuid.UUID) error
	GetPasswordResetByTokenID(ctx context.Context, tokenID uuid.UUID) (*entities.PasswordResetToken, error)
	UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error)
	GetApplicationByID(ctx context.Context, applicationID uuid.UUID) (*entities.Application, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.ApplicationUser, error)
	DeletePasswordResetFromUser(ctx context.Context, userID uuid.UUID) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) RevokeRefreshTokenFromUser(ctx context.Context, userID uuid.UUID) error {
	err := r.Store.RevokeRefreshTokenFromUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil
}

func (pr Repository) GetPasswordResetByTokenID(ctx context.Context, tokenID uuid.UUID) (*entities.PasswordResetToken, error) {
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

func (r Repository) UpdateUser(ctx context.Context, user *entities.ApplicationUser) (*entities.ApplicationUser, error) {
	now := time.Now().UTC()

	err := r.Store.UpdateUser(ctx, pgstore.UpdateUserParams{
		ID:               user.ID,
		Email:            user.Email,
		PasswordHash:     user.PasswordHash,
		UpdatedAt:        &now,
		IsActive:         user.IsActive,
		IsEmailConfirmed: user.IsEmailConfirmed,
		ShouldChangePass: user.ShouldChangePass,
	})

	return user, err
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

func (pr Repository) CreatePasswordReset(ctx context.Context, passwordResetToken *entities.PasswordResetToken) error {
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

func (pr Repository) DeletePasswordResetFromUser(ctx context.Context, userID uuid.UUID) error {
	err := pr.Store.DeletePasswordResetFromUser(ctx, userID)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.ApplicationUser, error) {
	user, err := r.Store.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.ApplicationUser{
		ID:                 user.ID,
		Email:              user.Email,
		PasswordHash:       user.PasswordHash,
		CreatedAt:          user.CreatedAt.Time,
		UpdatedAt:          user.UpdatedAt,
		IsActive:           user.IsActive,
		IsEmailConfirmed:   user.IsEmailConfirmed,
		ApplicationID:      user.ApplicationID,
		ShouldChangePass:   user.ShouldChangePass,
		Preferred2FAMethod: user.Preferred2faMethod,
	}, nil
}
