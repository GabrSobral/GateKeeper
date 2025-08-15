package externalloginprovider

import (
	"context"

	"github.com/gate-keeper/internal/domain/entities"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	GetExternalLoginByProviderKey(ctx context.Context, provider, providerKey string) (*entities.ExternalLogin, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.ApplicationUser, error)
	GetUserByEmail(ctx context.Context, userEmail string, applicationID uuid.UUID) (*entities.ApplicationUser, error)
	GetUserProfileByID(ctx context.Context, userID uuid.UUID) (*entities.UserProfile, error)
	AddUser(ctx context.Context, newUser *entities.ApplicationUser) error
	AddUserProfile(ctx context.Context, newUserProfile *entities.UserProfile) error
	AddExternalLogin(ctx context.Context, externalLogin *entities.ExternalLogin) error
}

type Repository struct {
	Store *pgstore.Queries
}

func (r Repository) AddExternalLogin(ctx context.Context, externalLogin *entities.ExternalLogin) error {
	err := r.Store.AddExternalLogin(ctx, pgstore.AddExternalLoginParams{
		UserID:      externalLogin.UserID,
		Email:       externalLogin.Email,
		Provider:    externalLogin.Provider,
		ProviderKey: externalLogin.ProviderKey,
	})

	return err
}

func (r Repository) AddUserProfile(ctx context.Context, newUserProfile *entities.UserProfile) error {
	err := r.Store.AddUserProfile(ctx, pgstore.AddUserProfileParams{
		UserID:      newUserProfile.UserID,
		DisplayName: newUserProfile.DisplayName,
		FirstName:   newUserProfile.FirstName,
		LastName:    newUserProfile.LastName,
		Address:     newUserProfile.Address,
		PhoneNumber: newUserProfile.PhoneNumber,
		PhotoUrl:    newUserProfile.PhotoURL,
	})

	return err
}

func (r Repository) AddUser(ctx context.Context, newUser *entities.ApplicationUser) error {
	err := r.Store.AddUser(ctx, pgstore.AddUserParams{
		ID:               newUser.ID,
		Email:            newUser.Email,
		ApplicationID:    newUser.ApplicationID,
		ShouldChangePass: newUser.ShouldChangePass,
		PasswordHash:     newUser.PasswordHash,
		CreatedAt:        pgtype.Timestamp{Time: newUser.CreatedAt, Valid: true},
		UpdatedAt:        newUser.UpdatedAt,
		IsActive:         newUser.IsActive,
		IsEmailConfirmed: newUser.IsEmailConfirmed,
	})

	return err
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

func (r Repository) GetExternalLoginByProviderKey(ctx context.Context, provider, providerKey string) (*entities.ExternalLogin, error) {
	externalLogin, err := r.Store.GetExternalLoginByProviderKey(ctx, pgstore.GetExternalLoginByProviderKeyParams{
		Provider:    provider,
		ProviderKey: providerKey,
	})

	if err != nil {
		return nil, err
	}

	return &entities.ExternalLogin{
		UserID:      externalLogin.UserID,
		Email:       externalLogin.Email,
		Provider:    externalLogin.Provider,
		ProviderKey: externalLogin.ProviderKey,
	}, nil
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
