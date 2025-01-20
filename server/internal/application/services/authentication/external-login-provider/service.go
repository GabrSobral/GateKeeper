package externalloginprovider

import (
	"context"
	"strings"
	"time"

	"github.com/gate-keeper/internal/domain/entities"
	"github.com/gate-keeper/internal/domain/errors"
	"github.com/gate-keeper/internal/infra/database/repositories"
	repository_handlers "github.com/gate-keeper/internal/infra/database/repositories/handlers"
	repository_interfaces "github.com/gate-keeper/internal/infra/database/repositories/interfaces"
	pgstore "github.com/gate-keeper/internal/infra/database/sqlc"
	"github.com/google/uuid"
)

type Request struct {
	Provider    string `json:"provider" validate:"required"`
	ProviderKey string `json:"provider_key" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
}

type Response struct {
	UserID           uuid.UUID  `json:"user_id"`
	UserEmail        string     `json:"user_email"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	CreatedAt        time.Time  `json:"created_at"`
	IsEmailConfirmed bool       `json:"is_email_confirmed"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type ExternalLoginProvider struct {
	UserRepository          repository_interfaces.IUserRepository
	UserProfileRepository   repository_interfaces.IUserProfileRepository
	ExternalLoginRepository repository_interfaces.IExternalLoginRepository
}

func New(q *pgstore.Queries) repositories.ServiceHandlerRs[Request, *Response] {
	return &ExternalLoginProvider{
		UserRepository:          repository_handlers.UserRepository{Store: q},
		UserProfileRepository:   repository_handlers.UserProfileRepository{Store: q},
		ExternalLoginRepository: repository_handlers.ExternalLoginRepository{Store: q},
	}
}

func (elp *ExternalLoginProvider) Handler(ctx context.Context, request Request) (*Response, error) {
	externalLogin, err := elp.ExternalLoginRepository.GetByProviderKey(ctx, request.Provider, request.ProviderKey)

	if err != nil {
		return nil, err
	}

	if externalLogin != nil {
		user, err := elp.UserRepository.GetUserByID(ctx, externalLogin.UserID)

		if err != nil {
			return nil, err
		}

		if user == nil {
			return nil, &errors.ErrUserNotFound
		}

		userProfile, err := elp.UserProfileRepository.GetUserById(ctx, user.ID)

		if err != nil {
			return nil, err
		}

		return &Response{
			UserID:           user.ID,
			UserEmail:        strings.ToLower(request.Email),
			FirstName:        userProfile.FirstName,
			LastName:         userProfile.LastName,
			CreatedAt:        user.CreatedAt,
			IsEmailConfirmed: user.IsEmailConfirmed,
			UpdatedAt:        nil,
		}, nil
	}

	user, err := elp.UserRepository.GetUserByEmail(ctx, request.Email)

	if err != nil {
		return nil, err
	}

	userProfile := &entities.UserProfile{}

	if user == nil {
		userID, err := uuid.NewV7()

		if err != nil {
			return nil, err
		}

		user = &entities.User{
			ID:               userID,
			Email:            request.Email,
			PasswordHash:     nil,
			CreatedAt:        time.Now(),
			UpdatedAt:        nil,
			IsActive:         true,
			IsEmailConfirmed: true,
			TwoFactorEnabled: false,
			TwoFactorSecret:  nil,
		}

		if err = elp.UserRepository.AddUser(ctx, user); err != nil {
			return nil, err
		}

		userProfile = &entities.UserProfile{
			UserID:      user.ID,
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			PhoneNumber: nil,
			Address:     nil,
			PhotoURL:    nil,
		}

		if err := elp.UserProfileRepository.AddUserProfile(ctx, userProfile); err != nil {
			return nil, err
		}
	} else {
		userProfileData, err := elp.UserProfileRepository.GetUserById(ctx, user.ID)

		if err != nil {
			return nil, err
		}

		userProfile = userProfileData
	}

	newExternalLogin := &entities.ExternalLogin{
		UserID:      user.ID,
		Email:       strings.ToLower(request.Email),
		Provider:    request.Provider,
		ProviderKey: request.ProviderKey,
	}

	if err = elp.ExternalLoginRepository.AddExternalLogin(ctx, newExternalLogin); err != nil {
		return nil, err
	}

	return &Response{
		UserID:           user.ID,
		UserEmail:        user.Email,
		FirstName:        userProfile.FirstName,
		LastName:         userProfile.LastName,
		CreatedAt:        user.CreatedAt,
		IsEmailConfirmed: user.IsEmailConfirmed,
		UpdatedAt:        nil,
	}, nil
}
