package login

import (
	"context"
	"strings"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/infra/database/repositories"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	GetApplicationByID(ctx context.Context, applicationID uuid.UUID) (*entities.Application, error)
	GetUserProfileByID(ctx context.Context, userID uuid.UUID) (*entities.UserProfile, error)
	GetUserByEmail(ctx context.Context, userEmail string, applicationID uuid.UUID) (*entities.ApplicationUser, error)
	RevokeAllChangePasswordCodeByUserID(ctx context.Context, userID uuid.UUID) error
	AddEmailMfaCode(ctx context.Context, emailMfaCode *entities.EmailMfaCode) error
	AddSessionCode(ctx context.Context, sessionCode *entities.SessionCode) error
	AddChangePasswordCode(ctx context.Context, changePasswordCode *entities.ChangePasswordCode) error
}

type Repository struct {
	Store *pgstore.Queries
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

func (r Repository) AddEmailMfaCode(ctx context.Context, emailMfaCode *entities.EmailMfaCode) error {
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

func (r Repository) AddChangePasswordCode(ctx context.Context, changePasswordCode *entities.ChangePasswordCode) error {
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

func (r Repository) RevokeAllChangePasswordCodeByUserID(ctx context.Context, userID uuid.UUID) error {
	err := r.Store.RevokeChangePasswordCodeByUserID(ctx, userID)

	return err
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

func (r Repository) GetUserProfileByID(ctx context.Context, userID uuid.UUID) (*entities.UserProfile, error) {
	userProfile, err := r.Store.GetUserProfileByUserId(ctx, userID)

	if err != nil {
		return nil, err
	}

	return &entities.UserProfile{
		UserID:      userProfile.UserID,
		DisplayName: userProfile.DisplayName,
		FirstName:   userProfile.FirstName,
		LastName:    userProfile.LastName,
		Address:     userProfile.Address,
		PhoneNumber: userProfile.PhoneNumber,
		PhotoURL:    userProfile.PhotoUrl,
	}, nil
}
